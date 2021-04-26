package main

import (
	"fmt"
	"strconv"
)

//Найти натуральное число от 1 до 10 000 с максимальной
//суммой делителей.

func main()  {
	n := Reader()
	task322(n)
}


func Reader() string {
	fmt.Print("Enter integer: ")
	var input string
	fmt.Scan(&input)
	return input
}


func task322(n string) int {
	new, _ := strconv.Atoi(n)
	if new >= 1 && new <= 10000 {
		var memory int = 0
		var CurrentNumber int = 0

		for i := 1; i <= new; i++ {
			SumDivides := 0
			for j := 1; j <= i; j++ {
				if i%j == 0 {
					SumDivides += j
				}
				if SumDivides > memory {
					memory = SumDivides
					CurrentNumber = i
				}
			}
		}
		fmt.Println("Number: ", CurrentNumber)
		return CurrentNumber
	} else {
		return 0
	}
}