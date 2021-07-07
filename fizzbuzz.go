package fizzbuzz

// Fizzbuzz takes number return either string and boolean
// true or empty string with boolean false.
//
// Callers are supposed to print numbers if it returns false.
func FizzBuzz(num int) (string, bool) {
	if num%15 == 0 {
		return "fizz-buzz", true
	}
	if num%5 == 0 {
		return "buzz", true
	}
	if num%3 == 0 {
		return "fizz", true
	}
	return "", false
}
