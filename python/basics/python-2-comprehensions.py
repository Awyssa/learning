# --- List Comprehension
# Instead of:
squares = []
for x in range(5):
    squares.append(x**2)

# You can write:
squares = [x**2 for x in range(5)]  # [0, 1, 4, 9, 16]

# Adding conditions, get only even squares:
even_squares = [x**2 for x in range(10) if x % 2 == 0]  # [0, 4, 16, 36, 64]

# Using with Strings:
message = "Hello World"
vowels = [char for char in message if char.lower() in "aeiou"]  # ['e', 'o', 'o']

# Nested List Comprehension, create a 3x3 matrix:
matrix = [[i + j for j in range(3)] for i in range(0, 9, 3)]
# Result: [[0, 1, 2], [3, 4, 5], [6, 7, 8]]

# Multiple Conditions, numbers divisible by both 2 and 3:
numbers = [x for x in range(20) if x % 2 == 0 if x % 3 == 0]  # [0, 6, 12, 18]

# If-Else in List Comprehension:
# Replace numbers: negative becomes zero, positive stays the same
numbers = [-4, -2, 0, 2, 4]
processed = [x if x > 0 else 0 for x in numbers]  # [0, 0, 0, 2, 4]

# --- Generator Expressions
# Generator expressions are similar to list/set comprehensions but are more memory-efficient since they generate values on-the-fly.

# List comprehension - creates entire list in memory
numbers_list = [x for x in range(1000000)]  # Immediately stores all numbers

# Generator expression - creates values as needed
numbers_gen = (x for x in range(1000000))  # Creates generator object
# Values are generated one at a time when needed


# Memory efficient way to process large files
def read_large_file(file_path):
    with open(file_path) as file:
        # Generate one line at a time instead of loading entire file
        line_lengths = (len(line.strip()) for line in file)
        for length in line_lengths:
            print(length)


# Generate even numbers
even_nums = (x for x in range(10) if x % 2 == 0)
print(list(even_nums))  # [0, 2, 4, 6, 8]

# Note: once consumed, generator is exhausted
print(list(even_nums))  # [] - generator is empty now


# Finding average without storing all numbers
def get_average(numbers):
    total = sum(num for num in numbers)
    count = sum(1 for _ in numbers)  # Counting items
    return total / count if count > 0 else 0
