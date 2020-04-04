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
        counter = -1
        lenX, lenY = len(grid), len(grid[0])
        rot = True
        while rot:
            print(grid, rot)
            changed = False
            for i in range(lenX):
                for j in range(lenY):
                    if grid[i][j] == 2:
                        changed = rotFreshNeighbours(i, j, grid) or changed
            rot = changed
            counter += 1
        for i in range(lenX):
            for j in range(lenY):
                if grid[i][j] == 1:
                    return -1
        return counter


def rotFreshNeighbours(x, y: int, grid: List[List[int]]) -> bool:
    rotted = False
    neighbours = [[x+1, y], [x-1, y], [x, y+1], [x, y-1]]
    for coords in neighbours:
        nX, nY = coords
        if nX < 0 or nX >= len(grid):
            continue
        if nY < 0 or nY >= len(grid[0]):
            continue
        if grid[nX][nY] == 1:
            rotted = True
            grid[nX][nY] = 2
    return rotted
