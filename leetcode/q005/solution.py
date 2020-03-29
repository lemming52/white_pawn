"""
Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.
"""

class Solution:
    def longestPalindrome(self, s: str) -> str:
        for i in range(len(s), 1, -1):
            for j in range(0, len(s) - i + 1):
                res = checkPalindrome(s[j:j+i])
                if res:
                    return s[j:j+i]
        return s[0]


def checkPalindrome(s: str) -> bool:
    for i in range(0, int(len(s) / 2)):
        if s[i] != s[-(i + 1)]:
            return False
    return True