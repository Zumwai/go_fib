package main

import (
	"fmt"
)

import "math/big"

func fib(n int64) *big.Int {
	if -2 <= n && n <= 2 {
		switch {
		case n > 0:
			return big.NewInt(1)
		case n < 0:
			return big.NewInt(-1)
		default:
			return big.NewInt(0)
		}
	}
	k := n / 2
	a := fib(k)

	var b *big.Int
	if k > 0 {
		b = fib(k + 1)
	} else {
		b = fib(k - 1)
	}

	if n%2 != 0 {
		return b.Add(a.Mul(a, a), b.Mul(b, b))
	}
	if n < 0 {
		return b.Mul(a, b.Add(b.Mul(b, big.NewInt(2)), a))
	}
	return b.Mul(a, b.Sub(b.Mul(b, big.NewInt(2)), a))
}

func main() {
	//or "how overflow terminal with digits"
	fmt.Println(fib(-8))
}
