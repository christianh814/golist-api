package entities

// Define what the properties for the Product
type Product struct {
	ID         uint    `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Desciption string  `json:"description"`
}
