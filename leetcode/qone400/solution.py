"""
Given a string s and an integer k. You should construct k non-empty palindrome strings using all the characters in s.

Return True if you can use all the characters in s to construct k palindrome strings or False otherwise.

Input: s = "annabelle", k = 2
Output: true
Explanation: You can construct two palindromes using all characters in s.
Some possible constructions "anna" + "elble", "anbna" + "elle", "anellena" + "b"
"""

class Solution:
    def canConstruct(self, s: str, k: int) -> bool:
        if k > len(s):
            return False
        characterCounts = {}
        for char in s:
            if char in characterCounts:
                characterCounts[char] += 1
            else:
                characterCounts[char] = 1

        odd = 0
        for char, count in characterCounts.items():
            odd += count % 2
        return k >= odd

