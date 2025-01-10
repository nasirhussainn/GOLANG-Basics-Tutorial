// Lecture 3: Control Structures

package main

import "fmt"

func main() {
	// 1. Conditional Statements (if, if-else, switch)
	fmt.Println("1. Conditional Statements")

	// if statement
	age := 20
	if age >= 18 {
		fmt.Println("You are an adult!")
	}

	// if-else statement
	if age >= 21 {
		fmt.Println("You are old enough to drink in the US.")
	} else {
		fmt.Println("You are not old enough to drink in the US.")
	}

	// if-else if-else statement
	weather := "sunny"
	if weather == "rainy" {
		fmt.Println("It's raining, take an umbrella!")
	} else if weather == "cloudy" {
		fmt.Println("It's cloudy, you might want a jacket.")
	} else {
		fmt.Println("It's sunny, enjoy the day!")
	}

	// 2. Switch Statement
	fmt.Println("\n2. Switch Statement")

	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	default:
		fmt.Println("Weekend")
	}

	// Switch with no condition (acts like if-else)
	season := "winter"
	switch {
	case season == "summer":
		fmt.Println("It's hot!")
	case season == "winter":
		fmt.Println("It's cold!")
	case season == "spring":
		fmt.Println("Flowers are blooming!")
	case season == "fall":
		fmt.Println("Leaves are falling!")
	default:
		fmt.Println("Season unknown!")
	}

	// 3. Loops (for, range)
	fmt.Println("\n3. Loops")

	// for loop
	fmt.Println("Using a for loop:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("Iteration %d\n", i)
	}

	// range with arrays/slices
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println("Using range with a slice:")
	for index, value := range numbers {
		fmt.Printf("Index %d: Value %d\n", index, value)
	}

	// 4. Break and Continue
	fmt.Println("\n4. Break and Continue")

	// break example
	fmt.Println("Using break in a loop:")
	for i := 1; i <= 5; i++ {
		if i == 3 {
			fmt.Println("Breaking out of the loop at i =", i)
			break
		}
		fmt.Println("i =", i)
	}

	// continue example
	fmt.Println("Using continue in a loop:")
	for i := 1; i <= 5; i++ {
		if i == 3 {
			fmt.Println("Skipping i =", i)
			continue
		}
		fmt.Println("i =", i)
	}
}
