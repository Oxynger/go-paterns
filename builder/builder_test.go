package builder

import (
	"testing"

	"../factory/car"
)

func TestBuilder(t *testing.T) {
	expect := []string{
		"BUeeeeeeeMmmmmVVV",
		"Drdrdrdrdrdrdrd",
		"Toyyyyyyooooota",
		"Bzzzzzzzzzzzzzz",
	}

	var builder Builder

	builder.Add(car.Models.Lada())

	cars := builder.Build()

	builder.Add(car.Models.Skoda()).Add(car.Models.Lada()).Add(car.Models.Toyota()).Add(car.Models.Tesla())

	cars = builder.Build()

	for index := range expect {
		roar := cars[index].Race()
		if expect[index] != roar {
			t.Errorf("Expect: \n %s \n but get: \n %s", expect[index], roar)
		}
	}
}

func TestNotCorrectBuilder(t *testing.T) {
	type myCar struct {
		car.Car
	}

	var builder Builder
	builder.Add(myCar{})

	wrongCar := builder.Build()

	if wrongCar != nil {
		t.Error("Expect: nil, but get: \n", wrongCar)
	}

}
