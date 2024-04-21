package scopes

import (
	"gorm.io/gorm"
)

// PreloadAssociations ...
func PreloadAssociations(associationNames []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for _, associationName := range associationNames {
			db = db.Preload(associationName)
		}

		return db
	}
}
