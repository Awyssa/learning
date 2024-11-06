# Introduction to Python Basics
# This file covers fundamental concepts including variables, functions, and data structures

# -------- Variables and Basic Data Types --------
# Numbers (integers and floats)
age = 25              # Integer
height = 1.75         # Float

# Strings
name = "Alice"
message = 'Hello, World!'  # Single or double quotes work

# Booleans
is_student = True
has_license = False

# -------- Basic Operations --------
# Arithmetic
sum_result = 10 + 5
difference = 10 - 5
product = 10 * 5
division = 10 / 5
integer_division = 10 // 3  # Returns 3 (floors the result)
remainder = 10 % 3         # Returns 1 (modulo operation)

# String operations
full_name = "Alice" + " " + "Smith"  # String concatenation
greeting = "Hello" * 3     # Repeats string 3 times: "HelloHelloHello"

# -------- Functions --------
def greet(name, greeting="Hello"):
    """
    A simple function that returns a greeting message.
    Args:
        name (str): The name of the person to greet
        greeting (str): The greeting to use (default is "Hello")
    Returns:
        str: The complete greeting message
    """
    return f"{greeting}, {name}!"

# Function with multiple returns
def calculate_statistics(numbers):
    """Calculate basic statistics for a list of numbers."""
    if not numbers:
        return 0, 0, 0
    
    total = sum(numbers)
    average = total / len(numbers)
    maximum = max(numbers)
    return total, average, maximum
