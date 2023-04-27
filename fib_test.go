package main

import "testing"
import "math/big"

type input struct {
	arg      int64
	expected *big.Int
}

var thousand = new(big.Int)
var bigminus = new(big.Int)

func init() {
	thousand, _ = thousand.SetString("43466557686937456435688527675040625802564660517371780402481729089536555417949051890403879840079255169295922593080322634775209689623239873322471161642996440906533187938298969649928516003704476137795166849228875", 10)
	bigminus, _ = bigminus.SetString("-135301852344706746049", 10)
}

var toTest = []input{
	{0, big.NewInt(0)},
	{1, big.NewInt(1)},
	{5, big.NewInt(5)},
	{12, big.NewInt(144)},
	{1000, thousand},
	{-1, big.NewInt(-1)},
	{-2, big.NewInt(-1)},
	{-19, big.NewInt(4181)},
	{-40, big.NewInt(-102334155)},
	{-43, big.NewInt(433494437)},
	{-75, big.NewInt(2111485077978050)},
	{-98, bigminus},
}

func TestFib(t *testing.T) {
	for _, inp := range toTest {
		out := fib(inp.arg)
		res := out.Cmp(inp.expected)

		if res != 0 {
			t.Errorf("Output %q is not equal to expected %q with %d as input", out, inp.expected, inp.arg)
		}
	}
}

func BenchmarkFibdoubling(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fib(1000)
	}
}

func BenchmarkFibonaccibinary(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fib(1000)
	}
}
