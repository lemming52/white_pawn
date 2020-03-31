"""
For example, two is written as II in Roman numeral, just two one's added together. Twelve is written as, XII, which is simply X + II. The number twenty seven is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.
Given an integer, convert it to a roman numeral. Input is guaranteed to be within the range from 1 to 3999.
"""

from typing import Dict, List

STRUCTURE_MAP = {
    "0": "",
    "1": "{base}",
    "2": "{base}{base}",
    "3": "{base}{base}{base}",
    "4": "{base}{half}",
    "5": "{half}",
    "6": "{half}{base}",
    "7": "{half}{base}{base}",
    "8": "{half}{base}{base}{base}",
    "9": "{base}{full}"
}

BASIS_MAP = {
    0: ["I", "V", "X"],
    1: ["X", "L", "C"],
    2: ["C", "D", "M"],
    3: ["M", "?", "?"]
}

class Solution:
    def intToRoman(self, num: int) -> str:
        num = str(num)
        max_power = len(num) - 1
        res = []
        for i, digit in enumerate(num):
            res.append(convert_digit(digit, max_power - i, STRUCTURE_MAP, BASIS_MAP))
        return "".join(res)


def convert_digit(digit: str, power: int, structure: Dict[str, str], basis: Dict[int, List[str]]) -> str:
    base, half, full = basis[power]
    return structure[digit].format(base=base, half=half, full=full)