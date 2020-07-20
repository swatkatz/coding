from queue import PriorityQueue
from typing import List


class Solution:
    def topkFrequent(self, nums: List[int], k: int) -> List[int]:
        d = dict()
        res = []
        # frequency of occurrence of each element in nums
        for e in nums:
            if e in d:
                d[e] = d[e] + 1
            else:
                d[e] = 1
        pq = PriorityQueue()
        for key, value in d.items():
            pq.put((-value, key))

        for i in range(0, k):
            res.append(pq.get()[1])
        return res


sol = Solution()
print(sol.topkFrequent([1, 2], 2))

