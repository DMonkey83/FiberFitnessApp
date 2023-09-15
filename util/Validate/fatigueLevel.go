package util

type Fatiguelevel string

const (
	FatiguelevelVeryLight Fatiguelevel = "Very Light"
	FatiguelevelLight     Fatiguelevel = "Light"
	FatiguelevelModerate  Fatiguelevel = "Moderate"
	FatiguelevelHeavy     Fatiguelevel = "Heavy"
	FatiguelevelVeryHeavy Fatiguelevel = "Very Heavy"
)

// IsSupportedFatigueLevel returns true if the fatigue level is supported
func IsSupportedFatigueLevel(fatigueLevel Fatiguelevel) bool {
	switch fatigueLevel {
	case FatiguelevelHeavy, FatiguelevelLight, FatiguelevelModerate, FatiguelevelVeryHeavy, FatiguelevelVeryLight:
		return true
	}
	return false
}
