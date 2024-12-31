package main

import "fmt"

// counter will return func() then func() will return int
func counter() func() int {
	var count int = 0
	fmt.Println(count)
	
	return func() int {
		fmt.Println(count)
		
		count += 1
		return count
	}
}

func main() {
	increment := counter()

	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	
}