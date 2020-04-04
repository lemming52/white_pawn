"""
Given an integer n. Each number from 1 to n is grouped according to the sum of its digits.

Return how many groups have the largest size.

Input: n = 13
Output: 4
Explanation: There are 9 groups in total, they are grouped according sum of its digits of numbers from 1 to 13:
[1,10], [2,11], [3,12], [4,13], [5], [6], [7], [8], [9]. There are 4 groups with largest size.
"""

class Solution:
    def countLargestGroup(self, n: int) -> int:
        counts = {}
        largestSize = 0
        for i in range(1, n + 1):
            digitSum = sumDigits(i)
            if digitSum in counts:
                counts[digitSum] += 1
            else:
                counts[digitSum] = 1
            if counts[digitSum] > largestSize:
                largestSize = counts[digitSum]
        counter = 0
        for k, v in counts.items():
            if v == largestSize:
                counter += 1
        return counter


def sumDigits(n: int) -> int:
    total = 0
    while n:
        total += n % 10
        n = n // 10
    return total