def parse_e(e):
    if e.startswith('0x'):
        return int(e[2:], 16), 16
    if e.startswith('0b'):
        return int(e[2:], 2), 2
    return int(e), 10


def parse_polynomial(line):
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
        power = parse_e(member.split('^')[1].strip())[0]
        res ^= (1 << power)
    return res


def parse_coeff(line):
    return [parse_e(e)[0] for e in line.split()]
