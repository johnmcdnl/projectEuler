package problems

import (
	"os"
	"sort"
	"strings"
)

//Names scores

/**

https://projecteuler.net/project/resources/p022_names.txt

Using names.txt (right click and 'Save Link/Target As...'), a 46K text file containing over five-thousand first names,
begin by sorting it into alphabetical order. Then working out the alphabetical value for each name, multiply this value
by its alphabetical position in the list to obtain a name score.

For example, when the list is sorted into alphabetical order, COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53,
is the 938th name in the list. So, COLIN would obtain a score of 938 × 53 = 49714.

What is the total of all the name scores in the file?
*/

func Problem022() int {
	var sum = 0
	for index, name := range sortedDataset() {
		nameVal := nameValue(name)
		sum += (index + 1) * nameVal
	}
	return sum
}

func nameValue(name string) int {
	nameRune := []rune(name)

	var value = 0
	for _, r := range nameRune {
		value += int(r) - 64
	}
	return value
}

func sortedDataset() []string {
	data := dataset()
	sort.Slice(data, func(i, j int) bool {
		return data[i] < data[j]
	})
	return data
}

func dataset() []string {
	data, err := os.ReadFile("./data/p022_names.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.ReplaceAll(string(data), `"`, ``), ",")
}
