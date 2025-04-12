package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrNotImplemented = errors.New("not implemented")
	ErrTruckNotFound  = errors.New("truck not found")
)

type Truck struct {
	id string
}

func (t *Truck) LoadCargo() error {
	return ErrTruckNotFound
}

func (t *Truck) UnloadCargo() error {
	return nil
}

// processTruck handles the loading and unloading of a truck
// add this
// add this again
// one more time
func processTruck(truck Truck) error {
	fmt.Printf("Processing Truck %s\n", truck.id)

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("Error truck did not load correctly %w", err)
	}

	if err := truck.UnloadCargo(); err != nil {
		return fmt.Errorf("Error truck did not unload correctly %w", err)
	}
	return ErrNotImplemented
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
			if errors.Is(err, ErrNotImplemented) {
				// we do this
			}

			if errors.Is(err, ErrTruckNotFound) {
				// we do this
			}

			log.Fatalf("Error processing truck %s\n", err)
		}
	}
}
