package debug

import "testing"

type Address struct {
	City  string
	State string
}

type Person struct {
	Name    string
	Age     int
	Address *Address
	Friends []*Person
}

func TestDeepPrint(t *testing.T) {
	p := Person{
		Name: "Alice",
		Age:  30,
		Address: &Address{
			City:  "Wonderland",
			State: "Fantasy",
		},
		Friends: []*Person{
			{Name: "Bob", Age: 25, Address: &Address{City: "Springfield", State: "Illinois"}},
			{Name: "Charlie", Age: 28, Address: &Address{City: "Metropolis", State: "New York"}},
		},
	}

	// Print the struct and its fields deeply
	DeepPrint(&p, "  ")

}
