# Example: Constructing a 2D Range Tree with Bridges

Let's illustrate how to build a 2D range tree with bridges using a small set of points step by step. We will use 6 points (each with an x and y coordinate) and follow the method of sorting by x for leaves, building a balanced tree bottom-up, maintaining sorted y-lists at each node, and adding bridge pointers between parent and child nodes.

1. **Sort points by x-coordinate and create leaf nodes:**  
   Suppose we have the following 6 points:  
   - P1 = (1, 5)  
   - P2 = (3, 7)  
   - P3 = (4, 2)  
   - P4 = (6, 9)  
   - P5 = (7, 1)  
   - P6 = (9, 6)  
   First, sort these points by their x-coordinate in ascending order. In sorted order by x, we get: **P1, P2, P3, P4, P5, P6** (corresponding x’s: 1 < 3 < 4 < 6 < 7 < 9). Each sorted point becomes a **leaf node** of the range tree. At this stage, we have 6 leaf nodes (one per point) arranged left-to-right in increasing x order.

2. **Build a balanced binary tree bottom-up from the leaves:**  
   Next, we connect the leaf nodes into a balanced binary tree based on their x-order. We pair up or group the leaves and create parent internal nodes until one root covers all points. For our 6 leaves:  
   - **Left subtree:** Take the first half of the leaves (P1, P2, P3) as the left subtree. Create an internal node **B** to be the parent covering points {P1, P2, P3}. To keep the tree roughly balanced, split this group further: P1 and P2 will form a pair under a new node **D**, and P3 will stand as a single child on the other side. Thus, Node B’s left child is **Node D** (which covers P1 and P2) and Node B’s right child is the leaf **P3**. Node D in turn has two children: the leaf **P1** (left) and leaf **P2** (right).  
   - **Right subtree:** Similarly, take the second half of the leaves (P4, P5, P6) as the right subtree. Create an internal node **C** covering points {P4, P5, P6}. Split this group: P4 and P5 form a pair under an internal **Node F**, and P6 is the other child. So Node C’s left child is **Node F** (covers P4 and P5) and its right child is the leaf **P6**. Node F has two children: leaf **P4** (left) and leaf **P5** (right).  
   - **Root:** Finally, create the **root node A** with Node B as its left child and Node C as its right child. The root A covers all points {P1, P2, P3, P4, P5, P6}. Now we have a balanced binary tree: the leaves are the sorted points, and each internal node represents the set of points in its subtree.

3. **Construct and store the sorted Y-list at each node:**  
   Each node will maintain a list of all points in its subtree, sorted by the y-coordinate. We build these lists from the bottom up, merging child lists (which are already sorted by y). For each node in our example:  
   - **Leaf nodes:** Each leaf just contains a single point, so its y-sorted list is just that point. (For example, leaf P1’s list = [P1(1,5)].)  
   - **Node D (covers P1,P2):** Merge the y-lists of leaf P1 and P2. P1’s y=5 and P2’s y=7, so Node D’s Y-list sorted by y is **[P1(1,5), P2(3,7)]** (ascending by y-value 5, 7).  
   - **Node F (covers P4,P5):** Merge leaf P4 and P5 lists. P4’s y=9 and P5’s y=1, so sorted by y it becomes **[P5(7,1), P4(6,9)]** (since 1 < 9).  
   - **Node B (covers P1,P2,P3):** Merge Node D’s list with leaf P3’s list. Node D’s Y-list is [P1(1,5), P2(3,7)] and P3’s list is [P3(4,2)]. Combining and sorting by y: P3 has y=2 (the smallest), then P1 y=5, then P2 y=7. So **Node B’s Y-list = [P3(4,2), P1(1,5), P2(3,7)].**  
   - **Node C (covers P4,P5,P6):** Merge Node F’s list with leaf P6’s list. Node F’s Y-list is [P5(7,1), P4(6,9)] and P6’s list is [P6(9,6)]. Combining them: the y-values are 1 (P5), 6 (P6), 9 (P4). Sorted by y gives **Node C’s Y-list = [P5(7,1), P6(9,6), P4(6,9)].**  
   - **Node A (root, covers P1–P6):** Merge Node B’s list with Node C’s list. Node B’s Y-list is [P3(4,2), P1(1,5), P2(3,7)] and Node C’s is [P5(7,1), P6(9,6), P4(6,9)]. Merging all six points by y-coordinate: the y-values in this combined set are 1, 2, 5, 6, 7, 9. In sorted order we get **Node A’s Y-list = [P5(7,1), P3(4,2), P1(1,5), P6(9,6), P2(3,7), P4(6,9)].** Each internal node now stores a sorted list of the y-coordinates (with associated points) of all points in its subtree.

