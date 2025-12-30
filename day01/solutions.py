import os


def read_input(filename="input.txt"):
    # Ensure the file is read relative to the script location
    with open(os.path.join(os.path.dirname(__file__), filename), "r") as f:
        return f.read().strip().split("\n")


def part1(data):
    cur: int = 50
    n_point_at_zero: int = 0
    for rotation in data:
        cur, _ = calculate(rotation, cur)
        if cur == 0: n_point_at_zero += 1
    return n_point_at_zero


# input L0 -> L99 , R0 -> R99
# cur: number
def calculate(rotation: str, cur: int):
    rotation_number = int(rotation[1:]) % 100
    pass_over_zero = int(rotation[1:]) // 100
    lamda = cur - rotation_number if rotation[0] == "L" else cur + rotation_number
    if cur != 0:
        if lamda <= 0 or lamda > 99: pass_over_zero += 1
    if lamda < 0:
        lamda += 100
    if lamda > 99:
        lamda -= 100
    return lamda, pass_over_zero


def part2(data):
    cur: int = 50
    total_pass_over_zero: int = 0
    for rotation in data:
        cur, pass_over_zero = calculate(rotation, cur)
        total_pass_over_zero += pass_over_zero
    return total_pass_over_zero


if __name__ == "__main__":
    data = read_input()

    print("--- Part 1 ---")
    # print(part1(data))

    print("--- Part 2 ---")
    print(part2(data))
