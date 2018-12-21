package car

import (
	"../components"
	"../components/body"
	"../components/engine"
	"../components/steeringwheel"
	"../components/transmition"
	"../components/wheel"
)

type (
	// Cars is car slice
	Cars []Car
)

type models struct {
}

func (m models) Toyota() toyota {
	return toyota{}
}
func (m models) Tesla() tesla {
	return tesla{}
}

func (m models) Skoda() skoda {
	return skoda{}
}

func (m models) Lada() lada {
	return lada{}
}

// Models of car
var Models models

// Car interface of cars
type Car interface {
	Race() string
	GetInfo() string
}

// imp is implementation Car interface
type imp struct {
	Car
	engine        engine.Imp
	body          body.Imp
	transmition   transmition.Imp
	wheel         wheel.Imp
	steeringwheel steeringwheel.Imp
}

// Race retrun roar of engine
func (c imp) Race() string {
	return c.engine.Roar()
}

// Race retrun roar of engine
func (c imp) GetInfo() string {
	return (c.engine.GetFullInfo() +
		c.body.GetFullInfo() +
		c.steeringwheel.GetFullInfo() +
		c.transmition.GetFullInfo() +
		c.steeringwheel.GetFullInfo())
}

// Constructor of car object
func Constructor(carEngine, carBody, carTransmition, carWheel, carSteeringWheel components.Component) Car {

	return imp{
		engine:        carEngine.(engine.Imp),
		body:          carBody.(body.Imp),
		transmition:   carTransmition.(transmition.Imp),
		wheel:         carWheel.(wheel.Imp),
		steeringwheel: carSteeringWheel.(steeringwheel.Imp),
	}
}

type (
	// Toyota is car model
	toyota imp

	// Tesla is car model
	tesla imp

	// Skoda is car model
	skoda imp

	// Lada is car model
	lada imp
)
