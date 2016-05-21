package problems


//Number letter counts

/*
If the numbers 1 to 5 are written out in words:
one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.

If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?

NOTE: Do not count spaces or hyphens.
For example, 342 (three hundred and forty-two) contains 23 letters and 115 (one hundred and fifteen) contains 20 letters.
The use of "and" when writing out numbers is in compliance with British usage.
 */
func Problem017() int {
	oneToHundred := []string{}
	oneToThousand := []string{}
	singles := []string{"", "one", "two", "three", "four", "five", "six", "seven", "eigth", "nine"}
	tenTwenty := []string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	tens := []string{"", "ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eigthy", "ninety"}

	for i := 1; i < 10; i++ {
		oneToHundred = append(oneToHundred, singles[i])
	}

	for i := 10; i < 20; i++ {
		oneToHundred = append(oneToHundred, tenTwenty[i - 10])
	}

	for i := 20; i < 100; i += 10 {
		oneToHundred = append(oneToHundred, tens[i / 10])
		for j := 1; j < 10; j++ {
			oneToHundred = append(oneToHundred, tens[i / 10] + singles[j])
		}
	}

	oneToThousand = oneToHundred
	for i := 1; i < 10; i++ {
		oneToThousand = append(oneToThousand, singles[i] + "hundred")
		asString := singles[i] + "hundredand"
		for _, val := range oneToHundred {
			oneToThousand = append(oneToThousand, asString + val)
		}
	}
	oneToThousand = append(oneToThousand, "onethousand")
	asString := ""
	for _, value := range oneToThousand {
		asString = asString + value
	}

	return len(asString)
}
