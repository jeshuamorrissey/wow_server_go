package srp

import "math/big"

const (
	// G is the SRP Generator; the base of many mathematical expressions.
	G = 7

	// K is the SRP Verifier Scale Factor; used to scale the verifier which
	// is stored in the database.
	K = 3
)

// N is the SRP Modulus; all operations are performed in base N.
func N() *big.Int {
	n := big.NewInt(0)
	n.SetString("62100066509156017342069496140902949863249758336000796928566441170293728648119", 10)
	return n
}
