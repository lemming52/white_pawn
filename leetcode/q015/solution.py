"""
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note:

The solution set must not contain duplicate triplets.
"""

from typing import Dict, List

class Solution:

    def threeSum(self, nums: List[int]) -> List[List[int]]:
        nums = sorted(nums)
        found = {}
        results = []
        for i, a in enumerate(nums):
            j = i + 1
            k = len(nums) - 1
            if a > 0:
                break
            while j < k:
                b = nums[j]
                c = nums[k]
                total = a + b + c
                if total == 0:
                    key = f"{a}{b}{c}"
                    if not key in found:
                        found[key] = True
                        results.append([a, b, c])
                    j += 1
                    k -= 1
                elif total > 0:
                    k -= 1
                else:
                    j += 1
        return results