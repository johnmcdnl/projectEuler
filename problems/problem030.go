package problems

import (
	"github.com/johnmcdnl/projectEuler/utils"
)

func Problem030() int {
	powers := 5
	max := powers
	for i := 1; i <= powers; i++ {
		max *= 9
	}
	numMap := generateSets(max)
	validAnswers := []int{}
	for i:=2; i<max; i++ {
		powSlice := powOfSlice(numMap[i], powers)
		if powSlice == i {
			validAnswers = append(validAnswers, i)
		}
	}

	return utils.SumSlice(validAnswers)
}

func powOfSlice(array []int, power int) int {

	sumPowers := 0
	for i := 0; i < len(array); i++ {
		powI := 1
		for j := 0; j < power; j++ {
			powI = powI * array[i]
		}

		sumPowers += powI
	}

	return sumPowers
}

func generateSets(max int) [][]int {

	index := 0

	numMap := make([][]int, max)
	outLoop:
	for a := 0; a <= 9; a++ {
		for b := 0; b <= 9; b++ {
			for c := 0; c <= 9; c++ {
				for d := 0; d <= 9; d++ {
					for e := 0; e <= 9; e++ {
						for f := 0; f <= 9; f++ {
							if a != 0 {
								numMap[index] = []int{a, b, c, d, e, f}
							} else {
								if b != 0 {
									numMap[index] = []int{b, c, d, e, f}
								} else {
									if c != 0 {
										numMap[index] = []int{c, d, e, f}
									} else {
										if d != 0 {
											numMap[index] = []int{d, e, f}
										} else {
											if e != 0 {
												numMap[index] = []int{e, f}
											} else {
												numMap[index] = []int{f}
											}
										}
									}

								}
							}

							index++
							if index >= max {
								break outLoop
							}
						}
					}
				}
			}
		}

	}
	return numMap
}