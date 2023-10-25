package api

import (
	"github.com/gofiber/fiber/v2"

	auth "github.com/DMonkey83/FiberFitnessApp/middleware"
)

func (server *Server) SetupRouter(app *fiber.App) {
	routes := app.Group("/api")
	// user SetupRouter
	routes.Post("users", server.createUser)
	routes.Post("users/login", server.LoginUser)
	routes.Post("/token/refresh", server.renewAccessToken)
	routes.Post("user_profile", server.createUserProfile)

	authRoutes := routes.Use(auth.AuthMiddleware(server.tokenMaker))

	// available_workouts SetupRouter
	authRoutes.Post("available_workouts", server.createAvailablePlan)
	authRoutes.Get("available_workouts/:id", server.getAvailablePlan)
	authRoutes.Patch("available_workouts", server.updateAvailablePlan)
	authRoutes.Get("available_workouts", server.listAllAvailablePlans)
	// available_workouts_exercises SetupRouter
	authRoutes.Post("available_workout_exercises", server.createAvailablePlanExercise)
	authRoutes.Get("available_workout_exercises/:id", server.getAvailablePlanExercise)
	authRoutes.Patch("available_workout_exercises", server.updateAvailablePlanExercise)
	authRoutes.Get("available_workout_exercises", server.listAllAvailablePlansExercises)
	// exercises SetupRouter
	authRoutes.Post("exercises", server.createExercise)
	authRoutes.Get("exercises/:exercise_name", server.getExercise)
	authRoutes.Patch("exercises", server.updateExercise)
	authRoutes.Get("exercises", server.listAllEquipmentExercises)
	authRoutes.Get("exercises", server.listAllMuscleGroupExercises)
	// max_rep_goal SetupRouter
	authRoutes.Post("max_rep_goal", server.createMaxRepGoal)
	authRoutes.Get("max_rep_goal/:goal_id/:username/:exercise_name", server.getMaxRepGoal)
	authRoutes.Patch("max_rep_goal", server.updateMaxRepGoal)
	authRoutes.Get("max_rep_goal", server.listMaxRepGoals)
	// max_weight_goal SetupRouter
	authRoutes.Post("max_weight_goal", server.createMaxWeightGoal)
	authRoutes.Get("max_weight_goal/:goal_id/:username/:exercise_name", server.getMaxWeightGoal)
	authRoutes.Patch("max_weight_goal", server.updateMaxWeightGoal)
	authRoutes.Get("max_weight_goal", server.listMaxWeightGoals)
	// one_off_workout_exercises SetupRouter
	authRoutes.Post("one_off_workout_exercises", server.createOneOfExercise)
	authRoutes.Get("one_off_workout_exercises/:workout_id/:id", server.getOneOfExercise)
	authRoutes.Patch("one_off_workout_exercises", server.updateOneOfExercise)
	authRoutes.Get("one_off_workout_exercises", server.listAllOneOfExercises)
	// set SetupRouter
	authRoutes.Post("set", server.createSet)
	authRoutes.Get("set/:set_id", server.getSet)
	authRoutes.Patch("set", server.updateSet)
	authRoutes.Get("set", server.listSets)
	// weight_entry SetupRouter
	authRoutes.Post("weight_entry", server.createWeightEntry)
	authRoutes.Get("weight_entry/:weight_entry_id/:username", server.getWeightEntry)
	authRoutes.Patch("weight_entry", server.updateWeightEntry)
	authRoutes.Get("weight_entry", server.listWeightEntries)
	// workout SetupRouter
	authRoutes.Post("workout", server.createWorkout)
	authRoutes.Get("workout/:workout_id", server.getWorkout)
	authRoutes.Patch("workout", server.updateWorkout)
	authRoutes.Get("workout", server.listWorkout)
	// workout_log SetupRouter
	authRoutes.Post("workout_log", server.createWorkoutLog)
	authRoutes.Get("workout_log/:log_id", server.getWorkoutLog)
	authRoutes.Patch("workout_log", server.updateWorkoutLog)
	authRoutes.Get("workout_log", server.listWorkoutLogs)
	// workout_plan SetupRouter
	authRoutes.Post("workout_plan", server.createWorkoutPlan)
	authRoutes.Get("workout_plan/:plan_id/:username", server.getWorkoutPlan)
	authRoutes.Patch("workout_plan", server.updateWorkoutPlan)
	// user_profile SetupRouter
	authRoutes.Get("user_profile/:username", server.getUserProfile)
	authRoutes.Patch("user_profile", server.updateUserProfile)

	server.app = app
}
