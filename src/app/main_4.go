package main

import "fmt"

func main() {
	// Looping (For Statement)
	for i := 0; i < 5; i++ { // i is scoped to the for loop
		fmt.Println(i)
	}

	for i, j := 0, 0; i < 8; i, j = i+2, j+3 {
		fmt.Println(i, j)
	}

	// Defer
	fmt.Println("Start")
	defer fmt.Println("Middle")
	fmt.Println("End")

	// Pointers
	var a int = 42
	var b *int = &a
	fmt.Println(a, *b)
	a = 27
	fmt.Println(a, *b)
	*b = 14
	fmt.Println(a, *b)

	var c = [3]int{1, 2, 3}
	var d *int = &c[0]
	var e *int = &c[1]
	fmt.Printf("%v, %p, %p\n", c, d, e)
}
