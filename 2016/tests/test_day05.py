from day05.solution import starts_with_5_zeros, is_valid, part1, part2

def test_starts_with_5_zeros():
	assert starts_with_5_zeros("000005") == True
	assert starts_with_5_zeros("1") == False
	assert starts_with_5_zeros("1000005") == False
	assert starts_with_5_zeros("00000000001") == True

def test_is_valid():
	assert is_valid("", set()) == False
	assert is_valid("12345", set()) == False
	assert is_valid("000005", set()) == False
	assert is_valid("00000", set()) == False
	assert is_valid("0000506", set()) == False
	assert is_valid("0000516", set("1")) == False
	assert is_valid("0000016", set("1")) == True
	assert is_valid("0000036", set("12")) == False
	assert is_valid("0000026", set("12")) == True

def test_part1():
	assert part1("abc") == "18f47a30"

def test_part2():
	assert part2("abc") == "05ace8e3"
