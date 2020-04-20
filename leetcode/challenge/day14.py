"""
You are given a string s containing lowercase English letters, and a matrix shift, where shift[i] = [direction, amount]:

direction can be 0 (for left shift) or 1 (for right shift).
amount is the amount by which string s is to be shifted.
A left shift by 1 means remove the first character of s and append it to the end.
Similarly, a right shift by 1 means remove the last character of s and add it to the beginning.
Return the final string after all operations.



Example 1:

Input: s = "abc", shift = [[0,1],[1,2]]
Output: "cab"
Explanation:
[0,1] means shift to left by 1. "abc" -> "bca"
[1,2] means shift to right by 2. "bca" -> "cab"
Example 2:
"""

class Solution:
    def stringShift(self, s: str, shift: List[List[int]]) -> str:
        length = len(s)
        for direction, amount in shift:
            amount = amount % len(s)
            if direction == 1:
                s = s[-amount:] + s[:-amount]
            else:
                s = s[amount:] + s[:amount]
        return s