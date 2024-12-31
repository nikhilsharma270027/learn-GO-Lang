package main

import "fmt"

// in GO, ...interface{} is like any type in GO

// ...int - we can get n number of parmeters of int type
func sum(nums ...int) int {
	total := 0

	for _, num :=  range nums {
		total = total + num
	}

	return total
}

func main() {
	nums := []int{1, 2, 3, 4}
	result := sum(nums...)

	fmt.Println(result)
	
}