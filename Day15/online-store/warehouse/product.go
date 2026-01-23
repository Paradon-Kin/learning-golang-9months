package warehouse

import "fmt"

type Product struct {
	Name  string
	Stock int
	price float64 //private
}

func NewProduct(n string, s int, p float64) *Product {
	return &Product{
		Name:  n,
		Stock: s,
		price: p,
	}
}

func (p *Product) GetPrice() float64 {
	return p.price
}

func (p *Product) ShowDetails() {
	fmt.Printf("[Warehouse] Product: %s | Stock: %d\n", p.Name, p.Stock)
}
