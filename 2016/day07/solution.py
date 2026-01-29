import re
from dataclasses import dataclass

@dataclass
class IP:
	superseqs: list[str]
	hypernets: list[str]

def contains_ABBA(seq: str) -> bool:
	for i in range(len(seq)-3):
		if seq[i] != seq[i+1] and seq[i] == seq[i+3] and seq[i+1] == seq[i+2]:
			return True
	return False

def supports_TLS(ip: IP) -> bool:
	if any(map(contains_ABBA, ip.hypernets)):
		return False
	if any(map(contains_ABBA, ip.superseqs)):
		return True
	return False

def supports_SSL(ip: IP) -> bool:
	ABAs = set()
	for seq in ip.superseqs:
		for i in range(len(seq)-2):
			if seq[i] != seq[i+1] and seq[i] == seq[i+2]:
				ABAs.add((seq[i], seq[i+1]))
	for seq in ip.hypernets:
		for i in range(len(seq)-2):
			if seq[i] != seq[i+1] and seq[i] == seq[i+2]:
				if (seq[i+1], seq[i]) in ABAs:
					return True
	return False

def part1(ips: list[IP]) -> int:
	return len(list(filter(supports_TLS, ips)))

def part2(ips: list[IP]) -> int:
	return len(list(filter(supports_SSL, ips)))

def run() -> None:
	ips: list[IP] = []
	ob = re.compile(r"[a-z]+|\[[a-z]+\]")
	with open("./day07/input.txt", "r") as file:
		for line in file:
			ips.append(IP([], []))
			for res in ob.findall(line.strip()):
				if res.startswith("["):
					ips[-1].hypernets.append(res[1:-1])
				else:
					ips[-1].superseqs.append(res)

	print("Part 1:", part1(ips))
	print("Part 2:", part2(ips))
