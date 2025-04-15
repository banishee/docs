ST table (Sparse Table) is a data structure used to solve the problem of repeatable contribution .

RMQ is the abbreviation of Range Maximum/Minimum Query, which means the maximum (minimum) value of a range. 

The ST table is based on the idea of ​​multiplication , and can do \Theta(n\log n) preprocessing and \Theta(1) answering each query. However, it does not support modification operations.

Based on the doubling idea, we consider how to find the maximum value of the interval. It can be found that if we follow the general doubling process and jump 2^i steps each time, the complexity of the query is still \Theta(\log n) , which is not better than the segment tree. On the contrary, the preprocessing step is slower than the segment tree.

We found that \max(x,x)=x , that is, interval maximum is a problem with the property of "repeatable contribution". Even if the preprocessing intervals used to solve the problem overlap, as long as the union of these intervals is the interval being solved, the final calculated answer is correct.

If we simulate manually, we can find that we can use at most two preprocessed intervals to cover the query interval, which means that the time complexity of the query can be reduced to \Theta(1) , which is very effective when dealing with questions with a large number of queries.