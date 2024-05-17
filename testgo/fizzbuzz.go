package main

import (
	"fmt"
	"strconv"
	"strings"
)

func FizzBuzz(n int) string {
	if n%15 == 0 {
		return "FizzBuzz"
	} else if n%5 == 0 {
		return "Buzz"
	} else if n%3 == 0 {
		return "Fizz"
	} else {
		return strconv.Itoa(n)
	}
}

func FizzBuzzLevl2(n int) string {
	Fizz := n%3 == 0
	Buzz := n%5 == 0
	FizzBuzz := Fizz && Buzz
	number := Fizz || Buzz || FizzBuzz

	result := fmt.Sprintf("%s%s%s",
		ternary(Fizz, "Fizz", ""),
		ternary(Buzz, "Buzz", ""),
		ternary(number, "", strconv.Itoa(n)),
	)
	result = strings.ReplaceAll(result, " ", "")

	return result
}

func ternary(condition bool, trueVal, falseVal string) string {
	if condition {
		return trueVal
	}
	return falseVal
}
