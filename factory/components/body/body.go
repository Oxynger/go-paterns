package body

import (
	"fmt"
	"image/color"

	"../../components"
)

type Imp struct {
	material string
	color    color.Color
}

func (body Imp) GetFullInfo() string {
	r, g, b, a := body.color.RGBA()

	return fmt.Sprintf("Body: \n\t material: %s \n\t color: (%d, %d, %d, %d) \n",
		body.material, r, g, b, a)
}

// Constructor of whell
func Constructor(material string, color color.Color) components.Component {
	return Imp{
		material: material,
		color:    color,
	}

}
