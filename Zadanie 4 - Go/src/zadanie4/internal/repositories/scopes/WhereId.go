package scopes

import (
	"gorm.io/gorm"
)

// WhereId ...
func WhereId(id string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	}
}
