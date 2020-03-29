
MAX_SUBSTRING_LENGTH = 26

class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        """
        for a given string, at each starting position within the string,
            check how far you can go without repeating a character
        """
        max_length = 0
        total_length = len(s)

        for i in range(0, total_length):
            if max_length > (total_length - i):
                return max_length
            length = findUniqueLength(s[i:])
            if length > max_length:
                max_length = length
            if max_length == MAX_SUBSTRING_LENGTH:
                return max_length
        return max_length

def findUniqueLength(s: str) -> int:
    chars = {}
    for char in s:
        if char in chars:
            return len(chars)
        chars[char] = True
    return len(chars)

