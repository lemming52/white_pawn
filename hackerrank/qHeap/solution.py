"""
This question is designed to help you get a better understanding of basic heap operations.
You will be given queries of  types:

" " - Add an element  to the heap.
" " - Delete the element  from the heap.
"" - Print the minimum of all the elements in the heap.
NOTE: It is guaranteed that the element to be deleted will be there in the heap. Also, at any instant, only distinct elements will be in the heap.

Input Format

The first line contains the number of queries, .
Each of the next  lines contains a single query of any one of the  above mentioned types.

Constraints


Output Format

For each query of type , print the minimum value on a single line.

Sample Input

5
1 4
1 9
3
2 4
3
Sample Output

4
9
"""

import heapq

class MinHeap:
    def __init__(self):
        self.heap = []

    def insert(self, item):
        heapq.heappush(self.heap, item)

    def pop(self):
        return heapq.heappop(self.h)

    def heapMin(self):
        return self.heap[0]

    def delete(self, val):
        self.heap.remove(val)
        heapq.heapify(self.heap)

    def __getitem__(self, item):
        return self.heap[item]

    def __len__(self):
        return len(self.h)

if __name__ == '__main__':
    heap = MinHeap()
    t = int(input())
    for line in range(t):
        values = input().split()

        if len(values) == 2:
            if values[0] == '1':
                heap.insert(values[1])
            else:
                heap.delete(values[1])
        elif len(values) == 1:
            print(heap.heapMin())


