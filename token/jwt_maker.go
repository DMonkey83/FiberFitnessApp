package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"

	"github.com/DMonkey83/FiberFitnessApp/config"
)

const minSecretKeySize = 32

// JWTMaker is a JSON Web Token maker
type JWTMaker struct {
	secretKey string
}

// define the name of the cookies, so that there is a reference + no possible grammar mistakes
var (
	AccessCookieName  = "access_cookie"
	RefreshCookieName = "refresh_cookie"
)

// NewJWTMaker creates a new JWTMaker
func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}
	return &JWTMaker{secretKey}, nil
}

// CreateToken creates a new token for a specific username and duration
func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", payload, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(maker.secretKey))
	return token, payload, err
}

// VerifyToken checks if the token is valid or not
func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}

// The time defined in the jwt claim needs to have the .UTC().Unix() so that the commented out function works
// ------------------
//
//	if !token.Valid {
//		return c.JSON(fiber.Map{
//			"jwt_exp": "The token has expired!",
//		})
//	}
//
// ------------------
func (maker *JWTMaker) GenerateRefreshCookie(token string, config config.Config, Username string, ctx *fiber.Ctx) error {
	RefreshTokenExpirationTime := RefreshTokenExpirationTime(config.RefreshTokenDuration)
	//
	// jwtKey := []byte(config.TokenSymmetricKey)
	//
	// refreshToken := jwt.New(jwt.SigningMethodHS256)
	// rtClaims := refreshToken.Claims.(jwt.MapClaims)
	//
	// // define the fields that are stored inside the refresh jwt payload
	// rtClaims["username"] = Username
	// rtClaims["exp"] = RefreshTokenExpirationTime.UTC().Unix()
	//
	// rt, err := refreshToken.SignedString(jwtKey)
	// if err != nil {
	// 	fmt.Println("Error creting refreshToken!")
	// }

	refresh_cookie := GenerateCookie(RefreshCookieName, token, RefreshTokenExpirationTime, ctx)
	ctx.Cookie(&refresh_cookie)

	// if err != nil {
	// 	fmt.Println("Error")
	// }
	return nil
}

func (maker *JWTMaker) GenerateAccessCookie(token string, config config.Config, Username string, ctx *fiber.Ctx) error {
	AccessTokenExpirationTime := AccessTokenExpirationTime(config.AccessTokenDuration)

	// jwtKey := []byte(config.TokenSymmetricKey)
	//
	// accessToken := jwt.New(jwt.SigningMethodHS256)
	// atClaims := accessToken.Claims.(jwt.MapClaims)
	//
	// // define the info that is stored inside the access jwt payload
	// atClaims["username"] = Username
	// atClaims["exp"] = AccessTokenExpirationTime.UTC().Unix()
	// at, err := accessToken.SignedString(jwtKey)
	// if err != nil {
	// 	fmt.Println("Error creting accessToken!")
	//}

	access_cookie := GenerateCookie(AccessCookieName, token, AccessTokenExpirationTime, ctx)
	ctx.Cookie(&access_cookie) // send cookie to the client

	return nil
}

func GenerateCookie(cookie_name string, cookie_value string, exp_time time.Time, c *fiber.Ctx) fiber.Cookie {
	return fiber.Cookie{
		Name:     cookie_name,
		Value:    cookie_value,
		HTTPOnly: true,
		Secure:   true,
		Expires:  exp_time,
	}
}

func AccessTokenExpirationTime(duration time.Duration) time.Time {
	expiration_time := time.Now().Add(duration * time.Minute)
	return expiration_time
}

func RefreshTokenExpirationTime(duration time.Duration) time.Time {
	expiration_time := time.Now().Add(duration)
	return expiration_time
}
