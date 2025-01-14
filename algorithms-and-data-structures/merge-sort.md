Merge sort is a highly efficient sorting algorithm that uses a "divide and conquer" strategy.

Here's how merge sort works:

1. Divide: Split the array into smaller subarrays until each subarray contains just one element
2. Conquer: Repeatedly merge these subarrays to produce new sorted subarrays until only one remains
3. The merging process compares elements from two subarrays and combines them in sorted order

Here's a Python implementation to demonstrate:

```python
def merge_sort(arr):
    # Base case - if array has 1 or fewer elements, it's already sorted
    if len(arr) <= 1:
        return arr
    
    # Divide array into two halves
    mid = len(arr) // 2
    left = arr[:mid]
    right = arr[mid:]
    
    # Recursively sort each half
    left = merge_sort(left)
    right = merge_sort(right)
    
    # Merge the sorted halves
    return merge(left, right)

def merge(left, right):
    result = []
    i = j = 0
    
    # Compare elements from both arrays and merge in sorted order
    while i < len(left) and j < len(right):
        if left[i] <= right[j]:
            result.append(left[i])
            i += 1
        else:
            result.append(right[j])
            j += 1
    
    # Add remaining elements from left array, if any
    result.extend(left[i:])
    # Add remaining elements from right array, if any
    result.extend(right[j:])
    
    return result

# Example usage
arr = [64, 34, 25, 12, 22, 11, 90]
sorted_arr = merge_sort(arr)
```

Let's walk through how this sorts [64, 34, 25, 12]:

1. First divide:
   ```
   [64, 34, 25, 12]
   [64, 34] [25, 12]
   [64] [34] [25] [12]
   ```

2. Then merge back while sorting:
   ```
   [64] [34] -> [34, 64]
   [25] [12] -> [12, 25]
   [34, 64] [12, 25] -> [12, 25, 34, 64]
   ```

The time complexity is O(n log n) in all cases, which makes it very efficient. The space complexity is O(n) because it needs additional space to store the subarrays during the sorting process.

Key advantages of merge sort:
- Stable sort (maintains relative order of equal elements)
- Guaranteed O(n log n) performance
- Works well for linked lists
- Predictable performance regardless of input data
