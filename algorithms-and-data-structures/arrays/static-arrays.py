# A static array is a fixed-size array where the size is determined at compile time and cannot be 
# changed during program execution. When you declare a static array, 
# you must specify its size upfront, and that memory space is allocated permanently.

# 1. Sorted set
class Solution:
    def removeDuplicates(self, nums: list[int]) -> int:
        unique = sorted(set(nums))
        nums[:len(unique)] = unique
        return len(unique)
      
# 2. Two pointer
class Solution:
    def removeDuplicates(self, nums: list[int]) -> int:
        left_pointer = 1
        for right_pointer in range(1, len(nums)):
            if nums[right_pointer] != nums[right_pointer - 1]:
                nums[left_pointer] = nums[right_pointer]
                left_pointer += left_pointer
        return left_pointer