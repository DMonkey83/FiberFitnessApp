package util

type Completionenum string

const (
	CompletionenumCompleted  Completionenum = "Completed"
	CompletionenumIncomplete Completionenum = "Incomplete"
	CompletionenumNotStarted Completionenum = "NotStarted"
)

// IsSupportedCompletion returns true if the completion is supported
func IsSupportedCompletion(completion Completionenum) bool {
	switch completion {
	case CompletionenumCompleted, CompletionenumIncomplete, CompletionenumNotStarted:
		return true
	}
	return false
}
