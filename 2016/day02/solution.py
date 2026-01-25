def part1(data: list[str]) -> str:
	ans: str = ""
	i, j = 1, 1
	for line in data:
		for c in line:
			if c == "U":
				i = max(0, i-1)
			elif c == "R":
				j = min(2, j+1)
			elif c == "D":
				i = min(2, i+1)
			else:
				j = max(0, j-1)
		ans += str(3*i + j + 1)
	return ans

def part2(data: list[str]) -> str:
	keypad: list[list[str]] = [
		[" ", " ", "1", " ", " "],
		[" ", "2", "3", "4", " "],
		["5", "6", "7", "8", "9"],
		[" ", "A", "B", "C", " "],
		[" ", " ", "D", " ", " "]
	]
	ans: str = ""
	i, j = 2, 0
	for line in data:
		for c in line:
			ni, nj = i, j
			if c == "U":
				ni = max(0, i-1)
			elif c == "R":
				nj = min(4, j+1)
			elif c == "D":
				ni = min(4, i+1)
			else:
				nj = max(0, j-1)
			if keypad[ni][nj] != " ":
				i, j = ni, nj
		ans += keypad[i][j]
	return ans

def run() -> None:
	data: list[str] = []
	with open("./day02/input.txt", "r") as file:
		for line in file.readlines():
			data.append(line.strip())

	print("Part 1:", part1(data))
	print("Part 2:", part2(data))
