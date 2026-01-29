from collections import Counter

def part1(data: list[str]) -> str:
	ans = ""
	for j in range(len(data[0])):
		f = Counter(map(lambda x: x[j], data))
		ans += max(f.keys(), key=f.__getitem__)
	return ans

def part2(data: list[str]) -> str:
	ans = ""
	for j in range(len(data[0])):
		f = Counter(map(lambda x: x[j], data))
		ans += min(f.keys(), key=f.__getitem__)
	return ans

def run() -> None:
	data: list[str] = []
	with open("./day06/input.txt", "r") as file:
		for line in file.readlines():
			data.append(line.strip())

	print("Part 1:", part1(data))
	print("Part 2:", part2(data))
