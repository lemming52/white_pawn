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
            if x < self.minimumX:
                self.minimumX = x
            self.output.insert(x - self.minimumX, [[val, y]])
            self.knownX[x] = True
        else:
            self.output[x - self.minimumX].append([val, y])


class Solution:
    def verticalTraversal(self, root: TreeNode) -> List[List[int]]:
        output = Output()
        targets = [[0, 0, root]]
        while len(targets) != 0:
            nextTargets = []
            for target in targets:
                x, y, node = target
                output.Add(x, y, node.val)
                if node.left:
                    nextTargets.append([x-1, y-1, node.left])
                if node.right:
                    nextTargets.append([x+1, y-1, node.right])
            targets = nextTargets
        sortedOutput = [sorted(subset, key=lambda x: (-x[1], x[0])) for subset in output.output]
        results = []
        for subset in sortedOutput:
            results.append([x[0] for x in subset])
        return results

