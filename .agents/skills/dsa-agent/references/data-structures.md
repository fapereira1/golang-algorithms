# Data Structures — Complete Implementation Reference

## Table of Contents

1. [Array / Dynamic Array](#dynamic-array)
2. [Linked List](#linked-list)
3. [Stack](#stack)
4. [Queue / Deque](#queue)
5. [Hash Table](#hash-table)
6. [Binary Search Tree](#bst)
7. [Heap / Priority Queue](#heap)
8. [Graph Representations](#graph)
9. [Trie](#trie)
10. [Union-Find](#union-find)
11. [Segment Tree](#segment-tree)
12. [Fenwick Tree](#fenwick-tree)
13. [Balanced BSTs (conceptual)](#balanced-bst)
14. [Bloom Filter](#bloom-filter)

---

## <a name="dynamic-array"></a> 1. Dynamic Array

### Key Insight
Amortized O(1) append by doubling capacity when full. The cost of N appends is O(N) total
because resize operations form a geometric series: 1 + 2 + 4 + ... + N = 2N - 1.

### Operations & Complexity
| Operation      | Average | Worst  | Note                              |
|----------------|---------|--------|-----------------------------------|
| Access [i]     | O(1)    | O(1)   | Direct memory offset              |
| Append         | O(1)*   | O(n)   | *amortized — resize costs O(n)    |
| Insert at i    | O(n)    | O(n)   | Shift elements right              |
| Delete at i    | O(n)    | O(n)   | Shift elements left               |
| Search         | O(n)    | O(n)   | Linear scan (unsorted)            |

### Implementation Pattern
```
class DynamicArray:
    capacity = INITIAL_CAPACITY
    size = 0
    data = allocate(capacity)

    append(value):
        if size == capacity:
            resize(capacity * 2)  // Growth factor 2 is standard; 1.5 saves memory
        data[size] = value
        size += 1

    resize(new_capacity):
        new_data = allocate(new_capacity)
        copy data[0..size-1] to new_data
        data = new_data
        capacity = new_capacity
```

### When to Use
- Default choice for ordered collections with frequent random access
- When you need cache-friendly iteration (contiguous memory)
- Avoid when frequent insertions/deletions at arbitrary positions are needed

---

## <a name="linked-list"></a> 2. Linked List

### Variants
- **Singly linked**: each node has `value` and `next`
- **Doubly linked**: each node has `value`, `next`, and `prev`
- **Circular**: tail.next points to head

### Operations & Complexity
| Operation           | Singly | Doubly | Note                           |
|---------------------|--------|--------|--------------------------------|
| Access [i]          | O(n)   | O(n)   | Must traverse from head        |
| Insert at head      | O(1)   | O(1)   | Pointer manipulation           |
| Insert at tail      | O(1)*  | O(1)   | *with tail pointer             |
| Insert after node   | O(1)   | O(1)   | Given reference to node        |
| Delete node         | O(n)   | O(1)   | Singly: need predecessor       |
| Search              | O(n)   | O(n)   | Linear traversal               |

### Essential Patterns

**Reverse a singly linked list (iterative):**
```
reverse(head):
    prev = null
    curr = head
    while curr != null:
        next_temp = curr.next
        curr.next = prev
        prev = curr
        curr = next_temp
    return prev  // prev is new head
```

**Detect cycle (Floyd's tortoise and hare):**
```
has_cycle(head):
    slow = head
    fast = head
    while fast != null AND fast.next != null:
        slow = slow.next
        fast = fast.next.next
        if slow == fast:
            return true
    return false
```

**Find cycle start:**
```
find_cycle_start(head):
    // Phase 1: detect meeting point
    slow = fast = head
    while fast != null AND fast.next != null:
        slow = slow.next
        fast = fast.next.next
        if slow == fast:
            break
    if fast == null OR fast.next == null:
        return null  // no cycle

    // Phase 2: find entry point
    slow = head
    while slow != fast:
        slow = slow.next
        fast = fast.next
    return slow  // cycle start
```

**Why Phase 2 works:** If the non-cycle part has length `a` and the cycle has length `c`,
at the meeting point, slow traveled `a + b` and fast traveled `a + b + kc`.
Since fast = 2*slow: `a + b + kc = 2(a + b)` → `a = kc - b`.
So moving `a` steps from meeting point lands at cycle start.

---

## <a name="stack"></a> 3. Stack

### Operations & Complexity
All operations O(1): push, pop, peek, isEmpty.

### Monotonic Stack Pattern
Used for "next greater/smaller element" problems.

```
next_greater_element(arr):
    result = [-1] * len(arr)
    stack = []  // stores indices
    for i in range(len(arr)):
        while stack is not empty AND arr[stack.top()] < arr[i]:
            idx = stack.pop()
            result[idx] = arr[i]
        stack.push(i)
    return result
```

**Key insight:** The stack maintains a decreasing sequence. When we find a larger element,
we pop all smaller elements — they've found their "next greater". O(n) because each element
is pushed and popped at most once.

### Classic Applications
- Expression evaluation (infix → postfix, parenthesis matching)
- DFS iteration (explicit stack instead of recursion)
- Undo operations
- Histogram problems (largest rectangle)

---

## <a name="queue"></a> 4. Queue / Deque

### Operations & Complexity
All operations O(1): enqueue, dequeue, peek (for both queue and deque).

### Circular Queue Pattern
```
class CircularQueue:
    data = allocate(capacity)
    head = 0
    tail = 0
    size = 0

    enqueue(value):
        if size == capacity: error("full")
        data[tail] = value
        tail = (tail + 1) % capacity
        size += 1

    dequeue():
        if size == 0: error("empty")
        value = data[head]
        head = (head + 1) % capacity
        size -= 1
        return value
```

### Deque Applications
- Sliding window maximum (monotonic deque)
- BFS with 0-1 edge weights (0-1 BFS)
- Palindrome checking

---

## <a name="hash-table"></a> 5. Hash Table

### Collision Resolution

**Chaining (linked lists at each bucket):**
- Simple to implement
- Performance degrades gracefully
- Extra memory for pointers

**Open Addressing (linear/quadratic probing, double hashing):**
- Better cache performance
- Sensitive to load factor (keep α < 0.7)
- Deletion requires tombstones

### Operations & Complexity
| Operation | Average | Worst | Note                            |
|-----------|---------|-------|---------------------------------|
| Insert    | O(1)    | O(n)  | Worst case: all collisions      |
| Search    | O(1)    | O(n)  | Depends on hash quality + α     |
| Delete    | O(1)    | O(n)  | Chaining: easy. Open addr: tomb |

### Good Hash Function Properties
- Deterministic
- Uniform distribution across buckets
- Avalanche effect (small input change → big hash change)
- Fast to compute

### Rehashing
When load factor α = n/m exceeds threshold (typically 0.75):
1. Allocate new table with 2× buckets
2. Reinsert every element (rehash with new modulus)
3. O(n) operation, but amortized O(1) per insert

---

## <a name="bst"></a> 6. Binary Search Tree

### Invariant
For every node: all values in left subtree < node.value < all values in right subtree.

### Operations & Complexity
| Operation    | Average    | Worst (degenerate) |
|-------------|------------|---------------------|
| Search      | O(log n)   | O(n)                |
| Insert      | O(log n)   | O(n)                |
| Delete      | O(log n)   | O(n)                |
| Min/Max     | O(log n)   | O(n)                |
| In-order    | O(n)       | O(n)                |

### Deletion Cases
1. **Leaf node:** simply remove
2. **One child:** replace node with its child
3. **Two children:** replace with in-order successor (smallest in right subtree)
   or in-order predecessor (largest in left subtree), then delete the successor/predecessor

### Traversals
```
in_order(node):     // Sorted order
    if node == null: return
    in_order(node.left)
    visit(node)
    in_order(node.right)

pre_order(node):    // Copy tree, serialize
    if node == null: return
    visit(node)
    pre_order(node.left)
    pre_order(node.right)

post_order(node):   // Delete tree, postfix expr
    if node == null: return
    post_order(node.left)
    post_order(node.right)
    visit(node)

level_order(root):  // BFS
    queue = [root]
    while queue not empty:
        node = queue.dequeue()
        visit(node)
        if node.left: queue.enqueue(node.left)
        if node.right: queue.enqueue(node.right)
```

---

## <a name="heap"></a> 7. Heap / Priority Queue

### Properties (Min-Heap)
- Complete binary tree stored as array
- Parent ≤ both children
- parent(i) = (i-1)/2, left(i) = 2i+1, right(i) = 2i+2

### Operations & Complexity
| Operation      | Time       | Note                          |
|----------------|------------|-------------------------------|
| Insert         | O(log n)   | Add at end, bubble up         |
| Extract-min    | O(log n)   | Remove root, bubble down      |
| Peek-min       | O(1)       | Return root                   |
| Heapify array  | O(n)       | Bottom-up, NOT O(n log n)     |
| Decrease-key   | O(log n)   | Update value, bubble up       |

### Why Heapify is O(n), not O(n log n)
Bottom-up heapify: start from last non-leaf, sift down each node.
Nodes at height h do O(h) work. There are ≤ n/2^(h+1) nodes at height h.
Total work = Σ (h=0 to log n) of n/2^(h+1) × h = O(n).
The sum converges because most nodes are near the bottom and do little work.

### Implementation Pattern
```
sift_up(i):
    while i > 0 AND heap[parent(i)] > heap[i]:
        swap(heap[parent(i)], heap[i])
        i = parent(i)

sift_down(i):
    while left(i) < size:
        smallest = i
        if heap[left(i)] < heap[smallest]:
            smallest = left(i)
        if right(i) < size AND heap[right(i)] < heap[smallest]:
            smallest = right(i)
        if smallest == i: break
        swap(heap[i], heap[smallest])
        i = smallest
```

---

## <a name="graph"></a> 8. Graph Representations

### Adjacency List vs Adjacency Matrix

| Aspect          | Adj. List      | Adj. Matrix   |
|-----------------|----------------|---------------|
| Space           | O(V + E)       | O(V²)         |
| Add edge        | O(1)           | O(1)          |
| Remove edge     | O(degree)      | O(1)          |
| Check edge      | O(degree)      | O(1)          |
| Iterate neighbors| O(degree)     | O(V)          |
| Best for        | Sparse graphs  | Dense graphs  |

### When to Use Which
- **Adj. list**: most real-world graphs (social networks, web graphs, road networks)
- **Adj. matrix**: dense graphs, Floyd-Warshall, need O(1) edge lookup
- **Edge list**: Kruskal's MST, simple storage when you just need to iterate all edges

---

## <a name="trie"></a> 9. Trie (Prefix Tree)

### Operations & Complexity
| Operation  | Time     | Note                         |
|-----------|----------|-------------------------------|
| Insert    | O(m)     | m = length of word            |
| Search    | O(m)     | Exact match                   |
| Prefix    | O(m)     | Check if any word starts with |
| Delete    | O(m)     | Remove + prune empty branches |

### Implementation Pattern
```
class TrieNode:
    children = {}  // map char → TrieNode
    is_end = false

insert(word):
    node = root
    for char in word:
        if char not in node.children:
            node.children[char] = new TrieNode()
        node = node.children[char]
    node.is_end = true

search(word):
    node = root
    for char in word:
        if char not in node.children:
            return false
        node = node.children[char]
    return node.is_end
```

### Memory Optimization
- Use array[26] instead of hash map for lowercase English letters
- Compressed trie (Patricia/Radix tree): merge chains of single-child nodes

---

## <a name="union-find"></a> 10. Union-Find (Disjoint Set Union)

### Operations & Complexity (with both optimizations)
| Operation | Time          | Note                                    |
|-----------|---------------|-----------------------------------------|
| Find      | O(α(n)) ≈ O(1)| α = inverse Ackermann, practically ≤ 4 |
| Union     | O(α(n)) ≈ O(1)| Via find + link                         |

### Implementation Pattern
```
parent = [0..n-1]  // each element is its own parent initially
rank = [0] * n

find(x):
    if parent[x] != x:
        parent[x] = find(parent[x])  // path compression
    return parent[x]

union(x, y):
    root_x = find(x)
    root_y = find(y)
    if root_x == root_y: return  // already connected

    // union by rank
    if rank[root_x] < rank[root_y]:
        parent[root_x] = root_y
    elif rank[root_x] > rank[root_y]:
        parent[root_y] = root_x
    else:
        parent[root_y] = root_x
        rank[root_x] += 1
```

### Why Path Compression + Union by Rank
Either optimization alone gives O(log n). Together they give O(α(n)),
which is effectively constant for all practical input sizes (α(n) ≤ 4 for n ≤ 10^80).

---

## <a name="segment-tree"></a> 11. Segment Tree

### Use Case
Range queries + point updates in O(log n), or range updates with lazy propagation.

### Operations & Complexity
| Operation          | Time     | Note                        |
|--------------------|----------|-----------------------------|
| Build              | O(n)     | Bottom-up                   |
| Point update       | O(log n) | Update leaf, propagate up   |
| Range query        | O(log n) | Combine relevant segments   |
| Range update (lazy)| O(log n) | Deferred propagation        |

### Implementation Pattern (iterative, 0-indexed)
```
// Array-based, size 2*n. Leaves at indices [n, 2n-1].
tree = array of size 2 * n

build(arr):
    // Copy leaves
    for i in [0, n): tree[n + i] = arr[i]
    // Build internal nodes bottom-up
    for i in [n-1, 0) step -1: tree[i] = combine(tree[2*i], tree[2*i+1])

update(pos, value):
    pos += n
    tree[pos] = value
    while pos > 1:
        pos /= 2
        tree[pos] = combine(tree[2*pos], tree[2*pos+1])

query(left, right):  // [left, right) half-open
    result = identity_element
    left += n
    right += n
    while left < right:
        if left is odd:  result = combine(result, tree[left]); left += 1
        if right is odd:  right -= 1; result = combine(result, tree[right])
        left /= 2
        right /= 2
    return result
```

`combine` is the associative operation (sum, min, max, gcd, etc.).

---

## <a name="fenwick-tree"></a> 12. Fenwick Tree (Binary Indexed Tree)

### Advantage over Segment Tree
Simpler to implement, less memory (n+1 vs 2n), faster constant factor.
Limited to operations with an inverse (sum works; min does not).

### Operations & Complexity
| Operation    | Time     |
|-------------|----------|
| Point update | O(log n) |
| Prefix query | O(log n) |
| Range query  | O(log n) | // query(l, r) = prefix(r) - prefix(l-1)

### Implementation Pattern (1-indexed)
```
tree = array of size n + 1, initialized to 0

update(i, delta):
    while i <= n:
        tree[i] += delta
        i += i & (-i)  // add lowest set bit

prefix_query(i):
    sum = 0
    while i > 0:
        sum += tree[i]
        i -= i & (-i)  // remove lowest set bit
    return sum

range_query(l, r):
    return prefix_query(r) - prefix_query(l - 1)
```

### Why `i & (-i)` Works
In two's complement, `-i` flips all bits and adds 1. `i & (-i)` isolates the lowest set bit.
This bit determines the range of responsibility for each index in the tree.

---

## <a name="balanced-bst"></a> 13. Balanced BSTs (Conceptual)

### AVL Tree
- **Invariant:** For every node, |height(left) - height(right)| ≤ 1
- **Rotations:** Single (LL, RR) and double (LR, RL) rotations to rebalance
- **Guarantee:** O(log n) worst case for all operations
- **When to prefer:** When you need strict balance (more lookups than modifications)

### Red-Black Tree
- **Invariant:** Nodes are red/black, root is black, red nodes have black children,
  all paths from root to null have equal black-node count
- **Guarantee:** O(log n) worst case, height ≤ 2 log(n+1)
- **When to prefer:** When modifications are frequent (fewer rotations on average)
- Most standard library implementations use red-black trees (e.g., TreeMap, std::map)

### Comparison
| Aspect       | AVL           | Red-Black     |
|-------------|---------------|---------------|
| Max height  | 1.44 log n    | 2 log n       |
| Lookup      | Faster        | Slightly slower|
| Insert/Delete| More rotations| Fewer rotations|
| Use case    | Read-heavy    | Write-heavy    |

---

## <a name="bloom-filter"></a> 14. Bloom Filter

### Properties
- Probabilistic set membership: "definitely not in set" or "probably in set"
- False positives possible, false negatives impossible
- Space-efficient: ~10 bits per element for 1% false positive rate

### Parameters
- m = number of bits in the array
- k = number of hash functions
- n = number of elements
- Optimal k = (m/n) × ln(2)
- False positive rate ≈ (1 - e^(-kn/m))^k

### Operations
| Operation | Time | Note                               |
|-----------|------|------------------------------------|
| Insert    | O(k) | Set k bits                         |
| Query     | O(k) | Check k bits, all 1 = "maybe yes" |
| Delete    | N/A  | Standard bloom filter can't delete |

### When to Use
- Pre-filtering before expensive lookups (database, disk, network)
- Spell checkers, URL deduplication, cache filtering
- When you can tolerate false positives but not false negatives