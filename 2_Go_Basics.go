// Lecture 2: Go Basics

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 1. Variables and Constants
	fmt.Println("1. Variables and Constants")
	var name string = "GoLang"        // Explicit type declaration
	age := 14                        // Implicit type declaration
	const version string = "1.20.5"  // Constant declaration

	fmt.Printf("Name: %s, Age: %d, Version: %s\n\n", name, age, version)

	// 2. Data Types
	fmt.Println("2. Data Types")
	var (
		isAwesome bool    = true
		rating    float32 = 4.5
	)
	fmt.Printf("Is Go awesome? %t, Rating: %.1f\n\n", isAwesome, rating)

	// 3. Input and Output
	fmt.Println("3. Input and Output")
	fmt.Print("Enter your favorite programming language: ")
	reader := bufio.NewReader(os.Stdin)
	favLang, _ := reader.ReadString('\n')
	fmt.Printf("You love: %s\n\n", favLang)

	// 4. Comments in Go
	// Single-line comment
	/*
		Multi-line comment
		This is often used for longer explanations.
	*/

	fmt.Println("4. Comments in Go")
	fmt.Println("// Single-line comments are preceded by //")
	fmt.Println("/* Multi-line comments are wrapped in /* */\n")

	// 5. Keywords and Identifiers
	fmt.Println("5. Keywords and Identifiers")
	// Keywords like `if`, `for`, `func` are reserved and can't be used as variable names.
	// Identifiers are names used for variables, constants, functions, etc.
	var identifierExample int = 42
	fmt.Printf("Example Identifier: %d\n\n", identifierExample)

	// 6. fmt Package for Printing
	fmt.Println("6. fmt Package for Printing")
	fmt.Println("fmt.Println() - Prints with a newline")
	fmt.Printf("fmt.Printf() - Formats output: %s is %d years old.\n\n", "Alice", 25)

	// 7. Basic Arithmetic Operations
	fmt.Println("7. Basic Arithmetic Operations")
	a, b := 10, 3
	fmt.Printf("Addition: %d + %d = %d\n", a, b, a+b)
	fmt.Printf("Subtraction: %d - %d = %d\n", a, b, a-b)
	fmt.Printf("Multiplication: %d * %d = %d\n", a, b, a*b)
	fmt.Printf("Division: %d / %d = %d\n", a, b, a/b)
	fmt.Printf("Modulus: %d %% %d = %d\n\n", a, b, a%b)

	// 8. Type Conversion
	fmt.Println("8. Type Conversion")
	var x int = 42
	var y float64 = float64(x) // Explicit conversion
	fmt.Printf("Converted int to float64: %.2f\n\n", y)

	// 9. Arrays
	fmt.Println("9. Arrays")
	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", numbers)
	fmt.Printf("First Element: %d, Length: %d\n\n", numbers[0], len(numbers))

	// 10. Slices
	fmt.Println("10. Slices")
	var slice = []int{10, 20, 30, 40, 50}
	fmt.Println("Slice:", slice)
	fmt.Printf("First Element: %d, Length: %d, Capacity: %d\n\n", slice[0], len(slice), cap(slice))

	// 11. Maps
	fmt.Println("11. Maps")
	var userInfo = map[string]string{"name": "Alice", "country": "Wonderland"}
	fmt.Printf("User Info: Name - %s, Country - %s\n\n", userInfo["name"], userInfo["country"])
}
