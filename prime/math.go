package prime

import (
	"math/big"
	"strings"
)

//Operation .
type Operation struct {
	result *big.Float
}

//math generate a instance of operation that contains a pointer to big.Float
func math() *Operation {

	result := bigNumber("0")
	operation := &Operation{
		result: result,
	}
	return operation
}

//Add .
func (op *Operation) Add(a, b string) *big.Float {
	return op.result.Add(bigNumber(a), bigNumber(b))
}

//Sub .
func (op *Operation) Sub(a, b string) *big.Float {
	return op.result.Sub(bigNumber(a), bigNumber(b))
}

//Mul .
func (op *Operation) Mul(a, b string) *big.Float {
	return op.result.Mul(bigNumber(a), bigNumber(b))
}

//Div .
func (op *Operation) Div(a, b string) (*big.Float, *big.Float) {

	//remainder = dividend - (divisor*quotation)
	// r = a - (b*q)

	quotient := op.result.Quo(bigNumber(a), bigNumber(b))
	remainder := bigNumber("0")
	dividend := bigNumber(a)

	//given that quotient could have an integral part and decimal part
	//it's important to get just the integral part to calculate the remainder.
	//For that reason we could get this value using splitting over quotient
	//result given that it's a string.
	intDec := strings.Split(quotient.String(), ".")
	integralQuo := bigNumber(intDec[0])
	mult := op.Mul(b, integralQuo.String())
	remainder = remainder.Sub(dividend, mult)
	return quotient, remainder
}

//Pow .
func (op *Operation) Pow(b string, e string) *big.Float {

	exponent := bigNumber(e)
	base := bigNumber(b)
	result := bigNumber("1")
	operations := bigNumber("1")
	for {
		result = op.Mul(base.String(), result.String())
		if operations.String() == exponent.String() {
			break
		}
		operations = operations.Add(operations, bigNumber("1"))
	}
	return result
}
