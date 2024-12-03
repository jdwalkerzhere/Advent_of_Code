from sys import argv
import re


token_spec = [
        ('MUL', r"mul\(\d{1,3},\d{1,3}\)"),
        ('DO', r"do\(\)"),
        ('DONT', r"don't\(\)")
]
tok_regex = '|'.join('(?P<%s>%s)' % pair for pair in token_spec)


def parts(filepath: str) -> tuple[int, int]:
    with open(filepath, 'r') as data:
        read_data = data.read()

    do = True
    p_one_total = 0
    total = 0
    for mo in re.finditer(tok_regex, read_data):
        kind = mo.lastgroup
        value = mo.group()
        match kind:
            case 'MUL':
                lhs, rhs = re.findall(r'\d{1,3}', value)
                product = int(lhs) * int(rhs)
                p_one_total += product
                if do:
                    total += product
            case 'DO':
                do = True
            case 'DONT':
                do = False
    return p_one_total, total

def main():
    filepath = argv[-1]
    ans_part_one, ans_part_two = parts(filepath)
    print(f"Non-corrupted mul-commands sum: {ans_part_one}")
    print(f'Non-corrupted with dos/donts: {ans_part_two}')

if __name__ == '__main__':
    main()
