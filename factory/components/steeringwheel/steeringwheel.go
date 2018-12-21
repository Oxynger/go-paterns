package steeringwheel

import (
	"fmt"

	"../../components"
)

type Imp struct {
	material string
	raduis   float32
}

func (s Imp) GetFullInfo() string {
	return fmt.Sprintf("Streering wheel: \n\t material: %s \n\t radius: %f \n",
		s.material, s.raduis)
}

// Constructor of whell
func Constructor(material string, radius float32) components.Component {
	return Imp{
		material: material,
		raduis:   radius,
	}
}
