package product_struct

type Product struct {
	Title string
	Id    int
	Price float64
}

func New(title string, id int, price float64) *Product {
	return &Product{title, id, price}
}
