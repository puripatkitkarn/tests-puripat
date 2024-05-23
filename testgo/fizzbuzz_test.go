package main

import "testing"

func TestFizzBuzzShouldReturn1WhenInput1(t *testing.T) {
	input := 1

	got := FizzBuzzLevl4(input)

	want := "1"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func TestFizzBuzzShouldReturn2WhenInput2(t *testing.T) {
	input := 2

	got := FizzBuzzLevl4(input)

	want := "2"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func TestFizzBuzzShouldReturnFizzWhenInput3(t *testing.T) {
	input := 3

	got := FizzBuzzLevl4(input)

	want := "Fizz"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func TestFizzBuzzShouldReturn4WhenInput4(t *testing.T) {
	input := 4

	got := FizzBuzzLevl4(input)

	want := "4"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func TestFizzBuzzShouldReturnBuzzWhenInput5(t *testing.T) {
	input := 5

	got := FizzBuzzLevl4(input)

	want := "Buzz"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func TestFizzBuzzShouldReturnFizzWhenInput6(t *testing.T) {
	input := 6

	got := FizzBuzzLevl4(input)

	want := "Fizz"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func TestFizzBuzzShouldReturn7WhenInput7(t *testing.T) {
	input := 7

	got := FizzBuzzLevl4(input)

	want := "7"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func TestFizzBuzzShouldReturn8WhenInput8(t *testing.T) {
	input := 8

	got := FizzBuzzLevl4(input)

	want := "8"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func TestFizzBuzzShouldReturnFizzWhenInput9(t *testing.T) {
	input := 9

	got := FizzBuzzLevl4(input)

	want := "Fizz"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

// func TestFizzBuzzShouldReturnBuzzWhenInput10(t *testing.T) {
// 	input := 10

// 	got := FizzBuzzLevl3(input)

// 	want := "Buzz"
// 	if got != want {
// 		t.Errorf("got %q but want %q", got, want)
// 	}
// }

// func TestFizzBuzzShouldReturn11WhenInput11(t *testing.T) {
// 	input := 11

// 	got := FizzBuzzLevl3(input)

// 	want := "11"
// 	if got != want {
// 		t.Errorf("got %q but want %q", got, want)
// 	}
// }

// func TestFizzBuzzShouldReturnFizzWhenInput12(t *testing.T) {
// 	input := 12

// 	got := FizzBuzzLevl3(input)

// 	want := "Fizz"
// 	if got != want {
// 		t.Errorf("got %q but want %q", got, want)
// 	}
// }

// func TestFizzBuzzShouldReturn13WhenInput13(t *testing.T) {
// 	input := 13

// 	got := FizzBuzzLevl3(input)

// 	want := "13"
// 	if got != want {
// 		t.Errorf("got %q but want %q", got, want)
// 	}
// }

// func TestFizzBuzzShouldReturn14WhenInput14(t *testing.T) {
// 	input := 14

// 	got := FizzBuzzLevl3(input)

// 	want := "14"
// 	if got != want {
// 		t.Errorf("got %q but want %q", got, want)
// 	}
// }

// func TestFizzBuzzShouldReturnFizzBuzzWhenInput15(t *testing.T) {
// 	input := 156

// 	got := FizzBuzzLevl3(input)

// 	want := "FizzBuzz"
// 	if got != want {
// 		t.Errorf("got %q but want %q", got, want)
// 	}
// }
