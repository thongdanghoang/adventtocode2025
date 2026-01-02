import os


def read_input(filename="input.txt") -> list[list[str]]:
    with open(os.path.join(os.path.dirname(__file__), filename), "r") as f:
        return [list(line) for line in f.read().strip().split("\n")]


def part1(data):
    total: int = 0
    for rowIndex, row in enumerate(data):
        for colIndex, cell in enumerate(row):
            if cell == "@" and fewer_than_four_rolls_of_paper(rowIndex, colIndex, data):
                total += 1
    return total


def fewer_than_four_rolls_of_paper(row_index: int, col_index: int, input: list[list[str]]) -> bool:
    row_lenth = len(input)
    col_lenth = len(input[0])

    rolls_of_paper: int = -1
    for row in range(max(0, row_index - 1), min(row_index + 2, row_lenth)):
        for col in range(max(0, col_index - 1), min(col_index + 2, col_lenth)):
            cell = input[row][col]
            if cell == "@":
                rolls_of_paper += 1
    return rolls_of_paper < 4


def part2(data):
    total: int = 0
    while True:
        cur_total, data = ride_until_no_more_changes(data)
        total += cur_total
        if cur_total == 0:
            break
    return total


def ride_until_no_more_changes(data: list[list[str]]) -> tuple[int, list[list[str]]]:
    total: int = 0
    to_be_replaced: list[tuple[int, int]] = []
    for rowIndex, row in enumerate(data):
        for colIndex, cell in enumerate(row):
            if cell == "@" and fewer_than_four_rolls_of_paper(rowIndex, colIndex, data):
                total += 1
                to_be_replaced.append((rowIndex, colIndex))
    for rowIndex, row in enumerate(data):
        for colIndex, cell in enumerate(row):
            if (rowIndex, colIndex) in to_be_replaced:
                data[rowIndex][colIndex] = "x"
    return total, data


if __name__ == "__main__":
    # data = read_input("sample_input.txt")
    data = read_input()

    print("--- Part 1 ---")
    print(part1(data))

    print("--- Part 2 ---")
    print(part2(data))
