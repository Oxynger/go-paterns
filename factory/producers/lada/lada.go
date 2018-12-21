package lada

import (
	"image/color"

	"../../components"
	"../../components/body"
	"../../components/engine"
	"../../components/steeringwheel"
	"../../components/transmition"
	"../../components/wheel"
)

var (
	// Transmition is transmition of produced by lada
	Transmition = func() components.Component {
		return transmition.Constructor(true, false, false)
	}

	// Engine is engine of produced by lada
	Engine = func() components.Component {
		return engine.Constructor("Drdrdrdrdrdrdrd", 146)
	}

	// Wheel is wheel of produced by lada
	Wheel = func() components.Component {
		return wheel.Constructor("cast-iron", 10, color.RGBA{128, 0, 0, 0})
	}

	// Body is body of produced by lada
	Body = func() components.Component {
		return body.Constructor("cast-iron", color.RGBA{30, 30, 30, 0})
	}

	// SteeringWheel is street wheel of produced by lada
	SteeringWheel = func() components.Component {
		return steeringwheel.Constructor("Lather", 8)
	}
)
