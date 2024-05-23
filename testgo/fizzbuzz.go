package main

import "fmt"

type FizzBuzz struct {
	number int
}

func CreateFizzBuzz(n int) *FizzBuzz {
	return &FizzBuzz{number: n}
}

func (fb *FizzBuzz) Generate() string {
	switch {
	case fb.number%15 == 0:
		return "FizzBuzz"
	case fb.number%3 == 0:
		return "Fizz"
	case fb.number%5 == 0:
		return "Buzz"
	default:
		return fmt.Sprintf("%d", fb.number)
	}
}

func FizzBuzzLevl4(n int) string {
	fb := CreateFizzBuzz(n)
	result := fb.Generate()
	return result
}
