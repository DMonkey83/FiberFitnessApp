package db

import (
	"context"
	"testing"

	"github.com/DMonkey83/MyFitnessApp/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomUserProfile(t *testing.T) Userprofile {
	user := CreateRandomUser(t)

	arg := CreateUserProfileParams{
		Username:      user.Username,
		FullName:      util.GetRandomUsername(15),
		Age:           int32(util.GetRandomAmount(16, 50)),
		Gender:        "male",
		HeightCm:      int32(util.GetRandomAmount(150, 220)),
		PreferredUnit: WeightunitKg,
	}

	userProfile, err := testStore.CreateUserProfile(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, userProfile)

	require.Equal(t, arg.Username, userProfile.Username)
	require.Equal(t, arg.FullName, userProfile.FullName)
	require.Equal(t, arg.Age, userProfile.Age)
	require.Equal(t, arg.PreferredUnit, userProfile.PreferredUnit)

	require.NotZero(t, user.Username)
	return userProfile
}

func TestCreateUserProfile(t *testing.T) {
	CreateRandomUserProfile(t)
}

func TestGetUserProfile(t *testing.T) {
	userP1 := CreateRandomUserProfile(t)
	userP2, err := testStore.GetUserProfile(context.Background(), userP1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, userP2)

	require.Equal(t, userP1.Username, userP2.Username)
	require.Equal(t, userP1.UserProfileID, userP2.UserProfileID)
	require.Equal(t, userP1.FullName, userP2.FullName)
	require.Equal(t, userP1.Gender, userP2.Gender)
	require.Equal(t, userP1.Age, userP2.Age)
}

func TestUpdateUserProfile(t *testing.T) {
	user1 := CreateRandomUserProfile(t)

	arg := UpdateUserProfileParams{
		Username:      user1.Username,
		FullName:      util.GetRandomEmail(8),
		Age:           int32(util.GetRandomAmount(16, 50)),
		Gender:        "female",
		HeightCm:      int32(util.GetRandomAmount(150, 220)),
		PreferredUnit: WeightunitKg,
	}

	user2, err := testStore.UpdateUserProfile(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, arg.Username, user2.Username)
	require.Equal(t, arg.FullName, user2.FullName)
	require.Equal(t, arg.Age, user2.Age)
}

func TestDeleteUserProfile(t *testing.T) {
	user1 := CreateRandomUserProfile(t)
	err := testStore.DeleteUserProfile(context.Background(), user1.Username)
	require.NoError(t, err)

	user2, err := testStore.GetUserProfile(context.Background(), user1.Username)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, user2)
}
