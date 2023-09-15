package util

type Musclegroupenum string

const (
	MusclegroupenumChest     Musclegroupenum = "Chest"
	MusclegroupenumBack      Musclegroupenum = "Back"
	MusclegroupenumLegs      Musclegroupenum = "Legs"
	MusclegroupenumShoulders Musclegroupenum = "Shoulders"
	MusclegroupenumArms      Musclegroupenum = "Arms"
	MusclegroupenumAbs       Musclegroupenum = "Abs"
	MusclegroupenumCardio    Musclegroupenum = "Cardio"
)

// IsSupportedMuscleGroup returns true if the muscle group is supported
func IsSupportedMuscleGroup(mGroup Musclegroupenum) bool {
	switch mGroup {
	case MusclegroupenumAbs, MusclegroupenumArms, MusclegroupenumBack, MusclegroupenumCardio, MusclegroupenumChest, MusclegroupenumLegs, MusclegroupenumShoulders:
		return true
	}
	return false
}
