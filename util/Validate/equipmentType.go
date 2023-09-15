package util

type Equipmenttype string

const (
	EquipmenttypeBarbell    Equipmenttype = "Barbell"
	EquipmenttypeDumbbell   Equipmenttype = "Dumbbell"
	EquipmenttypeMachine    Equipmenttype = "Machine"
	EquipmenttypeBodyweight Equipmenttype = "Bodyweight"
	EquipmenttypeOther      Equipmenttype = "Other"
)

// IsSupportedEquipment returns true if the equipment is supported
func IsSupportedEquipment(equipment Equipmenttype) bool {
	switch equipment {
	case EquipmenttypeBarbell, EquipmenttypeBodyweight, EquipmenttypeDumbbell, EquipmenttypeMachine, EquipmenttypeOther:
		return true
	}
	return false
}
