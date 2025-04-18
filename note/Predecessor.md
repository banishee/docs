
Maintain a dynamic set $S$ with support for:
- `predecessor(x)`: Largest element in \( S \) that is ≤ x
- `successor(x)`: Smallest element in \( S \) that is ≥ x
- `insert(x)` and `delete(x)`

## Bitvector

- **Solution 1: Simple Bitvector**
  - Walk left from `x` in a bitvector of size `u`
  - Time complexity: **O(u)**

- **Solution 2: Two-Level Bitvector**
  - Top and bottom bitvectors with some improvements
  - Still **O(u)** time due to linear scans

- **Solutions 3 & 4: Optimized Two-Level Bitvector**
  - Introduce min/max per block and recursive division
  - Leads to **O(√u)** and eventually **O(u^1/4)** time

## van Emde Boas

## Tries