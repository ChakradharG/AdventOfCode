import numpy as np
import re
from numpy.typing import NDArray
import cvxpy as cp
from typing import Any


class Machine:
	def __init__(self, line: str) -> None:
		l: int = line.find("{")
		r: int = line.find("}", l+1)
		N: int = line.count(",", l+1, r) + 1
		self.Y: NDArray[np.float64] = np.zeros(N, dtype=np.float64)
		for i, jolt in enumerate(line[l+1:r].split(",")):
			self.Y[i] = int(jolt)
		matches: list[str] = re.findall(r"\([\d,]+\)", line)
		M: int = len(matches)
		self.X: NDArray[np.float64] = np.zeros((N, M), dtype=np.float64)
		for j, match in enumerate(matches):
			for i in map(int, match[1:-1].split(",")):
				self.X[i, j] = 1

	def solve(self) -> int:
		w: Any = cp.Variable(self.X.shape[1], integer=True)
		cp.Problem(
			cp.Minimize(cp.sum(w)), 
			constraints=[self.X @ w == self.Y, w >= 0],
		).solve(solver="GLPK_MI")
		return int(w.value.sum())


def main():
	ans: int = 0
	with open("./input.txt", "r") as f:
		for line in f.readlines():
			ans += Machine(line).solve()
	print(ans)

main()
