package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func cancelPanic() {
	if err := recover(); err != nil {
		fmt.Fprintln(os.Stderr, "Maybe wrong input format, check for extra spaces")
	}
}

func init2(polyStr, coeffStr string) {
	polynomial = parsePolynomial(polyStr)
	polynomialBitLength = polynomial.BitLen()
	coordinateFmt = fmt.Sprintf("%%0%ds", polynomialBitLength-1)
	a, b, c = parseABC(coeffStr)
	Mod(a)
	Mod(b)
	Mod(c)
}

func main() {
	flag.UintVar(
		&baseType,
		"base",
		0,
		"Система счисления в которой будут подаваться всегда числа, для char K = 2 координаты точки всегда считываются в двоичном виде",
	)
	flag.StringVar(
		&inputFileName,
		"i",
		"",
		"Файл с входными данными",
	)
	flag.StringVar(
		&outputFileName,
		"o",
		"",
		"Результирующий файл",
	)
	flag.Parse()
	if baseType == 0 {
		fmt.Fprintln(os.Stderr, "Нужен флаг -base")
		flag.Usage()
		os.Exit(1)
	}
	if baseType != binType && baseType != decType && baseType != hexType {
		fmt.Fprintln(os.Stderr, "Поддерживаемые системы счисления 2, 10 и 16")
		os.Exit(1)
	}
	if inputFileName == "" {
		fmt.Fprintln(os.Stderr, "Нужен флаг -i")
		flag.Usage()
		os.Exit(1)
	}
	if outputFileName == "" {
		fmt.Fprintln(os.Stderr, "Нужен флаг -o")
		flag.Usage()
		os.Exit(1)
	}
	configLines := make([]string, 3)

	inputFile, err := os.Open(inputFileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Не могу открыть входной файл: %v\n", err)
		os.Exit(1)
	}
	defer inputFile.Close()

	outputFile, err = os.Create(outputFileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Не могу создать выходной файл: %v\n", err)
		os.Exit(1)
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	for i := 0; i < cap(configLines); i++ {
		if scanner.Scan() {
			configLines[i] = scanner.Text()
		} else {
			fmt.Fprintln(os.Stderr, "Неполный ввод")
			os.Exit(1)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка ввода: %v", err)
		os.Exit(1)
	}

	switch configLines[0] {
	case nonSuperSingular:
		init2(configLines[1], configLines[2])
	case superSingular:
		init2(configLines[1], configLines[2])
	default:
		P = parseBigNum(configLines[0], false)
		aP, bP = parseAB(configLines[2])
		aP.Mod(aP, P)
		bP.Mod(bP, P)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода")
		return
	}

	for scanner.Scan() {
		task := scanner.Text()
		switch configLines[0] {
		case nonSuperSingular:
			solveTask2N(task)
		case superSingular:
			solveTask2S(task)
		default:
			solveTaskP(task)
		}
	}
}
