"""
In a given grid, each cell can have one of three values:

the value 0 representing an empty cell;
the value 1 representing a fresh orange;
the value 2 representing a rotten orange.
Every minute, any fresh orange that is adjacent (4-directionally) to a rotten orange becomes rotten.

Return the minimum number of minutes that must elapse until no cell has a fresh orange.  If this is impossible, return -1 instead.
"""
class Solution:
    def orangesRotting(self, grid: List[List[int]]) -> int:
        lenX, lenY = len(grid), len(grid[0])
        counter = 0
        rottable = []
        for i, row in enumerate(grid):
            for j, value in enumerate(row):
                if value == 2:
                    rottable.append([i, j])
        while rottable:
            nextRot = []
            for coords in rottable:
                nextRot.extend(getRottableNeighbours(coords[0], coords[1], lenX, lenY, grid))
            rottable = nextRot
            for coords in rottable:
                grid[coords[0]][coords[1]] = 2
            if rottable:
                counter += 1

        if any(1 in row for row in grid):
            return -1
        return counter


def getRottableNeighbours(x, y, lenX, lenY: int, grid: List[List[int]]) -> List[List[int]]:
    neighbours = [[x+1, y], [x-1, y], [x, y+1], [x, y-1]]
    rottable = []
    for coords in neighbours:
        nX, nY = coords
        if nX < 0 or nX >= lenX:
            continue
        if nY < 0 or nY >= lenY:
            continue
        if grid[nX][nY] == 1:
            rottable.append([nX, nY])
    return rottable
