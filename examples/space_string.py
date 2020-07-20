import math


class Solution:
    def __init__(self, v: set):
        self.v = v

    def check_words(self, spaced_str: str) -> bool:
        for w in spaced_str.split(" ", -1):
            if w not in self.v:
                return False
        return True

    def get_valid_words(self, given_str: str):
        n = len(given_str)
        final = ""
        perms = int(math.pow(2, n - 1))
        for pos in range(0, perms):
            if self.check_words(final):
                print(final)
            final = given_str[0]
            for i in range(0, n - 1):
                if pos & (1 << i):
                    final += " " + given_str[i + 1]
                else:
                    final += given_str[i + 1]

    def append_str(self, given_str: str, pos: int, final: str):
        if pos == len(given_str):
            if self.check_words(final):
                print(final)
            return
        self.append_str(given_str, pos + 1, final + given_str[pos])
        self.append_str(given_str, pos + 1, final + " " + given_str[pos])

    def get_valid_words_recursive(self, given_str: str):
        self.append_str(given_str, 1, given_str[0])


vocab = {"the", "me", "park", "theme", "pa", "rk"}
sol = Solution(vocab)
print("----non recursive answer----")
sol.get_valid_words("themepark")
print("----recursive answer----")
sol.get_valid_words_recursive("themepark")
