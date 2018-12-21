package car_test

import (
	"testing"

	"../producers/bmw"
	"../producers/lada"
	"../producers/skoda"
	"../producers/toyota"
	"./"
)

func TestCar(t *testing.T) {
	expect := []string{
		"Drdrdrdrdrdrdrd",
		"Toyyyyyyooooota",
	}

	cars := car.Cars{
		car.Constructor(
			lada.Engine(),
			lada.Body(),
			lada.Transmition(),
			lada.Wheel(),
			lada.SteeringWheel(),
		),
		car.Constructor(
			toyota.Engine(),
			toyota.Body(),
			bmw.Transmition(),
			toyota.Wheel(),
			skoda.SteeringWheel(),
		),
	}

	for index := range expect {
		roar := cars[index].Race()
		if expect[index] != roar {
			t.Errorf("Expect: \n %s \n but get: \n %s", expect[index], roar)
		}
	}
}
