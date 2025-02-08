package productbuilder

import simpleproduct "all-things-basics/design-patterns/creational/builder/product/simpleProduct"

type ProductBuilder struct {
	product simpleproduct.SimpleProduct
}

func (pb *ProductBuilder) SetName(name string) *ProductPriceBuilder {
	pb.product.SetName(name)
	return &ProductPriceBuilder{
		productBuilder: *pb,
	}
}

type ProductPriceBuilder struct {
	productBuilder ProductBuilder
}

func (ppb *ProductPriceBuilder) SetPrice(price float64) *ProductInStockCountBuilder {
	ppb.productBuilder.product.SetPrice(price)
	return &ProductInStockCountBuilder{
		productBuilder: ppb.productBuilder,
	}
}

// maybe these fields are optional, so we implement the Build() method on the optional field struct
type ProductInStockCountBuilder struct {
	productBuilder ProductBuilder
}

func (piscb *ProductInStockCountBuilder) SetInStockCount(count int, sign bool) *ProductInStockCountBuilder {
	if sign {
		piscb.productBuilder.product.SetInStock(
			piscb.productBuilder.product.GetInStock() + count,
		)
		return piscb
	}

	piscb.productBuilder.product.SetInStock(
		piscb.productBuilder.product.GetInStock() - count,
	)
	return piscb
}

func (piscb *ProductInStockCountBuilder) Build() *simpleproduct.SimpleProduct {
	return &piscb.productBuilder.product
}
