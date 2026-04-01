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

type Profile struct {
	Name    string
	Age     int
	Address string
}

func isAdult(age int) bool {
	if age >= 18 {
		return true
	}
	return false
}
func birthDay(user *Profile) {
	user.Age++
}

func charelocateUser(user *Profile, newCity string) {
	user.Address = newCity
}

func GetProfile() {
	fmt.Println("enter your name, age, and city :")
	var pr Profile
	fmt.Scan(&pr.Name, &pr.Age, &pr.Address)
	adult := isAdult(pr.Age)
	if adult {
		fmt.Printf("Welcome, %s! You are an adult.\n", pr.Name)
	}

	fmt.Println("change city ?")
	var verify string
	fmt.Scan(&verify)
	if verify == "yes" || verify == "y" {
		fmt.Println("enter new city name :")
		var newCity string
		fmt.Scan(&newCity)
		charelocateUser(&pr, newCity)
	}
	birthDay(&pr)
	fmt.Printf("Profile updated: Name: %s, Age: %d, Address: %s\n", pr.Name, pr.Age, pr.Address)

}
