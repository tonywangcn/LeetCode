// Given an array of integers, return indices of the two numbers such that they add up to a specific target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// Example:

// Given nums = [2, 7, 11, 15], target = 9,

// Because nums[0] + nums[1] = 2 + 7 = 9,
// return [0, 1].

/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(nums, target) {
  var ret = {};
  for (let i = 0; i < nums.length; i++) {
    if (typeof ret[target - nums[i]] != "undefined") {
      return [ret[target - nums[i]], i];
    } else {
      ret[nums[i]] = i;
    }
  }
  return [];
};

nums = [2, 7, 11, 15];
target = 13;

console.log(twoSum(nums, target));
