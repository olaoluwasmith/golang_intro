package main

import "fmt"

func main() {
	// Array
	grades := [3]int{97, 85, 93}     // method 1
	grades_2 := [...]int{97, 85, 93} // method 2
	fmt.Printf("Grades: %v\n", grades)
	fmt.Printf("Grades: %v\n", grades_2)

	var students [3]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "Racheal Iduh,"
	students[1] = "Ayomide Adu,"
	students[2] = "Temitope Imisi"
	fmt.Printf("Students: %v\n", students)

	// Slices
	a := []int{}
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
	a = append(a, 1)
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))

	b := make([]int, 10)
	c := make([]int, 10, 100)
	fmt.Printf("Length: %v\n", len(b))
	fmt.Printf("Capacity: %v\n", cap(b))
	fmt.Printf("Length: %v\n", len(c))
	fmt.Printf("Capacity: %v\n", cap(c))

	//Maps
	statePopulations := map[string]int{
		"Lagos":  234567,
		"Rivers": 201000,
		"Abuja":  120000,
	}
	fmt.Println(statePopulations)
}
