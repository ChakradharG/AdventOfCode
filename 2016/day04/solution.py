from dataclasses import dataclass
import re
from collections import Counter
from functools import partial

@dataclass
class Room:
	encrypted_name: str
	sector_id: int
	checksum: str

def is_real(room: Room) -> bool:
	d = Counter(room.encrypted_name.replace("-", ""))
	checksum = "".join(
		k for k, _ in sorted(
			d.items(), 
			key=lambda x: (-x[1], x[0])
		)[:5]
	)
	return room.checksum == checksum

def decrypt(c: str, k: int) -> str:
	if c == "-":
		return " "
	offset = ord("a")
	return chr(((ord(c) - offset + k) % 26) + offset).lower()

def part1(rooms: list[Room]) -> int:
	return sum(room.sector_id for room in rooms if is_real(room))

def part2(rooms: list[Room]) -> int:
	for room in rooms:
		if is_real(room):
			func = partial(decrypt, k=room.sector_id % 26)
			if "northpole" in "".join(map(func, room.encrypted_name)):
				return room.sector_id
	return 0 # no solution found

def run() -> None:
	rooms: list[Room] = []
	ob = re.compile(r"([a-z-]+)(\d+)\[(\w+)\]")
	with open("./day04/input.txt", "r") as file:
		for line in file:
			if (res := ob.search(line)) is None:
				continue
			name, sector, checksum = res.groups()
			rooms.append(Room(
				encrypted_name=name,
				sector_id=int(sector),
				checksum=checksum
			))

	print("Part 1:", part1(rooms))
	print("Part 2:", part2(rooms))
