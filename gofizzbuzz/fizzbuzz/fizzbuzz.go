package fizzbuzz

import (
	"fmt"
	"math/rand"
)

func Run() (num int) {
	var result string
	for {
		num = rand.Intn(100) + 1
		result = check(num)
		fmt.Printf("%d %s\n", num, result)
		if result == "FizzBuzz" {
			return num
		}
	}
}

func check(num int) string {
	if (num%3 == 0) && (num%5 == 0) {
		return "FizzBuzz"
	} else if num%3 == 0 {
		return "Fizz"
	} else if num%5 == 0 {
		return "Buzz"
	}
	return ""
}
