package errors

import (
	"errors"

	"gorm.io/gorm"
)

// HandleDatabaseError ...
func HandleDatabaseError(sqliteError error) error {
	switch {
	case errors.Is(sqliteError, gorm.ErrRecordNotFound):
		return &ResourceNotFoundError{}
	}

	return sqliteError
}
