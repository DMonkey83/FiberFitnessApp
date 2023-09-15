package util

type Weightunit string

const (
	WeightunitKg Weightunit = "kg"
	WeightunitLb Weightunit = "lb"
)

// IsSupportedWeightUnit returns true if the weight unit is supported
func IsSupportedWeightUnit(wUnit Weightunit) bool {
	switch wUnit {
	case WeightunitKg, WeightunitLb:
		return true
	}
	return false
}
