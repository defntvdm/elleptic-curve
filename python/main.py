# -*- coding: cp1251 -*-

from argparse import ArgumentParser
import os
from sys import argv, exit

from parsers import parse_polynomial, parse_coeff, parse_e
from p import solveP
from n import solveN
from s import solveS

def parse_args():
    parser = ArgumentParser(
        epilog='(C) Николаев Вадим КБ-501 2018',
        description='Сложение точек на эллипитической кривой',
        usage='{} [-i|--input FILENAME] [-o|--output FILENAME]',
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
        line = line.strip()
        line = line.replace('\ufeff', '')
        if not line:
            continue
        yield line


def ntvdm():
    args = parse_args()
    inp, out = args.input, args.output
    
    try:
        inp_file = open(inp, 'rt', encoding='utf-8')
    except Exception as exc:
        print('Error opening file {}: {}'.format(inp, str(exc)))
        exit(1)

    lines = read_lines(inp_file)

    try:
        curve_type = next(lines)
        polynomial, a, b, c, p = None, None, None, None, None
        if curve_type in ('2N', '2S'):
            polynomial = parse_polynomial(next(lines))
            a, b, c = parse_coeff(next(lines))
        else:
            p = parse_e(curve_type)[0]
            a, b = parse_coeff(next(lines))

        with open(out, 'wt', encoding='cp1251') as res_f:
            for line in lines:
                tokens = [e.strip() for e in line.split() if e]
                if curve_type == '2N':
                    res_f.write(solveN(tokens, polynomial, a, b, c))
                    res_f.write(os.linesep)
                elif curve_type == '2S':
                    res_f.write(solveS(tokens,  polynomial, a, b, c))
                    res_f.write(os.linesep)
                else:
                    res_f.write(solveP(tokens, p, a, b))
                    res_f.write(os.linesep)

    except StopIteration:
        print('Неполный ввод')
        exit(1)

    finally:
        inp_file.close()
    

if __name__ == '__main__':
    ntvdm()
