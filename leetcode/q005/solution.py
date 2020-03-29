"""
Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.
"""

class Solution:
    def longestPalindrome(self, s: str) -> str:
        total_length = len(s)
        if total_length == 0:
            return ""
        result = ""
        for i in range(0, total_length):
            odd_length = checkCentre(s, i, i)
            even_length = checkCentre(s, i, i + 1)
            length = max([odd_length, even_length])
            if length > len(result):
                result = s[i - int((length - 1) / 2): i + int(length / 2) + 1]
        return result


def checkCentre(s: str, left, right: int) -> int:
    i = left
    j = right
    length = 1
    while i >= 0 and j < len(s):
        if s[i] != s[j]:
            return length - 1
        i -= 1
        j += 1
        length = j - i
    return length - 1
