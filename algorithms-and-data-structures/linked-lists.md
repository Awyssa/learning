A linked list is a data structure that consists of a sequence of elements, called nodes, where each node contains both data and a reference (or link) to the next node in the sequence. Unlike arrays, linked lists don't store elements in contiguous memory locations, which gives them some unique advantages and disadvantages.

Here's a basic implementation of a singly linked list in Python:

```python
class Node:
    def __init__(self, data):
        self.data = data
        self.next = None

class LinkedList:
    def __init__(self):
        self.head = None
    
    def append(self, data):
        new_node = Node(data)
        
        if self.head is None:
            self.head = new_node
            return
            
        current = self.head
        while current.next:
            current = current.next
        current.next = new_node
    
    def display(self):
        current = self.head
        while current:
            print(current.data, end=" -> ")
            current = current.next
        print("None")
```

Let's see how to use it:

```python
# Create a new linked list
my_list = LinkedList()

# Add some elements
my_list.append(1)
my_list.append(2)
my_list.append(3)

# Display the list
my_list.display()  # Output: 1 -> 2 -> 3 -> None
```

Key characteristics of linked lists:

1. Dynamic size - They can grow or shrink during runtime
2. Non-contiguous memory - Elements can be stored anywhere in memory
3. Efficient insertion and deletion at the beginning - O(1) time complexity
4. Linear access time - To reach the nth element, you need to traverse from the beginning

Let's add some more common operations to make this more practical:

```python
class LinkedList:
    # ... (previous methods remain the same)
    
    def prepend(self, data):
        """Add element at the beginning"""
        new_node = Node(data)
        new_node.next = self.head
        self.head = new_node
    
    def delete(self, data):
        """Delete first occurrence of data"""
        if self.head is None:
            return
            
        if self.head.data == data:
            self.head = self.head.next
            return
            
        current = self.head
        while current.next:
            if current.next.data == data:
                current.next = current.next.next
                return
            current = current.next
    
    def search(self, data):
        """Return True if data is found"""
        current = self.head
        while current:
            if current.data == data:
                return True
            current = current.next
        return False
```

You might use these new methods like this:

```python
my_list = LinkedList()

# Add elements at the end
my_list.append(2)
my_list.append(3)

# Add element at the beginning
my_list.prepend(1)

my_list.display()  # Output: 1 -> 2 -> 3 -> None

# Search for elements
print(my_list.search(2))  # Output: True
print(my_list.search(4))  # Output: False

# Delete an element
my_list.delete(2)
my_list.display()  # Output: 1 -> 3 -> None
```

This is a singly linked list, where each node only points to the next node. There are other variants like:
- Doubly linked lists (nodes have references to both next and previous nodes)
- Circular linked lists (last node points back to the first node)
- Multiply linked lists (nodes have multiple references to other nodes)