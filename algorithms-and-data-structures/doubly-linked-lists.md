A Doubly Linked List is a data structure where each node contains data and two pointers/references: one pointing to the next node and another pointing to the previous node. This bidirectional linking allows us to traverse the list in both forward and backward directions.

Here's a Python implementation with explanations:

```python
class Node:
    def __init__(self, data):
        self.data = data    # The actual data stored in the node
        self.next = None    # Reference to the next node
        self.prev = None    # Reference to the previous node

class DoublyLinkedList:
    def __init__(self):
        self.head = None    # Points to the first node
        self.tail = None    # Points to the last node
    
    def append(self, data):
        new_node = Node(data)
        
        # If list is empty
        if not self.head:
            self.head = new_node
            self.tail = new_node
            return
            
        # Add node to the end
        new_node.prev = self.tail
        self.tail.next = new_node
        self.tail = new_node
    
    def print_forward(self):
        current = self.head
        while current:
            print(current.data, end=" → ")
            current = current.next
        print("None")
    
    def print_backward(self):
        current = self.tail
        while current:
            print(current.data, end=" → ")
            current = current.prev
        print("None")
```

Let's see how to use this implementation:

```python
# Create a new doubly linked list
dll = DoublyLinkedList()

# Add some elements
dll.append(1)
dll.append(2)
dll.append(3)

# Print in both directions
print("Forward:")
dll.print_forward()    # Output: 1 → 2 → 3 → None
print("Backward:")
dll.print_backward()   # Output: 3 → 2 → 1 → None
```

Key advantages of Doubly Linked Lists:
1. We can traverse in both directions (forward and backward)
2. Deletion of nodes is more efficient than in singly linked lists since we have direct access to the previous node
3. We can quickly insert before a given node

Let me show you how to implement node deletion:

```python
def delete_node(self, key):
    current = self.head
    
    # Find the node to delete
    while current and current.data != key:
        current = current.next
    
    # If node not found
    if not current:
        return
    
    # If node is head
    if current == self.head:
        self.head = current.next
        if self.head:
            self.head.prev = None
        return
    
    # If node is tail
    if current == self.tail:
        self.tail = current.prev
        self.tail.next = None
        return
    
    # If node is in the middle
    current.prev.next = current.next
    current.next.prev = current.prev
```
