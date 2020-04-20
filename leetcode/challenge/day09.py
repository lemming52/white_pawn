"""
Given two strings S and T, return if they are equal when both are typed into empty text editors. # means a backspace character.

Note that after backspacing an empty text, the text will continue empty.

Example 1:

Input: S = "ab#c", T = "ad#c"
Output: true
Explanation: Both S and T become "ac".
Example 2:

Input: S = "ab##", T = "c#d#"
Output: true
Explanation: Both S and T become "".
Example 3:

Input: S = "a##c", T = "#a#c"
Output: true
Explanation: Both S and T become "c".
"""

class Solution:
    def backspaceCompare(self, S: str, T: str) -> bool:
        s = []
        backspace = '#'
        for c in S:
            if c == backspace:
                if len(s) > 0:
                    s.pop(-1)
            else:
                s.append(c)

        t = []
        for c in T:
            if c == backspace:
                if len(t) > 0:
                    t.pop(-1)
            else:
                t.append(c)

        return ''.join(s) == ''.join(t)



