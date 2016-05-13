package problems

import (
	"strconv"
)

//Problem 43
//Sub-string divisibility

//The number, 1406357289, is a 0 to 9 PanDigital number because it is made up of each of the digits 0 to 9 in some order,
// but it also has a rather interesting sub-string divisibility property.

//Let d1 be the 1st digit, d2 be the 2nd digit, and so on. In this way, we note the following:

// d2d3d4=406 is divisible by 2
// d3d4d5=063 is divisible by 3
// d4d5d6=635 is divisible by 5
// d5d6d7=357 is divisible by 7
// d6d7d8=572 is divisible by 11
// d7d8d9=728 is divisible by 13
// d8d9d10=289 is divisible by 17
// Find the sum of all 0 to 9 PanDigital numbers with this property.

func Problem043() int {

	max := 9
	panDigital := [][]int{}

	for a1 := 1; a1 <= max; a1++ {
		for a2 := 0; a2 <= max; a2++ {
			if a2 != a1 {
				for a3 := 0; a3 <= max; a3++ {
					if a3 != a1 && a3 != a2 {
						for a4 := 0; a4 <= max; a4++ {
							if a4 != a3 && a4 != a2 && a4 != a1 {
								if (a4 % 2 == 0) {
									for a5 := 0; a5 <= max; a5++ {
										if a5 != a4 && a5 != a3 && a5 != a2 && a5 != a1 {
											if (a3 + a4 + a5) % 3 == 0 {
												for a6 := 0; a6 <= max; a6++ {
													if a6 != a5 && a6 != a4 && a6 != a3 && a6 != a2 && a6 != a1 {
														if a6 == 5 || a6 == 0 {
															for a7 := 0; a7 <= max; a7++ {
																if a7 != a6 && a7 != a5 && a7 != a4 && a7 != a3 && a7 != a2 && a7 != a1 {
																	if dividesBy(a5, a6, a7, 7) {
																		for a8 := 0; a8 <= max; a8++ {
																			if a8 != a7 && a8 != a6 && a8 != a5 && a8 != a4 && a8 != a3 && a8 != a2 && a8 != a1 {
																				if dividesBy(a6, a7, a8, 11) {
																					for a9 := 0; a9 <= max; a9++ {
																						if a9 != a8 && a9 != a7 && a9 != a6 && a9 != a5 && a9 != a4 && a9 != a3 && a9 != a2 && a9 != a1 {
																							if dividesBy(a7, a8, a9, 13) {

																								for a10 := 0; a10 <= max; a10++ {
																									if a10 != a9 && a10 != a8 && a10 != a7 && a10 != a6 && a10 != a5 && a10 != a4 && a10 != a3 && a10 != a2 && a10 != a1 {
																										if dividesBy(a8, a9, a10, 17) {
																											var values = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
																											values[0] = a1
																											values[1] = a2
																											values[2] = a3
																											values[3] = a4
																											values[4] = a5
																											values[5] = a6
																											values[6] = a7
																											values[7] = a8
																											values[8] = a9
																											values[9] = a10
																											panDigital = append(panDigital, values)
																										}
																									}
																								}

																							}
																						}

																					}

																				}
																			}
																		}

																	}
																}
															}

														}
													}
												}

											}
										}
									}
								}
							}
						}
					}
				}
			}

		}
	}

	return sumAnswer(panDigital)
}

func dividesBy(a, b, c, target int) bool {
	val, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b) + strconv.Itoa(c))
	return val % target == 0
}

func sumAnswer(values [][]int) int {
	sum := 0
	for _, value := range values {
		asNum, _ := strconv.Atoi(strconv.Itoa(value[0]) + strconv.Itoa(value[1]) + strconv.Itoa(value[2]) + strconv.Itoa(value[3]) + strconv.Itoa(value[4]) + strconv.Itoa(value[5]) + strconv.Itoa(value[6]) + strconv.Itoa(value[7]) + strconv.Itoa(value[8]) + strconv.Itoa(value[9]))
		sum = sum + asNum
	}
	return sum
}

