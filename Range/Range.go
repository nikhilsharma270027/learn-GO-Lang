package main

import "fmt"

func main() {

	// Defination
	// range in Go
	// The range keyword in Go is used in a for loop to iterate over elements of different
	// data structures like arrays, slices, maps, strings, and channels.

	// It provides a convenient way to loop while automatically giving access to index/key and value.

	// 1. Syntax
	// for index, value := range collection {
	// 	// use index and value+++++++++++++++++++++
	// }
	// index / key → Position (arrays, slices, strings) or Key (maps).
	// value → The element at that position.

	// nums := []int{6, 7, 8}
	//     key, value
	// for i, num := range nums { // _ is index
	// 	fmt.Println(num, i)
	// }
	// o/p: 6 0
	// 7 1
	// 8 2

	// unicode/ASCII value if done on string
	// for i, c := range "golang" {
	// 	fmt.Println(i, c)
	// }
	// 0 103
	// 1 111
	// 2 108
	// 3 97
	// 4 110
	// 5 103

	// for i, c := range "golang" {
	// 	fmt.Println(i, string(c))
	// }
	// 0 g
	// 1 o
	// 2 l
	// 3 a
	// 4 n
	// 5 g

	// New

	nums := []int{10, 20, 30, 40}

	for i, v := range nums {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}
}
