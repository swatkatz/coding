from typing import List

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        d = dict()
        for i in range(0, len(nums)):
            num = nums[i]
            pair = target - num
            d[pair] = i

        for i in range(0, len(nums)):
            num = nums[i]
            if num in d and not d[num] == i:
                return [i, d[num]]
        return [-1, -1]

sol = Solution()
print(sol.twoSum([3, 2, 4], 6))
