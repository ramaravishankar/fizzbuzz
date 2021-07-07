package fizzbuzz_test

import (
	"fmt"
	"testing"

	"github.com/ramaravishankar/fizzbuzz"
)

func TestFizz(t *testing.T) {
	_, ok := fizzbuzz.FizzBuzz(1)
	if ok {
		t.Logf("The value %v should not have returned true", 1)
		t.Fail()
	}
	result, ok := fizzbuzz.FizzBuzz(3)
	if !ok {
		t.Logf("The value %v should have returned true", 3)
		t.Fail()
	}
	if result != "fizz" {
		t.Log("Result should have been fizz")
		t.Fail()
	}
	result, ok = fizzbuzz.FizzBuzz(5)
	if !ok {
		t.Logf("The value %v should have returned true", 5)
		t.Fail()
	}
	if result != "buzz" {
		t.Log("Result should have been buzz")
		t.Fail()
	}
	result, ok = fizzbuzz.FizzBuzz(15)
	if !ok {
		t.Logf("The value %v should have returned true", 15)
		t.Fail()
	}
	if result != "fizz-buzz" {
		t.Log("Result should have been fizz-buzz")
		t.Fail()
	}
}

func ExampleFizzBuzz() {
	result, _ := fizzbuzz.FizzBuzz(3)
	fmt.Println(result)
	// Output: fizz
}
