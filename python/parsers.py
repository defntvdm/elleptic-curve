from typing import Tuple


def parse_polynomial(line: str, base: int) -> int:
    res = 0
    members = line.split('+')
    for member in members:
        if '+' in member:
            continue
        striped = member.strip()
        if striped == '1':
            res ^= 1
            continue
        if striped == 'x':
            res ^= 0b10
            continue
        power = int(member.split('^')[1].strip(), base)
        res ^= (1 << power)
    return res


def parse_coeff(line: str, base: int) -> Tuple[int]:
    return [int(e, base) for e in line.split()]
