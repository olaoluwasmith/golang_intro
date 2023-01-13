package main

import (
	"fmt"
)

// Functions and methods
func main() {
	for i := 0; i < 5; i++ {
		sayMessage("Hello Go!", i)
	}

	s := sum(1, 2, 3, 4, 5)
	fmt.Println("The sum is ", *s)

	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()
}

func sayMessage(msg string, idx int) {
	fmt.Println(msg)
	fmt.Println("The value of the index is", idx)
}

func sum(values ...int) *int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return &result
}

type greeter struct { // struct for method greet()
	greeting string
	name     string
}

func (g greeter) greet() { // method
	fmt.Println(g.greeting, g.name)
}
