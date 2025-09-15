package main

import "fmt"

// numbered squence of specific number
func main() {
	// var nums [4]int    arry-> 0  string -> " "  bool -> false defalut values created
	// var bool [4]bool

	// nums[0] = 1
	// fmt.Println(len(nums))

	// fmt.Println(nums)
	// [1 0 0 0]

	// var name [3]string
	// name[0] = "golang"
	// fmt.Println(name)
	// [golang  ]

	// fmt.Println(bool)
	// [false false false false]

	// Array with value
	// name := [3]int{1, 2, 3}
	// fmt.Println(name)
	// [1 2 3]

	// 2D ARRAY

	// nums:= [2][2]int{{2, 4}, { 5, 6}}

	// fmt.Println(nums)
	// [[2 4] [5 6]]

	// use when
	// - fixed size, that is predictable
	// - memory optimization
	// - contant time access

	// new

	var nums [5]int
	nums[0] = 1
	nums[2] = 2
	fmt.Print(nums)
	// [1 0 2 0 0]

	arr := [3]string{"a", "b", "c"}
	fmt.Println(arr)

	//
	arrbool := [3]bool{true}
	fmt.Println(arrbool)

	// deletion using while loop
	delarr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(delarr)

	for i := 2; i < len(delarr)-1; i++ {
		delarr[i] = delarr[i+1]
	}
	delarr[4] = 0
	fmt.Println(delarr)
	// [1 2 4 5 5]

	// appending values in the arr

	number := []int{1, 2, 3}
	number = append(number, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(number) // [1 2 3 4 5 6 7 8 9 10]

	// capcity and length
	fmt.Println("Length:", len(number))   // Length: 10
	fmt.Println("Capacity:", cap(number)) // Capacity: 16 (capacity may vary)

	// Slicing
	slice := number[5:6]
	fmt.Println("Slice:", slice) // Slice: [6]

}
