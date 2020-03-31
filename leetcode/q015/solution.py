"""
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note:

The solution set must not contain duplicate triplets.
"""

from typing import Dict, List

class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        neg, zero, pos = [], [], []
        results = []
        for n in nums:
            if n < 0:
                neg.append(n)
            elif n > 0:
                pos.append(n)
            else:
                zero.append(0)

        n_zero = len(zero)
        if n_zero > 0:
            if n_zero >= 3:
                results.append([0,0,0])
            check_with_zero(neg, pos, results)
        for i, a in enumerate(neg):
            check_two_pos(pos, a * -1, results)
            check_one_pos(neg, pos, a * -1, results)
        return results

def check_with_zero(neg, pos: List[int], res: List[List[int]]) -> None:
    found = {}
    for a in neg:
        for b in pos:
            if a + b == 0:
                key = f'{a}{b}'
                if not key in found:
                    res.append([a, 0, b])
                    found[key] = True

def check_two_pos(pos: List[int], a: int, res: List[List[int]]) -> None:
    found = {}
    for i, b in enumerate(pos):
        for c in pos[i+1:]:
            if (b + c) == a:
                key = f'{b}{c}'
                if not key in found:
                    res.append([a * -1, b, c])
                    found[key] = True

def check_one_pos(neg, pos: List[int], a: int, res: List[List[int]]) -> None:
    found = {}
    for b in neg:
        for c in pos:
            if b + c == a:
                key = f'{b}{c}'
                if not key in found:
                    res.append([a * -1, b, c])
                    found[key] = True