package main

import (
	"cad/prime/prime"
	"cad/prime/server"
)

//IsPrime is a wrapper to handle CustomPrimalyTest as an external library through FFI
func IsPrime(a string, sample float32, umbral float32) (bool, float32, float32) {
	return prime.CustomPrimalyTest(a, sample, umbral)
}

func main() {

	s := server.Server{Port: ":8000"}
	s.NewHTTP()
}
