package util

type Workoutgoalenum string

const (
	WorkoutgoalenumBuildMuscle      Workoutgoalenum = "Build Muscle"
	WorkoutgoalenumLoseWeight       Workoutgoalenum = "Lose Weight"
	WorkoutgoalenumImproveEndurance Workoutgoalenum = "Improve Endurance"
	WorkoutgoalenumMaintainFitness  Workoutgoalenum = "Maintain Fitness"
	WorkoutgoalenumToneBody         Workoutgoalenum = "Tone Body"
	WorkoutgoalenumCustom           Workoutgoalenum = "Custom"
)

// IsSupportedWeightUnit returns true if the weight unit is supported
func IsSupportedGoal(goal Workoutgoalenum) bool {
	switch goal {
	case WorkoutgoalenumBuildMuscle, WorkoutgoalenumCustom, WorkoutgoalenumImproveEndurance, WorkoutgoalenumLoseWeight, WorkoutgoalenumMaintainFitness, WorkoutgoalenumToneBody:
		return true
	}
	return false
}
