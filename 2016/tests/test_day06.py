from day06.solution import part1, part2

data = [
	"eedadn",
	"drvtee",
	"eandsr",
	"raavrd",
	"atevrs",
	"tsrnev",
	"sdttsa",
	"rasrtv",
	"nssdts",
	"ntnada",
	"svetve",
	"tesnvt",
	"vntsnd",
	"vrdear",
	"dvrsen",
	"enarar"
]

def test_part1():
	assert part1(data) == "easter"

def test_part2():
	assert part2(data) == "advent"
