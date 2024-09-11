package main

import (
	"fmt"
	lcsc "lcscapi/api"
)

func main() {
	client := lcsc.NewClient()

	product, err := client.GetProduct("C2911376")
	if err != nil {
		fmt.Printf("Error getting product: %v\n", err)
		return
	}
	fmt.Printf("Product: %s\n", product.ProductModel)

	categories, err := client.GetCategories()
	if err != nil {
		fmt.Printf("Error getting product: %v\n", err)
		return
	}
	fmt.Printf("category0: %s\n", (*categories)[0].CatalogName)
}
