package builder

import (
	"../factory"
	"../factory/car"
)

// Builder store array of cars with his model, add new car in order and build car park
type Builder struct {
	orderedCars car.Cars
}

// Add new car in order
func (b *Builder) Add(model car.Car) *Builder {
	b.orderedCars = append(b.orderedCars, model)

	return b
}

// Build car pack with ordered cars
func (b *Builder) Build() car.Cars {
	var carPark car.Cars

	for _, currentModel := range b.orderedCars {
		currentCar := factory.MakeCar(currentModel)

		if currentCar == nil {
			return nil
		}

		carPark = append(carPark, currentCar)
	}
	b.orderedCars = car.Cars{}
	return carPark
}
