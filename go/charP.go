package main

import (
	"fmt"
	"math/big"
	"os"
	"strings"
)

type Point struct {
	x *big.Int
	y *big.Int
}

// x^3 + a*x*z^2 + b*z^3 - y^2*z = 0
func solveTaskP(task string) {
	if task == "" {
		return
	}

	tokens := strings.Split(task, " ")
	switch tokens[0] {
	case MUL:
		num, point := parseTokensForMulP(tokens[1:])
		result := mulP(num, point)
		fmt.Fprintf(
			outputFile,
			"%s * %s = %s\n",
			formatPPoint(point),
			num.Text(int(baseType)),
			formatPPoint(result),
		)
	case SUM:
		p1, p2 := parseTokensForSumP(tokens[1:])
		result := sumP(p1, p2)
		fmt.Fprintf(
			outputFile,
			"%s + %s = %s\n",
			formatPPoint(p1),
			formatPPoint(p2),
			formatPPoint(result),
		)
	default:
		fmt.Fprintln(
			os.Stderr,
			"Ожидалась операция сложения двух точек 'С' или умножения точки на число 'У'",
		)
		os.Exit(1)
	}
}

func mulP(n *big.Int, point Point) Point {
	switch n.Cmp(zero) {
	case -1:
		negN := new(big.Int).Neg(n)
		negPoint := Point{
			x: point.x,
			y: new(big.Int).Sub(P, point.y),
		}
		return mulP(negN, negPoint)
	case 0:
		return fieldZero
	}
	result := point
	for _, bit := range n.Text(2)[1:] {
		result = sumP(result, result)
		if bit == '1' {
			result = sumP(result, point)
		}
	}
	return result
}

func sumP(p1, p2 Point) Point {
	if p1.x == nil {
		return p2
	}
	if p2.x == nil {
		return p1
	}
	var k *big.Int = new(big.Int)
	if p1.x.Cmp(p2.x) != 0 {
		subX := new(big.Int).Sub(p2.x, p1.x)
		subY := new(big.Int).Sub(p2.y, p1.y)

		subX.ModInverse(subX, P)

		k.Mul(subX, subY)
	} else {
		if p1.y.Cmp(p2.y) != 0 || p1.y.Cmp(zero) == 0 {
			return fieldZero
		}
		numerator := new(big.Int).Exp(p1.x, two, P)
		denominator := new(big.Int).SetInt64(2)

		numerator.Mul(numerator, three).Add(numerator, aP)

		denominator.Mul(denominator, p1.y).ModInverse(denominator, P)

		k.Mul(numerator, denominator)
	}

	k.Mod(k, P)

	k2 := new(big.Int).Exp(k, two, P)

	d := new(big.Int).Mul(k, p1.x)
	d.Neg(d).Add(d, p1.y)

	x3 := new(big.Int).Sub(k2, p1.x)
	x3.Sub(x3, p2.x).Mod(x3, P)

	y3 := new(big.Int).Mul(k, x3)
	y3.Add(y3, d).Neg(y3).Mod(y3, P)

	return Point{x3, y3}
}

func formatPPoint(p Point) string {
	if p.x == nil {
		return "E"
	}
	return fmt.Sprintf("(%s, %s)", p.x.Text(int(baseType)), p.y.Text(int(baseType)))
}
