package models

// Product ...
type Product struct {
	CommonFields

	Name        string
	Description *string
	Categories  []*Category `gorm:"many2many:products_categories;"`
}
