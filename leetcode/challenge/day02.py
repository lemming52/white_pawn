"""
Write an algorithm to determine if a number is "happy".

A happy number is a number defined by the following process: Starting with any positive integer, replace the number by the sum of the squares of its digits,
and repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
Those numbers for which this process ends in 1 are happy numbers.
"""
from typing import Dict

class Solution:
    def isHappy(self, n: int) -> bool:
        found = {}
        return checkNumber(n, found)


def checkNumber(n: int, found: Dict[int, bool]) -> bool:
    if n in found:
        return False
    found[n] = True
    total = 0
    while n:
        digit = n % 10
        total += digit * digit
        n = n // 10
    if total == 1:
        return True
    return checkNumber(total, found)