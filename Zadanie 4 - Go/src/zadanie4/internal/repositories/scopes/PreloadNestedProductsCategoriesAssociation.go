package scopes

import "gorm.io/gorm"

// PreloadNestedProductsCategoriesAssociation ...
func PreloadNestedProductsCategoriesAssociation(db *gorm.DB) *gorm.DB {
	return db.Preload("Products.Categories")

}
