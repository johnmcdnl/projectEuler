package main

import (
	"github.com/johnmcdnl/projectEuler/problems"
	"fmt"
	"github.com/johnmcdnl/projectEuler/utils"
)

func main() {
	fmt.Print(problems.Problem007())

	fmt.Println()
	fmt.Println()
	fmt.Println()

	fmt.Println(utils.IsPrime(uint64(18446744073709551557 )))

}