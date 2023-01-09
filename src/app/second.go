package main

import "fmt"

// Struct (Declaration for line 13...)
type Doctor struct {
	number     int
	actorName  string
	companions []string
}

func main() {
	aDoctor := Doctor{
		number:    3,
		actorName: "John Dumelo",
		companions: []string{
			"Liz Johnson, ",
			"Mercy Aigbe",
		},
	}
	fmt.Println(aDoctor)

	// If Statement
	num := 50
	guess := -5
	if guess < 1 {
		fmt.Println("The guess must be greater than 1!")
	} else if guess > 100 {
		fmt.Println("The guess must be less than 100!")
	} else {
		if guess < num {
			fmt.Println("Too low")
		}
		if guess > num {
			fmt.Println("Too high")
		}
		if guess == num {
			fmt.Println("You got it!")
		}
	}

	// Switch Statement
	i := 10
	switch {
	case i <= 10:
		fmt.Println("Less than or equal to ten")
	case i >= 20:
		fmt.Println("Less than or equal to twenty")
	default:
		fmt.Println("Greater than twenty")
	}

	// Looping (For Statement)
	for i := 0; i < 5; i++ { // i is scoped to the for loop
		fmt.Println(i)
	}

	for i, j := 0, 0; i < 8; i, j = i+2, j+3 {
		fmt.Println(i, j)
	}

	// Defer, Panic, Recover
	fmt.Println("Start")
	defer fmt.Println("Middle")
	fmt.Println("End")

	fmt.Println("Start")
	panic("Something bad happened!")
	fmt.Println("End")
}
