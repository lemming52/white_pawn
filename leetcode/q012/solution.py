"""
For example, two is written as II in Roman numeral, just two one's added together. Twelve is written as, XII, which is simply X + II. The number twenty seven is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.
Given an integer, convert it to a roman numeral. Input is guaranteed to be within the range from 1 to 3999.
"""

from typing import Dict

CHARACTER_MAP = {
    0: {
        "0": "",
        "1": "I",
        "2": "II",
        "3": "III",
        "4": "IV",
        "5": "V",
        "6": "VI",
        "7": "VII",
        "8": "VIII",
        "9": "IX"
    },
    1: {
        "0": "",
        "1": "X",
        "2": "XX",
        "3": "XXX",
        "4": "XL",
        "5": "L",
        "6": "LX",
        "7": "LXX",
        "8": "LXXX",
        "9": "XC"
    },
    2: {
        "0": "",
        "1": "C",
        "2": "CC",
        "3": "CCC",
        "4": "CD",
        "5": "D",
        "6": "DC",
        "7": "DCC",
        "8": "DCCC",
        "9": "CM"
    },
    3: {
        "0": "",
        "1": "M",
        "2": "MM",
        "3": "MMM",
    }

}

class Solution:
    def intToRoman(self, num: int) -> str:
        num = str(num)
        max_power = len(num) - 1
        res = []
        for i, digit in enumerate(num):
            res.append(convert_digit(digit, max_power - i, CHARACTER_MAP))
        return "".join(res)


def convert_digit(digit: str, power: int, characters: Dict[int, Dict[str, str]]) -> str:
    return characters[power][digit]