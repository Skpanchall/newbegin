package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ArrayAndSlices() {
	fmt.Println("Select which task you want to run :")
	fmt.Println("1. print all numbers")
	fmt.Println("2. student mark system")
	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		PrintAllNumberbySlice()
	case 2:
		StudentMarkSystem()
	}

}
func PrintAllNumberbySlice() {
	fmt.Println("how many numbers you want to print ?")
	var n int
	fmt.Scanln(&n)
	fmt.Println("give me a numbers through one by one ")

	var numbers []int
	for i := 1; i <= n; i++ {
		var number int
		fmt.Scanln(&number)
		numbers = append(numbers, number)
	}

	for _, nu := range numbers {
		println(nu)
	}
}

func StudentMarkSystem() {
	fmt.Println("howm many students have ?")
	var n int
	fmt.Scanln(&n)
	fmt.Println("enter marks for those students by comma ex 88,44,")
	var allmarks string
	fmt.Scan(&allmarks)
	marks := strings.Split(allmarks, ",")
	if len(marks) != n {
		fmt.Println("number of marks entered does not match the number of students ")
		return
	}
	var highMarks int
	var passCount int = 0
	for i, mark := range marks {
		markee, err := strconv.Atoi(strings.TrimSpace(mark))
		if err != nil {
			fmt.Printf("invalid mark for student %d: %s\n", i+1, mark)
			continue
		}
		if highMarks < markee {
			highMarks = markee
		}
		if markee >= 40 {
			passCount++
		}
		fmt.Println("", i+1, ". student marks is ", mark)
	}

	fmt.Println("total pass students are ", passCount)
	fmt.Printf("the highest marks is %d\n", highMarks)

}
