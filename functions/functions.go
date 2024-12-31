package main

// this is return type for int
// func add(a, b int) int { can be written like this
// func add(a int, b int) int {
// 	return a + b
// }

// func getLanguages() (string, string, string){
// 	return "go", "js","c"
// }

// func processIt(fn func(a int) int){
// 	fn(1)
// }
func processIt() func(a int) int{
	return func(a int) int {
		return 4
	}
}
func main(){
	// result := add(3, 5)
	// fmt.Println(result)

	// l1, l2, l3 := getLanguages()
	
	// fmt.Println(l1, l2, l3) 
	// fmt.Println(getLanguages())

	fn := func(a int) int {
		return 2
	}
	processIt(fn)
}