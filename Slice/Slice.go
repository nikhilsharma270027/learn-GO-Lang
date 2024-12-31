// Slices are similar to arrays, but are more powerful and flexible.
// Like arrays, slices are also used to store multiple values of the same type in a single variable.
// However, unlike arrays, the length of a slice can grow and shrink as you see fit.

// In Go, there are several ways to create a slice:

// Using the []datatype{values} format
// Create a slice from an array
// Using the make() function

// var myarray = [length]datatype{values} // An array
// myslice := myarray[start:end] // A slice made from the array
// if the capacity is reached it go algo doubles the capacity
package main

import "fmt"

func main() {
    // unitialized slice  
	// var nums []int

	// fmt.Println(nums == nil)
	// fmt.Println(len(nums)) o/p: 0
	
	// var nums = make([]int, 0, 5) // ([]int - type, size, capacity)
	// capacity -> maximum numbers of elements can fit

	// fmt.Println(cap(nums)) o/p: 2 ifcapacity not used
	// fmt.Println(cap(nums)) o/p: 5 if capacity is used
	// fmt.Println(nums) o/p: [0 0]
	// fmt.Println(nums)      //o/p: [0 0 1]
	// nums[0] = 1
	// nums[1] = 2
	// nums = append(nums, 3)
	// nums = append(nums, 4)
	// nums = append(nums, 5)
	// nums = append(nums, 6)
	// fmt.Println(nums)      //o/p: [0 0 1]
	// fmt.Println(cap(nums)) //0/p: 5

	// var nums = []int{1, 2, 3}
	// fmt.Println(nums[0:1])
	// fmt.Println(nums[:1])
	// fmt.Println(nums[1:])
	// [1]
	// [1]
	// [2 3]

	// 2D Slice

	var nums = [][]int{{1, 2}, {3, 4, 5}}
	fmt.Println(nums)
	
}
