"""
Given an array nums of n integers and an integer target, find three integers in nums such that the sum is closest to target.
Return the sum of the three integers. You may assume that each input would have exactly one solution.
"""

class Solution:
    def threeSumClosest(self, nums: List[int], target: int) -> int:
        nums = sorted(nums)
        closest = None
        closestDelta = None
        for i in range(len(nums) - 1, 0):
            if nums[i] < target:
                nums = nums[:i]
                break
        for i, a in enumerate(nums[:-2]):
            j = i + 1
            k = len(nums) - 1
            while j < k:
                b = nums[j]
                c = nums[k]
                delta = target - (a + b + c)
                if delta == 0:
                    return target
                elif delta < 0:
                    k -= 1
                else:
                    j += 1
                if closestDelta is None or abs(delta) < closestDelta:
                    closest = a + b + c
                    closestDelta = abs(delta)
        return closest