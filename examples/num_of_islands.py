from typing import List


class Solution:
    def dfs(self, a, v, m, n, i, j):
        # outside the matrix
        if i < 0 or j < 0 or i == m or j == n:
            return
        # current cell is a 1, hence continue doing dfs
        if a[i][j] == 1 and v[i][j] != 1:
            v[i][j] = 1
            self.dfs(a, v, m, n, i + 1, j)
            self.dfs(a, v, m, n, i - 1, j)
            self.dfs(a, v, m, n, i, j + 1)
            self.dfs(a, v, m, n, i, j - 1)

    # Brute force, didn't work
    def numIslands2BruteForce(self, m: int, n: int, positions: List[List[int]]) -> List[int]:
        a = [[0] * n for _ in range(m)]
        res = list()
        for pos in positions:
            # set the position to be 1
            a[pos[0]][pos[1]] = 1
            # initialize visited
            v = [[0] * n for _ in range(m)]

            # iterate over the matrix to calculate the number of islands for this position
            count = 0
            for i in range(len(a)):
                for j in range(len(a[i])):
                    # if already visited, continue
                    if v[i][j] == 1:
                        continue
                    if a[i][j] == 1:
                        self.dfs(a, v, m, n, i, j)
                        count = count + 1
            res.append(count)
        return res

    def get_adjacent_keys(self, pos, m, n) -> List[int]:
        res = list()
        row = pos[0]
        col = pos[1]
        if row - 1 >= 0:
            res.append(self.convert_point_to_key([row - 1, col], n))
        if col - 1 >= 0:
            res.append(self.convert_point_to_key([row, col - 1], n))
        if row + 1 < m:
            res.append(self.convert_point_to_key([row + 1, col], n))
        if col + 1 < n:
            res.append(self.convert_point_to_key([row, col + 1], n))
        return res

    def convert_point_to_key(self, point, n):
        return point[0] * n + point[1]

    # Try something else
    def numIslands2(self, m: int, n: int, positions: List[List[int]]) -> List[int]:
        # point to island_id
        islands = dict()
        res = list()
        num_islands = 0
        island_counter = 0
        for pos in positions:
            island_ids = set()
            # get key
            curr_pos_key = self.convert_point_to_key(pos, n)
            adjacent_keys = self.get_adjacent_keys(pos, m, n)

            # find the min island_id
            for key in adjacent_keys:
                # found an adjacent key in islands
                if key in islands:
                    island_ids.add(islands[key])

            # island_ids now contain all the island_ids for the adjacent cells
            # no adjacent cell belongs to an island
            if len(island_ids) == 0:
                num_islands = num_islands + 1
                island_counter = island_counter + 1
                islands[curr_pos_key] = island_counter
            elif len(island_ids) == 1:
                islands[curr_pos_key] = island_ids.pop()
            else:
                # multiple islands need to be connected by this one and their ids changed
                root_island_id = island_ids.pop()
                island_ids.add(root_island_id)
                for point, island_id in islands.items():
                    if island_id in island_ids:
                        islands[point] = root_island_id
                islands[curr_pos_key] = root_island_id
                num_islands = num_islands - len(island_ids) + 1
            res.append(num_islands)

        return res


sol = Solution()
print(sol.numIslands2(3, 3, [[0,0],[2,0],[0,1],[2,1],[0,2],[2,2],[0,1],[1,2]]))
