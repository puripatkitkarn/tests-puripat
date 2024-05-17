package main

import "testing"

func TestFizzBuzzShouldReturn1WhenInput1(t *testing.T) {
	input := 1

	got := FizzBuzzLevl2(input)

	want := "1"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func TestFizzBuzzShouldReturn2WhenInput2(t *testing.T) {
	input := 2

	got := FizzBuzzLevl2(input)

	want := "2"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func TestFizzBuzzShouldReturnFizzWhenInput3(t *testing.T) {
	input := 3

	got := FizzBuzzLevl2(input)

	want := "Fizz"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func TestFizzBuzzShouldReturn4WhenInput4(t *testing.T) {
	input := 4

	got := FizzBuzzLevl2(input)

	want := "4"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func TestFizzBuzzShouldReturnBuzzWhenInput5(t *testing.T) {
	input := 5

	got := FizzBuzzLevl2(input)

	want := "Buzz"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func TestFizzBuzzShouldReturnFizzWhenInput6(t *testing.T) {
	input := 6

	got := FizzBuzzLevl2(input)

	want := "Fizz"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func TestFizzBuzzShouldReturn7WhenInput7(t *testing.T) {
	input := 7

	got := FizzBuzzLevl2(input)

	want := "7"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func TestFizzBuzzShouldReturn8WhenInput8(t *testing.T) {
	input := 8

	got := FizzBuzzLevl2(input)

	want := "8"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
