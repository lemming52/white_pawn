"""
Markov takes out his Snakes and Ladders game, stares at the board and wonders: "If I can always roll the die to whatever number I want, what would be the least number of rolls to reach the destination?"

Rules The game is played with a cubic die of  faces numbered  to .

Starting from square , land on square  with the exact roll of the die. If moving the number rolled would place the player beyond square , no move is made.

If a player lands at the base of a ladder, the player must climb the ladder. Ladders go up only.

If a player lands at the mouth of a snake, the player must go down the snake and come out through the tail. Snakes go down only.

Function Description

Complete the quickestWayUp function in the editor below. It should return an integer that represents the minimum number of moves required.

quickestWayUp has the following parameter(s):

ladders: a 2D integer array where each  contains the start and end cell numbers of a ladder
snakes: a 2D integer array where each  contains the start and end cell numbers of a snake
Input Format

The first line contains the number of tests, .

For each testcase:
- The first line contains , the number of ladders.
- Each of the next  lines contains two space-separated integers, the start and end of a ladder.
- The next line contains the integer , the number of snakes.
- Each of the next  lines contains two space-separated integers, the start and end of a snake.

Constraints



The board is always  with squares numbered  to .
Neither square  nor square  will be the starting point of a ladder or snake.
A square will have at most one endpoint from either a snake or a ladder.

Output Format

For each of the t test cases, print the least number of rolls to move from start to finish on a separate line. If there is no solution, print -1.

Sample Input

2
3
32 62
42 68
12 98
7
95 13
97 25
93 37
79 27
75 19
49 47
67 17
4
8 52
6 80
26 42
2 72
9
51 19
39 11
37 29
81 3
59 5
79 23
53 7
43 33
77 21
Sample Output

3
5
"""
#!/bin/python3

import math
import os
import random
import re
import sys
from collections import defaultdict

class Graph:

    def __init__(self):
        self.neighbours=defaultdict(list)

    def add_edge(self,u,v,dist):
        if dist >= 0:
            self.neighbours[u].append([v, dist])
        else:
            self.neighbours[u] = [[v, 0]]

    def add_node(self, a):
        self.nodes[a] = []

    def shortest_path(self):
        queue = []
        visited = {}
        queue.append([0, 0])

        while queue:
            index, rolls = queue.pop(0)
            if index in visited:
                continue
            visited[index] = rolls

            if index == 99:
                break

            for neighbour in self.neighbours[index]:
                if neighbour[0] not in visited:
                    queue.append([neighbour[0], rolls + neighbour[1]])

        if 99 in visited:
            return visited[99]
        else:
            return -1

# Complete the quickestWayUp function below.
def quickestWayUp(ladders, snakes):
    g = Graph()
    for i in range(99):
        for j in range(1, 7):
            g.add_edge(i, i + j, 1)

    for ladder in ladders:
        g.add_edge(ladder[0]-1, ladder[1]-1, 0)

    for snake in snakes:
        g.add_edge(snake[0]-1, snake[1]-1, 0)

    return g.shortest_path()




if __name__ == '__main__':
    fptr = sys.stdout

    t = int(input())

    for t_itr in range(t):
        n = int(input())

        ladders = []

        for _ in range(n):
            ladders.append(list(map(int, input().rstrip().split())))

        m = int(input())

        snakes = []

        for _ in range(m):
            snakes.append(list(map(int, input().rstrip().split())))

        result = quickestWayUp(ladders, snakes)

        fptr.write(str(result) + '\n')

    fptr.close()
