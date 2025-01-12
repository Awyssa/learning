A stack is a linear data structure that follows the Last In, First Out (LIFO) principle. Think of it like a stack of plates - you can only add or remove plates from the top.

Here are the key operations of a stack:

1. Push - adds an element to the top
```python
stack = []
stack.push(5)    # or stack.append(5) in Python
```

2. Pop - removes and returns the top element
```python
top_element = stack.pop()
```

3. Peek/Top - views the top element without removing it
```python
top_element = stack[-1]    # in Python
```

Common use cases:
- Function call stack (managing function calls/returns)
- Undo/Redo operations in applications
- Expression evaluation (like checking balanced parentheses)
- Browser's back/forward history
- Syntax parsing in compilers

Example implementation in Python:
```python
class Stack:
    def __init__(self):
        self.items = []
    
    def push(self, item):
        self.items.append(item)
    
    def pop(self):
        if not self.is_empty():
            return self.items.pop()
    
    def peek(self):
        if not self.is_empty():
            return self.items[-1]
    
    def is_empty(self):
        return len(self.items) == 0
```

The main characteristics are:
- Constant time O(1) for push and pop operations
- Elements can only be accessed from one end
- No random access to elements in the middle