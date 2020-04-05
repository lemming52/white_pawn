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

class QueueHeap(object):
    def __init__(self):
        self.heap = []
        self.size = 0

    def parent(self, i):
        return i / 2

    def left(self, i):
        return 2 * i

    def right(self, i):
        return 2 * i + 1

    def min_heapify(self, i):
        left = self.left(i)
        right = self.right(i)
        smallest = i
        if left <= self.size and self.heap[left] < self.heap[i]:
            smallest = left
        if right <= self.size and self.heap[right] < self.heap[smallest]:
            smallest = right
        if smallest != i:
            self.heap[i], self.heap[smallest] = self.heap[smallest], self.heap[i]
            self.min_heapify(smallest)

    def insert(self, key):
        self.size += 1
        self.heap.append(10**10)
        self.decreaseKey(self.size - 1, key)

    def decreaseKey(self, i, key):
        self.heap[i] = key
        while i > 1 and self.heap[self.parent(i)] > self.heap[i]:
            self.heap[self.parent(i)],  self.heap[i] = self.heap[i], self.heap[self.parent(i)]
            i = self.parent(i)

    def heapMin(self):
        return self.heap[0]

    def delete(self, key):
        for i, n in enumerate(self.heap):
            if n == key:
                self.heap[i] = self.heap.pop()
                self.size -= 1
                self.min_heapify(i)
                break

if __name__ == '__main__':
    heap = QueueHeap()
    t = int(input())
    for line in range(t):
        values = input().split()

        if len(values) == 2:
            if values[0] == '1':
                heap.insert(values[1])
            else:
                heap.delete(values[1])
        elif len(values) == 1:
            print(heap.heapMin)


