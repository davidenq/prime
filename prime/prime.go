package prime

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

//primes contains the first 4 primes numbers
//it's used to use with Custom Fertmat's Little Test
var primes = []uint{
	2, 3, 5, 7, 11, 13, 17, 19, 23, 29,
}

//CustomPrimalyTest based on Fermat Primaly Test
//p is a pseudo prime number to be proved
//s is the number of samples to test the pseudprime number p.
//If s is a big number the margin of error will be less.
func CustomPrimalyTest(p string, samples float32, umbral float32) (bool, float32, float32) {

	/*
		Fermat Little Theorem says:
		=> p divides b^(p - 1) in another words b^(prime-1) congruent 1 (mod p)
		Where:
		p is a pseudoprime, which must be proved.
		b is a random number between 0 < b < p - 1

		But since that p could be a very large number to test, instead of using
		the exponencial as a value to be proved, why not change the logic about
		b and p. It means b could be the pseudoprime to be proved and p a prime
		number that it already had been proved before.
		So, having the above idea in mind, considere:
		p  a prime number that it already had been proved before
		b  a pseudo-prome to be proved
		q  quotient of b^p/p
		r  remainder of b^p = q mod prime
		p < b

		we can write the equation as bellow:
			     b^(p-1)
			   ---------- q, r  [dividend/divisor = q, r]
					    p

		if r is 0, then p is probably a prime number, otherwise is a compositive number.

		This consideration will be a conjecture to be proved.

		In order to reduce de error margin we will prove the pseudoprime with s samples specified
		by the user and also reinforce with trial_division test to asure initialize that the numbers
		doesn't has more than 2 divisible elementes.
	*/
	var trustedValues float32
	var operations int
	var limit int

	firstTest := trailDivision(p, int(samples))

	if firstTest {
		return false, 0, 100
	}

	if len(p) < 100 {
		limit = 2
	} else {
		limit = 10
	}

	//s tests
	for {
		if float32(operations) < samples {
			//custom fermat primaly test
			rand.Seed(time.Now().UnixNano())

			prime := primes[rand.Intn(limit)]
			result := conjectureFermatPrimalyTest(p, prime)
			if result {
				trustedValues++
			}
		} else {
			break
		}
		operations++
	}

	marginError := (trustedValues / samples) * 100

	return umbral < marginError, trustedValues, marginError / umbral
}

func conjectureFermatPrimalyTest(pseudoPrime string, prime uint) bool {
	primeMinusOne := math().Sub(fmt.Sprint(prime), "1")         //p-1
	dividend := math().Pow(pseudoPrime, primeMinusOne.String()) //b^(p-1)
	_, remainder := math().Div(dividend.String(), pseudoPrime)  //[b^(p-1)]/p
	if remainder.String() == "0" {
		return true
	}
	return false
}

func trailDivision(p string, samples int) bool {

	divisibledElements := 0
	count := 1
	pInteger, _ := strconv.Atoi(p)
	if samples > 100 {
		samples = 100
	}
	for {
		if count == samples || count == pInteger {
			break
		}
		_, remainder := math().Div(p, strconv.Itoa(count))
		if remainder.String() == "0" {
			divisibledElements++
		}
		count++
	}

	if divisibledElements > 2 {
		return true
	}
	return false
}

/*
====================================
Incomplete Miller Rabin Primality Test  (implementation in progress)
======================================
*/

//MillerRabinPrimalityTest @p is a pseudo-prime number that will be prove if probably is a prime number.
func MillerRabinPrimalityTest(p string) bool {
	//Miller-Rabin primality test has 3 steps
	//1. p - 1 = m * a^k where
	//p is a pseudo-prime number
	//k and m are whole numbers
	//a is a random between 1 < a < n -1. Generally a = 2 for a better performance.

	//first calculating k and m
	// p - 1
	k := 1
	base := 2
	m := "0"

	for {

		primeMinusOne := math().Sub(p, "1")
		quotient, remainder := math().Div(primeMinusOne.String(), strconv.Itoa(base))

		if remainder.String() != "0" {
			k = k - 1
			break
		}
		base = base * 2
		k++
		m = quotient.String()
	}
	if m == "0" {
		m = p
	}
	fmt.Println("p=", p, "k=", k, "m=", m)
	return true

}
