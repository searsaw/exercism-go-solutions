package triangle

// The Kind of triangle
type Kind int

// The Kinds of triangles that are tested for
const (
	Equ Kind = iota
	Iso
	Sca
	NaT
)

// KindFromSides returns the kind of triangle
// based on the lengths of the sides passed
// into the function
func KindFromSides(a, b, c float64) Kind {
	if isTriangle(a, b, c) {
		if isEquilateral(a, b, c) {
			return Equ
		} else if isIsosceles(a, b, c) {
			return Iso
		}

		return Sca
	}

	return NaT
}

func isTriangle(a, b, c float64) bool {
	return a+b > c && a+c > b && b+c > a
}

func isEquilateral(a, b, c float64) bool {
	return a == b && b == c
}

func isIsosceles(a, b, c float64) bool {
	return a == b || a == c || b == c
}
