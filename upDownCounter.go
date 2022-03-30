package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Insert code here
	var number1, number2 string
	fmt.Println("Please enter first number: ")
	fmt.Scanln(&number1)
	fmt.Println("Please enter second number: ")
	fmt.Scanln(&number2)

	num1, _ := strconv.ParseInt(number1, 10, 0)
	num2, _ := strconv.ParseInt(number2, 10, 0)

	if num1 > num2 {
		num1, num2 = num2, num1
	}

	for i := num1; i <= num2; i++ {
		if num1 == num2 {
			fmt.Println("Number 1 and number 2 are the same. Please try again!")
			break
		}
		fmt.Println((i))
	}

	for i := num2 - 1; i >= num1; i-- {
		if num1 == num2 {
			// fmt.Println("Number 1 and number 2 are the same. Please try again!")
			break
		}
		fmt.Println((i))
	}
}
