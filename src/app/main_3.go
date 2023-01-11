package main

import "fmt"

// Struct (Declaration for line 13...)
type Actor struct {
	number     int
	actorName  string
	companions []string
}

func main() {
	aActor := Actor{ // refer to the struct in line 6
		number:    3,
		actorName: "Oga Bello",
		companions: []string{
			"Mr. Government,",
			"Ishaya Bello, ",
			"Adufe",
		},
	}
	fmt.Println(aActor)

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
}
