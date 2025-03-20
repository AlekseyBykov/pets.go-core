package receivers

import (
	"fmt"
	"testing"
)

func TestProduct_GetProductInfo(t *testing.T) {
	p := Product{Name: "Laptop", Price: 1500}

	fmt.Println(Product.getProductInfo(p))

	info := p.getProductInfo()
	expected := "Laptop 1500"

	if info != expected {
		t.Errorf("expected '%s', got '%s'", expected, info)
	}

	if p.Name != "Laptop" {
		t.Errorf("expected Name to be 'Laptop', got '%s'", p.Name)
	}
	if p.Price != 1500 {
		t.Errorf("expected Price to be 1500, got %d", p.Price)
	}
}
