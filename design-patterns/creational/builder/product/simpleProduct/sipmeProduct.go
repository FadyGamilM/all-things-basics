package simpleproduct

type SimpleProduct struct {
	name         string
	price        float64
	countInStock int
}

func (sp *SimpleProduct) SetName(name string) {
	sp.name = name
}

func (sp *SimpleProduct) SetPrice(price float64) {
	sp.price = price
}

func (sp *SimpleProduct) SetInStock(inStock int) {
	sp.countInStock = inStock
}

func (sp *SimpleProduct) GetInStock() int {
	return sp.countInStock
}

// func (sp SimpleProduct) Build() productbuilder.ProductBuilder {
// 	return sp
// }
