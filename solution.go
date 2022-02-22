package square

import (
	"math"
)

type intCustomType int

const SidesCircle = intCustomType(0)
const SidesTriangle = intCustomType(3)
const SidesSquare = intCustomType(4)

func CalcSquare(sideLen float64, sidesNum intCustomType) float64 {
	var result float64

	switch sidesNum {
	case intCustomType(0):
		result = calcAreaCircle(sideLen)
	case intCustomType(3):
		result = calcAreaTriangle(sideLen)
	case intCustomType(4):
		result = calcAreaSquare(sideLen)
	default:
		result = 0
	}

	return result
}

func calcAreaTriangle(sideLen float64) float64 {
	var halfPerimetr float64
	halfPerimetr = (3 * sideLen) / 2
	return math.Sqrt(halfPerimetr * (halfPerimetr - sideLen) * (halfPerimetr - sideLen) * (halfPerimetr - sideLen))
}

func calcAreaSquare(sideLen float64) float64 {
	return sideLen * sideLen
}

func calcAreaCircle(sideLen float64) float64 {
	return math.Pi * sideLen * sideLen
}
