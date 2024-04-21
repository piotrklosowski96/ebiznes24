package scopes

import (
	"fmt"

	"gorm.io/gorm"
)

// SkipAssociationsUpsert ...
func SkipAssociationsUpsert(associationNames []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for _, associationName := range associationNames {
			db = db.Omit(fmt.Sprintf("%s.*", associationName))
		}

		return db
	}
}
