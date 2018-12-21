package factory

import (
	"reflect"

	"./car"
	"./producers/bmw"
	"./producers/lada"
	"./producers/skoda"
	"./producers/tesla"
	"./producers/toyota"
)

func buildToyota() car.Car {
	return car.Constructor(
		toyota.Engine(),
		toyota.Body(),
		bmw.Transmition(),
		toyota.Wheel(),
		skoda.SteeringWheel(),
	)
}

func buildLada() car.Car {
	return car.Constructor(
		lada.Engine(),
		lada.Body(),
		lada.Transmition(),
		lada.Wheel(),
		lada.SteeringWheel(),
	)
}

func buildSkoda() car.Car {
	return car.Constructor(
		bmw.Engine(),
		skoda.Body(),
		bmw.Transmition(),
		toyota.Wheel(),
		skoda.SteeringWheel(),
	)
}

func buildTesla() car.Car {
	return car.Constructor(
		tesla.Engine(),
		tesla.Body(),
		bmw.Transmition(),
		toyota.Wheel(),
		tesla.SteeringWheel(),
	)
}

// MakeCar Factory method
func MakeCar(currentCar car.Car) car.Car {

	switch reflect.TypeOf(currentCar) {
	case reflect.TypeOf(car.Models.Toyota()):
		return buildToyota()
	case reflect.TypeOf(car.Models.Lada()):
		return buildLada()
	case reflect.TypeOf(car.Models.Tesla()):
		return buildTesla()
	case reflect.TypeOf(car.Models.Skoda()):
		return buildSkoda()
	}

	return nil
}
