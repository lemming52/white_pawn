# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def middleNode(self, head: ListNode) -> ListNode:
        d = 1
        current = head.next
        middle = head
        while current:
            d += 1
            if d % 2 == 0:
                middle = middle.next
            current = current.next

        return middle