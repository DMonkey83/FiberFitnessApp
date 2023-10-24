package db

import (
	"context"
	"testing"

	"github.com/DMonkey83/MyFitnessApp/util"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func CreateRandomExerciseSet(t *testing.T) Exerciseset {
	exercise := CreateRandomExerciseLog(t)
	arg := CreateExerciseSetParams{
		ExerciseLogID:        exercise.ExerciseLogID,
		SetNumber:            int32(util.GetRandomAmount(1, 10)),
		WeightLifted:         int32(util.GetRandomAmount(1, 100)),
		RepetitionsCompleted: int32(util.GetRandomAmount(1, 2000)),
	}

	exSet, err := testStore.CreateExerciseSet(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, exSet)

	require.Equal(t, arg.ExerciseLogID, exSet.ExerciseLogID)
	require.Equal(t, arg.RepetitionsCompleted, exSet.RepetitionsCompleted)
	require.Equal(t, arg.SetNumber, exSet.SetNumber)
	require.Equal(t, arg.WeightLifted, exSet.WeightLifted)

	require.NotEmpty(t, exSet.SetID)

	return exSet
}

func TestCreateExerciseSet(t *testing.T) {
	CreateRandomExerciseSet(t)
}

func TestGetExerciseSet(t *testing.T) {
	set1 := CreateRandomExerciseSet(t)
	set2, err := testStore.GetExerciseSet(context.Background(), set1.SetID)
	require.NoError(t, err)
	require.NotEmpty(t, set2)

	require.Equal(t, set1.ExerciseLogID, set2.ExerciseLogID)
	require.Equal(t, set1.SetNumber, set2.SetNumber)
	require.Equal(t, set1.WeightLifted, set2.WeightLifted)
	require.Equal(t, set1.RepetitionsCompleted, set2.RepetitionsCompleted)
}

func TestUpdateExerciseSet(t *testing.T) {
	set1 := CreateRandomExerciseSet(t)

	arg := UpdateExerciseSetParams{
		SetID:                set1.SetID,
		RepetitionsCompleted: pgtype.Int4{Int32: int32(util.GetRandomAmount(1, 100)), Valid: true},
		WeightLifted:         pgtype.Int4{Int32: int32(util.GetRandomAmount(1, 2000)), Valid: true},
	}

	set2, err := testStore.UpdateExerciseSet(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, set2)

	require.Equal(t, set1.ExerciseLogID, set2.ExerciseLogID)
	require.Equal(t, set1.SetID, set2.SetID)
	require.Equal(t, arg.RepetitionsCompleted.Int32, set2.RepetitionsCompleted)
	require.Equal(t, arg.WeightLifted.Int32, set2.WeightLifted)
}

func TestDeleteExerciseSet(t *testing.T) {
	set1 := CreateRandomExerciseSet(t)
	err := testStore.DeleteExerciseSet(context.Background(), set1.SetID)
	require.NoError(t, err)

	set2, err := testStore.GetExerciseSet(context.Background(), set1.SetID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, set2)
}

func TestListExerciseSets(t *testing.T) {
	lastSet := CreateRandomExerciseSet(t)
	for i := 0; i < 10; i++ {
		lastSet = CreateRandomExerciseSet(t)
	}

	arg := ListExerciseSetsParams{
		Limit:         5,
		Offset:        0,
		ExerciseLogID: lastSet.ExerciseLogID,
	}

	exercises, err := testStore.ListExerciseSets(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, exercises)

	for _, exercise := range exercises {
		require.NotEmpty(t, exercise)
		require.Equal(t, lastSet.ExerciseLogID, exercise.ExerciseLogID)
	}
}
