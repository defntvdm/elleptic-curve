package main

import (
	"fmt"
	"math/big"
	"os"
	"strings"
)

// y^2*z + a*y*z^2 + x^3 + b*x*z^2 + c*z^3 = 0 - суперсингулярная кривая
func solveTask2S(task string) {
	if task == "" {
		return
	}

	tokens := strings.Split(task, " ")
	switch tokens[0] {
	case MUL:
		num, point := parseTokensForMul2(tokens[1:])
		result := mul2S(num, point)
		fmt.Fprintf(
			outputFile,
			"%s * %s = %s\n",
			format2Point(point),
			num.Text(int(baseType)),
			format2Point(result),
		)
	case SUM:
		p1, p2 := parseTokensForSum2(tokens[1:])
		result := sum2S(p1, p2)
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

func mul2S(n *big.Int, point Point) Point {
	switch n.Cmp(zero) {
	case -1:
		negN := new(big.Int).Neg(n)
		negPoint := Point{
			x: point.x,
			y: new(big.Int).Xor(point.y, a),
		}
		return mul2S(negN, negPoint)
	case 0:
		return fieldZero
	}
	result := point
	for _, b := range n.Text(2)[1:] {
		result = sum2S(result, result)
		if b == '1' {
			result = sum2S(result, point)
		}
	}
	return result
}

func sum2S(p1, p2 Point) Point {
	var k *big.Int = new(big.Int)
	var x3, y3 *big.Int
	if p1.x.Cmp(p2.x) != 0 {
		subX := new(big.Int).Xor(p2.x, p1.x)
		subY := new(big.Int).Xor(p2.y, p1.y)

		subX = Inverse(subX)

		Mul(k, subX, subY)
		x3 = Mul(new(big.Int), k, k)
		x3.Xor(x3, p1.x).Xor(x3, p2.x)
	} else {
		if p1.y.Cmp(p2.y) != 0 {
			return fieldZero
		}
		numerator := Mul(new(big.Int), p1.x, p1.x)
		numerator.Xor(numerator, b)
		denominator := Inverse(a)

		k := Mul(new(big.Int), numerator, denominator)

		x3 = Mul(new(big.Int), k, k)
	}
	y3 = new(big.Int).Xor(p1.x, x3)
	Mul(y3, y3, k).Xor(y3, p1.y).Xor(y3, a)
	return Point{x3, y3}
}
