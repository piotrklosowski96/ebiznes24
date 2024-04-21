package errors

import "fmt"

// ResourceNotFoundError ...
type ResourceNotFoundError struct {
	ResourceID string `json:"resource_id"`
}

// Error ...
func (e *ResourceNotFoundError) Error() string {
	return fmt.Sprintf("resource not found (id: '%s')", e.ResourceID)
}
