import re
from collections import namedtuple, defaultdict

Instruction = namedtuple("Instruction", ["chip", "bot_id"])
TARGET_LO = 17
TARGET_HI = 61

class Output_Bin:
	def __init__(self) -> None:
		self.chips: list[int] = []
	def propagate_chip(self, chip: int) -> None:
		self.chips.append(chip)

class Bot:
	def __init__(self) -> None:
		self.next_lo: Bot | Output_Bin | None = None
		self.next_hi: Bot | Output_Bin | None = None
		self.chip_lo: int = -1
		self.chip_hi: int = -1
		self.responsible: bool = False
	def propagate_chip(self, chip: int) -> None:
		if self.chip_lo == -1:
			self.chip_lo = chip
		else:
			self.chip_hi = max(self.chip_lo, chip)
			self.chip_lo = min(self.chip_lo, chip)
			if (self.chip_lo == TARGET_LO and self.chip_hi == TARGET_HI):
				self.responsible = True
			if self.next_lo is not None:
				self.next_lo.propagate_chip(self.chip_lo)
			if self.next_hi is not None:
				self.next_hi.propagate_chip(self.chip_hi)
			self.chip_lo = -1
			self.chip_hi = -1

def part1(bots: dict[str, Bot], instructions: list[Instruction]) -> str:
	for chip, bot_id in instructions:
		bots[bot_id].propagate_chip(chip)
	for bot_id, bot in bots.items():
		if bot.responsible:
			return bot_id
	return "-1"	# should not reach this point

def part2(output_bins: dict[str, Output_Bin]) -> int:
	result = 1
	for bin_id in ["0", "1", "2"]:
		result *= output_bins[bin_id].chips[0]
	return result

def run() -> None:
	ob1 = re.compile(r"value (\d+) goes to bot (\d+)")
	ob2 = re.compile(r"bot (\d+) gives low to (bot|output) (\d+) and high to (bot|output) (\d+)")
	output_bins: dict[str, Output_Bin] = defaultdict(Output_Bin)
	bots: dict[str, Bot] = defaultdict(Bot)
	instructions: list[Instruction] = []
	with open("./day10/input.txt", "r") as file:
		for line in file:
			if line.startswith("v"):
				if (mth := ob1.search(line)) is not None:
					instructions.append(Instruction(int(mth.groups()[0]), mth.groups()[1]))
			else:
				if (mth := ob2.search(line)) is not None:
					bot_id, type_lo, next_lo_id, type_hi, next_hi_id = mth.groups()
					bots[bot_id].next_lo = bots[next_lo_id] if type_lo == "bot" else output_bins[next_lo_id]
					bots[bot_id].next_hi = bots[next_hi_id] if type_hi == "bot" else output_bins[next_hi_id]

	print("Part 1:", part1(bots, instructions))
	print("Part 2:", part2(output_bins))
