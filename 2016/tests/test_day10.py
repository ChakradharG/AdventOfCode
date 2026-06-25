from day10.solution import Bot, OutputBin, part1, Instruction

def test_bot_propagate_chip():
	bot = Bot()
	output_bin_lo = OutputBin()
	output_bin_hi = OutputBin()
	bot.next_lo = output_bin_lo
	bot.next_hi = output_bin_hi

	bot.propagate_chip(5)
	assert bot.chip_lo == 5
	assert bot.chip_hi == -1
	assert not bot.responsible

	bot.propagate_chip(10)
	assert bot.chip_lo == -1
	assert bot.chip_hi == -1
	assert not bot.responsible
	assert output_bin_lo.chips == [5]
	assert output_bin_hi.chips == [10]

	bot.propagate_chip(17)
	bot.propagate_chip(61)
	assert bot.responsible

def test_part1():
	bots = {
		"0": Bot(),
		"1": Bot(),
		"2": Bot()
	}
	output_bins = {
		"0": OutputBin(),
		"1": OutputBin(),
		"2": OutputBin()
	}

	bots["0"].next_lo = output_bins["2"]
	bots["0"].next_hi = output_bins["0"]
	bots["1"].next_lo = output_bins["1"]
	bots["1"].next_hi = bots["0"]
	bots["2"].next_lo = bots["1"]
	bots["2"].next_hi = bots["0"]

	instructions = [
		Instruction(61, "2"),
		Instruction(19, "1"),
		Instruction(17, "2")
	]

	assert part1(bots, instructions) == "2"