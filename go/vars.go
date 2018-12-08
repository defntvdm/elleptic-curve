package main

import (
	"math/big"
	"os"
)

// common
var (
	zero           = big.NewInt(0)
	one            = big.NewInt(1)
	two            = big.NewInt(2)
	three          = big.NewInt(3)
	fieldZero      = Point{nil, nil}
	baseType       uint
	inputFileName  string
	outputFileName string
	outputFile     *os.File
)

// For char(F) = 2
var (
	a, b, c             *big.Int
	polynomial          *big.Int
	polynomialBitLength int
	coordinateFmt       string
)

// For char(F) = P
var (
	aP, bP, P *big.Int
)
