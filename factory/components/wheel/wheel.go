package wheel

import (
	"fmt"
	"image/color"

	"../../components"
)

type Imp struct {
	material string
	radius   float32
	color    color.Color
}

func (w Imp) GetFullInfo() string {
	r, g, b, a := w.color.RGBA()

	return fmt.Sprintf("Wheel: \n\t material: %s \n\t radius: %f, \n\t color: (%d, %d, %d, %d) \n",
		w.material, w.radius, r, g, b, a)
}

// Constructor of whell
func Constructor(material string, radius float32, color color.Color) components.Component {
	return Imp{
		material: material,
		radius:   radius,
		color:    color,
	}

}
