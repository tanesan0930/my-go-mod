// Package hsd provides a function to calculate the Hamming distance between two strings.
package hsd

// StringDistance returns the Hamming distance between two strings.
func StringDistance(lhs, rhs string) int {
	return Distance([]rune(lhs), []rune(rhs))
}

// Distance returns the Hamming distance between two rune slices.
func Distance(a, b []rune) int {
	dist := 0
	if len(a) != len(b) {
		return -1
	}
	for i := range a {
		if a[i] != b[i] {
			dist++
		}
	}
	return dist
}
