package main

import (
	"fmt"

	"./builder"
	"./factory/car"
)

func main() {
	var builder builder.Builder
	for index := 0; index < 10; index++ {
		builder.Add(
			car.Models.Toyota(),
		).Add(
			car.Models.Lada(),
		)
	}

	for index := 0; index < 3; index++ {
		builder.Add(
			car.Models.Tesla(),
		)
	}

	carPark := builder.Add(car.Models.Skoda()).Build()

	for _, currentValue := range carPark {
		fmt.Print(currentValue.GetInfo())
	}

	fmt.Print(carPark[16].Race())
}
