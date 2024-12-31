package main

import "fmt"

func chnageNum(num *int) {
	*num = 5

	fmt.Println("in changeNum", num)
}
func main() {
	num := 1

	chnageNum(&num)

	fmt.Println("memory address", num)

}