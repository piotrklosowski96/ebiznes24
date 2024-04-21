package errors

// UnknownError ...
type UnknownError struct{}

// Error ...
func (e *UnknownError) Error() string {
	return "unknown database error"
}
