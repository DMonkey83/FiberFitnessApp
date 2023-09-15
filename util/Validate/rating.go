package util

type Rating string

const (
	Rating1 Rating = "1"
	Rating2 Rating = "2"
	Rating3 Rating = "3"
	Rating4 Rating = "4"
	Rating5 Rating = "5"
)

// IsSupportedRating returns true if the rating is supported
func IsSupportedRating(rating Rating) bool {
	switch rating {
	case Rating1, Rating2, Rating3, Rating4, Rating5:
		return true
	}
	return false
}
