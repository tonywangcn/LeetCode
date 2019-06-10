"""
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
"""
import unittest


class Solution:
    def twoSum(self, nums, target):
        res = {}
        for k, v in enumerate(nums):
            if res.get(target - v, None) != None:
                return [res[target - v], k]
            else:
                res[v] = k
        return []


class TwoSumTest(unittest.TestCase):
    def test_sum1(self):
        nums = [2, 7, 11, 15]
        target = 9
        self.assertEqual(Solution().twoSum(nums, target), [0, 1])

    def test_sum2(self):
        nums = [2, 7, 11, 15]
        target = 13
        self.assertEqual(Solution().twoSum(nums, target), [0, 2])


if __name__ == "__main__":
    unittest.main()