4. **Add bridge pointers from each parent’s Y-list to its children’s Y-lists:**  
   Now we augment the structure with **bridges**: for each entry in an internal node’s Y-list, we add pointers to the **corresponding entry in the left child’s list and the right child’s list**. A “corresponding” entry here means the same point appearing in the child’s subtree. This allows quick navigation from a parent’s list to the same point in a child’s list. For each internal node in our example, we establish the following bridge connections:  
   - **Node D (parent of leaves P1 and P2):** Node D’s Y-list is [P1, P2]. Its left child is leaf P1 and right child is leaf P2.  
     - The entry for **P1** in Node D points to **P1** in the left child’s list (left child P1 has [P1]) and has no corresponding entry in the right child’s list (since P1 is not in P2’s subtree).  
     - The entry for **P2** in Node D points to **P2** in the right child’s list (right child P2’s list [P2]) and no entry in the left child’s list.  
   - **Node F (parent of leaves P4 and P5):** Node F’s Y-list is [P5, P4]. Left child is P4, right child is P5.  
     - **P5** in Node F’s list points to **P5** in the right child (leaf P5) and not to any entry in the left (P5 is not in P4’s list).  
     - **P4** in Node F’s list points to **P4** in the left child (leaf P4) and has no corresponding entry in the right child’s list.  
   - **Node B (parent of Node D and leaf P3):** Node B’s Y-list is [P3, P1, P2]. The left child is Node D (covers P1,P2) and the right child is leaf P3.  
     - Entries for **P1** and **P2** (which come from the left subtree) point to **P1** and **P2** in Node D’s Y-list, respectively. They have no entries in the right child’s list (P3’s list) because P1 and P2 are not in the right subtree.  
     - The entry for **P3** (which comes from the right subtree) points to **P3** in the right child’s list (leaf P3) and has no corresponding entry in Node D’s list.  
   - **Node C (parent of Node F and leaf P6):** Node C’s Y-list is [P5, P6, P4]. Left child is Node F (covers P4,P5) and right child is leaf P6.  
     - **P5** and **P4** are from the left subtree, so their entries point to **P5** and **P4** in Node F’s list. They have no pointers into the right child’s list (leaf P6) since those points aren’t in the right subtree.  
     - **P6** is from the right subtree, so its entry points to **P6** in the right child’s list. (P6 does not appear in Node F’s left list.)  
   - **Node A (root, parent of Node B and Node C):** Node A’s Y-list is [P5, P3, P1, P6, P2, P4]. Left child is Node B (covers P1,P2,P3) and right child is Node C (covers P4,P5,P6).  
     - Entries for **P3, P1, P2** belong to the left subtree, so each of these points’ entries in Node A point to the same point in **Node B’s** Y-list. They have no corresponding entries in Node C’s list. For example, P3’s entry at the root links to P3 in Node B’s list (and not to Node C).  
     - Entries for **P5, P6, P4** belong to the right subtree, so they each have a bridge to their entry in **Node C’s** Y-list, and no entry in the left child. For example, P5’s entry in Node A points to P5 in Node C’s list (P5 is not in Node B’s list).  

With these bridge pointers in place, every point in an internal node’s y-sorted list can quickly lead us to that same point in either the left or right subtree’s list. We have now constructed a complete 2D range tree with bridges: the tree is balanced by x-coordinate, each node stores a sorted list of points by y-coordinate, and each of those list entries has pointers “bridging” to the corresponding entries in the child nodes. This structure is ready to support efficient 2D range queries.


# k-D Tree

kD Tree (KDT, k-Dimension Tree) is a tree that can efficiently process$k$ The data structure of dimensional space information.

When the number of nodes n is much larger than $2^k$ , the time efficiency of applying kD Tree is very good.

kD Tree has the form of a binary search tree. Each node on the binary search tree corresponds to a point in the k -dimensional space. The points in each of its subtrees are in a k -dimensional hypercube, and all the points in this hypercube are also in this subtree.