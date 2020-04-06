"""
Given a linked list, remove the n-th node from the end of list and return its head.

Example:

Given linked list: 1->2->3->4->5, and n = 2.

After removing the second node from the end, the linked list becomes 1->2->3->5.
Note:

Given n will always be valid.
"""

# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def removeNthFromEnd(self, head: ListNode, n: int) -> ListNode:
        d = 1
        current = head.next
        while current:
            d += 1
            current = current.next

        removal_index = d - n
        if removal_index <= 0:
            return head.next

        counter = 1
        prior = head
        current = head.next
        while counter < removal_index:
            prior = current
            current = prior.next
            counter += 1
        if current.next is None:
            prior.next = None
        else:
            following = current.next
            prior.next = following

        return head




