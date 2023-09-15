package db

import (
	"context"
	"testing"

	"github.com/DMonkey83/MyFitnessApp/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomExercise(t *testing.T) Exercise {
	arg := CreateExerciseParams{
		ExerciseName:      util.GetRandomUsername(5),
		EquipmentRequired: EquipmenttypeBarbell,
		MuscleGroupName:   MusclegroupenumAbs,
		Description:       util.GetRandomUsername(100),
	}

	exerc, err := testStore.CreateExercise(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, exerc)

	require.Equal(t, arg.EquipmentRequired, exerc.EquipmentRequired)
	require.Equal(t, arg.MuscleGroupName, exerc.MuscleGroupName)
	require.Equal(t, arg.Description, exerc.Description)
	require.Equal(t, arg.ExerciseName, exerc.ExerciseName)

	return exerc
}

func TestCreateExercise(t *testing.T) {
	CreateRandomExercise(t)
}

func TestGetExercise(t *testing.T) {
	ex1 := CreateRandomExercise(t)
	ex2, err := testStore.GetExercise(context.Background(), ex1.ExerciseName)
	require.NoError(t, err)
	require.NotEmpty(t, ex2)

	require.Equal(t, ex1.ExerciseName, ex2.ExerciseName)
	require.Equal(t, ex1.EquipmentRequired, ex2.EquipmentRequired)
	require.Equal(t, ex1.MuscleGroupName, ex2.MuscleGroupName)
	require.Equal(t, ex1.Description, ex2.Description)
}

func TestUpdateExercise(t *testing.T) {
	ex1 := CreateRandomExercise(t)

	arg := UpdateExerciseParams{
		ExerciseName:      ex1.ExerciseName,
		MuscleGroupName:   MusclegroupenumBack,
		Description:       util.GetRandomUsername(100),
		EquipmentRequired: EquipmenttypeMachine,
	}

	ex2, err := testStore.UpdateExercise(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, ex2)

	require.Equal(t, arg.EquipmentRequired, ex2.EquipmentRequired)
	require.Equal(t, arg.Description, ex2.Description)
	require.Equal(t, arg.MuscleGroupName, ex2.MuscleGroupName)
}

func TestDeleteExercise(t *testing.T) {
	ex1 := CreateRandomExercise(t)
	err := testStore.DeleteExercise(context.Background(), ex1.ExerciseName)
	require.NoError(t, err)

	ex2, err := testStore.GetExercise(context.Background(), ex1.ExerciseName)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, ex2)
}

func TestListExercises(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomExercise(t)
	}

	arg := ListAllExercisesParams{
		Limit:  5,
		Offset: 0,
	}

	exercises, err := testStore.ListAllExercises(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, exercises)

	for _, exercise := range exercises {
		require.NotEmpty(t, exercise)
	}
}

func TestEquipmentExercises(t *testing.T) {
	var lastAccount Exercise
	for i := 0; i < 10; i++ {
		lastAccount = CreateRandomExercise(t)
	}

	arg := ListEquipmentExercisesParams{
		EquipmentRequired: EquipmenttypeBarbell,
		Limit:             5,
		Offset:            0,
	}

	exercises, err := testStore.ListEquipmentExercises(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, exercises)

	for _, exercise := range exercises {
		require.NotEmpty(t, exercise)
		require.Equal(t, lastAccount.EquipmentRequired, exercise.EquipmentRequired)
	}
}
