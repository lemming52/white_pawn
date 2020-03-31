"""
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note:

The solution set must not contain duplicate triplets.
"""

from typing import Dict, List

class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        results = []
        length = len(nums)
        found = {}, {}
        for i, a in enumerate(nums[:-2]):
            for j in range(i+1, length-1):
                b = nums[j]
                for k in range(j+1, length):
                    c = nums[k]
                    print(a, i, b, j, c, k)
                    if (a + b + c) == 0:
                        res = [a, b, c]
                        key = "".join(sorted(res))
                        if not key in found:
                            results.append(res)
                            found[key] = True
        return results


def alreadyFound(a, b, c: int, products, sums: Dict[int, bool]):
    return (a * b * c in products and a + b + c in sums)