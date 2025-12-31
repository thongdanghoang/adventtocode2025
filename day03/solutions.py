import os
from collections import deque
from itertools import islice


def read_input(filename="input.txt"):
    with open(os.path.join(os.path.dirname(__file__), filename), "r") as f:
        return f.read().strip().split("\n")


def part1(data):
    total = 0
    for e in data:
        result = deque()
        capacity = len(e) - 2
        for digit in e:
            result, capacity = append(result, int(digit), capacity)
        first_2_lagest_number: list[str] = list(islice(result, 2))
        join: int = int("".join(map(str, first_2_lagest_number)))
        total += join
    return total


def append(stack: deque, value: int, capacity: int) -> tuple[deque, int]:
    while True:
        if len(stack) == 0:
            stack.append(value)
            break
        if stack[-1] >= value:
            stack.append(value)
            break
        if capacity == 0:
            stack.append(value)
            break
        stack.pop()
        capacity -= 1
    return stack, capacity



def part2(data):
    total = 0
    for e in data:
        result = deque()
        capacity = len(e) - 12
        if capacity < 0: continue
        for digit in e:
            result, capacity = append(result, int(digit), capacity)
        first_2_lagest_number: list[str] = list(islice(result, 12))
        join: int = int("".join(map(str, first_2_lagest_number)))
        total += join
    return total


if __name__ == "__main__":
    # data = ["987654321111111", "811111111111119", "234234234234278", "818181911112111"]
    data = read_input()

    print("--- Part 1 ---")
    print(part1(data))

    print("--- Part 2 ---")
    print(part2(data))
