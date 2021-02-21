#!/usr/bin/env python
# -*- coding: utf-8 -*-

#  input =[1,2,3,2,2]
#  out=4

class Solution:
    def atMostK(self, nums, K):
        counter = collections.Counter()
        res = i = 0
        for j in range(len(nums)):
            if counter[nums[j]] == 0:
                K -= 1
            counter[nums[j]] += 1
            while K < 0:
                counter[nums[i]] -= 1
                if counter[nums[i]] == 0:
                    K += 1
                i += 1
            res = max(res, j - i + 1)
        return res

    def totalFruit(self, tree: List[int]) -> int:
        return self.atMostK(tree, 2)


if __name__ == '__main__':
    s = Solution()
    print(s.totalFruit([0,1,2,2]))
