package problems

import (
	"io/ioutil"
	"strings"
	"fmt"
)
//Code Triangle Numbers

/*
The nth term of the sequence of triangle numbers is given by, tn = ½n(n+1);
so the first ten triangle numbers are:

1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...

By converting each letter in a word to a number corresponding to
its alphabetical position and adding these values we form a word value.
For example, the word value for SKY is 19 + 11 + 25 = 55 = t10.
If the word value is a triangle number then we shall call the word a triangle word.

Using words.txt (right click and 'Save Link/Target As...'),
a 16K text file containing nearly two-thousand common English words, how many are triangle words?
 */

func init() {
	generateAlphabet()
	fmt.Println()
	fmt.Println()
	fmt.Println(getWordValue("SKY"))
	fmt.Println()
	fmt.Println()
}

var alphabetMap map[byte]int

func generateAlphabet() {
	alphabet := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	alphabetMap = make(map[byte]int)
	for index, value := range alphabet {
		alphabetMap[value] = index + 1
	}

}

func Problem042() int {
	words := readFile()
	triangles := getTriangleNumbers(longestWordInSlice(words))
	triangleNumbers := 0
	for _, word := range words {
		wordValue := getWordValue(word)
		if _, ok := triangles[wordValue]; ok {
			triangleNumbers++
		}
	}
	return triangleNumbers
}

func getWordValue(word string) int {
	asBytes := []byte(word)
	sum := 0
	for _, val := range asBytes {
		sum += alphabetMap[val]
	}
	return sum
}

func readFile() []string {
	words, _ := ioutil.ReadFile("./problems/problem042_words.txt")
	wordsStr := string(words)
	wordsStr = strings.Replace(wordsStr, `"`, "", -1)
	return strings.SplitAfter(wordsStr, ",")
}

func longestWordInSlice(words []string) int {
	longest := 0
	for _, word := range words {
		length := len([]byte(word))
		if length > longest {
			longest = length
		}
	}
	return longest
}

func getTriangleNumbers(n int) map[int]int {
	max := n * 26
	triangles := make(map[int]int)

	sum := 0
	i := 1
	for sum < max {
		sum += i
		i++
		triangles[sum] = sum
	}

	return triangles

}
