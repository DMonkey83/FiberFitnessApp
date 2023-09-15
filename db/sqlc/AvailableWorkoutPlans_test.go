package db

import (
	"context"
	"testing"

	"github.com/DMonkey83/MyFitnessApp/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomAvailableWorkoutPlan(t *testing.T) Availableworkoutplan {
	user := CreateRandomUser(t)

	arg := CreateAvailablePlanParams{
		CreatorUsername: user.Username,
		PlanName:        util.GetRandomUsername(8),
		Description:     util.GetRandomUsername(60),
		Goal:            WorkoutgoalenumLoseWeight,
		IsPublic:        VisibilityPrivate,
		Difficulty:      DifficultyIntermediate,
	}

	plan, err := testStore.CreateAvailablePlan(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, plan)

	require.Equal(t, arg.CreatorUsername, plan.CreatorUsername)
	require.Equal(t, arg.PlanName, plan.PlanName)
	require.Equal(t, arg.Description, plan.Description)
	require.Equal(t, arg.Goal, plan.Goal)
	require.Equal(t, arg.Difficulty, plan.Difficulty)

	require.NotZero(t, plan.PlanID)
	return plan
}

func TestAvailableWorkoutPlan(t *testing.T) {
	CreateRandomAvailableWorkoutPlan(t)
}

func TestGetAvailableWorkoutPlan(t *testing.T) {
	plan1 := CreateRandomAvailableWorkoutPlan(t)
	plan2, err := testStore.GetAvailablePlan(context.Background(), plan1.PlanID)
	require.NoError(t, err)
	require.NotEmpty(t, plan2)

	require.Equal(t, plan1.CreatorUsername, plan2.CreatorUsername)
	require.Equal(t, plan1.Description, plan2.Description)
	require.Equal(t, plan1.Goal, plan2.Goal)
	require.Equal(t, plan1.Difficulty, plan2.Difficulty)
	require.Equal(t, plan1.IsPublic, plan2.IsPublic)
}

func TestUpdateAvailableWorkoutPlan(t *testing.T) {
	plan1 := CreateRandomAvailableWorkoutPlan(t)

	arg := UpdateAvailablePlanParams{
		PlanID:      plan1.PlanID,
		PlanName:    util.GetRandomUsername(6),
		Description: util.GetRandomUsername(49),
		Difficulty:  DifficultyBeginner,
		Goal:        WorkoutgoalenumBuildMuscle,
		IsPublic:    VisibilityPublic,
	}

	plan2, err := testStore.UpdateAvailablePlan(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, plan2)

	require.Equal(t, plan1.CreatorUsername, plan2.CreatorUsername)
	require.Equal(t, arg.PlanID, plan2.PlanID)
	require.Equal(t, arg.Description, plan2.Description)
	require.Equal(t, arg.Goal, plan2.Goal)
	require.Equal(t, arg.Difficulty, plan2.Difficulty)
	require.Equal(t, arg.IsPublic, plan2.IsPublic)
}

func TestDeletAvailableWorkoutPlan(t *testing.T) {
	plan1 := CreateRandomAvailableWorkoutPlan(t)
	err := testStore.DeleteAvailablePlan(context.Background(), plan1.PlanID)
	require.NoError(t, err)

	plan2, err := testStore.GetAvailablePlan(context.Background(), plan1.PlanID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, plan2)
}

func TestListAllAvailableWorkoutPlans(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomAvailableWorkoutPlan(t)
	}

	arg := ListAllAvailablePlansParams{
		Limit:  5,
		Offset: 0,
	}

	plans, err := testStore.ListAllAvailablePlans(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, plans)

	for _, plan := range plans {
		require.NotEmpty(t, plan)
	}
}

func TestListAvailableWorkoutPlansByCreator(t *testing.T) {
	var lastPlan Availableworkoutplan
	for i := 0; i < 10; i++ {
		lastPlan = CreateRandomAvailableWorkoutPlan(t)
	}

	arg := ListAvailablePlansByCreatorParams{
		CreatorUsername: lastPlan.CreatorUsername,
		Limit:           5,
		Offset:          0,
	}

	plans, err := testStore.ListAvailablePlansByCreator(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, plans)

	for _, plan := range plans {
		require.NotEmpty(t, plan)
		require.Equal(t, lastPlan.CreatorUsername, plan.CreatorUsername)
	}
}
