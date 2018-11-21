package axis

// Axis ...
type Axis struct {
	X, Y int32
}

// Axes is the list Axis
var Axes = [8]Axis{
	Axis{X: 1, Y: 0},   // →
	Axis{X: 1, Y: 1},   // ↘
	Axis{X: 0, Y: 1},   // ↓
	Axis{X: -1, Y: 1},  // ↙
	Axis{X: -1, Y: 0},  // ←
	Axis{X: -1, Y: -1}, // ↖
	Axis{X: 1, Y: -1},  // ↑
	Axis{X: 1, Y: -1},  // ↗
}

func inverse(nb int32) int32 {
	switch nb {
	case -1:
		return 1
	case 1:
		return -1
	default:
		return 0
	}
}

// Inverse ...
func (axis *Axis) Inverse() Axis {
	return Axis{X: inverse(axis.X), Y: inverse(axis.Y)}

}

// Multiply multiply axis by multiplier
func (axis *Axis) Multiply(multiplier int32) Axis {
	return Axis{X: axis.X * multiplier, Y: axis.Y * multiplier}
}

// Divide divide axis by divisor
func (axis *Axis) Divide(divisor int32) Axis {
	return Axis{X: axis.X / divisor, Y: axis.Y / divisor}
}

// Add Addition axis by divisor
func (axis *Axis) Add(add Axis) Axis {

	return Axis{X: axis.X + add.X, Y: axis.Y + add.Y}
}

// Sub subtract axis by divisor
func (axis *Axis) Sub(sub Axis) Axis {
	return Axis{X: axis.X - sub.X, Y: axis.Y - sub.Y}
}
