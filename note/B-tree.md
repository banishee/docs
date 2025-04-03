# B-tree

B-tree is a self-balancing search tree in which each node can contain **more than one key** and can have **more than two children**.

B-tree can decreases the height significantly allowing fewer disk accesses.

## Properties

1. For each node `x`, keys are stored in increasing order.
2. For each node except root can have at most `n` children and at least `n/2` children.
3. All leaves have the same depth.
4. The root has at least 2 children and contains a minimum of 1 key.