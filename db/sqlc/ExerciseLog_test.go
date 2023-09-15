package db

import (
	"context"
	"testing"

	"github.com/DMonkey83/MyFitnessApp/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomExerciseLog(t *testing.T) Exerciselog {
	exercise := CreateRandomExercise(t)
	log := CreateRandomWorkoutLog(t)
	arg := CreateExerciseLogParams{
		ExerciseName:         exercise.ExerciseName,
		LogID:                log.LogID,
		SetsCompleted:        int32(util.GetRandomAmount(1, 10)),
		RepetitionsCompleted: int32(util.GetRandomAmount(1, 100)),
		WeightLifted:         int32(util.GetRandomAmount(1, 2000)),
		Notes:                util.GetRandomUsername(100),
	}

	exLog, err := testStore.CreateExerciseLog(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, exLog)

	require.Equal(t, arg.ExerciseName, exLog.ExerciseName)
	require.Equal(t, arg.LogID, exLog.LogID)
	require.Equal(t, arg.SetsCompleted, exLog.SetsCompleted)
	require.Equal(t, arg.RepetitionsCompleted, exLog.RepetitionsCompleted)
	require.Equal(t, arg.WeightLifted, exLog.WeightLifted)
	require.Equal(t, arg.Notes, exLog.Notes)

	require.NotEmpty(t, exLog.ExerciseLogID)

	return exLog
}

func TestCreateExerciseLog(t *testing.T) {
	CreateRandomExerciseLog(t)
}

func TestGetExerciseLog(t *testing.T) {
	ex1 := CreateRandomExerciseLog(t)
	ex2, err := testStore.GetExerciseLog(context.Background(), ex1.ExerciseLogID)
	require.NoError(t, err)
	require.NotEmpty(t, ex2)

	require.Equal(t, ex1.ExerciseName, ex2.ExerciseName)
	require.Equal(t, ex1.LogID, ex2.LogID)
	require.Equal(t, ex1.SetsCompleted, ex2.SetsCompleted)
	require.Equal(t, ex1.RepetitionsCompleted, ex2.RepetitionsCompleted)
	require.Equal(t, ex1.WeightLifted, ex2.WeightLifted)
	require.Equal(t, ex1.Notes, ex2.Notes)
}

func TestUpdateExerciseLog(t *testing.T) {
	ex1 := CreateRandomExerciseLog(t)

	arg := UpdateExerciseLogParams{
		ExerciseLogID:        ex1.ExerciseLogID,
		SetsCompleted:        int32(util.GetRandomAmount(1, 10)),
		RepetitionsCompleted: int32(util.GetRandomAmount(1, 100)),
		WeightLifted:         int32(util.GetRandomAmount(1, 2000)),
		Notes:                util.GetRandomUsername(100),
	}

	ex2, err := testStore.UpdateExerciseLog(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, ex2)

	require.Equal(t, ex1.ExerciseName, ex2.ExerciseName)
	require.Equal(t, ex1.LogID, ex2.LogID)
	require.Equal(t, arg.SetsCompleted, ex2.SetsCompleted)
	require.Equal(t, arg.RepetitionsCompleted, ex2.RepetitionsCompleted)
	require.Equal(t, arg.WeightLifted, ex2.WeightLifted)
	require.Equal(t, arg.Notes, ex2.Notes)
}

func TestDeleteExerciseLog(t *testing.T) {
	ex1 := CreateRandomExerciseLog(t)
	err := testStore.DeleteExerciseLog(context.Background(), ex1.ExerciseLogID)
	require.NoError(t, err)

	ex2, err := testStore.GetExerciseLog(context.Background(), ex1.ExerciseLogID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, ex2)
}

func TestListExerciseLogs(t *testing.T) {
	lastLog := CreateRandomExerciseLog(t)
	for i := 0; i < 10; i++ {
		lastLog = CreateRandomExerciseLog(t)
	}

	arg := ListExerciseLogParams{
		LogID:  lastLog.LogID,
		Limit:  5,
		Offset: 0,
	}

	exercises, err := testStore.ListExerciseLog(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, exercises)

	for _, exercise := range exercises {
		require.NotEmpty(t, exercise)
		require.Equal(t, lastLog.LogID, exercise.LogID)
	}
}
