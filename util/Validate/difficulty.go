package util

type Difficulty string

const (
	DifficultyBeginner     Difficulty = "Beginner"
	DifficultyIntermediate Difficulty = "Intermediate"
	DifficultyAdvanced     Difficulty = "Advanced"
)

// IsSupportedDifficulty returns true if the difficulty is supported
func IsSupportedDifficulty(difficulty Difficulty) bool {
	switch difficulty {
	case DifficultyAdvanced, DifficultyBeginner, DifficultyIntermediate:
		return true
	}
	return false
}
