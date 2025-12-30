import os


def read_input(filename="input.txt"):
    with open(os.path.join(os.path.dirname(__file__), filename), "r") as f:
        return f.read().strip().split("\n")


def part1(data):
    splittedByComma = data.split(",")
    total: int = 0
    for pair in splittedByComma:
        pairArr = pair.split("-")
        firstID = int(pairArr[0])
        secondID = int(pairArr[1])
        for id in range(firstID, secondID + 1):
            if valid_id_by_n_digits_repeated_twice(str(id)): total += id
    return total

def valid_id_by_n_digits_repeated_twice(id: str) -> bool:
    if id[0] == "0": return False
    if len(id) % 2 != 0: return False
    ids = [id[i:i + len(id) // 2] for i in range(0, len(id), len(id) // 2)]
    return len(set(ids)) == 1 and len(ids) > 1


def part2(data):
    splittedByComma = data.split(",")
    total = 0
    for pair in splittedByComma:
        pairArr = pair.split("-")
        firstID = int(pairArr[0])
        secondID = int(pairArr[1])
        for id in range(firstID, secondID + 1):
            if valid_id(str(id)): total += id
    return total


def valid_id(id: str) -> bool:
    for i in range(1, (len(id) // 2) + 1):
        if valid_id_by_n_digits_repeated(id, i): return True
    return False


def valid_id_by_n_digits_repeated(id: str, n_digits_repeated: int) -> bool:
    if id[0] == "0": return False
    if len(id) % n_digits_repeated != 0: return False
    ids = [id[i:i + n_digits_repeated] for i in range(0, len(id), n_digits_repeated)]
    return len(set(ids)) == 1 and len(ids) > 1


if __name__ == "__main__":
    # data = ["11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"]
    data = read_input()

    print("--- Part 1 ---")
    print(part1(data[0]))

    print("--- Part 2 ---")
    print(part2(data[0]))
