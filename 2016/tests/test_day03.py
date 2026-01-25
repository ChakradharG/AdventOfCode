import pytest
from day03.solution import part1, part2

@pytest.mark.parametrize("triangles, expected", [
    ([[5, 10, 25]], 0),
    ([[5, 10, 12]], 1),
    ([[5, 10, 25], [5, 10, 12]], 1)
])
def test_part1(triangles, expected):
    assert part1(triangles) == expected

def test_part2():
    triangles = [
        [101, 301, 501],
        [102, 302, 502],
        [103, 303, 503],
        [201, 401, 601],
        [202, 402, 602],
        [203, 403, 603]
    ]
    assert part2(triangles) == 6
