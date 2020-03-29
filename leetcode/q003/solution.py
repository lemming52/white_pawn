class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        """
        for a given string, at each starting position within the string,
            check how far you can go without repeating a character
        """
        max_length = 0
        i = 0
        total_length = len(s)

        while max_length < (total_length - i):
            length, nextpos = findUniqueLength(s[i:])
            i += nextpos
            if length > max_length:
                max_length = length
        return max_length

def findUniqueLength(s: str) -> (int, int):
    chars = {}
    for i, char in range(0, len(s)):
        char = s[i]
        if char in chars:
            nextStart = chars[char]
            return len(chars), nextStart + 1
        chars[char] = i
    return len(chars), len(s)

