package main

import "fmt"

func BioMaker() {
	var name string
	var age int
	var city string
	var isWorking bool
	var salary int
	fmt.Println("what is your name ?")
	fmt.Scanln(&name)
	fmt.Println("age ? ")
	fmt.Scanln(&age)
	fmt.Println("city ? ")
	fmt.Scanln(&city)
	fmt.Println("working ? ")
	fmt.Scanln(&isWorking)
	fmt.Println("salary ? ")
	fmt.Scanln(&salary)
	fmt.Printf("so now your bio is : \nname : %s \nage:%d  \ncity:%s  \n")
}

// func main() {
// 	bioMaker()
// }

// day 1
// learn a variable how to declar
//
