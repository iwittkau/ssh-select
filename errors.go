package sshselect

// Error is the sshs error type
type Error string

// Error implements the golang builtin error type
func (e Error) Error() string {
	return string(e)
}

// Tmux errors
const (
	ErrTmuxNotAttached = Error("not attached to tmux session; a new sshs tmux session was created, attach with 'tmux a -t sshs'")
)
