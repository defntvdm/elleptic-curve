# -*- coding: cp1251 -*-

from char2utils import formatMul2, formatSum2, mul, mod, inverse
from parsers import parse_e


class Point:
    polynomial = None
    a = None
    b = None
    c = None
    
    def __init__(self, x: int=None, y: int=None):
        self.x = x
        self.y = y

    def __add__(self, other):
        if self.x == None:
            return other
        if other.x == None:
            return self
        if self.x != other.x:
            numerator = self.y ^ other.y
            denominator = self.x ^ other.x
            denominator = inverse(denominator, Point.polynomial)
            k = mul(numerator, denominator, Point.polynomial)
            x3 = mul(k, k, Point.polynomial) ^ self.x ^ other.x
        else:
            if self.y != other.y:
                return Point()
            numerator = mul(self.x, other.x, Point.polynomial) ^ Point.b
            denominator = inverse(Point.a, Point.polynomial)
            k = mul(numerator, denominator, Point.polynomial)
            x3 = mul(k, k, Point.polynomial)
        y3 = mul(self.x ^ x3, k, Point.polynomial) ^ self.y ^Point.a
        return Point(x3, y3)

    def __mul__(self, num: int):
        if num < 0:
            new_point = Point(self.x, self.y ^ Point.a)
            return new_point * (-num)
        if num == 0:
            return Point()
        result = self
        for c in bin(num)[3:]:
            result += result
            if c == '1':
                result += self
        return result


def solveS(
    tokens,
    polynomial,
    a,
    b,
    c,
):
    Point.polynomial = polynomial
    Point.a = a
    Point.b = b
    Point.c = c
    if tokens[0] == 'У':
        if len(tokens) != 4:
            print('Неверный ввод')
            exit(1)
        num, base = parse_e(tokens[3])
        x = parse_e(tokens[1])[0]
        y = parse_e(tokens[2])
        point = Point(x, y)
        return formatMul2(point, num, point * num, base, polynomial.bit_length() - 1)
    elif tokens[0] == 'С':
        if len(tokens) != 5:
            print('Неверный ввод')
            exit(1)
        x1 = parse_e(tokens[1])[0]
        y1 = parse_e(tokens[2])[0]
        x2 = parse_e(tokens[3])[0]
        y2, base = parse_e(tokens[4])
        point1 = Point(x1, y1)
        point2 = Point(x2, y2)
        return formatSum2(point1, point2, point1 + point2, polynomial.bit_length() - 1)
    else:
        print("Ожидалась операция сложения двух точек 'С' или умножения точки на число 'У'")
        exit(1)
