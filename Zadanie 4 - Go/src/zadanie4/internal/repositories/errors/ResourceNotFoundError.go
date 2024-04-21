package errors

// ResourceNotFoundError ...
type ResourceNotFoundError struct{}

// Error ...
func (e *ResourceNotFoundError) Error() string {
	return "resource not found"
}
