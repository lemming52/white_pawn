"""
Given a circle represented as (radius, x_center, y_center) and an axis-aligned rectangle represented as (x1, y1, x2, y2),
where (x1, y1) are the coordinates of the bottom-left corner, and (x2, y2) are the coordinates of the top-right corner of the rectangle.

Return True if the circle and rectangle are overlapped otherwise return False.

In other words, check if there are any point (xi, yi) such that belongs to the circle and the rectangle at the same time.

Input: radius = 1, x_center = 0, y_center = 0, x1 = 1, y1 = -1, x2 = 3, y2 = 1
Output: true
Explanation: Circle and rectangle share the point (1,0)
"""
import math

class Solution:
    def checkOverlap(self, radius: int, x_center: int, y_center: int, x1: int, y1: int, x2: int, y2: int) -> bool:
        corners = [[x1, y1], [x1, y2], [x2, y1], [x2, y2]]
        withinYRange = y1 <= y_center <= y2
        withinXRange = x1 <= x_center <= x2
        if withinXRange and withinYRange:"""
Given a circle represented as (radius, x_center, y_center) and an axis-aligned rectangle represented as (x1, y1, x2, y2),
where (x1, y1) are the coordinates of the bottom-left corner, and (x2, y2) are the coordinates of the top-right corner of the rectangle.

Return True if the circle and rectangle are overlapped otherwise return False.

In other words, check if there are any point (xi, yi) such that belongs to the circle and the rectangle at the same time.

Input: radius = 1, x_center = 0, y_center = 0, x1 = 1, y1 = -1, x2 = 3, y2 = 1
Output: true
Explanation: Circle and rectangle share the point (1,0)
"""
import math

class Solution:
    def checkOverlap(self, radius: int, x_center: int, y_center: int, x1: int, y1: int, x2: int, y2: int) -> bool:
        corners = [[x1, y1], [x1, y2], [x2, y1], [x2, y2]]
        withinYRange = y1 <= y_center <= y2
        withinXRange = x1 <= x_center <= x2
        if withinXRange and withinYRange:
            return True
        if withinYRange:
            if abs(x1 - x_center) <=  radius or abs(x2- x_center) <= radius:
                return True
        if withinXRange:
            if abs(y1 - y_center) <=  radius or abs(y2- y_center) <= radius:
                return True
        radius = radius**2
        for corner in corners:
            if distance(corner[0], corner[1], x_center, y_center) < radius:
                return True
        return False

def distance(x1, y1, x2, y2: int) -> float:
    return(x2 - x1)**2 + (y2 - y1)**2
