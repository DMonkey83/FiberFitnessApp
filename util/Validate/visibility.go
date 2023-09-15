package util

type Visibility string

const (
	VisibilityPublic  Visibility = "Public"
	VisibilityPrivate Visibility = "Private"
)

// IsSupportedVisibility returns true if the visibility is supported
func IsSupportedVisibility(visibility Visibility) bool {
	switch visibility {
	case VisibilityPublic, VisibilityPrivate:
		return true
	}
	return false
}
