package main

import "fmt"

func Agechecker() {
	fmt.Println("hello what is you name ?")
	var name string
	fmt.Scanln(&name)
	fmt.Println("age ? ")
	var age int
	fmt.Scanln(&age)
	fmt.Printf("Hello %s\n", name)
	if age < 18 {
		fmt.Println("You are a kid")
	} else if age >= 18 && age < 30 {
		fmt.Println("You are a younger but a mature")
	} else {
		fmt.Println("You are an adult")
	}
}

func WorkStatusChecker() {
	fmt.Println("hello what is you name ?")
	var name string
	fmt.Scanln(&name)
	fmt.Println("age ? ")
	var age int
	fmt.Scanln(&age)
	fmt.Println("working ? (yes/no) ")
	var working string
	fmt.Scanln(&working)
	var isWorking bool
	if working == "yes" || working == "y" {
		isWorking = true
	}
	fmt.Printf("Hello %s\n", name)
	if age < 18 {
		fmt.Println("You are a kid ")
	} else if age >= 18 && age < 30 && isWorking {
		fmt.Println("You are a younger but a mature \nStatus : working Professional")
	} else if isWorking {
		fmt.Println("status : working Professional")
	} else {
		fmt.Println("you are just adult hood")
	}
}

// func main() {
// 	// agechecker()
// 	workStatusChecker()
// }
