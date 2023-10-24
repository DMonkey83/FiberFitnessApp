package db

import (
	"context"
	"testing"

	"github.com/DMonkey83/MyFitnessApp/util"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func CreateRandomMaxWeightGoal(t *testing.T) Maxweightgoal {
	ru := CreateRandomUser(t)
	ex := CreateRandomExercise(t)
	arg := CreateMaxWeightGoalParams{
		Username:     ru.Username,
		ExerciseName: ex.ExerciseName,
		GoalWeight:   int32(util.GetRandomAmount(1, 100)),
		Notes:        util.GetRandomUsername(100),
	}

	maxWG, err := testStore.CreateMaxWeightGoal(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, maxWG)

	require.Equal(t, arg.ExerciseName, maxWG.ExerciseName)
	require.Equal(t, arg.Notes, maxWG.Notes)
	require.Equal(t, arg.GoalWeight, maxWG.GoalWeight)
	require.Equal(t, arg.Username, maxWG.Username)

	require.NotZero(t, maxWG.GoalID)
	return maxWG
}

func TestCreateMaxWeightGoal(t *testing.T) {
	CreateRandomMaxWeightGoal(t)
}

func TestGetMaxWeightGoal(t *testing.T) {
	mwg1 := CreateRandomMaxWeightGoal(t)
	arg := GetMaxWeightGoalParams{
		ExerciseName: mwg1.ExerciseName,
		Username:     mwg1.Username,
		GoalID:       mwg1.GoalID,
	}
	mwg2, err := testStore.GetMaxWeightGoal(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, mwg1)

	require.Equal(t, mwg1.ExerciseName, mwg2.ExerciseName)
	require.Equal(t, mwg1.Notes, mwg2.Notes)
	require.Equal(t, mwg1.Username, mwg2.Username)
	require.Equal(t, mwg1.GoalWeight, mwg2.GoalWeight)
}

func TestUpdateMaxWeightGoal(t *testing.T) {
	mwg1 := CreateRandomMaxWeightGoal(t)

	arg := UpdateMaxWeightGoalParams{
		ExerciseName: mwg1.ExerciseName,
		GoalWeight:   pgtype.Int4{Int32: int32(util.GetRandomAmount(2, 20)), Valid: true},
		Notes:        pgtype.Text{String: util.GetRandomUsername(20), Valid: true},
		Username:     mwg1.Username,
		GoalID:       mwg1.GoalID,
	}

	mwg2, err := testStore.UpdateMaxWeightGoal(context.Background(), arg)
	require.Equal(t, mwg1.ExerciseName, mwg2.ExerciseName)
	require.Equal(t, arg.Notes.String, mwg2.Notes)
	require.Equal(t, mwg1.Username, mwg2.Username)
	require.Equal(t, arg.GoalWeight.Int32, mwg2.GoalWeight)
	require.NoError(t, err)
	require.NotEmpty(t, mwg2)
}

func TestDeleteMaxWeightGoal(t *testing.T) {
	mwg1 := CreateRandomMaxWeightGoal(t)
	arg1 := DeleteMaxWeightGoalParams{
		ExerciseName: mwg1.ExerciseName,
		Username:     mwg1.Username,
		GoalID:       mwg1.GoalID,
	}

	err := testStore.DeleteMaxWeightGoal(context.Background(), arg1)
	require.NoError(t, err)

	arg2 := GetMaxWeightGoalParams{
		ExerciseName: mwg1.ExerciseName,
		Username:     mwg1.Username,
		GoalID:       mwg1.GoalID,
	}

	mwg2, err := testStore.GetMaxWeightGoal(context.Background(), arg2)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, mwg2)
}
