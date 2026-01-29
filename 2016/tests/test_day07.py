import pytest
from day07.solution import IP, contains_ABBA, supports_TLS, supports_SSL

@pytest.mark.parametrize("seq, expected", [
	("ABBA", True),
	("ABBB", False),
	("abcd", False),
	("abca", False),
])
def test_contains_ABBA(seq: str, expected: bool):
	assert contains_ABBA(seq) == expected

@pytest.mark.parametrize("ip, expected", [
	(IP(["abba", "qrst"], ["mnop"]), True),
	(IP(["abcd", "xyyx"], ["bddb"]), False),
	(IP(["aaaa", "tyui"], ["qwer"]), False),
	(IP(["ioxxoj", "zxcvbn"], ["asdfgh"]), True)
])
def test_supports_TLS(ip: IP, expected: bool):
	assert supports_TLS(ip) == expected

@pytest.mark.parametrize("ip, expected", [
	(IP(["aba", "xyz"], ["bab"]), True),
	(IP(["xyx", "xyx"], ["xyx"]), False),
	(IP(["aaa", "eke"], ["kek"]), True),
	(IP(["zazbz", "cdb"], ["bzb"]), True)
])
def test_supports_SSL(ip: IP, expected: bool):
	assert supports_SSL(ip) == expected
