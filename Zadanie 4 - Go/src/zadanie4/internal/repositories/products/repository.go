package products

// Product ...
type Product struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Repository ...
type Repository struct {
	products []*Product
}

// New ...
func New() *Repository {
	return &Repository{
		products: []*Product{
			{
				Id:          "0",
				Name:        "product_name_0",
				Description: "product_description_0",
			},
			{
				Id:          "1",
				Name:        "product_name_1",
				Description: "product_description_1",
			},
			{
				Id:          "2",
				Name:        "product_name_2",
				Description: "product_description_2",
			},
		},
	}
}

// Add ...
func (r *Repository) Add(product *Product) {
	r.products = append(r.products, product)
}

// GetAll ...
func (r *Repository) GetAll() []*Product {
	return r.products
}

// GetById ...
func (r *Repository) GetById(id string) *Product {
	for _, p := range r.products {
		if p.Id == id {
			return p
		}
	}

	return nil
}

// Update ...
func (r *Repository) Update(id string, updatedProduct *Product) *Product {
	for idx, p := range r.products {
		if p.Id == id {
			r.products[idx] = updatedProduct
			return r.products[idx]
		}
	}

	return nil
}

// Remove ...
func (r *Repository) Remove(id string) {
	for idx, p := range r.products {
		if p.Id == id {
			r.products[idx] = r.products[len(r.products)-1]
			r.products = r.products[:len(r.products)-1]
		}
	}
}
