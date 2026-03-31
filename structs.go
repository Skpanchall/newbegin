package main

import "fmt"

type Product struct {
	Name  string
	Price float64
}

func ProductTotalValue() float64 {
	fmt.Println("enter a name and price of the product :")
	var pr Product
	fmt.Scan(&pr.Name, &pr.Price)
	fmt.Println("how much Quantity you want ?")
	var qty int
	fmt.Scan(&qty)
	total := pr.Price * float64(qty)
	return total
}

func GetProductTotalValues() {
	total := ProductTotalValue()
	if total > 2000 {
		fmt.Println("Congratulations! You are eligible for a discount.")
		fmt.Println("Discount Applied : 10%", total*0.1)
		total = total * 0.9
	}
	fmt.Println("total :", total)

}
