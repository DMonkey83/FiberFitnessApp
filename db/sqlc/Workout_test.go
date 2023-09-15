package db

import (
	"context"
	"testing"
	"time"

	"github.com/DMonkey83/MyFitnessApp/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomWorkout(t *testing.T) Workout {
	user := CreateRandomUser(t)

	arg := CreateWorkoutParams{
		Username:            user.Username,
		Notes:               util.GetRandomUsername(100),
		WorkoutDuration:     "0h10m",
		WorkoutDate:         time.Now(),
		FatigueLevel:        FatiguelevelHeavy,
		TotalCaloriesBurned: int32(util.GetRandomAmount(1, 3000)),
		TotalDistance:       int32(util.GetRandomAmount(1, 10000)),
		TotalRepetitions:    int32(util.GetRandomAmount(1, 100)),
		TotalSets:           int32(util.GetRandomAmount(1, 20)),
		TotalWeightLifted:   int32(util.GetRandomAmount(1, 2000)),
	}

	workout, err := testStore.CreateWorkout(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, workout)

	require.Equal(t, arg.Username, workout.Username)
	require.Equal(t, arg.Notes, workout.Notes)
	require.Equal(t, arg.WorkoutDuration, workout.WorkoutDuration)

	require.NotZero(t, workout.WorkoutID)
	return workout
}

func TestCreateWorkout(t *testing.T) {
	CreateRandomWorkout(t)
}

func TestGetWorkout(t *testing.T) {
	wkout1 := CreateRandomWorkout(t)
	wkout2, err := testStore.GetWorkout(context.Background(), wkout1.WorkoutID)
	require.NoError(t, err)
	require.NotEmpty(t, wkout2)

	require.Equal(t, wkout1.Username, wkout2.Username)
	require.Equal(t, wkout1.Notes, wkout2.Notes)
	require.Equal(t, wkout1.WorkoutDuration, wkout2.WorkoutDuration)
}

func TestUpdateWorkout(t *testing.T) {
	wkout1 := CreateRandomWorkout(t)

	arg := UpdateWorkoutParams{
		WorkoutID:           wkout1.WorkoutID,
		Notes:               wkout1.Notes,
		WorkoutDuration:     "1h",
		WorkoutDate:         time.Now(),
		FatigueLevel:        FatiguelevelLight,
		TotalCaloriesBurned: int32(util.GetRandomAmount(1, 3000)),
		TotalDistance:       int32(util.GetRandomAmount(1, 10000)),
		TotalRepetitions:    int32(util.GetRandomAmount(1, 100)),
		TotalSets:           int32(util.GetRandomAmount(1, 20)),
		TotalWeightLifted:   int32(util.GetRandomAmount(1, 2000)),
	}

	wkout2, err := testStore.UpdateWorkout(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, wkout2)

	require.Equal(t, arg.WorkoutID, wkout2.WorkoutID)
	require.Equal(t, arg.WorkoutDuration, wkout2.WorkoutDuration)
}

func TestDeleteWorkout(t *testing.T) {
	wkout1 := CreateRandomWorkout(t)
	err := testStore.DeleteWorkout(context.Background(), wkout1.WorkoutID)
	require.NoError(t, err)

	wkout2, err := testStore.GetWorkout(context.Background(), wkout1.WorkoutID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, wkout2)
}

func TestListWorkouts(t *testing.T) {
	lastWorkout := CreateRandomWorkout(t)
	arg := ListWorkoutsParams{
		Username: lastWorkout.Username,
		Limit:    5,
		Offset:   0,
	}

	workouts, err := testStore.ListWorkouts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, workouts)

	for _, workout := range workouts {
		require.NotEmpty(t, workout)
		require.Equal(t, lastWorkout.Username, workout.Username)
	}
}
