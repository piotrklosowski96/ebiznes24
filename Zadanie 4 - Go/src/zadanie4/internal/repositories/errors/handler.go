package errors

import (
	"errors"

	"github.com/mattn/go-sqlite3"
	"gorm.io/gorm"
)

// HandleDatabaseError ...
func HandleDatabaseError(sqliteError error) error {
	// NOTE(Piotr Kłosowski): This need more robust handling...

	var internalErr sqlite3.Error
	if errors.As(sqliteError, &internalErr) {
		switch {
		case errors.Is(internalErr.Code, sqlite3.ErrConstraint):
			return &ForeignKeyConstraintViolated{}
		case errors.Is(internalErr.Code, sqlite3.ErrNotFound):
			return &ResourceNotFoundError{}
		}
	}

	switch {
	case errors.Is(sqliteError, gorm.ErrRecordNotFound):
		return &ResourceNotFoundError{}
	case errors.Is(sqliteError, gorm.ErrForeignKeyViolated):
		return &ResourceNotFoundError{}
	}

	return &UnknownError{}
}
