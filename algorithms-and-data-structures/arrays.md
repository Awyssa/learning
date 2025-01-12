-- Static Array:
A static array is a fixed-size array where the size is determined at compile time and cannot be changed during program execution. When you declare a static array, you must specify its size upfront, and that memory space is allocated permanently.
Here are the key characteristics:

Example in C
int numbers[5]; // Static array of size 5

- Fixed Size: Once declared, you cannot add or remove elements to change its length. If you need to store more elements than the declared size, you'll need to create a new array.
- Memory Allocation: The memory is allocated on the stack (rather than the heap) and remains fixed throughout the program's execution.
- Fast Access: Since the memory is contiguous and the size is known, accessing elements is very fast - O(1) time complexity.
- Memory Efficiency: The memory usage is predictable and constant, but it might waste memory if you allocate more space than needed.

This contrasts with dynamic arrays (like ArrayList in Java or vector in C++), which can grow or shrink during runtime as needed.


-- Dynamic Array:
A dynamic array is a resizable array that can grow or shrink in size during program execution. Unlike static arrays, dynamic arrays automatically handle memory allocation and reallocation when elements are added or removed.
Here's how they work:

Initial Allocation: When first created, the array allocates some initial capacity (often larger than immediately needed).

In Python, lists are dynamic arrays
numbers = []  # Creates a dynamic array with initial capacity

Growth Strategy: When the array fills up and a new element is added:
- Creates a new, larger array (typically 1.5x or 2x the current size)
- Copies existing elements to the new array
- Deallocates the old array
- Adds the new element

For example, in pseudocode:
    def append(element):
        if size == capacity:
            # Create new array with double capacity
            new_capacity = capacity * 2
            new_array = allocate_memory(new_capacity)
            # Copy elements
            copy_elements(old_array, new_array)
            # Switch to new array
            deallocate(old_array)
            array = new_array
        
        array[size] = element
        size += 1

Key characteristics:
- Flexible size
- Amortized O(1) time for append operations
- More memory overhead than static arrays
- Slightly slower access times due to additional memory management
- Common implementations include Python's list, Java's ArrayList, and C++'s vector

The main advantage is convenience and flexibility, as you don't need to predict the exact size needed beforehand.

-- Array methods
Here are the most common array methods in Python (using lists, which are Python's dynamic arrays):

1. `append(x)`: Adds an element at the end of the list
   numbers = [1, 2, 3]
   numbers.append(4)  # Result: [1, 2, 3, 4]

2. `extend(iterable)`: Adds all elements from an iterable to the end
   numbers.extend([5, 6])  # Result: [1, 2, 3, 4, 5, 6]

3. `insert(i, x)`: Inserts element x at position i
   numbers.insert(0, 0)  # Result: [0, 1, 2, 3, 4, 5, 6]

4. `remove(x)`: Removes the first occurrence of value x
   numbers.remove(3)  # Result: [0, 1, 2, 4, 5, 6]

5. `pop([i])`: Removes and returns element at index i (or last element if i not specified)
   last = numbers.pop()  # Removes and returns 6

6. `clear()`: Removes all elements
   numbers.clear()  # Result: []

7. `index(x)`: Returns index of first occurrence of x
   numbers = [1, 2, 3, 2]
   pos = numbers.index(2)  # Returns 1

8. `count(x)`: Returns number of occurrences of x
   twos = numbers.count(2)  # Returns 2

9. `sort()`: Sorts the list in-place
   numbers.sort()  # Result: [1, 2, 2, 3]

10. `reverse()`: Reverses elements in-place
    numbers.reverse()  # Result: [3, 2, 2, 1]

11. `copy()`: Returns a shallow copy of the list
    new_numbers = numbers.copy()