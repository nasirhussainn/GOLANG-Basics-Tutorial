// Lecture 4: Functions in Go

package main

import (
	"fmt"
)

// 1. Defining and Calling Functions
func greet(name string) {
	fmt.Println("Hello, " + name)
}

// 2. Function Parameters and Return Values
func add(a int, b int) int {
	return a + b
}

// 3. Named Return Values
func multiply(a, b int) (result int) {
	result = a * b
	return // named return value used here
}

// 4. Variadic Functions
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// 5. Anonymous Functions and Closures
func main() {
	// 1. Defining and Calling Functions
	fmt.Println("1. Defining and Calling Functions:")
	greet("Go Developer")

	// 2. Function Parameters and Return Values
	fmt.Println("\n2. Function Parameters and Return Values:")
	result := add(10, 5)
	fmt.Println("Addition Result:", result)

	// 3. Named Return Values
	fmt.Println("\n3. Named Return Values:")
	result2 := multiply(4, 3)
	fmt.Println("Multiplication Result:", result2)

	// 4. Variadic Functions
	fmt.Println("\n4. Variadic Functions:")
	result3 := sum(1, 2, 3, 4, 5)
	fmt.Println("Sum Result:", result3)

	// 5. Anonymous Functions and Closures
	fmt.Println("\n5. Anonymous Functions and Closures:")
	// Defining an anonymous function and calling it immediately
	func(name string) {
		fmt.Println("Hello from anonymous function, " + name)
	}("Anonymous")

	// Closure example
	counter := createCounter()
	fmt.Println("Closure Counter:", counter()) // 1
	fmt.Println("Closure Counter:", counter()) // 2
	fmt.Println("Closure Counter:", counter()) // 3

	// 6. Defer, Panic, and Recover
	fmt.Println("\n6. Defer, Panic, and Recover:")

	// Using Defer
	defer fmt.Println("This message is printed last.")
	fmt.Println("This message is printed first.")

	// Using Panic and Recover
	// Uncomment the next two lines to see panic and recover in action:
	// fmt.Println("Calling panic function.")
	// causePanic()

	// To demonstrate panic and recover, we define a function that causes a panic.
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered from panic:", r)
			}
		}()
		fmt.Println("Causing a panic!")
		panic("Something went wrong!")
	}()
}

// Closure function that returns a counter
func createCounter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// Panic function for demonstration
func causePanic() {
	fmt.Println("This will cause a panic!")
	panic("Oops! A panic occurred!")
}
