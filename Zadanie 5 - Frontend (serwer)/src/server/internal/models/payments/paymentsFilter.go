package payments

// PaymentsFilter ...
type PaymentsFilter struct {
	Offset int64
	Limit  int64
	Status *string
}
