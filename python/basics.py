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

# Function with multiple returns
def calculate_statistics(numbers):
    """Calculate basic statistics for a list of numbers."""
    if not numbers:
        return 0, 0, 0
    
    total = sum(numbers)
    average = total / len(numbers)
    maximum = max(numbers)
    return total, average, maximum

# -------- Lists --------
# Creating and modifying lists
fruits = ["apple", "banana", "orange"]
fruits.append("grape")           # Add to end
fruits.insert(1, "mango")       # Insert at position
first_fruit = fruits[0]         # Access by index
fruits[-1] = "kiwi"             # Modify last element

# List operations
numbers = [1, 2, 3, 4, 5]
subset = numbers[1:4]           # Slicing: [2, 3, 4]
reversed_list = numbers[::-1]   # Reverse list
combined = fruits + numbers     # List concatenation

# -------- Dictionaries --------
# Creating and using dictionaries
person = {
    "name": "Alice",
    "age": 25,
    "city": "New York"
}

# Accessing and modifying dictionary values
person["email"] = "alice@example.com"  # Add new key-value pair
age = person.get("age")               # Safe way to get values
person["age"] = 26                    # Modify existing value

# -------- Tuples --------
# Immutable sequences
coordinates = (10, 20)
x, y = coordinates              # Tuple unpacking

# -------- Sets --------
# Unique collections of items
unique_numbers = {1, 2, 3, 3, 4, 4, 5}  # Creates {1, 2, 3, 4, 5}
unique_numbers.add(6)           # Add element
unique_numbers.remove(1)        # Remove element

# -------- Basic Program Flow --------
# Example using various concepts
def process_user_data(users):
    """Process a list of user dictionaries and return summary statistics."""
    ages = []
    cities = set()
    
    for user in users:
        if "age" in user:
            ages.append(user["age"])
        if "city" in user:
            cities.add(user["city"])
    
    total_age, avg_age, max_age = calculate_statistics(ages)
    
    return {
        "total_users": len(users),
        "average_age": avg_age,
        "max_age": max_age,
        "unique_cities": len(cities)
    }

# Example usage
if __name__ == "__main__":
    users = [
        {"name": "Alice", "age": 25, "city": "New York"},
        {"name": "Bob", "age": 30, "city": "Boston"},
        {"name": "Charlie", "age": 35, "city": "New York"}
    ]
    
    stats = process_user_data(users)
    print("User Statistics:", stats)
    
    # Example function calls
    print(greet("World"))
    print(greet("Python", greeting="Welcome to"))