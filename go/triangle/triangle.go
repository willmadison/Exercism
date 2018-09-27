package triangle

import (
	"math"
)

const testVersion = 3

// Kind represents the a particular kind of triangle.
type Kind int

const (
	// NaT not a triangle
	NaT Kind = iota
	// Equ equilateral
	Equ
	// Iso isosceles
	Iso
	// Sca scalene
	Sca
)

// KindFromSides determines based on the sides given the type of triangle it represents.
func KindFromSides(a, b, c float64) Kind {
	kind := NaT

	if meetsTriangleInequalityTheorum(a, b, c) &&
		areOfValidLength(a, b, c) {
		switch {
		case a == b && a == c:
			kind = Equ
		case (a == b || a == c) && b != c ||
			(b == c || b == a) && a != c ||
			(c == b || c == a) && a != b:
			kind = Iso
		case a != b && a != c && c != b:
			kind = Sca
		}
	}

	return kind
}

func meetsTriangleInequalityTheorum(a, b, c float64) bool {
	return a+b >= c && a+c >= b && b+c >= a
}

func areOfValidLength(a, b, c float64) bool {
	sides := []float64{a, b, c}

	for _, side := range sides {
		if side <= 0 {
			return false
		}

		if math.IsInf(side, 0) {
			return false
		}

		if math.IsNaN(side) {
			return false
		}
	}

	return true
}
