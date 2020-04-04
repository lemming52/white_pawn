"""
Given a binary tree, return the vertical order traversal of its nodes values.

For each node at position (X, Y), its left and right children respectively will be at positions (X-1, Y-1) and (X+1, Y-1).

Running a vertical line from X = -infinity to X = +infinity, whenever the vertical line touches some nodes, we report the values of the nodes in order from top to bottom (decreasing Y coordinates).

If two nodes have the same position, then the value of the node that is reported first is the value that is smaller.

Return an list of non-empty reports in order of X coordinate.  Every report will have a list of values of nodes.
"""

# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Output:
    def __init__(self):
        self.output = []
        self.minimumX = 0
        self.knownX = {}

    def Add(self, x, y, val: int) -> None:
        if x not in self.knownX:
            self.output.insert(x - self.minimumX, [val])
            self.knownX[x] = True
            if x < self.minimumX:
                self.minimumX = x
        else:
            self.output[x - self.minimumX].append(val)

    def Export(self) -> List[List[int]]:
        return [sorted(x) for x in self.output]


class Solution:
    def verticalTraversal(self, root: TreeNode) -> List[List[int]]:
        output = Output()
        traverseAndSort(root, 0, 0, output)
        return output.Export()


def traverseAndSort(node: TreeNode, x, y: int, output: Output) -> None:
    output.Add(x, y, node.val)
    if node.left:
        traverseAndSort(node.left, x-1, y-1, output)
    if node.right:
        traverseAndSort(node.right, x+1, y-1, output)

