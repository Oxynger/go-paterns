package bmw

import (
	"../../components"
	"../../components/engine"
	"../../components/transmition"
)

var (
	// Transmition is transmition of produced by bmw
	Transmition = func() components.Component {
		return transmition.Constructor(true, false, false)
	}

	// Engine is engine of produced by bmw
	Engine = func() components.Component {
		return engine.Constructor("BUeeeeeeeMmmmmVVV", 100500)
	}
)
