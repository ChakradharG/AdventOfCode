from hashlib import md5

def starts_with_5_zeros(hsh: str) -> bool:
	if len(hsh) < 5:
		return False
	return hsh[:5] == "00000"

def is_valid(hsh: str, empty_pos: set[str]) -> bool:
	if not starts_with_5_zeros(hsh):
		return False
	if len(hsh) < 7:
		return False
	return hsh[5] in empty_pos

def part1(door_id: str) -> str:
	index = 0
	password = ""
	while len(password) < 8:
		hsh = md5(f"{door_id}{index}".encode()).hexdigest()
		if starts_with_5_zeros(hsh):
			password += hsh[5]
		index += 1
	return password

def part2(door_id: str) -> str:
	index = 0
	password = ["" for _ in range(8)]
	empty_pos = set("01234567")
	while empty_pos:
		hsh = md5(f"{door_id}{index}".encode()).hexdigest()
		if is_valid(hsh, empty_pos):
			empty_pos.remove(hsh[5])
			password[int(hsh[5])] = hsh[6]
		index += 1
	return "".join(password)

def run() -> None:
	with open("./day05/input.txt", "r") as file:
		door_id = file.readline().strip()

	print("Part 1:", part1(door_id))
	print("Part 2:", part2(door_id))
