# -*- coding: cp1251 -*-

from char2utils import formatMul2, formatSum2, mul, mod, inverse
from parsers import parse_e

class Point:
    polynomial = None
    a = None
    b = None
    c = None
    
    def __init__(self, x=None, y=None):
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
            x3 = mul(k, k, Point.polynomial) ^ self.x ^ other.x ^ mul(Point.a, k, Point.polynomial) ^ Point.b
        else:
            if self.y != other.y or self.x == 0:
                return Point()
            numerator = mul(self.x, self.x, Point.polynomial) ^ mul(Point.a, self.y, Point.polynomial)
            denominator = mul(self.x, Point.a, Point.polynomial)
            denominator = inverse(denominator, Point.polynomial)
            k = mul(numerator, denominator, Point.polynomial)
            x3 = mul(k, k, Point.polynomial) ^ Point.b ^ mul(k, Point.a, Point.polynomial)
        y3 = self.y ^ mul(k, self.x ^ x3, Point.polynomial) ^ mul(Point.a, x3, Point.polynomial)
        return Point(x3, y3)

    def __mul__(self, num):
        if num < 0:
            new_point = Point(self.x, mul(self.x, Point.a, Point.polynomial) ^ self.y)
            return new_point * (-num)
        if num == 0:
            return Point()
        result = self
        for c in bin(num)[3:]:
            result += result
            if c == '1':
                result += self
        return result


def solveN(
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
    if tokens[0] == '�':
        if len(tokens) != 4:
            print('�������� ����')
            exit(1)
        num, base = parse_e(tokens[3])
        x = parse_e(tokens[1])[0]
        y = parse_e(tokens[2])[0]
        point = Point(x, y)
        return formatMul2(point, num, point * num, base, polynomial.bit_length() - 1)
    elif tokens[0] == '�':
        if len(tokens) != 5:
            print('�������� ����')
            exit(1)
        x1 = parse_e(tokens[1])[0]
        y1 = parse_e(tokens[2])[0]
        x2 = parse_e(tokens[3])[0]
        y2, base = parse_e(tokens[4])
        point1 = Point(x1, y1)
        point2 = Point(x2, y2)
        return formatSum2(point1, point2, point1 + point2, polynomial.bit_length() - 1)
    else:
        print("��������� �������� �������� ���� ����� '�' ��� ��������� ����� �� ����� '�'")
        exit(1)
