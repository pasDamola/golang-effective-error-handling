package main

import (
	"fmt"
)

type Truck struct {
	id string
}

// processTruck handles the loading and unloading of a truck
func processTruck(truck Truck) error {
	fmt.Printf("Processing Truck %s\n", truck.id)
	return nil
}

func main() {
	trucks := []Truck{
		{id: "Truck-1"},
		{id: "Truck-2"},
		{id: "Truck-3"},
		{id: "Truck-4"},
	}

	for _, truck := range trucks {
		fmt.Printf("Truck %s has arrived\n", truck.id)
		if err := processTruck(truck); err != nil {
			fmt.Println("error")
		}
	}
}
