"""
Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai).
n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container,
such that the container contains the most water.

Note: You may not slant the container and n is at least 2.
"""

class Solution:
    def maxArea(self, height: List[int]) -> int:
        entries = len(height)
        res = 0
        for i in range(entries):
            for j in range(entries - 1, i, -1):
                if height[i] * (j - i) < res:
                    break
                area = getArea(i, height[i], j, height[j])
                if area > res:
                    res = area
        return res

def getArea(x1, y1, x2, y2) -> int:
    height = y1 if y1 < y2 else y2
    return (x2 - x1) * height