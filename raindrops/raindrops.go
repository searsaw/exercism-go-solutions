package raindrops

import (
	"bytes"
	"strconv"
)

// Convert converts a number to a string, the
// contents of which depends on the number's prime factors.
// - If the number contains 3 as a prime factor, output 'Pling'.
// - If the number contains 5 as a prime factor, output 'Plang'.
// - If the number contains 7 as a prime factor, output 'Plong'.
// - If the number does not contain 3, 5, or 7 as a prime factor,
//   just pass the number's digits straight through.
func Convert(num int) string {
	var buffer bytes.Buffer

	if num%3 == 0 {
		buffer.WriteString("Pling")
	}

	if num%5 == 0 {
		buffer.WriteString("Plang")
	}

	if num%7 == 0 {
		buffer.WriteString("Plong")
	}

	if buffer.Len() <= 0 {
		return strconv.Itoa(num)
	}

	return buffer.String()
}
