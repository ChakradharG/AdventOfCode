def part1(instructions: list[str]) -> int:
	x, y = 0, 0
	dx, dy = 0, 1
	for ins in instructions:
		d, m = ins[0], int(ins[1:])
		if d == 'R':
			dx, dy = dy, -dx
		else:
			dx, dy = -dy, dx
		x, y = x + m*dx, y + m*dy
	return abs(x) + abs(y)

def part2(instructions: list[str]) -> int:
	x, y = 0, 0
	dx, dy = 0, 1
	vis = {(0, 0)}
	for ins in instructions:
		d, m = ins[0], int(ins[1:])
		if d == 'R':
			dx, dy = dy, -dx
		else:
			dx, dy = -dy, dx
		for _ in range(m):
			x, y = x + dx, y + dy
			if (x, y) in vis:
				return abs(x) + abs(y)
			vis.add((x, y))
	return abs(x) + abs(y)

def run() -> None:
	with open("./day01/input.txt", "r") as file:
		instructions: list[str] = file.readline().strip().split(", ")

	print("Part 1:", part1(instructions))
	print("Part 2:", part2(instructions))
