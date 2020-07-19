package prime

import "math/big"

//bigNumber is used to generate very large numbers from string
func bigNumber(n string) *big.Float {
	bf := new(big.Float)
	number, _ := bf.SetString(n)
	//number.SetPrec(256)
	return number
}
