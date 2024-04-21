package models

// Product ...
type Product struct {
	CommonFields

	Name        string  `gorm:""`
	Description *string `gorm:""`
}
