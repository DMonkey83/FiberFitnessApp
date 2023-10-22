package api

import (
	"errors"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	db "github.com/DMonkey83/FiberFitnessApp/db/sqlc"
	val "github.com/DMonkey83/FiberFitnessApp/util/Validate"
	"github.com/DMonkey83/FiberFitnessApp/util/auth"
	res "github.com/DMonkey83/FiberFitnessApp/util/response"
)

// createUserRequest struct
type createUserRequest struct {
	Username      string         `json:"username" validate:"required,alphanum,min=6,max=50"`
	FullName      string         `json:"full_name" validate:"required,min=6,max=50"`
	Age           int32          `json:"age" validate:"required,min=1,max=120"`
	Gender        string         `json:"gender" validate:"required,oneof=female male"`
	HeightCm      int32          `json:"height_cm" validate:"required,min=1,max=300"`
	HeightFtIn    string         `json:"height_ft_in"`
	PreferredUnit val.Weightunit `json:"preferred_unit" validate:"required"`
	Password      string         `json:"password" validate:"required,min=6,max=50"`
	Email         string         `json:"email" validate:"required,email"`
}

type getUserRequest struct {
	Username string `uri:"username" binding:"required,min=1"`
}

// loginUserRequest struct
type loginUserRequest struct {
	Username string `json:"username" validate:"required,alphanum,min=6,max=50"`
	Password string `json:"password" validate:"required,min=6,max=50"`
}

// userResponse struct
type userResponse struct {
	Username          string    `json:"username"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

// newUserResponse function
func newUserResponse(user db.User) userResponse {
	return userResponse{
		Username:          user.Username,
		Email:             user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}
}

// loginUserResponse struct  
type loginUserResponse struct {
	SessionID             uuid.UUID    `json:"session_id"`
	AccessToken           string       `json:"access_token"`
	AccessTokenExpiresAt  time.Time    `json:"access_token_expires_at"`
	RefreshToken          string       `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time    `json:"refresh_token_expires_at"`
	User                  userResponse `json:"user"`
}

// createUser method
func (server *Server) createUser(ctx *fiber.Ctx) error {
	req := new(createUserRequest)
	log.Print("createUser")
	if err := val.ValidatePayload(ctx, req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}

	hashedPasswrd, err := auth.HashPassword(req.Password)
	if err != nil {
		return res.ResponseError(ctx, nil, "Error hashing the password")
	}

	arg := db.CreateUserParams{
		Username:     req.Username,
		PasswordHash: hashedPasswrd,
		Email:        req.Email,
	}
	user, err := server.store.CreateUser(ctx.Context(), arg)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			return res.ResponseError(ctx, nil, "Username, or email already taken!")
		}
		return res.ResponseError(ctx, nil, "Eror accured during registration!")
	}

	arg1 := db.CreateUserProfileParams{
		Username:      req.Username,
		FullName:      req.FullName,
		Age:           req.Age,
		Gender:        req.Gender,
		HeightCm:      req.HeightCm,
		HeightFtIn:    req.HeightFtIn,
		PreferredUnit: db.Weightunit(req.PreferredUnit),
	}
	profile, err := server.store.CreateUserProfile(ctx.Context(), arg1)
	if err != nil {
		errorCode := db.ErrorCode(err)
		if errorCode == db.UniqueViolation {
			return res.ResponseUnauthenticated(ctx, nil, err.Error())
		}
		return res.ResponseError(ctx, nil, "Error while creating user profile")
	}
	log.Print(profile)

	rsp := newUserResponse(user)

	return res.ResponseSuccess(ctx, rsp, res.CreatedMessage("User created!"))
}

func (server *Server) LoginUser(ctx *fiber.Ctx) error {
	req := new(loginUserRequest)
	if err := val.ValidatePayload(ctx, req); err != nil {
		return res.ResponseError(ctx, nil, err.Error())
	}

	user, err := server.store.GetUser(ctx.Context(), req.Username)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return res.ResponseError(ctx, nil, "User not found!")
		}
		return res.ResponseError(ctx, nil, "Eror accured during Login!")
	}
	accessToken, _, err := server.tokenMaker.CreateToken(
		user.Username,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		return res.ResponseError(ctx, nil, "Eror accured while creating access token!")
	}

	refreshToken, refresherPayload, err := server.tokenMaker.CreateToken(
		user.Username,
		server.config.RefreshTokenDuration,
	)
	if err != nil {
		return res.ResponseError(ctx, nil, "Eror accured while creating refresh token!")
	}
	err = auth.CheckPassword(req.Password, user.PasswordHash)
	if err != nil {
		return res.ResponseUnauthenticated(ctx, nil, "Incorrect password!")
	}

	log.Print(refresherPayload, ctx)
	// if the username exists and the passwords match, set JWT Auth cookies with the user details.
	server.tokenMaker.GenerateAccessCookie(accessToken, server.config, user.Username, ctx)
	server.tokenMaker.GenerateRefreshCookie(refreshToken, server.config, user.Username, ctx)
	session, err := server.store.CreateSession(ctx.Context(), db.CreateSessionParams{
		ID:           refresherPayload.ID,
		Username:     user.Username,
		RefreshToken: refreshToken,
		UserAgent:    ctx.Get("User-Agent"),
		ClientIp:     ctx.IP(),
		IsBlocked:    false,
		ExpiresAt:    refresherPayload.ExpiredAt,
	})
	if err != nil {
		return res.ResponseUnauthenticated(ctx, nil, "Error accured while creating session!")
	}

	rsp := loginUserResponse{
		SessionID:            session.ID,
		RefreshToken:         refreshToken,
		AccessToken:          accessToken,
		AccessTokenExpiresAt: refresherPayload.ExpiredAt,
		User:                 newUserResponse(user),
	}

	return res.ResponseSuccess(ctx, rsp, "Login succesful!")
}
