from queue import PriorityQueue
from typing import List


class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

    def __lt__(self, other):
        return self.val < other.val


class Solution:
    def mergeKLists(self, lists: List[ListNode]) -> ListNode:
        # iterate over all lists and initialize list_nodes
        list_nodes = PriorityQueue()
        for l in lists:
            list_nodes.put(l)

        # initialize res
        res = head = ListNode
        while not list_nodes.empty():
            list_node = list_nodes.get()
            if not res:
                res = list_node
                head = res
            else:
                res.next = list_node
                res = res.next
            # if the next thing is not null, put it in the heap
            if list_node.next:
                list_nodes.put(list_node.next)

        head = head.next
        return head



sol = Solution()
l5 = ListNode(5)
ll4 = ListNode(4)
ll4.next = l5
ll1 = ListNode(1)
ll1.next = ll4

l1 = ListNode(1)
l3 = ListNode(3)
l4 = ListNode(4)
l3.next = l4
l1.next = l3

l6 = ListNode(6)
l2 = ListNode(2)
l2.next = l6

sol.mergeKLists([ll1, l1, l2])
