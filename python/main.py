#!/usr/bin/env python3

from argparse import ArgumentParser
import os
from sys import argv, exit

from parsers import parse_polynomial, parse_coeff
from p import solveP
from n import solveN
from s import solveS

def parse_args():
    parser = ArgumentParser(
        epilog='(C) Николаев Вадим КБ-501 2018',
        description='Сложение точек на эллипитической кривой',
        usage='{} [-b|--base BASE] [-i|--input FILENAME] [-o|--output FILENAME]',
    )
    parser.add_argument(
        '-b', '--base',
        type=int,
        choices=[2, 10, 16],
        default=10,
        help='Система счисления, в которой даются числа [10]',
    )
    parser.add_argument(
        '-i', '--input',
        type=str,
        default='input.txt',
        help='Входной файл в формате, указанном в README.md [input.txt]',
    )
    parser.add_argument(
        '-o', '--output',
        type=str,
        default='output.txt',
        help='Выходной файл в формате, указанном в README.md [output.txt]',
    )
    return parser.parse_args()


def read_lines(f):
    for line in f:
        if not line:
            continue
        yield line[:-1] if line[-1] == '\n' else line

def ntvdm():
    args = parse_args()
    inp, out, base = args.input, args.output, args.base
    try:
        inp_file = open(inp, 'rt')
    except Exception as exc:
        print('Error opening file {}: {}'.format(inp, str(exc)))
        exit(1)

    lines = read_lines(inp_file)

    try:
        curve_type = next(lines)
        polynomial, a, b, c, p = None, None, None, None, None
        if curve_type in ('2N', '2S'):
            polynomial = parse_polynomial(next(lines), base)
            a, b, c = parse_coeff(next(lines), base)
        else:
            p = int(curve_type, base)
            a, b = parse_coeff(next(lines), base)

        with open(out, 'wt') as res_f:
            for line in lines:
                tokens = [e.strip() for e in line.split() if e]
                if curve_type == '2N':
                    res_f.write(solveN(tokens, base, polynomial, a, b, c))
                elif curve_type == '2S':
                    res_f.write(solveS(tokens, base,  polynomial, a, b, c))
                else:
                    res_f.write(solveP(tokens, base, p, a, b))

    except StopIteration:
        print('Неполный ввод')
        exit(1)
    finally:
        inp_file.close()
    

if __name__ == '__main__':
    ntvdm()
