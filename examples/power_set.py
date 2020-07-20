import math


class Solution:
    def choose_val(self, s: [], i: int, ps: set):
        # we have looked at all elements
        if i == len(s):
            print(ps)
            return

        val = s[i]
        self.choose_val(s, i + 1, ps)
        ps.add(val)
        self.choose_val(s, i + 1, ps)
        ps.remove(val)

    def recursive_power_set(self, s: []):
        self.choose_val(s, 0, set())

    def power_set(self, s: []):
        n = len(s)
        perms = int(math.pow(2, n))
        print(perms)
        for pos in range(0, perms):
            ps = set()
            print("new_pos : ", pos)
            for i in range(0, n):
                if pos & (1 << i):
                    ps.add(s[i])
            print(ps)


sol = Solution()
sol.recursive_power_set([1, 2, 3])
