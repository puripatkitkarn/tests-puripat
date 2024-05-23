package main

import "fmt"

func FizzBuzzLevl4(n int) string {
	result := map[bool]string{true: "Fizz", false: ""}[n%3 == 0]
	result += map[bool]string{true: "Buzz", false: ""}[n%5 == 0]
	return map[bool]string{true: result, false: fmt.Sprint(n)}[result != ""]
}
