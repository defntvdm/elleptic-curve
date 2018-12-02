package main

import (
	"fmt"
	"math/big"
)

func Mul(z, x, y *big.Int) *big.Int {
	var tmp *big.Int
	if y.Cmp(zero) == 0 || x.Cmp(zero) == 0 {
		z.Set(zero)
		return z
	}
	if x != z {
		z.Set(x)
		tmp = x
	} else {
		tmp = new(big.Int).Set(x)
	}
	for _, c := range y.Text(2)[1:] {
		z.Lsh(z, 1)
		if c == '1' {
			z.Xor(z, tmp)
		}
	}
	return Mod(z)
}

func Mod(x *big.Int) *big.Int {
	bitLen := x.BitLen()
	for bitLen >= polynomialBitLength {
		x.Xor(x, new(big.Int).Lsh(polynomial, uint(bitLen-polynomialBitLength)))
		bitLen = x.BitLen()
	}
	return x
}

func Inverse(x *big.Int) *big.Int {
	value := []*big.Int{x, polynomial}
	xCol := []*big.Int{one, zero}

	length := 2

	for value[length-1].Cmp(one) != 0 {
		v1len := value[length-1].BitLen()
		v2len := value[length-2].BitLen()
		if v1len == v2len {
			value = append(value, new(big.Int).Xor(value[length-1], value[length-2]))
			xCol = append(xCol, new(big.Int).Xor(xCol[length-1], xCol[length-2]))
			length++
		} else {
			if v1len < v2len {
				value = append(value, new(big.Int).Lsh(value[length-1], uint(v2len-v1len)))
				xCol = append(xCol, new(big.Int).Lsh(xCol[length-1], uint(v2len-v1len)))
				length++
				value[length-1].Xor(value[length-1], value[length-3])
				xCol[length-1].Xor(xCol[length-1], xCol[length-3])
			} else {
				value = append(value, new(big.Int).Lsh(value[length-2], uint(v1len-v2len)))
				xCol = append(xCol, new(big.Int).Lsh(xCol[length-2], uint(v1len-v2len)))
				length++
				value[length-1].Xor(value[length-1], value[length-2])
				xCol[length-1].Xor(xCol[length-1], xCol[length-2])
			}
		}
	}
	return Mod(xCol[length-1])
}

func formatCoordinate(coordinate *big.Int) string {
	return fmt.Sprintf(coordinateFmt, coordinate.Text(2))
}

func format2Point(p Point) string {
	if p.x == nil {
		return "E"
	}
	return fmt.Sprintf(
		"(%s, %s)",
		formatCoordinate(p.x),
		formatCoordinate(p.y),
	)
}
