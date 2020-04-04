"""
Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Example:

Input: [0,1,0,3,12]
Output: [1,3,12,0,0]
Note:

You must do this in-place without making a copy of the array.
Minimize the total number of operations.
"""

class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        i, j = 0, len(nums) - 1
        while j > i:
            print(nums)
            a = nums[i]
            b = nums[j]
            if b == 0:
                j -= 1
                continue
            if a != 0:
                a += 1
                continue
            swap(nums, i, j)
            i += 1
            j -= 1

def swap(nums: List[int], i, j: int) -> None:
    while i < j:
        nums[i] = nums[i+1]
        i += 1
    nums[j] = 0

