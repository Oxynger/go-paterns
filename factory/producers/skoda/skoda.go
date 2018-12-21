package skoda

import (
	"image/color"

	"../../components"
	"../../components/body"
	"../../components/steeringwheel"
)

var (
	// Body is body of produced by Skoda
	Body = func() components.Component {
		return body.Constructor("Plastic", color.RGBA{250, 0, 0, 0})
	}

	// SteeringWheel is street wheel of produced by Skoda
	SteeringWheel = func() components.Component {
		return steeringwheel.Constructor("Carbone", 4)
	}
)
