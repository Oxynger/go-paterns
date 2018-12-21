package tesla

import (
	"image/color"

	"../../components"
	"../../components/body"
	"../../components/engine"
	"../../components/steeringwheel"
)

var (
	// Engine is engine of produced by Tesla
	Engine = func() components.Component {
		return engine.Constructor("Bzzzzzzzzzzzzzz", 9000)
	}

	// Body is body of produced by Tesla
	Body = func() components.Component {
		return body.Constructor("Plastic", color.RGBA{0, 0, 0, 0})
	}

	// SteeringWheel is street wheel of produced by Tesla
	SteeringWheel = func() components.Component {
		return steeringwheel.Constructor("Lather", 5)
	}
)
