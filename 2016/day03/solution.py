def part1(triangles: list[list[int]]) -> int:
	ans: int = 0
	for triangle in triangles:
		a, b, c = sorted(triangle)
		if c < a + b:
			ans += 1
	return ans

def part2(triangles: list[list[int]]) -> int:
	ans: int = 0
	for i in range(0, len(triangles), 3):
		for j in range(3):
			a, b, c = sorted([triangles[i][j], triangles[i+1][j], triangles[i+2][j]])
			if c < a + b:
				ans += 1
	return ans

def run() -> None:
	triangles: list[list[int]] = []
	with open("./day03/input.txt", "r") as file:
		for line in file.readlines():
			triangles.append([*map(int, line.split())])

	print("Part 1:", part1(triangles))
	print("Part 2:", part2(triangles))
