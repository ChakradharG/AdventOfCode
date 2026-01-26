import pytest
from day04.solution import part1, Room, is_real, decrypt

def test_is_real():
    assert is_real(Room("aaaaa-bbb-z-y-x-", 123, "abxyz")) == True
    assert is_real(Room("a-b-c-d-e-f-g-h-", 987, "abcde")) == True
    assert is_real(Room("not-a-real-room-", 404, "oarel")) == True
    assert is_real(Room("totally-real-room-", 200, "decoy")) == False

def test_decrypt():
    name = "qzmt-zixmtkozy-ivhz"
    shifted = "".join([decrypt(c, 343) for c in name])
    assert shifted == "very encrypted name"

def test_part1():
    rooms = [
        Room("aaaaa-bbb-z-y-x-", 123, "abxyz"),
        Room("a-b-c-d-e-f-g-h-", 987, "abcde"),
        Room("not-a-real-room-", 404, "oarel"),
        Room("totally-real-room-", 200, "decoy")
    ]
    assert part1(rooms) == 1514
