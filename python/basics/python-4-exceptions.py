# Introduction to Try/Except Exception Handling in Python
# This file covers error handling concepts including try/except blocks, custom exceptions, and best practices

# -------- Basic Try/Except Structure --------
try:
    result = 10 / 0  # This will raise a ZeroDivisionError
except ZeroDivisionError:
    print("Cannot divide by zero!")


# -------- Multiple Exception Types --------
def process_data(data):
    try:
        number = int(data)  # Might raise ValueError
        result = 100 / number  # Might raise ZeroDivisionError
        return result
    except ValueError:
        print("Please provide a valid number")
    except ZeroDivisionError:
        print("Number cannot be zero")
    except Exception as e:  # Catch any other exceptions
        print(f"Unexpected error: {e}")


# -------- Try/Except/Else/Finally --------
def read_file(filename):
    try:
        file = open(filename, "r")  # Might raise FileNotFoundError
    except FileNotFoundError:
        print(f"File {filename} not found")
    else:  # Runs if try block succeeds
        content = file.read()
        file.close()
        return content
    finally:  # Always runs, regardless of success/failure
        print("File operation attempted")


# -------- Custom Exceptions --------
class InvalidAgeError(Exception):
    """Custom exception for invalid age values"""

    pass


def verify_age(age):
    try:
        age = int(age)
        if age < 0:
            raise InvalidAgeError("Age cannot be negative")
        if age > 150:
            raise InvalidAgeError("Age seems unrealistic")
    except ValueError:
        print("Please enter a valid number")
    except InvalidAgeError as e:
        print(f"Invalid age: {e}")


# -------- Context Managers (with statement) --------
def process_file_safe(filename):
    try:
        with open(filename, "r") as file:  # Automatically handles file closing
            content = file.read()
            return content
    except FileNotFoundError:
        print("File not found")
    except PermissionError:
        print("Permission denied")


# -------- Practical Examples --------
# Example 1: Dictionary access
def get_user_info(user_dict, key):
    try:
        return user_dict[key]
    except KeyError:
        return f"Key '{key}' not found in user data"


# Example 2: Type conversion with validation
def convert_to_number(value):
    try:
        float_value = float(value)
        return int(float_value) if float_value.is_integer() else float_value
    except ValueError:
        return f"'{value}' cannot be converted to a number"


# Example 3: Nested exception handling
def complex_operation(data):
    try:
        with open("config.txt") as file:
            config = file.read()
            try:
                processed_data = data / int(config)
                return processed_data
            except (ValueError, ZeroDivisionError) as e:
                print(f"Data processing error: {e}")
    except FileNotFoundError:
        print("Config file missing")
