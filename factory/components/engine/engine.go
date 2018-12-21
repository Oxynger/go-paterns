package engine

import (
	"fmt"

	"../../components"
)

type Imp struct {
	sound         string
	powerCapacity float32
}

func (e Imp) GetFullInfo() string {
	return fmt.Sprintf("Engine: \n\t powerCapacity: %f \n",
		e.powerCapacity)
}

// Roar retrun sound of engine
func (e Imp) Roar() string {
	return e.sound
}

// Constructor of whell
func Constructor(roar string, power float32) components.Component {
	return Imp{
		sound:         roar,
		powerCapacity: power,
	}
}
