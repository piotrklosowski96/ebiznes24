package models

// Category ...
type Category struct {
	CommonFields

	Name        string
	Description *string
	Categories  []*Category `gorm:"many2many:products_categories;"`
}
