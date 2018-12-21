package toyota

import (
	"image/color"

	"../../components"
	"../../components/body"
	"../../components/engine"
	"../../components/wheel"
)

var (
	// Engine is engine of produced by Toyota
	Engine = func() components.Component {
		return engine.Constructor("Toyyyyyyooooota", 10000)
	}

	// Wheel is wheel of produced by Toyota
	Wheel = func() components.Component {
		return wheel.Constructor("Steel", 10, color.RGBA{128, 0, 0, 0})
	}

	// Body is body of produced by Toyota
	Body = func() components.Component {
		return body.Constructor("Aluminum", color.RGBA{0, 128, 128, 0})
	}
)
