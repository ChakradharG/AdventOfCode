from day08.solution import Ins_type, Axis, Instruction, Screen, part1, parse

def test_Screen():
	screen = Screen(3, 7)
	assert screen.rect(3, 2) == 6
	assert screen.screen == [
		["#", "#", "#", ".", ".", ".", "."],
		["#", "#", "#", ".", ".", ".", "."],
		[".", ".", ".", ".", ".", ".", "."]
	]

	screen.shift_col(1, 1)
	assert screen.screen == [
		["#", ".", "#", ".", ".", ".", "."],
		["#", "#", "#", ".", ".", ".", "."],
		[".", "#", ".", ".", ".", ".", "."]
	]

	screen.shift_row(0, 4)
	assert screen.screen == [
		[".", ".", ".", ".", "#", ".", "#"],
		["#", "#", "#", ".", ".", ".", "."],
		[".", "#", ".", ".", ".", ".", "."]
	]

def test_parse_part1():
	instructions: list[Instruction] = []
	res = parse("rect 3x2")
	assert res == Instruction(Ins_type.rect, 3, 2, None)
	instructions.append(res) # type: ignore
	res = parse("rotate column x=1 by 1")
	assert res == Instruction(Ins_type.rotate, 1, 1, Axis.x)
	instructions.append(res) # type: ignore
	res = parse("rotate row y=0 by 4")
	assert res == Instruction(Ins_type.rotate, 0, 4, Axis.y)
	instructions.append(res) # type: ignore
	assert parse("fasf fasf") is None

	assert part1(instructions) == 6
