"""
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:
"""
from typing import Dict

class Solution:
    def convert(self, s: str, numRows: int) -> str:
        if numRows == 1:
            return s
        results = {i: "" for i in range(0, numRows)}
        marker = 0
        reset = 2 * numRows - 2
        row_map = get_row_map(numRows, reset)
        for i, char in enumerate(s):
            row = row_map[i % reset]
            results[row] += char
            marker += 1
        return "".join([x for k, x in results.items()])


def get_row_map(numRows, reset: int) -> Dict[int, int]:
    row_map = {i: i for i in range(0, numRows)}
    for i in range(numRows, reset):
        row_map[i] = reset - i
    return row_map