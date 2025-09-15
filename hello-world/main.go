package main

import (
	"errors"
	"fmt"
)

func main() {
	// all print functions are in fmt package
	// 1. fmt.Print()
	fmt.Printf("Format Printf: %.2f\n", 3.14159)

	// 2. fmt.Println()
	fmt.Println("Hello, World!")
	message := fmt.Sprintf("Pi value up to 2 decimals: %.2f", 3.14159)
	fmt.Println(message)

	// 3. fmt.Printf()
	name := "Nikhil"
	age := 23
	fmt.Printf("My name is %s and I am %d years old.\n", name, age)
	// 4. fmt.Sprintf()
	formattedMessage := fmt.Sprintf("My name is %s and I am %d years old.", name, age)
	fmt.Println(formattedMessage)
	// fmt.Sprintf() returns the formatted string instead of printing it
	// 5. fmt.Errorf()
	err := fmt.Errorf("something went wrong: %v", errors.New("connection failed"))
	fmt.Println(err)
}

// Documents
// 2. fmt.Println()

// Definition: Prints the arguments to the standard output with a space between arguments and a newline at the end.

// Usage: For printing human-readable outputs with automatic line breaks.

// Example:
// package main

// import "fmt"

// func main() {
//     fmt.Println("Hello")
//     fmt.Println("World")
// }

// Output:

// Hello
// World

// Scenario: Useful when you want clean, separate outputs in logs or debugging without manually adding \n.

// 3. fmt.Printf()

// Definition: Prints formatted output using format specifiers, similar to C's printf.

// Usage: When you want more control over the output format (integers, strings, floats, alignment, etc.).

// Example:
// package main

// import "fmt"

// func main() {
//     name := "Nikhil"
//     age := 23
//     fmt.Printf("My name is %s and I am %d years old.\n", name, age)
// }

// Output:

// My name is Nikhil and I am 23 years old.

// Scenario: Useful for formatted messages, such as printing variables with types, debugging structured data, or displaying UI-like console outputs.

// 4. fmt.Sprintf()

// Definition: Similar to Printf, but instead of printing, it returns the formatted string.

// Usage: When you need to store formatted text for later use (logging, saving to files, sending over a network).

// Example:
// package main

// import "fmt"

// func main() {
//     message := fmt.Sprintf("Pi value up to 2 decimals: %.2f", 3.14159)
//     fmt.Println(message)
// }

// Output:

// Pi value up to 2 decimals: 3.14

// Scenario: Useful for string construction before printing or writing to logs/files.

// 5. fmt.Errorf()

// Definition: Creates a formatted error message (returns an error object, not printed directly).

// Usage: For error handling with custom messages.

// Example:
// package main

// import (
//     "errors"
//     "fmt"
// )

// func main() {
//     err := fmt.Errorf("something went wrong: %v", errors.New("connection failed"))
//     fmt.Println(err)
// }

// Output:

// something went wrong: connection failed

// Scenario: Used for creating descriptive error messages in Go applications.
