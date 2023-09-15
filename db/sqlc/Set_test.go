package db

import (
	"context"
	"testing"

	"github.com/DMonkey83/MyFitnessApp/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomSet(t *testing.T) Set {
	ex := CreateRandomExercise(t)
	arg := CreateSetParams{
		ExerciseName: ex.ExerciseName,
		SetNumber:    int32(util.GetRandomAmount(1, 10)),
		Weight:       int32(util.GetRandomAmount(1, 200)),
		Notes:        util.GetRandomUsername(100),
		RestDuration: "1h",
	}

	set, err := testStore.CreateSet(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, set)

	require.Equal(t, arg.ExerciseName, set.ExerciseName)
	require.Equal(t, arg.Notes, set.Notes)
	require.Equal(t, arg.Weight, set.Weight)
	require.Equal(t, arg.Notes, set.Notes)

	require.NotZero(t, set.SetID)
	return set
}

func TestCreateSet(t *testing.T) {
	CreateRandomSet(t)
}

func TestGeSet(t *testing.T) {
	set1 := CreateRandomSet(t)
	set2, err := testStore.GetSet(context.Background(), set1.SetID)
	require.NoError(t, err)
	require.NotEmpty(t, set1)

	require.Equal(t, set1.ExerciseName, set2.ExerciseName)
	require.Equal(t, set1.Notes, set2.Notes)
	require.Equal(t, set1.Notes, set2.Notes)
	require.Equal(t, set1.Weight, set2.Weight)
}

func TestUpdateSet(t *testing.T) {
	set1 := CreateRandomSet(t)

	arg := UpdateSetParams{
		SetID:        set1.SetID,
		SetNumber:    int32(util.GetRandomAmount(1, 10)),
		Weight:       int32(util.GetRandomAmount(1, 200)),
		Notes:        util.GetRandomUsername(100),
		RestDuration: "1m",
	}

	set2, err := testStore.UpdateSet(context.Background(), arg)
	require.Equal(t, arg.Notes, set2.Notes)
	require.Equal(t, arg.Weight, set2.Weight)
	require.Equal(t, arg.SetNumber, set2.SetNumber)
	require.NoError(t, err)
	require.NotEmpty(t, set2)
}

func TestDeleteSet(t *testing.T) {
	set1 := CreateRandomSet(t)
	err := testStore.DeleteSet(context.Background(), set1.SetID)
	require.NoError(t, err)

	set2, err := testStore.GetSet(context.Background(), set1.SetID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, set2)
}

func TestSet(t *testing.T) {
	lastSet := CreateRandomSet(t)
	for i := 0; i < 10; i++ {
		lastSet = CreateRandomSet(t)
	}

	arg := ListSetsParams{
		ExerciseName: lastSet.ExerciseName,
		Limit:        5,
		Offset:       0,
	}

	exercises, err := testStore.ListSets(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, exercises)

	for _, exercise := range exercises {
		require.NotEmpty(t, exercise)
		require.Equal(t, lastSet.ExerciseName, exercise.ExerciseName)
	}
}
