package srp_test

import (
	"math/big"
	"testing"

	"github.com/jeshuamorrissey/wow_server_go/authserver/srp"
	"gotest.tools/assert"
)

func TestGenerateEphemeralPair(t *testing.T) {
	var v, expectedb, expectedB big.Int

	v.SetString("37510828772889775988011936555774753323884663064643682200527651910267083044538", 10)
	expectedb.SetString("3679141816495610969398422835318306156547245306", 10)
	expectedB.SetString("16630279820182697578309394812726193457375869535456855997552735653810818403718", 10)

	b, B := srp.GenerateEphemeralPair(&v)
	assert.Equal(t, b.Cmp(&expectedb), 0)
	assert.Equal(t, B.Cmp(&expectedB), 0)
}

func TestGenerateVerifier(t *testing.T) {
	var s, expectedV big.Int

	s.SetString("66759882342950727220130969932663635787137805713109467932708165413389947953699", 10)
	expectedV.SetString("37510828772889775988011936555774753323884663064643682200527651910267083044538", 10)

	v := srp.GenerateVerifier("JESHUA", "JESHUA", &s)
	assert.Equal(t, v.Cmp(&expectedV), 0)
}

func TestCalculateSessionKey(t *testing.T) {
	var A, B, b, v, s, expectedK, expectedM big.Int

	A.SetString("1234344069974946706941181551060269688256096998192437644043961152849307948728", 10)
	B.SetString("16630279820182697578309394812726193457375869535456855997552735653810818403718", 10)
	b.SetString("3679141816495610969398422835318306156547245306", 10)
	v.SetString("37510828772889775988011936555774753323884663064643682200527651910267083044538", 10)
	s.SetString("66759882342950727220130969932663635787137805713109467932708165413389947953699", 10)

	expectedK.SetString("1223778727786691224255566132121120158338166041153346746306820190174949498228440143950889596323712", 10)
	expectedM.SetString("1278405643266187066239549723718271591736372958987", 10)

	K, M := srp.CalculateSessionKey(&A, &B, &b, &v, &s, "JESHUA")

	assert.Equal(t, K.Cmp(&expectedK), 0)
	assert.Equal(t, M.Cmp(&expectedM), 0)
}

func TestCalculateServerProof(t *testing.T) {
	var A, M, K, expectedProof big.Int

	A.SetString("1234344069974946706941181551060269688256096998192437644043961152849307948728", 10)
	M.SetString("1278405643266187066239549723718271591736372958987", 10)
	K.SetString("1223778727786691224255566132121120158338166041153346746306820190174949498228440143950889596323712", 10)
	expectedProof.SetString("1284245613498486112994244042115912960631626548879", 10)

	proof := srp.CalculateServerProof(&A, &M, &K)
	assert.Equal(t, proof.Cmp(&expectedProof), 0)
}
