from typing import List


class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        d = dict()
        for s in strs:
            key = str(''.join(sorted(s)))
            if key in d:
                val = d.get(key)
                val.append(s)

            else:
                d[key] = [s]
        res = []
        for key in d.keys():
            res.append(d[key])

        print(res)


sol = Solution()
sol.groupAnagrams(["eat", "tea", "tan", "ate", "nat", "bat"])
