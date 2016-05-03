package utils

func RemoveDuplicateInts(elements []int) []int {
	found := map[int]bool{}
	result := []int{}

	for e := range elements {
		if found[elements[e]] == false {
			found[elements[e]] = true
			result = append(result, elements[e])
		}
	}
	return result
}
