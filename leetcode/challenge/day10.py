"""
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

push(x) -- Push element x onto stack.
pop() -- Removes the element on top of the stack.
top() -- Get the top element.
getMin() -- Retrieve the minimum element in the stack.

"""

class MinStack:

    def __init__(self):
        """
        initialize your data structure here.
        """
        self.heap = []
        self.minVal = None


    def push(self, x: int) -> None:
        if self.minVal is None:
            self.minVal = x
        elif x <= self.minVal:
            self.heap.append(self.minVal)
            self.minVal = x
        self.heap.append(x)

    def pop(self) -> None:
        if len(self.heap) > 0 and self.heap[-1] == self.minVal:
            self.heap.pop(-1)
            if len(self.heap) > 0:
                self.minVal = self.heap[-1]
            else:
                self.minVal = None

        if len(self.heap) > 0:
            self.heap.pop(-1)

    def top(self) -> int:
        return self.heap[-1]


    def getMin(self) -> int:
        return self.minVal



# Your MinStack object will be instantiated and called as such:
# obj = MinStack()
# obj.push(x)
# obj.pop()
# param_3 = obj.top()
# param_4 = obj.getMin()