package db

import (
	"context"
	"testing"

	"github.com/DMonkey83/MyFitnessApp/util"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func CreateRandomWorkoutLog(t *testing.T) Workoutlog {
	plan := CreateRandomWorkoutPlan(t)

	arg := CreateWorkoutLogParams{
		Username:            plan.Username,
		PlanID:              plan.PlanID,
		TotalSets:           int32(util.GetRandomAmount(1, 20)),
		FatigueLevel:        FatiguelevelLight,
		TotalWeightLifted:   int32(util.GetRandomAmount(1, 2000)),
		TotalRepetitions:    int32(util.GetRandomAmount(1, 200)),
		TotalDistance:       int32(util.GetRandomAmount(1, 20000)),
		TotalCaloriesBurned: int32(util.GetRandomAmount(1, 2000)),
		WorkoutDuration:     "4h",
		Rating:              Rating1,
		Comments:            util.GetRandomUsername(100),
		OverallFeeling:      "gool",
	}

	wlog, err := testStore.CreateWorkoutLog(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, wlog)

	require.Equal(t, arg.Username, wlog.Username)
	require.Equal(t, arg.WorkoutDuration, wlog.WorkoutDuration)
	require.Equal(t, arg.PlanID, wlog.PlanID)
	require.Equal(t, arg.FatigueLevel, wlog.FatigueLevel)
	require.Equal(t, arg.TotalCaloriesBurned, wlog.TotalCaloriesBurned)
	require.Equal(t, arg.TotalDistance, wlog.TotalDistance)
	require.Equal(t, arg.TotalRepetitions, wlog.TotalRepetitions)
	require.Equal(t, arg.TotalWeightLifted, wlog.TotalWeightLifted)
	require.Equal(t, arg.TotalSets, wlog.TotalSets)
	require.Equal(t, arg.Rating, wlog.Rating)
	require.Equal(t, arg.Comments, wlog.Comments)
	require.Equal(t, arg.OverallFeeling, wlog.OverallFeeling)

	require.NotZero(t, wlog.LogID)
	return wlog
}

func TestCreateWorkoutLog(t *testing.T) {
	CreateRandomWorkoutLog(t)
}

func TestGetWorkoutLog(t *testing.T) {
	wlog1 := CreateRandomWorkoutLog(t)
	wlog2, err := testStore.GetWorkoutLog(context.Background(), wlog1.LogID)
	require.NoError(t, err)
	require.NotEmpty(t, wlog2)

	require.Equal(t, wlog1.Username, wlog2.Username)
	require.Equal(t, wlog1.WorkoutDuration, wlog2.WorkoutDuration)
	require.Equal(t, wlog1.PlanID, wlog2.PlanID)
	require.Equal(t, wlog1.FatigueLevel, wlog2.FatigueLevel)
	require.Equal(t, wlog1.TotalCaloriesBurned, wlog2.TotalCaloriesBurned)
	require.Equal(t, wlog1.TotalDistance, wlog2.TotalDistance)
	require.Equal(t, wlog1.TotalRepetitions, wlog2.TotalRepetitions)
	require.Equal(t, wlog1.TotalWeightLifted, wlog2.TotalWeightLifted)
	require.Equal(t, wlog1.TotalSets, wlog2.TotalSets)
	require.Equal(t, wlog1.Rating, wlog2.Rating)
	require.Equal(t, wlog1.Comments, wlog2.Comments)
	require.Equal(t, wlog1.OverallFeeling, wlog2.OverallFeeling)
}

func TestUpdateWorkoutLog(t *testing.T) {
	wlog1 := CreateRandomWorkoutLog(t)

	arg := UpdateWorkoutLogParams{
		LogID:               wlog1.LogID,
		TotalSets:           pgtype.Int4{Int32: int32(util.GetRandomAmount(1, 20)), Valid: true},
		FatigueLevel:        NullFatiguelevel{Fatiguelevel: FatiguelevelHeavy, Valid: true},
		TotalWeightLifted:   pgtype.Int4{Int32: int32(util.GetRandomAmount(1, 2000)), Valid: true},
		TotalRepetitions:    pgtype.Int4{Int32: int32(util.GetRandomAmount(1, 200)), Valid: true},
		TotalDistance:       pgtype.Int4{Int32: int32(util.GetRandomAmount(1, 20000)), Valid: true},
		TotalCaloriesBurned: pgtype.Int4{Int32: int32(util.GetRandomAmount(1, 2000)), Valid: true},
		WorkoutDuration:     pgtype.Text{String: "4h", Valid: true},
		Rating:              NullRating{Rating: Rating1, Valid: true},
		Comments:            pgtype.Text{String: util.GetRandomUsername(100), Valid: true},
		OverallFeeling:      pgtype.Text{String: "gool", Valid: true},
	}

	wlog2, err := testStore.UpdateWorkoutLog(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, wlog2)

	require.Equal(t, wlog1.Username, wlog2.Username)
	require.Equal(t, wlog1.PlanID, wlog2.PlanID)
	require.Equal(t, arg.LogID, wlog2.LogID)
	require.Equal(t, arg.WorkoutDuration.String, wlog2.WorkoutDuration)
	require.Equal(t, arg.FatigueLevel.Fatiguelevel, wlog2.FatigueLevel)
	require.Equal(t, arg.TotalCaloriesBurned.Int32, wlog2.TotalCaloriesBurned)
	require.Equal(t, arg.TotalDistance.Int32, wlog2.TotalDistance)
	require.Equal(t, arg.TotalRepetitions.Int32, wlog2.TotalRepetitions)
	require.Equal(t, arg.TotalWeightLifted.Int32, wlog2.TotalWeightLifted)
	require.Equal(t, arg.TotalSets.Int32, wlog2.TotalSets)
	require.Equal(t, arg.Rating.Rating, wlog2.Rating)
	require.Equal(t, arg.Comments.String, wlog2.Comments)
	require.Equal(t, arg.OverallFeeling.String, wlog2.OverallFeeling)
}

func TestDeleteWorkoutLog(t *testing.T) {
	wlog1 := CreateRandomWorkoutLog(t)
	err := testStore.DeleteWorkoutLog(context.Background(), wlog1.LogID)
	require.NoError(t, err)

	wlog2, err := testStore.GetWorkoutLog(context.Background(), wlog1.LogID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, wlog2)
}

func TestListWorkoutsLogs(t *testing.T) {
	lastWorkout := CreateRandomWorkoutLog(t)
	arg := ListWorkoutLogsParams{
		PlanID: lastWorkout.PlanID,
		Limit:  5,
		Offset: 0,
	}

	workouts, err := testStore.ListWorkoutLogs(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, workouts)

	for _, workout := range workouts {
		require.NotEmpty(t, workout)
		require.Equal(t, lastWorkout.PlanID, workout.PlanID)
	}
}
