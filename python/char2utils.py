# -*- coding: cp1251 -*-

def mul(n1, n2, polynomial):
    if n1 == 0 or n2 == 0:
        return 0
    res = n1
    for c in bin(n2)[3:]:
        res <<= 1
        if c == '1':
            res ^= n1
    return mod(res, polynomial)


def mod(number, polynomial):
    bit_len = number.bit_length()
    polybl = polynomial.bit_length()
    while bit_len >= polybl:
        number ^= polynomial << (bit_len- polybl)
        bit_len = number.bit_length()
    return number


def inverse(number, polynomial):
    value = [number, polynomial]
    x_col = [1, 0]
    while value[-1] != 1:
        v1len = value[-1].bit_length()
        v2len = value[-2].bit_length()
        if v1len == v2len:
            value.append(value[-1] ^ value[-2])
            x_col.append(x_col[-1] ^ x_col[-2])
        else:
            if v1len < v2len:
                value.append(value[-1] << (v2len-v1len))
                x_col.append(x_col[-1] << (v2len-v1len))
                value[-1] = value[-1] ^ value[-3]
                x_col[-1] = x_col[-1] ^ x_col[-3]
            else:
                value.append(value[-2] << (v1len-v2len))
                x_col.append(x_col[-2] << (v1len-v2len))
                value[-1] = value[-1] ^ value[-2]
                x_col[-1] = x_col[-1] ^ x_col[-2]
    return mod(x_col[-1], polynomial)


def formatPoint(p, length):
    if p.x == None:
        return 'E'
    return '({}, {})'.format(
        bin(p.x)[2:].rjust(length, '0'),
        bin(p.y)[2:].rjust(length, '0'),
    )


def formatSum2(p1, p2, res, length):
    return '{} + {} = {}'.format(
        formatPoint(p1, length),
        formatPoint(p2, length),
        formatPoint(res, length),
    )


def formatMul2(p, num, res, base, length):
    if base == 10:
        num = str(num)
    elif base == 2:
        num = bin(num)[2:]
    else:
        num = hex(num)[2:]
    return '{} * {} = {}'.format(
        formatPoint(p, length),
        num,
        formatPoint(res, length),
    )
