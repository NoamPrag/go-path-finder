package spline

import (
	"math"
)

type Point struct {
	X         float64
	Y         float64
	Direction float64
}

type Spline interface {
	X(s float64) float64
	Y(s float64) float64
	DX(s float64) float64
	DY(s float64) float64
	DDX(s float64) float64
	DDY(s float64) float64
	SetOptimzationParams(params []float64)
	GetOptimizationParams() []float64
	GetOptimizationParamsLength() uint
	GetPoints() []Point
}

const (
	MinSRange float64 = 0
	MaxSRange float64 = 1
	SRange    float64 = MaxSRange - MinSRange
)

func Angle(spline Spline, s float64) float64 {
	return math.Atan(spline.DY(s) / spline.DX(s))
}

func Radius(spline Spline, s float64) (radius float64) {
	dx := spline.DX(s)
	ddx := spline.DDX(s)
	dy := spline.DY(s)
	ddy := spline.DDY(s)

	return (ddx*dy - ddy*dx) / math.Pow(math.Pow(dx, 2)+math.Pow(dy, 2), 1.5)
}
