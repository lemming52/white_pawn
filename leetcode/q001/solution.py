class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        found = {}
        for i, n in enumerate(nums):
            complement = target - n
            if complement in found:
                return [found[complement], i]
            found[n] = i
