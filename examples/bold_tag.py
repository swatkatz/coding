# Given a string s and a list of strings dict, you need to add a closed pair of bold tag <b> and
# </b> to wrap the substrings in s that exist in dict. If two such substrings overlap, you need
# to wrap them together by only one pair of closed bold tag. Also, if two substrings wrapped by
# bold tags are consecutive, you need to combine them
from typing import List

class Solution:
    def find_all(self, s: str, w: str) -> [int]:
        res = []
        val = 0
        while True:
            val = s.find(w, val)
            if val == -1:
                break
            res.append([val, val + len(w)])
            val = val + 1
        return res

    def max(self, i: int, j: int) -> int:
        if i < j:
            return j
        return i

    def add_bold_tag(self, s: str, dict: List[str]) -> str:
        intercepts = []
        for d in dict:
            intercepts += self.find_all(s, d)
        intercepts.sort(key=lambda intercept: intercept[0])

        res = []
        i = 0
        while i < len(intercepts):
            curr = intercepts[i]
            if len(res) == 0:
                res.append(curr)
            else:
                # get the last result and compare
                j = len(res) - 1
                # overlap
                if res[j][0] <= curr[0] <= res[j][1]:
                    val = self.max(res[j][1], curr[1])
                    res[j][1] = val
                else:
                    # non-overlap
                    res.append(curr)
            i += 1

        offset = 0
        for st, en in res:
            s = "{0}<b>{1}</b>{2}".format(
                s[:st+offset],
                s[st+offset:en+offset],
                s[en+offset:])
            # because the same string is being modified, we now have extra characters
            offset += len("<b></b>")

        return s



sol = Solution()
sol.add_bold_tag("abcxyz123",  ["abc","123"])
