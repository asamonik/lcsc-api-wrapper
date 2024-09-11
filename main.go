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
	fmt.Printf("Product: %+v\n", product)
}
