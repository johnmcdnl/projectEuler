package problems

import (
	"bytes"
	"strconv"
)

/**
Problem 40 - Champernowne's constant - https://projecteuler.net/problem=40

An irrational decimal fraction is created by concatenating the positive integers:
0.123456789101112131415161718192021...
It can be seen that the 12th digit of the fractional part is 1.

If dn represents the nth digit of the fractional part, find the value of the following expression.

d1 × d10 × d100 × d1000 × d10000 × d100000 × d1000000
 */
func Problem040() int {

	var buffer bytes.Buffer
	i := 0
	for len(buffer.Bytes()) <= 1000000 {
		buffer.WriteString(strconv.Itoa(i))
		i++
	}
	decimal := buffer.String()
	sum := 1
	for i := 1; i <= 1000000; i *= 10 {
		num, _ := strconv.Atoi(string(decimal[i]))
		sum *= num
	}

	return sum
}
