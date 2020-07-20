import random


class RandomizedSet:
    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.stack = list()
        self.d = dict()

    def insert(self, val: int) -> bool:
        """
        Inserts a value to the set. Returns true if the set did not already contain the specified element.
        """
        # val is present in d
        if val in self.d:
            return False

        # add to stack
        self.stack.append(val)
        # add to d
        self.d[val] = len(self.stack) - 1
        return True

    def remove(self, val: int) -> bool:
        """
        Removes a value from the set. Returns true if the set contained the specified element.
        """
        # set doesn't contain this element
        if not (val in self.d):
            return False

        # if set contains then perform the put the last element in index location and pop
        index = self.d[val]
        tos = self.stack[len(self.stack) - 1]
        self.stack[index] = tos
        self.d[tos] = index

        # remove the last element from stack and d
        self.stack.pop()
        self.d.pop(val)
        return True

    def getRandom(self) -> int:
        """
        Get a random element from the set.
        """
        index = random.randrange(0, len(self.stack))
        return self.stack[index]


