import pytest
from day01.solution import part1, part2

@pytest.mark.parametrize("instructions, expected", [
    (["R2", "L3"], 5),
    (["R2", "R2", "R2"], 2),
    (["R5", "L5", "R5", "R3"], 12),
])
def test_part1(instructions, expected):
    assert part1(instructions) == expected

def test_part2():
    assert part2(["R8", "R4", "R4", "R8"]) == 4
