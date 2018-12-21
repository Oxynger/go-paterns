package transmition

import (
	"fmt"

	"../../components"
)

type Imp struct {
	fourWheelDrive  bool
	rearDrive       bool
	frontWheelDrive bool
}

func (t Imp) GetFullInfo() string {
	return fmt.Sprintf("Transmition Drive: \n\t  four-wheel drive: %t \n\t rear drive: %t \n\t front-wheel drive %t \n",
		t.fourWheelDrive, t.rearDrive, t.frontWheelDrive)
}

// Constructor of whell
func Constructor(four, rear, front bool) components.Component {
	return Imp{
		fourWheelDrive:  four,
		rearDrive:       rear,
		frontWheelDrive: front,
	}
}
