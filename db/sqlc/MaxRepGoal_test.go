package db

import (
	"context"
	"testing"

	"github.com/DMonkey83/MyFitnessApp/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomMaxRepGoal(t *testing.T) Maxrepgoal {
	ru := CreateRandomUser(t)
	ex := CreateRandomExercise(t)
	arg := CreateMaxRepGoalParams{
		Username:     ru.Username,
		ExerciseName: ex.ExerciseName,
		GoalReps:     int32(util.GetRandomAmount(1, 100)),
		Notes:        util.GetRandomUsername(100),
	}

	maxRepG, err := testStore.CreateMaxRepGoal(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, maxRepG)

	require.Equal(t, arg.ExerciseName, maxRepG.ExerciseName)
	require.Equal(t, arg.Notes, maxRepG.Notes)
	require.Equal(t, arg.GoalReps, maxRepG.GoalReps)
	require.Equal(t, arg.Username, maxRepG.Username)

	require.NotZero(t, maxRepG.GoalID)
	return maxRepG
}

func TestCreateMaxRepGoal(t *testing.T) {
	CreateRandomMaxRepGoal(t)
}

func TestGetMaxRepGoal(t *testing.T) {
	mrg1 := CreateRandomMaxRepGoal(t)
	arg := GetMaxRepGoalParams{
		ExerciseName: mrg1.ExerciseName,
		Username:     mrg1.Username,
		GoalID:       mrg1.GoalID,
	}
	mrg2, err := testStore.GetMaxRepGoal(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, mrg1)

	require.Equal(t, mrg1.ExerciseName, mrg2.ExerciseName)
	require.Equal(t, mrg1.Notes, mrg2.Notes)
	require.Equal(t, mrg1.Username, mrg2.Username)
	require.Equal(t, mrg1.GoalReps, mrg2.GoalReps)
}

func TestUpdateMaxRepGoal(t *testing.T) {
	mrg1 := CreateRandomMaxRepGoal(t)

	arg := UpdateMaxRepGoalParams{
		ExerciseName: mrg1.ExerciseName,
		Username:     mrg1.Username,
		GoalReps:     mrg1.GoalReps,
		Notes:        mrg1.Notes,
		GoalID:       mrg1.GoalID,
	}

	mrg2, err := testStore.UpdateMaxRepGoal(context.Background(), arg)
	require.Equal(t, mrg1.ExerciseName, mrg2.ExerciseName)
	require.Equal(t, mrg1.Notes, mrg2.Notes)
	require.Equal(t, mrg1.Username, mrg2.Username)
	require.Equal(t, mrg1.GoalReps, mrg2.GoalReps)
	require.NoError(t, err)
	require.NotEmpty(t, mrg2)

}

func TestDeleteMaxRepGoal(t *testing.T) {
	mrg1 := CreateRandomMaxRepGoal(t)
	arg1 := DeleteMaxRepGoalParams{
		ExerciseName: mrg1.ExerciseName,
		Username:     mrg1.Username,
		GoalID:       mrg1.GoalID,
	}
	err := testStore.DeleteMaxRepGoal(context.Background(), arg1)
	require.NoError(t, err)

	arg2 := GetMaxRepGoalParams{
		ExerciseName: mrg1.ExerciseName,
		Username:     mrg1.Username,
		GoalID:       mrg1.GoalID,
	}

	mrg2, err := testStore.GetMaxRepGoal(context.Background(), arg2)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, mrg2)
}
