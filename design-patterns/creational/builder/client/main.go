package main

import (
	productbuilder "all-things-basics/design-patterns/creational/builder/productBuilder"
	"log"
)

func main() {
	// eventually we need to get an instance of the product
	productBuilderObj := &productbuilder.ProductBuilder{}
	product := productBuilderObj.SetName("nike v2k").SetPrice(float64(100.98)).Build()
	log.Println(product)
}
