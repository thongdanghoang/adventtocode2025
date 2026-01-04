import os


def read_input(filename="input.txt"):
    with open(os.path.join(os.path.dirname(__file__), filename), "r") as f:
        return f.read().strip().split("\n")


def part1(data):
    idx = data.index('')
    ranges = data[:idx]
    numbers = data[idx + 1:]

    available_ingredient_ids = available_Ingredient_Id(ranges, numbers)

    return len(available_ingredient_ids)


def available_Ingredient_Id(ranges, numbers) -> set:
    available_ids = set()
    for range in ranges:
        range = range.split('-')
        for number in numbers:
            if int(range[0]) <= int(number) <= int(range[1]):
                available_ids.add(int(number))
    return available_ids


def part2(data):
    idx = data.index('')
    ranges = data[:idx]
    ranges_map: list[tuple[int, int]] = []
    for range_str in ranges:
        range_str_split = range_str.split('-')
        start = range_str_split[0]
        end = range_str_split[1]
        ranges_map = merge_ranges(ranges_map, (int(start), int(end)))

    total = 0
    for range_tupple in ranges_map:
        total += range_tupple[1] - range_tupple[0] + 1
    return total


def merge_ranges(ranges: list[tuple[int, int]], new_range: tuple[int, int]) -> list[tuple[int, int]]:
    if len(ranges) == 0:
        ranges.append(new_range)
        return ranges

    overlaped_ranges = find_overlaped_ranges(ranges, new_range)

    if len(overlaped_ranges) == 0:
        ranges.append(new_range)
        return ranges

    return merge_overlaped_ranges(ranges, overlaped_ranges, new_range)


def find_overlaped_ranges(ranges: list[tuple[int, int]], new_range: tuple[int, int]) -> list[tuple[int, int]]:
    overlaped_ranges = []
    new_start = new_range[0]
    new_end = new_range[1]
    for cur_range in ranges:
        cur_start = int(cur_range[0])
        cur_end = int(cur_range[1])
        # if cur_start <= new_start <= cur_end or cur_start <= new_end <= cur_end:
        #     overlaped_ranges.append(cur_range)
        if cur_start <= new_end and new_start <= cur_end:
            overlaped_ranges.append(cur_range)
    return overlaped_ranges


def merge_overlaped_ranges(ranges: list[tuple[int, int]], overlaped_ranges: list[tuple[int, int]],
                           new_range: tuple[int, int]) -> list[tuple[int, int]]:
    for overlaped_range in overlaped_ranges:
        ranges.remove(overlaped_range)

    overlaped_ranges.append(new_range)
    start = min(overlaped_ranges, key=lambda x: x[0])[0]
    end = max(overlaped_ranges, key=lambda x: x[1])[1]
    ranges.append((start, end))
    return ranges


if __name__ == "__main__":
    data = read_input()
    # data = read_input("sample_input.txt")

    print("--- Part 1 ---")
    print(part1(data))

    print("--- Part 2 ---")
    print(part2(data))
