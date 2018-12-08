from sys import exit
from typing import List

class Point:
    modulus = None
    a = None
    b = None

    def __init__(self, x: int=None, y: int=None):
        self.x = x
        self.y = y

    def copy(self):
        return Point(self.x, self.y)

    def __add__(self, other):
        if self.x == None:
            return other
        if other.x == None:
            return self
        
        k, denominator, numerator = None, None, None
        
        if self.x != other.x:
            numerator = other.y - self.y
            denominator = other.x - self.x
        else:
            if self.y != other.y or self.y == 0:
                return Point()
            numerator = 3 * pow(self.x, 2) + self.a
            denominator = 2 * self.y
        
        denominator = inverse(denominator, self.modulus)
        k = (numerator * denominator) % self.modulus

        x3 = (pow(k, 2, self.modulus) - self.x - other.x) % self.modulus
        y3 = (k * (self.x - x3) - self.y) % self.modulus
        
        return Point(x3, y3)

    def __mul__(self, n: int):
        if n < 0:
            new_point = Point(self.x, self.modulus - self.y)
            return new_point * -n
        elif n == 0:
            return Point()
        result = self.copy()
        for c in bin(n)[3:]:
            result += result
            if c == '1':
                result += self
        return result


def inverse(number: int, modulus: int) -> int:
    number %= modulus
    prev_n = 0
    cur_n = 1
    prev_res = modulus
    cur_res = number
    while cur_res != 1:
        k = prev_res // cur_res
        prev_res, cur_res = cur_res, prev_res % cur_res
        prev_n, cur_n = cur_n, prev_n - k * cur_n
    return cur_n % modulus



def formatMulP(p: Point, num: int, res: Point, base: int) -> str:
    if base == 10:
        n = str(num)
    elif base == 16:
        n = hex(num)[2:]
    else:
        n = bin(num)[2:]
    return '{} * {} = {}\n'.format(formatPoint(p, base), n, formatPoint(res, base))


def formatPoint(p: Point, base: int):
    if p.x == None:
        return 'E'
    if base == 10:
        x, y = str(p.x), str(p.y)
    elif base == 16:
        x, y = hex(p.x)[2:], hex(p.y)[2:]
    else:
        x, y = bin(p.x)[2:], bin(p.y)[2:]
    return '({}, {})'.format(x, y)

def formatSumP(p1: Point, p2: Point, res: Point, base: int) -> str:
    return '{} + {} = {}\n'.format(formatPoint(p1, base), formatPoint(p2, base), formatPoint(res, base))


def solveP(tokens: List[str], base: int, p: int, a: int, b: int) -> str:
    Point.modulus = p
    Point.a = a
    Point.b = b
    if tokens[0] == 'У':
        if len(tokens) != 4:
            print('Неверный ввод')
            exit(1)
        num = int(tokens[3], base)
        point = Point(int(tokens[1], base), int(tokens[2], base))
        return formatMulP(point, num, point * num, base)
    elif tokens[0] == 'С':
        if len(tokens) != 5:
            print('Неверный ввод')
            exit(1)
        point1 = Point(int(tokens[1], base), int(tokens[2], base))
        point2 = Point(int(tokens[3], base), int(tokens[4], base))
        return formatSumP(point1, point2, point1 + point2, base)
    else:
        print("Ожидалась операция сложения двух точек 'С' или умножения точки на число 'У'")
        exit(1)
