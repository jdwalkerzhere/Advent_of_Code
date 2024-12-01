from collections import Counter
from sys import argv


def get_nums(filepath: str) -> tuple[list[int], list[int]]:
    with open(filepath, 'r') as file:
        nums = [[int(digits) for digits in line.split()] for line in file.readlines()]
        left, right = [], []
        for group in nums:
            left.append(group[0])
            right.append(group[1])
        return sorted(left), sorted(right)

def part_one(left: list[int], right: list[int]) -> int:
    return sum(abs(l - r) for l, r in zip(left, right))


def part_two(left: list[int], right: list[int]) -> int:
    num_map = Counter([n for n in right if n in set(left)])
    return sum(num*occurance for num, occurance in num_map.items())


if __name__ == '__main__':
    filepath = argv[-1]
    left, right = get_nums(filepath)
    print(f'Total of Differences: {part_one(left, right)}')
    print(f'Similarity Score: {part_two(left, right)}')
