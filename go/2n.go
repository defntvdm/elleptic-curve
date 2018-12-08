package main

import (
	"fmt"
	"math/big"
	"os"
	"strings"
)

// y^2*z + a*x*y* + x^3 + b*x^2*z + c*z^3 = 0 - несуперсингулярная кривая
func solveTask2N(task string) {
	if task == "" {
		return
	}

	tokens := strings.Split(task, " ")
	switch tokens[0] {
	case MUL:
		num, point := parseTokensForMul2(tokens[1:])
		result := mul2N(num, point)
		fmt.Fprintf(
			outputFile,
			"%s * %s = %s\n",
			format2Point(point),
			num.Text(int(baseType)),
			format2Point(result),
		)
	case SUM:
		p1, p2 := parseTokensForSum2(tokens[1:])
		result := sum2N(p1, p2)
		fmt.Fprintf(
			outputFile,
			"%s + %s = %s\n",
			format2Point(p1),
			format2Point(p2),
			format2Point(result),
		)
	default:
		fmt.Fprintln(
			os.Stderr,
			"Ожидалась операция сложения двух точек 'С' или умножения точки на число 'У'",
		)
		os.Exit(1)
	}
}

func mul2N(n *big.Int, point Point) Point {
	switch n.Cmp(zero) {
	case -1:
		negN := new(big.Int).Neg(n)
		newY := Mul(new(big.Int), a, point.x)
		negPoint := Point{
			x: point.x,
			y: newY.Xor(newY, point.y),
		}
		return mul2N(negN, negPoint)
	case 0:
		return fieldZero
	}
	result := point
	for _, bit := range n.Text(2)[1:] {
		result = sum2N(result, result)
		if bit == '1' {
			result = sum2N(result, point)
		}
	}
	return result
}

func sum2N(p1, p2 Point) Point {
	if p1.x == nil {
		return p2
	}
	if p2.x == nil {
		return p1
	}
	var k *big.Int = new(big.Int)
	var x3, y3 *big.Int
	if p1.x.Cmp(p2.x) != 0 {
		subX := new(big.Int).Xor(p2.x, p1.x)
		subY := new(big.Int).Xor(p2.y, p1.y)

		subX = Inverse(subX)

		Mul(k, subX, subY)

		x3 = Mul(new(big.Int), k, k)

		x3.Xor(x3, p1.x).Xor(x3, p2.x).Xor(x3, Mul(new(big.Int), a, k)).Xor(x3, b)
	} else {
		if p1.y.Cmp(p2.y) != 0 || p1.x.Cmp(zero) == 0 {
			return fieldZero
		}

		numerator := Mul(new(big.Int), p1.x, p1.x)
		denominator := Mul(new(big.Int), p1.x, a)

		numerator.Xor(numerator, Mul(new(big.Int), a, p1.y))
		denominator = Inverse(denominator)

		Mul(k, numerator, denominator)

		x3 = Mul(new(big.Int), k, k)

		x3.Xor(x3, b).Xor(x3, Mul(new(big.Int), a, k))
	}

	y3 = new(big.Int).Set(p1.y)
	y3.Xor(y3, Mul(new(big.Int), k, new(big.Int).Xor(x3, p1.x)))

	y3.Xor(y3, Mul(new(big.Int), a, x3))

	return Point{x3, y3}
}
