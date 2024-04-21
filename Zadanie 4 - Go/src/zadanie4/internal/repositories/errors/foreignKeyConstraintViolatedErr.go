package errors

// ForeignKeyConstraintViolated ...
type ForeignKeyConstraintViolated struct{}

// Error ...
func (e *ForeignKeyConstraintViolated) Error() string {
	return "foreign key constraint violated"
}
