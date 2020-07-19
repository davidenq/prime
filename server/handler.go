package server

import (
	math "cad/prime/prime"
	"fmt"
	"net/http"
	"strconv"
)

//CheckIfIsPrime .
func CheckIfIsPrime(w http.ResponseWriter, r *http.Request) {

	prime := r.URL.Query().Get("prime")
	sample := r.URL.Query().Get("samples")
	umbral := r.URL.Query().Get("umbral")

	if prime == "" {
		Reply(w, http.StatusBadRequest, "A number to test is required")
		return
	}
	if sample == "" {
		Reply(w, http.StatusBadRequest, "A samples number is required to test as much times as you specified and reduce the margin of error.")
		return
	}
	if umbral == "" {
		Reply(w, http.StatusBadRequest, "Umbral of error between 0 and 100 is required to consider if the number pass or not as a prime.")
		return
	}

	sample64, _ := strconv.ParseFloat(sample, 32)
	sample32 := float32(sample64)

	umbral64, _ := strconv.ParseFloat(umbral, 32)
	umbral32 := float32(umbral64)

	ok, trustedSamples, marginError := math.CustomPrimalyTest(prime, sample32, umbral32)
	if ok {
		message := fmt.Sprintf("%s is a prime number. Pass with %.0f samples. The margin error is %.2f", prime, trustedSamples, marginError)
		Reply(w, http.StatusOK, message)
		return
	}
	message := fmt.Sprintf("%s is not a prime number. Pass with %.0f samples. The margin error is %.2f", prime, trustedSamples, marginError)
	Reply(w, http.StatusOK, message)
	return

}

//GenerateLargePrime .
func GenerateLargePrime(w http.ResponseWriter, r *http.Request) {
}
