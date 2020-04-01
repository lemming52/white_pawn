"""
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.

A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.


Input: "23"
Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
"""

CHARACTER_MAP = {
    "2": ["a", "b", "c"],
    "3": ["d", "e", "f"],
    "4": ["g", "h", "i"],
    "5": ["j", "k", "l"],
    "6": ["m", "n", "o"],
    "7": ["p", "q", "r", "s"],
    "8": ["t", "u", "v"],
    "9": ["w", "x", "y", "z"],
}

class Solution:
    def letterCombinations(self, digits: str) -> List[str]:
        if digits == "":
            return []
        results = []
        addOptions("", digits, results)
        return results


def addOptions(option: str, number: str, results: List[str]):
    if number == "":
        results.append(option)
        return
    for character in CHARACTER_MAP[number[0]]:
        addOptions(option + character, number[1:], results)
