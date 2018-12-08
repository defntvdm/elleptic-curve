from typing import List

from char2utils import formatMul2, formatSum2, mul, mod, inverse


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

    def __mul__(self, num: int):
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
    tokens: List[str],
    base: int,
    polynomial: int,
    a: int,
    b: int,
    c: int,
) -> str:
    Point.polynomial = polynomial
    Point.a = a
    Point.b = b
    Point.c = c
    if tokens[0] == 'У':
        if len(tokens) != 4:
            print('Неверный ввод')
            exit(1)
        num = int(tokens[3], base)
        point = Point(int(tokens[1], 2), int(tokens[2], 2))
        return formatMul2(point, num, point * num, base, polynomial.bit_length() - 1)
    elif tokens[0] == 'С':
        if len(tokens) != 5:
            print('Неверный ввод')
            exit(1)
        point1 = Point(int(tokens[1], 2), int(tokens[2], 2))
        point2 = Point(int(tokens[3], 2), int(tokens[4], 2))
        return formatSum2(point1, point2, point1 + point2, polynomial.bit_length() - 1)
    else:
        print("Ожидалась операция сложения двух точек 'С' или умножения точки на число 'У'")
        exit(1)