package prime

import "errors"

// Nth returns the nth prime number
func Nth(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("Invalid input: given number must be >= 1.")
	}

	resultNum := 1
	for n > 0 {
		resultNum++
		
		// Check if resultNum is prime
		if isPrime(resultNum) {
			n--
		}
	}

	return resultNum, nil
}

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	if num == 2 {
		return true
	}
	if num%2 == 0 {
		return false
	}
	
	// Check divisibility by odd numbers up to sqrt(num)
	for i := 3; i*i <= num; i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}