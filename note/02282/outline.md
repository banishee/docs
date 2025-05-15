The oral exam is without preparation. That means that you come in and then we (examiner and external examiner) ask questions. The exam is approximately 20 minutes.  We will ask you questions about different subjects in the curriculum. The questions can fx be of the following form (but are not limited to):

* Which algorithms/data structures do you know that solves the following problem?
* How does algorithm/data structure X work?  
* Run algorithm/data structure X on the following example,
* Analyse the running time of algorithm/data structure X,
* Analyse the space usage/query time of /data structure X,
* Argue correctness of algorithm/data structure X.

# 1. Hashing:

Want a random, crazy, chaotic function that maps a large universe to a small range. The
function should distribute the items “evenly.”

• Low collision probability: for any x≠y, we want Pr(h(x) = h(y)) to be small.
• Fast evaluation.
• Small space.

## 1.1 Dictionary problem

A dynamic set of integers $S \subseteq U$ to support `lookup`, `insert` and `delete`.

We want a compact data structure with fast operations.

### Chained Hashing

<p align="center"><img src=".data/hashing_chained.png" alt="pic" width="40%" /></p>

> Choose universal hash function $h$ from $U$ to ${0, ..,m-1}$.
> 
> Initialize an array $A[0, ..., m-1]$.
> 
> $A[i]$ stores a linked list containing the keys in $S$ whose hash value is $i$.
>
> Space: $O(m + n) = O(n)$
> 
> Time: $O(1 + |A[h(x)]|)$
>

## 1.2 Static dictionary problem

A set $S \subseteq U$ of size $n$ to support LOOKUP(x).

We want a collision-free hash function on S with $O(n)$ space and $O(1)$ worst-case lookup time.

### JKS-Scheme

<p align="center"><img src=".data/hashing_FKS.png" alt="pic" width="40%" /></p>

> Two-level solution:
> * At level 1 use solution with many collisions and linear space.
> * Resolve each collisions at level 1 with collision-free solution at level 2.
>
> Spce: $O(n)$ and $O(n)$ expected preprocessing time.
> 
> Time: $O(1)$ worst-case time per lookup.
>

## 1.3 String matching

Given strings S and P, determine if P is a substring in S. 

|P| = m, |S| = n.


### Karp-Rabin Fingerprint

> Let S be a string of length n. 
>
> Let p is a prime number. Pick uniformly at random integer z ∈{0, ..., p-1}.
>
> The fingerprint of $S$ is the polynomial over the field $Z_p$ evaluated at the random integer $z$.
> 
>$$\varphi_{p,z}(S) = \left( \sum_{i=1}^{n} S[i] \cdot z^{n - i} \right) \mod p$$
> 

### Karp-Rabin Algorithm

> Pick $p >= m^2$
> 
> Compute $\varphi(P)$
>
> Compute and compare $\varphi(S[i, i+m-1])$ with $\varphi(P)$ for all i.
>
> If fingerprints match, verify using brute-force comparison. 
>
> Time: $O(m + n + Fm)$


# Week12:

https://priyadarshanghosh26.medium.com/christofides-algorithm-the-secret-weapon-for-route-optimization-d2b9ec68d66e