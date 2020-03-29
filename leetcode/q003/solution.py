class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        """
        for a given string, at each starting position within the string,
            check how far you can go without repeating a character
        """
        max_length = 0
        i, j = 0, 0
        total_length = len(s)

        positions = {}

        while j < total_length:
            char = s[j]
            if char in positions:
                if positions.get(char, 0) > i:
                    i = positions[char]
            if j - i + 1 > max_length:
                max_length = j - i + 1
            j += 1
            positions[char] = j
        return max_length
