from day02.solution import part1, part2

data: list[str] = [
	"ULL",
	"RRDDD",
	"LURDL",
	"UUUUD"
]

def test_part1():
	assert part1(data) == "1985"

def test_part2():
	assert part2(data) == "5DB3"
