package prime

import "errors"

var primes = []int{2}

func isPrime(n int) bool {
	for _, p := range primes {
        if n % p == 0 { return false }
    }
    return true
}

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n == 0 {
        return 0, errors.New("there is no zeroth prime")
    }

    if n == 1 {
        return 2, nil
    }

    for i := 3; len(primes) < n; i++ {
        if isPrime(i) == true {
            primes = append(primes, i)
        }
    }

    return primes[len(primes) - 1], nil
}
