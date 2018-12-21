package factory_test

import (
	"testing"

	"./"
	"./car"
)

func TestFactory(t *testing.T) {
	expect := []string{
		"BUeeeeeeeMmmmmVVV",
		"Drdrdrdrdrdrdrd",
		"Toyyyyyyooooota",
		"Bzzzzzzzzzzzzzz",
	}

	var cars []car.Car
	cars = append(cars, factory.MakeCar(car.Models.Skoda()))
	cars = append(cars, factory.MakeCar(car.Models.Lada()))
	cars = append(cars, factory.MakeCar(car.Models.Toyota()))
	cars = append(cars, factory.MakeCar(car.Models.Tesla()))

	for index := range expect {
		roar := cars[index].Race()
		if expect[index] != roar {
			t.Errorf("Expect: \n %s \n but get: \n %s", expect[index], roar)
		}
	}
}

func TestNotCorrectFactory(t *testing.T) {
	type myCar struct {
		car.Car
	}

	wrongCar := factory.MakeCar(myCar{})
	if wrongCar != nil {
		t.Error("Expect: nil, but get: \n", wrongCar)
	}

}
