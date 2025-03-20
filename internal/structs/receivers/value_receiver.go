package receivers

import "fmt"

type Product struct {
	Name  string
	Price int
}

func (p Product) getProductInfo() string {
	return fmt.Sprintf("%s %d", p.Name, p.Price)
}
