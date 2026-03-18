# Algorithm Patterns — Complete Implementation Reference

## Table of Contents

1. [Binary Search Patterns](#binary-search)
2. [Sorting Algorithms](#sorting)
3. [Graph Algorithms](#graph)
4. [Dynamic Programming](#dp)
5. [Greedy Algorithms](#greedy)
6. [Divide and Conquer](#divide-conquer)
7. [Sliding Window & Two Pointers](#sliding-window)
8. [Backtracking](#backtracking)
9. [String Algorithms](#string)
10. [Bit Manipulation](#bit-manipulation)
11. [Math & Number Theory](#math)

---

## <a name="binary-search"></a> 1. Binary Search Patterns

### Standard Binary Search
```
binary_search(arr, target):
    lo = 0
    hi = len(arr) - 1
    while lo <= hi:
        mid = lo + (hi - lo) / 2  // Avoid overflow vs (lo + hi) / 2
        if arr[mid] == target:
            return mid
        elif arr[mid] < target:
            lo = mid + 1
        else:
            hi = mid - 1
    return -1  // not found
```

### Lower Bound (first position ≥ target)
```
lower_bound(arr, target):
    lo = 0
    hi = len(arr)  // Note: hi = len, not len-1
    while lo < hi:  // Note: strict <, not <=
        mid = lo + (hi - lo) / 2
        if arr[mid] < target:
            lo = mid + 1
        else:
            hi = mid
    return lo
```

### Upper Bound (first position > target)
```
upper_bound(arr, target):
    lo = 0
    hi = len(arr)
    while lo < hi:
        mid = lo + (hi - lo) / 2
        if arr[mid] <= target:  // Only difference from lower_bound: <=
            lo = mid + 1
        else:
            hi = mid
    return lo
```

### Binary Search on Answer
When the answer space is monotonic (if `f(x)` is true, then `f(x+1)` is also true):
```
search_on_answer(lo, hi, condition):
    while lo < hi:
        mid = lo + (hi - lo) / 2
        if condition(mid):
            hi = mid  // mid might be the answer, search left
        else:
            lo = mid + 1  // mid is too small
    return lo
```

**Classic applications:** minimum capacity to ship packages, koko eating bananas,
split array largest sum, magnetic force between balls.

### Avoiding Infinite Loops
- `lo < hi` with `hi = mid`: converges because `mid < hi` when `lo < hi`
- `lo <= hi` with `hi = mid - 1`: converges because search space shrinks by ≥1
- DANGER: `lo < hi` with `lo = mid` can infinite loop when `hi - lo == 1`.
  Fix: use `mid = lo + (hi - lo + 1) / 2` (round up) when `lo = mid`.

---

## <a name="sorting"></a> 2. Sorting Algorithms

### Merge Sort
```
merge_sort(arr, lo, hi):
    if lo >= hi: return
    mid = (lo + hi) / 2
    merge_sort(arr, lo, mid)
    merge_sort(arr, mid + 1, hi)
    merge(arr, lo, mid, hi)

merge(arr, lo, mid, hi):
    left = arr[lo..mid]    // copy
    right = arr[mid+1..hi] // copy
    i = j = 0
    k = lo
    while i < len(left) AND j < len(right):
        if left[i] <= right[j]:  // <= for stability
            arr[k] = left[i]; i += 1
        else:
            arr[k] = right[j]; j += 1
        k += 1
    // Copy remaining
    while i < len(left): arr[k] = left[i]; i += 1; k += 1
    while j < len(right): arr[k] = right[j]; j += 1; k += 1
```
- Time: O(n log n) always. Space: O(n). Stable: yes.
- Bonus: count inversions by counting how many times right[j] is placed before left[i].

### Quick Sort (Hoare Partition)
```
quick_sort(arr, lo, hi):
    if lo >= hi: return
    pivot = partition(arr, lo, hi)
    quick_sort(arr, lo, pivot)      // Note: include pivot in left half (Hoare)
    quick_sort(arr, pivot + 1, hi)

partition_hoare(arr, lo, hi):
    pivot = arr[lo + (hi - lo) / 2]  // or random
    i = lo - 1
    j = hi + 1
    while true:
        do: i += 1  while arr[i] < pivot
        do: j -= 1  while arr[j] > pivot
        if i >= j: return j
        swap(arr[i], arr[j])
```
- Time: O(n log n) average, O(n²) worst. Space: O(log n) stack. Not stable.
- Randomized pivot makes worst case astronomically unlikely.

### Counting Sort
```
counting_sort(arr, max_val):
    count = array of size (max_val + 1), initialized to 0
    for x in arr: count[x] += 1
    i = 0
    for val in [0, max_val]:
        while count[val] > 0:
            arr[i] = val; i += 1
            count[val] -= 1
```
- Time: O(n + k) where k = range. Space: O(k). Stable (with modification).
- Use when k is small relative to n.

---

## <a name="graph"></a> 3. Graph Algorithms

### BFS (Breadth-First Search)
```
bfs(graph, start):
    visited = set()
    queue = [start]
    visited.add(start)
    distance = {start: 0}

    while queue not empty:
        node = queue.dequeue()
        for neighbor in graph[node]:
            if neighbor not in visited:
                visited.add(neighbor)
                distance[neighbor] = distance[node] + 1
                queue.enqueue(neighbor)
    return distance
```
- Time: O(V + E). Space: O(V).
- Gives shortest path in **unweighted** graphs.
- Multi-source BFS: add all sources to queue initially with distance 0.

### DFS (Depth-First Search)
```
dfs(graph, start):
    visited = set()
    _dfs_helper(graph, start, visited)

_dfs_helper(graph, node, visited):
    visited.add(node)
    for neighbor in graph[node]:
        if neighbor not in visited:
            _dfs_helper(graph, neighbor, visited)
```
- Time: O(V + E). Space: O(V) for recursion stack.
- Applications: cycle detection, topological sort, connected components,
  articulation points, bridges.

### Cycle Detection

**Undirected graph (DFS):**
```
has_cycle_undirected(graph):
    visited = set()
    for node in graph:
        if node not in visited:
            if _dfs_cycle(graph, node, -1, visited):
                return true
    return false

_dfs_cycle(graph, node, parent, visited):
    visited.add(node)
    for neighbor in graph[node]:
        if neighbor not in visited:
            if _dfs_cycle(graph, neighbor, node, visited):
                return true
        elif neighbor != parent:
            return true  // back edge = cycle
    return false
```

**Directed graph (colors: WHITE=unvisited, GRAY=in-progress, BLACK=done):**
```
has_cycle_directed(graph):
    color = {node: WHITE for node in graph}
    for node in graph:
        if color[node] == WHITE:
            if _dfs_cycle_directed(graph, node, color):
                return true
    return false

_dfs_cycle_directed(graph, node, color):
    color[node] = GRAY
    for neighbor in graph[node]:
        if color[neighbor] == GRAY:
            return true  // back edge = cycle
        if color[neighbor] == WHITE:
            if _dfs_cycle_directed(graph, neighbor, color):
                return true
    color[node] = BLACK
    return false
```

### Topological Sort

**Kahn's Algorithm (BFS-based):**
```
topological_sort_kahn(graph, n):
    in_degree = compute in-degrees for all nodes
    queue = [node for node if in_degree[node] == 0]
    order = []

    while queue not empty:
        node = queue.dequeue()
        order.append(node)
        for neighbor in graph[node]:
            in_degree[neighbor] -= 1
            if in_degree[neighbor] == 0:
                queue.enqueue(neighbor)

    if len(order) != n:
        return null  // cycle detected
    return order
```

**DFS-based:**
```
topological_sort_dfs(graph):
    visited = set()
    stack = []
    for node in graph:
        if node not in visited:
            _topo_dfs(graph, node, visited, stack)
    return reverse(stack)

_topo_dfs(graph, node, visited, stack):
    visited.add(node)
    for neighbor in graph[node]:
        if neighbor not in visited:
            _topo_dfs(graph, neighbor, visited, stack)
    stack.append(node)  // post-order
```

### Dijkstra's Algorithm
```
dijkstra(graph, start):
    dist = {node: INFINITY for node in graph}
    dist[start] = 0
    pq = min_heap([(0, start)])

    while pq not empty:
        d, u = pq.extract_min()
        if d > dist[u]: continue  // stale entry

        for (v, weight) in graph[u]:
            if dist[u] + weight < dist[v]:
                dist[v] = dist[u] + weight
                pq.insert((dist[v], v))

    return dist
```
- Time: O((V + E) log V) with binary heap. O(V² + E) with array (better for dense).
- **Only works with non-negative weights.** Why: Dijkstra assumes that once a node
  is extracted from the priority queue, its distance is final. Negative edges break this.

### Bellman-Ford
```
bellman_ford(edges, n, start):
    dist = [INFINITY] * n
    dist[start] = 0

    for i in [1, n-1]:  // n-1 relaxation passes
        for (u, v, w) in edges:
            if dist[u] + w < dist[v]:
                dist[v] = dist[u] + w

    // Check for negative cycles
    for (u, v, w) in edges:
        if dist[u] + w < dist[v]:
            return "negative cycle detected"

    return dist
```
- Time: O(V × E). Handles negative weights.

### Kruskal's MST
```
kruskal(edges, n):
    sort edges by weight
    uf = UnionFind(n)
    mst = []
    total_weight = 0

    for (u, v, w) in sorted_edges:
        if uf.find(u) != uf.find(v):
            uf.union(u, v)
            mst.append((u, v, w))
            total_weight += w
            if len(mst) == n - 1: break

    return mst, total_weight
```
- Time: O(E log E) dominated by sorting.

---

## <a name="dp"></a> 4. Dynamic Programming

### Framework
1. **Define the state:** What does `dp[i]` (or `dp[i][j]`) represent?
2. **Write the recurrence:** How do we compute `dp[i]` from previous states?
3. **Identify base cases:** What are the trivially solvable subproblems?
4. **Determine computation order:** Ensure dependencies are computed first.
5. **Optimize space if possible:** Often only need previous row/state.

### 1D DP Templates

**Climbing Stairs / Fibonacci-like:**
```
dp[0] = 1
dp[1] = 1
for i in [2, n]:
    dp[i] = dp[i-1] + dp[i-2]
// Space optimization: only need prev two values
```

**Coin Change (minimum coins):**
```
dp[0] = 0
dp[1..amount] = INFINITY
for i in [1, amount]:
    for coin in coins:
        if i >= coin:
            dp[i] = min(dp[i], dp[i - coin] + 1)
```

**Longest Increasing Subsequence (O(n log n)):**
```
tails = []  // tails[i] = smallest tail element of IS of length i+1
for num in arr:
    pos = lower_bound(tails, num)
    if pos == len(tails):
        tails.append(num)
    else:
        tails[pos] = num
return len(tails)
```

### 2D DP Templates

**Longest Common Subsequence:**
```
dp[i][j] = length of LCS of s1[0..i-1] and s2[0..j-1]

Base: dp[0][j] = dp[i][0] = 0
Recurrence:
    if s1[i-1] == s2[j-1]:
        dp[i][j] = dp[i-1][j-1] + 1
    else:
        dp[i][j] = max(dp[i-1][j], dp[i][j-1])
```

**Edit Distance:**
```
dp[i][j] = min edits to transform s1[0..i-1] to s2[0..j-1]

Base: dp[0][j] = j (all inserts), dp[i][0] = i (all deletes)
Recurrence:
    if s1[i-1] == s2[j-1]:
        dp[i][j] = dp[i-1][j-1]
    else:
        dp[i][j] = 1 + min(
            dp[i-1][j],    // delete
            dp[i][j-1],    // insert
            dp[i-1][j-1]   // replace
        )
```

**0/1 Knapsack:**
```
dp[i][w] = max value using items 0..i-1 with capacity w

Base: dp[0][w] = 0
Recurrence:
    dp[i][w] = dp[i-1][w]  // don't take item i
    if weight[i-1] <= w:
        dp[i][w] = max(dp[i][w], dp[i-1][w - weight[i-1]] + value[i-1])

// Space optimization: single row, iterate w from right to left
for i in [0, n):
    for w in [capacity, weight[i]) step -1:
        dp[w] = max(dp[w], dp[w - weight[i]] + value[i])
```

### Bitmask DP Template
```
// TSP: minimum cost to visit all n cities starting from 0
// State: dp[mask][i] = min cost to visit cities in mask, ending at city i
dp[1][0] = 0  // mask=1 means only city 0 visited

for mask in [1, 2^n):
    for i in [0, n) where bit i is set in mask:
        for j in [0, n) where bit j is NOT set in mask:
            new_mask = mask | (1 << j)
            dp[new_mask][j] = min(dp[new_mask][j], dp[mask][i] + cost[i][j])

answer = min over all i of (dp[(1<<n)-1][i] + cost[i][0])
```

---

## <a name="greedy"></a> 5. Greedy Algorithms

### When Greedy Works
The problem has **greedy choice property** (locally optimal choice leads to globally optimal)
AND **optimal substructure** (optimal solution contains optimal solutions to subproblems).

### Activity Selection
```
activity_selection(activities):
    sort activities by end time
    selected = [activities[0]]
    last_end = activities[0].end

    for i in [1, n):
        if activities[i].start >= last_end:
            selected.append(activities[i])
            last_end = activities[i].end

    return selected
```
**Correctness:** Choosing the earliest-ending activity leaves the most room for remaining activities.
This is the "greedy stays ahead" argument: at every step, our solution is at least as good as
any alternative.

### Interval Scheduling Variants
- **Max non-overlapping intervals:** sort by end time, pick greedily (above)
- **Min intervals to remove for non-overlap:** n - max_non_overlapping
- **Min meeting rooms:** sort by start time, use min-heap of end times
- **Merge overlapping intervals:** sort by start, merge if overlap

---

## <a name="divide-conquer"></a> 6. Divide and Conquer

### Master Theorem
For recurrences of form T(n) = aT(n/b) + O(n^d):
- If d < log_b(a): T(n) = O(n^(log_b(a)))
- If d = log_b(a): T(n) = O(n^d × log n)
- If d > log_b(a): T(n) = O(n^d)

**Examples:**
- Merge sort: T(n) = 2T(n/2) + O(n) → a=2, b=2, d=1 → case 2 → O(n log n)
- Binary search: T(n) = T(n/2) + O(1) → a=1, b=2, d=0 → case 2 → O(log n)
- Karatsuba: T(n) = 3T(n/2) + O(n) → a=3, b=2, d=1 → case 1 → O(n^1.585)

---

## <a name="sliding-window"></a> 7. Sliding Window & Two Pointers

### Fixed-Size Window
```
max_sum_subarray(arr, k):
    window_sum = sum(arr[0..k-1])
    max_sum = window_sum
    for i in [k, n):
        window_sum += arr[i] - arr[i - k]  // slide: add right, remove left
        max_sum = max(max_sum, window_sum)
    return max_sum
```

### Variable-Size Window (Shrinkable)
```
// Longest substring without repeating characters
longest_unique_substring(s):
    seen = {}  // char → last index
    left = 0
    max_len = 0
    for right in [0, len(s)):
        if s[right] in seen AND seen[s[right]] >= left:
            left = seen[s[right]] + 1
        seen[s[right]] = right
        max_len = max(max_len, right - left + 1)
    return max_len
```

### Variable-Size Window (Minimum)
```
// Minimum window substring containing all chars of target
min_window(s, t):
    need = frequency_count(t)
    have = 0
    required = len(need)
    left = 0
    result = (INFINITY, 0, 0)

    for right in [0, len(s)):
        char = s[right]
        if char in need:
            window_count[char] += 1
            if window_count[char] == need[char]:
                have += 1

        while have == required:
            // Update result
            if right - left + 1 < result[0]:
                result = (right - left + 1, left, right)
            // Shrink from left
            left_char = s[left]
            if left_char in need:
                window_count[left_char] -= 1
                if window_count[left_char] < need[left_char]:
                    have -= 1
            left += 1

    return result
```

### Two Pointers — Sorted Array
```
// Two Sum on sorted array
two_sum_sorted(arr, target):
    left = 0
    right = len(arr) - 1
    while left < right:
        current = arr[left] + arr[right]
        if current == target:
            return (left, right)
        elif current < target:
            left += 1
        else:
            right -= 1
    return null
```

---

## <a name="backtracking"></a> 8. Backtracking

### Template
```
backtrack(state, choices, result):
    if is_solution(state):
        result.append(copy(state))
        return

    for choice in choices:
        if is_valid(state, choice):
            state.apply(choice)        // CHOOSE
            backtrack(state, next_choices, result)  // EXPLORE
            state.undo(choice)         // UNCHOOSE (backtrack)
```

### Permutations
```
permutations(nums):
    result = []
    _perm(nums, 0, result)
    return result

_perm(nums, start, result):
    if start == len(nums):
        result.append(copy(nums))
        return
    for i in [start, len(nums)):
        swap(nums[start], nums[i])
        _perm(nums, start + 1, result)
        swap(nums[start], nums[i])  // backtrack
```

### Combinations (n choose k)
```
combinations(n, k):
    result = []
    _comb(1, n, k, [], result)
    return result

_comb(start, n, k, current, result):
    if len(current) == k:
        result.append(copy(current))
        return
    // Pruning: need k - len(current) more, so stop if not enough remaining
    for i in [start, n - (k - len(current)) + 1]:
        current.append(i)
        _comb(i + 1, n, k, current, result)
        current.pop()  // backtrack
```

---

## <a name="string"></a> 9. String Algorithms

### KMP (Knuth-Morris-Pratt)
```
// Build failure function
build_lps(pattern):
    lps = [0] * len(pattern)
    length = 0
    i = 1
    while i < len(pattern):
        if pattern[i] == pattern[length]:
            length += 1
            lps[i] = length
            i += 1
        elif length > 0:
            length = lps[length - 1]  // don't increment i
        else:
            lps[i] = 0
            i += 1
    return lps

kmp_search(text, pattern):
    lps = build_lps(pattern)
    i = 0  // text pointer
    j = 0  // pattern pointer
    matches = []
    while i < len(text):
        if text[i] == pattern[j]:
            i += 1; j += 1
        if j == len(pattern):
            matches.append(i - j)
            j = lps[j - 1]
        elif i < len(text) AND text[i] != pattern[j]:
            if j > 0:
                j = lps[j - 1]
            else:
                i += 1
    return matches
```
- Time: O(n + m). Space: O(m) for LPS array.

### Rabin-Karp (Rolling Hash)
```
rabin_karp(text, pattern):
    base = 256  // or 31 for lowercase only
    mod = 10^9 + 7
    m = len(pattern)
    n = len(text)

    // Compute hash of pattern and first window
    p_hash = 0
    t_hash = 0
    h = pow(base, m - 1) % mod  // for removing leading digit

    for i in [0, m):
        p_hash = (p_hash * base + pattern[i]) % mod
        t_hash = (t_hash * base + text[i]) % mod

    matches = []
    for i in [0, n - m]:
        if p_hash == t_hash:
            if text[i..i+m] == pattern:  // verify (avoid hash collision false positive)
                matches.append(i)
        if i < n - m:
            t_hash = (t_hash - text[i] * h) * base + text[i + m]
            t_hash = t_hash % mod

    return matches
```
- Time: O(n + m) average, O(nm) worst (many hash collisions). Space: O(1).

---

## <a name="bit-manipulation"></a> 10. Bit Manipulation

### Essential Operations
```
x & (x - 1)     // Clear lowest set bit. Useful for counting set bits.
x & (-x)         // Isolate lowest set bit. Used in Fenwick trees.
x | (x + 1)      // Set lowest unset bit.
x ^ y            // Bits that differ. XOR of all = find the unique element.
(x >> i) & 1     // Check if bit i is set.
x | (1 << i)     // Set bit i.
x & ~(1 << i)    // Clear bit i.
x ^ (1 << i)     // Toggle bit i.
```

### Count Set Bits (Brian Kernighan)
```
count_bits(n):
    count = 0
    while n > 0:
        n = n & (n - 1)  // clear lowest set bit
        count += 1
    return count
```
- Time: O(number of set bits), not O(32).

### Power of Two Check
```
is_power_of_two(n):
    return n > 0 AND (n & (n - 1)) == 0
```

---

## <a name="math"></a> 11. Math & Number Theory

### GCD (Euclidean Algorithm)
```
gcd(a, b):
    while b != 0:
        a, b = b, a % b
    return a
```
- Time: O(log(min(a, b)))

### Modular Exponentiation
```
mod_pow(base, exp, mod):
    result = 1
    base = base % mod
    while exp > 0:
        if exp is odd:
            result = (result * base) % mod
        exp = exp / 2
        base = (base * base) % mod
    return result
```
- Time: O(log exp)

### Sieve of Eratosthenes
```
sieve(n):
    is_prime = [true] * (n + 1)
    is_prime[0] = is_prime[1] = false
    for i in [2, sqrt(n)]:
        if is_prime[i]:
            for j in [i*i, n] step i:
                is_prime[j] = false
    return [i for i in [2, n] if is_prime[i]]
```
- Time: O(n log log n). Space: O(n).