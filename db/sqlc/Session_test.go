package db

import (
	"context"
	"testing"
	"time"

	"github.com/DMonkey83/MyFitnessApp/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func CreateRandomSession(t *testing.T) Session {
	user := CreateRandomUser(t)
	newUUid, err := uuid.NewUUID()
	require.NoError(t, err)
	arg := CreateSessionParams{
		ID:           newUUid,
		Username:     user.Username,
		RefreshToken: util.GetRandomUsername(30),
		UserAgent:    util.GetRandomUsername(8),
		ClientIp:     util.GetRandomUsername(10),
		IsBlocked:    false,
		ExpiresAt:    time.Now().Add(time.Hour),
	}

	session, err := testStore.CreateSession(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, session)

	require.Equal(t, arg.UserAgent, session.UserAgent)
	require.Equal(t, arg.Username, session.Username)
	require.Equal(t, arg.RefreshToken, session.RefreshToken)
	require.Equal(t, arg.IsBlocked, session.IsBlocked)

	require.NotZero(t, session.ID)
	return session
}

func TestCreateSession(t *testing.T) {
	CreateRandomSession(t)
}

func TestGeSession(t *testing.T) {
	session1 := CreateRandomSession(t)
	session2, err := testStore.GetSession(context.Background(), session1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, session1)

	require.Equal(t, session1.UserAgent, session2.UserAgent)
	require.Equal(t, session1.Username, session2.Username)
	require.Equal(t, session1.ClientIp, session2.ClientIp)
	require.Equal(t, session1.IsBlocked, session2.IsBlocked)
	require.Equal(t, session1.ExpiresAt, session2.ExpiresAt)
	require.Equal(t, session1.CreatedAt, session2.CreatedAt)
}
