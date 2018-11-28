package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func parseTokensForMulP(tokens []string) (*big.Int, Point) {
	point := Point{}
	point.x = parseBigNum(tokens[0], false)
	point.y = parseBigNum(tokens[1], false)

	return parseBigNum(tokens[2], false), point
}

func parseTokensForSumP(tokens []string) (Point, Point) {
	p1 := Point{}
	p1.x = parseBigNum(tokens[0], false)
	p1.y = parseBigNum(tokens[1], false)

	p2 := Point{}
	p2.x = parseBigNum(tokens[2], false)
	p2.y = parseBigNum(tokens[3], false)

	return p1, p2
}

func parseTokensForMul2(tokens []string) (*big.Int, Point) {
	point := Point{}
	point.x = parseBigNum(tokens[0], true)
	point.y = parseBigNum(tokens[1], true)

	return parseBigNum(tokens[2], false), point
}

func parseTokensForSum2(tokens []string) (Point, Point) {
	p1 := Point{}
	p1.x = parseBigNum(tokens[0], true)
	p1.y = parseBigNum(tokens[1], true)

	p2 := Point{}
	p2.x = parseBigNum(tokens[2], true)
	p2.y = parseBigNum(tokens[3], true)

	return p1, p2
}

func parseBigNum(num string, for2Point bool) *big.Int {
	baseTypeTmp := baseType
	var ok bool
	var res *big.Int
	n := new(big.Int)
	if for2Point {
		baseTypeTmp = 2
	}
	switch baseTypeTmp {
	case binType:
		res, ok = n.SetString(num, binType)
		if ok {
			return res
		}
		fmt.Fprintf(
			os.Stderr,
			"Не получилось распарсить число %q\nОжидалось число в двоичной системе счисления\n",
			num,
		)
		os.Exit(1)
	case hexType:
		res, ok = n.SetString(num, hexType)
		if ok {
			return res
		}
		fmt.Fprintf(
			os.Stderr,
			"Не получилось распарсить число %q\nОжидалось число в шестнадцатеричной системе счисления\n",
			num,
		)
		os.Exit(1)
	case decType:
		res, ok = n.SetString(num, decType)
		if ok {
			return res
		}
		fmt.Fprintf(
			os.Stderr,
			"Не получилось распарсить число %q\nОжидалось число в десятичной системе счисления\n",
			num,
		)
		os.Exit(1)
	}
	return nil
}

func parsePolynomial(poly string) *big.Int {
	result := new(big.Int)
	members := strings.Split(poly, "+")
	for _, member := range members {
		member := strings.Trim(member, " ")
		if strings.HasPrefix(member, "x^") {
			power := parseUint(member[2:])
			result.Xor(result, new(big.Int).Lsh(one, power))
		} else {
			if member == "x" {
				result.Xor(result, two)
			} else {
				number := parseUint(member)
				result.Xor(result, big.NewInt(int64(number%2)))
			}
		}
	}
	return result
}

func parseUint(num string) uint {
	switch baseType {
	case binType:
		result, err := strconv.ParseUint(num, binType, 0)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Степень члена должна быть в двоичной записи, получена %q\n", num)
			os.Exit(1)
		}
		return uint(result)
	case decType:
		result, err := strconv.ParseUint(num, decType, 0)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Степень члена должна быть в десятичной записи, получена %q", num)
			os.Exit(1)
		}
		return uint(result)
	case hexType:
		result, err := strconv.ParseUint(num, hexType, 0)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Степень члена должна быть в шетнадцатиричной записи, получена %q", num)
			os.Exit(1)
		}
		return uint(result)
	}
	return 0
}

func parseAB(ab string) (*big.Int, *big.Int) {
	nums := strings.Split(ab, " ")
	a := parseBigNum(nums[0], false)
	b := parseBigNum(nums[1], false)
	return a, b
}

func parseABC(abc string) (*big.Int, *big.Int, *big.Int) {
	nums := strings.Split(abc, " ")
	a := parseBigNum(nums[0], true)
	b := parseBigNum(nums[1], true)
	c := parseBigNum(nums[2], true)
	return a, b, c
}
