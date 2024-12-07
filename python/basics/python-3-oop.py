# Introduction to Object-Oriented Programming (OOP) in Python
# This file covers fundamental concepts including classes, objects, inheritance, and encapsulation


# -------- Basic Class Definition --------
class Student:
    # Class variable (shared among all instances)
    school_name = "Python High School"

    # Constructor (initializer method)
    def __init__(self, name, age, grade):
        # Instance variables (unique to each instance)
        self.name = name
        self.age = age
        self.grade = grade
        self._attendance = 0  # Protected variable (convention)

    # Instance method
    def study(self, subject):
        return f"{self.name} is studying {subject}"

    # Getter method (property decorator)
    @property
    def attendance(self):
        return self._attendance

    # Setter method
    @attendance.setter
    def attendance(self, days):
        if days >= 0:
            self._attendance = days


# -------- Inheritance Example --------
class GraduateStudent(Student):
    def __init__(self, name, age, grade, research_topic):
        # Call parent class constructor
        super().__init__(name, age, grade)
        self.research_topic = research_topic

    # Method overriding
    def study(self, subject):
        base_message = super().study(subject)
        return f"{base_message} and researching {self.research_topic}"


# -------- Multiple Classes Interaction --------
class Course:
    def __init__(self, name, teacher):
        self.name = name
        self.teacher = teacher
        self.students = []

    def add_student(self, student):
        self.students.append(student)

    def get_class_size(self):
        return len(self.students)


# -------- Usage Examples --------
# Creating objects (instances)
student1 = Student("Alice", 15, "10th")
grad_student = GraduateStudent("Bob", 22, "Graduate", "AI")

# Using methods and accessing attributes
print(student1.study("Math"))  # Alice is studying Math
print(grad_student.study("CS"))  # Bob is studying CS and researching AI

# Using properties
student1.attendance = 5  # Using the setter
print(student1.attendance)  # Using the getter

# Creating a course and adding students
python_course = Course("Python 101", "Dr. Smith")
python_course.add_student(student1)
python_course.add_student(grad_student)


# Class Methods in Python
class Library:
    total_books = 0  # Class variable

    def __init__(self, name):
        self.name = name
        self.books = []

    # Regular instance method - has access to instance attributes (self)
    def add_book(self, book):
        self.books.append(book)
        Library.total_books += 1

    # Class method - has access to class attributes (cls)
    @classmethod
    def from_book_list(cls, name, book_list):
        """Create a Library instance from a list of books"""
        library = cls(name)  # Create new instance
        for book in book_list:
            library.add_book(book)
        return library

    # Another class method - for working with class variables
    @classmethod
    def get_total_books(cls):
        return cls.total_books

    # Alternative constructor using class method
    @classmethod
    def create_empty(cls, name):
        """Create an empty library with just a name"""
        return cls(name)


# Usage example
books = ["1984", "Dune", "Foundation"]
city_library = Library.from_book_list("City Library", books)
empty_library = Library.create_empty("New Branch")

print(Library.get_total_books())  # Prints: 3


class ExampleClass:
    class_variable = 0

    def __init__(self, value):
        self.value = value

    # 1. Instance Method
    def instance_method(self, x):
        """Regular instance method - has access to instance through 'self'"""
        return self.value + x

    # 2. Class Method
    @classmethod
    def class_method(cls, x):
        """Class method - has access to class through 'cls'"""
        cls.class_variable += x
        return cls.class_variable

    # 3. Static Method
    @staticmethod
    def static_method(x, y):
        """Static method - no access to instance or class"""
        return x + y

    # 4. Property Method (Getter)
    @property
    def my_property(self):
        """Property getter - accessed like an attribute"""
        return self.value * 2

    # 5. Property Setter
    @my_property.setter
    def my_property(self, new_value):
        """Property setter - allows controlled attribute setting"""
        self.value = new_value // 2

    # 6. Property Deleter
    @my_property.deleter
    def my_property(self):
        """Property deleter - handles attribute deletion"""
        self.value = 0

    # 7. Magic/Dunder Methods
    def __str__(self):
        """Magic method for string representation"""
        return f"ExampleClass with value {self.value}"

    def __len__(self):
        """Magic method for len()"""
        return self.value

    # 8. Private Method (by convention)
    def _private_method(self):
        """Private method (single underscore) - convention for internal use"""
        return "This is meant to be private"

    # 9. Name-mangled Method
    def __really_private(self):
        """Name-mangled method (double underscore) - harder to access from outside"""
        return "This is name-mangled"
