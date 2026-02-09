import re
from enum import Enum
from collections import namedtuple
from typing import Optional

import numpy as np
import cv2
import ollama
import base64

Ins_type = Enum("Ins_type", [("rect", 0), ("rotate", 1)])
Axis = Enum("Axis", [("y", 0), ("x", 1)])
Instruction = namedtuple("Instruction", ["type", "a", "b", "axis"])

class Screen:
	def __init__(self, M=6, N=50) -> None:
		self.M, self.N = M, N
		self.screen = [["." for j in range(N)] for i in range(M)]
	def rect(self, a: int, b: int) -> int:
		count = 0
		for i in range(b % self.M):
			for j in range(a % self.N):
				if self.screen[i][j] == ".":
					count += 1
					self.screen[i][j] = "#"
		return count
	def shift_row(self, row: int, by: int) -> None:
		temp: list[str] = []
		for j in range(self.N):
			temp.append(self.screen[row][(j - by + self.N) % self.N])
		self.screen[row] = temp
	def shift_col(self, col: int, by: int) -> None:
		temp: list[str] = []
		for i in range(self.M):
			temp.append(self.screen[(i - by + self.M) % self.M][col])
		for i in range(self.M):
			self.screen[i][col] = temp[i]

def create_image(screen: Screen) -> np.ndarray:
	img = np.zeros((screen.M, screen.N + 10), dtype=np.uint8)
	for i in range(screen.M):
		for j in range(screen.N):
			if screen.screen[i][j] == "#":
				img[i][j + j//5] = 255
	kernel = np.ones((7, 7), dtype=np.uint8)
	img = cv2.resize(img, (screen.N * 20, screen.M * 20), interpolation=cv2.INTER_NEAREST)
	img = np.pad(img, ((20, 20), (20, 10)), mode="constant", constant_values=0)
	img = cv2.dilate(img, kernel, iterations=1)
	return img

def cv2_to_base64(img: np.ndarray) -> str:
    _, buffer = cv2.imencode(".png", img)
    img_bytes = buffer.tobytes()
    return base64.b64encode(img_bytes).decode("utf-8")

def part1(instructions: list[Instruction]) -> int:
	screen = Screen()
	ans = 0
	for ins in instructions:
		if ins.type == Ins_type.rect:
			ans += screen.rect(ins.a, ins.b)
		else:
			if ins.axis == Axis.y:
				screen.shift_row(ins.a, ins.b)
			else:
				screen.shift_col(ins.a, ins.b)
	return ans

def part2(instructions: list[Instruction]) -> str:
	screen = Screen()
	for ins in instructions:
		if ins.type == Ins_type.rect:
			screen.rect(ins.a, ins.b)
		else:
			if ins.axis == Axis.y:
				screen.shift_row(ins.a, ins.b)
			else:
				screen.shift_col(ins.a, ins.b)
	# return "\n" + "\n".join("".join(screen.screen[i]) for i in range(screen.M))
	# too easy, let's do something more fun
	return par2_alternate(screen)

def par2_alternate(screen: Screen) -> str:
	img = create_image(screen)
	ans = ollama.chat(
		model="glm-ocr", 
		messages=[{
			"role": "user", 
			"content": "Extract the text from the image. Be precise and do not add any extra characters.",
			"images": [cv2_to_base64(img)],
		}],
		options={"temperature": 0.0}
	)["message"]["content"]
	return ans

def parse(line: str) -> Optional[Instruction]:
	rect = re.search(r"(\d+)x(\d+)", line)
	rotate = re.search(r"(x|y)=(\d+) by (\d+)", line)
	if rect is not None:
		return Instruction(Ins_type.rect, int(rect.group(1)), int(rect.group(2)), None)
	elif rotate is not None:
		axis = Axis.x if rotate.group(1) == "x" else Axis.y
		return Instruction(Ins_type.rotate, int(rotate.group(2)), int(rotate.group(3)), axis)

def run() -> None:
	instructions: list[Instruction] = []
	with open("./day08/input.txt", "r") as file:
		for line in file:
			if (ins := parse(line)) is not None:
				instructions.append(ins)

	print("Part 1:", part1(instructions))
	print("Part 2:", part2(instructions))
