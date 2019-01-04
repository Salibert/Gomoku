package axis

import "github.com/Salibert/Gomoku/back/server/inter"

// Axis ...
type Axis struct {
	X, Y int
}

// Radius is a node array cut by the same axis with a length of 12 squares
type Radius [11]inter.Node

// DialRightAxes ...
var DialRightAxes = [4]Axis{
	Axis{X: 0, Y: -1}, // ↑
	Axis{X: 1, Y: -1}, // ↗
	Axis{X: 1, Y: 0},  // →
	Axis{X: 1, Y: 1},  // ↘
}

// DialLeftAxes ...
var DialLeftAxes = [4]Axis{
	Axis{X: 0, Y: 1},   // ↓
	Axis{X: -1, Y: 1},  // ↙
	Axis{X: -1, Y: 0},  // ←
	Axis{X: -1, Y: -1}, // ↖
}

func inverse(nb int) int {
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
func (axis *Axis) Multiply(multiplier int) Axis {
	return Axis{X: axis.X * multiplier, Y: axis.Y * multiplier}
}

// Divide divide axis by divisor
func (axis *Axis) Divide(divisor int) Axis {
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
