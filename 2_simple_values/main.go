package main

import (
	"fmt"
)

func main() {
	// var name string = "golang"
    // fmt.Println(name)

	// shorthand syntax
	// name := "golang"
	// fmt.Println(name)

	// for- only construct in go for looping //

	// while loop
	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	// for i:= 0;i < 3;i++ {
	// 	fmt.Println(i)
	// }
	// break contine

	// 1.22 range
	// for i:= range 3{
	// 	fmt.Println(i)
	// }

	// if else
	// age := 17

	// if age >= 10 {
	// 	fmt.Println("person is an adult")
	// } else if age == 17 {
	// 	fmt.Println("not adult")
	// } else {
	// 	fmt.Println("not adult")
	// }

	// types of switch
	// multiple swicth condition
	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday :
	// 	fmt.Println("Weekend")
	//  default: 
	// 	fmt.Println("Default")
	// }

	// type switch
	whoAmI := func(i interface{}){
		switch t := i.(type) {
		case int:
			fmt.Println("integer")
		case string:
			fmt.Println("string")
		case bool:
			fmt.Println("Boolean")
		default: 
			fmt.Println("Default", t)
		}
	}

	whoAmI(true)
}
