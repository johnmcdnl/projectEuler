package problems

/**
Problem 9 - Special Pythagorean triplet
 Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
a2 + b2 = c2
For example, 32 + 42 = 9 + 16 = 25 = 52.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
 */
func Problem009() int {

	var a, b, c int

	loop:
	for c = 3; c <= 998; c++ {
		for b = c - 1; b > 2; b-- {
			for a = b - 1; a > 1; a-- {
				if a + b + c == 1000 {
					if a * a + b * b == c * c {
						break loop
					}
				}
			}
		}
	}

	return a * b * c
}
