package triangle

import "math"

type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

func KindFromSides(a, b, c float64) Kind {
	if a <= 0 || b <= 0 || c <= 0 || math.IsNaN(a) || math.IsNaN(b) ||
		math.IsNaN(c) || math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return NaT
	}
	if a == b && b == c {
		return Equ
	}
	if (a+b) < c || (b+c) < a || (c+a) < b {
		return NaT
	}

	sideLengthCount := make(map[float64]int)
	sideLengthCount[a] = 1
	sideLengthCount[b] = sideLengthCount[b] + 1
	sideLengthCount[c] = sideLengthCount[c] + 1

	if len(sideLengthCount) == 1 {
		return Equ
	}
	if len(sideLengthCount) == 2 {
		return Iso
	}
	if len(sideLengthCount) == 3 {
		return Sca
	}

	return NaT
}
