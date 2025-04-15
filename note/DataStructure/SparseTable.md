ST table (Sparse Table) is a data structure used to solve the problem of repeatable contribution .

RMQ is the abbreviation of Range Maximum/Minimum Query, which means the maximum (minimum) value of a range. 

The ST table is based on the idea of ​​multiplication , and can do \Theta(n\log n) preprocessing and \Theta(1) answering each query. However, it does not support modification operations.

Based on the doubling idea, we consider how to find the maximum value of the interval. It can be found that if we follow the general doubling process and jump 2^i steps each time, the complexity of the query is still \Theta(\log n) , which is not better than the segment tree. On the contrary, the preprocessing step is slower than the segment tree.

We found that \max(x,x)=x , that is, interval maximum is a problem with the property of "repeatable contribution". Even if the preprocessing intervals used to solve the problem overlap, as long as the union of these intervals is the interval being solved, the final calculated answer is correct.

If we simulate manually, we can find that we can use at most two preprocessed intervals to cover the query interval, which means that the time complexity of the query can be reduced to \Theta(1) , which is very effective when dealing with questions with a large number of queries.

```c
#include <algorithm>
#include <iostream>
using namespace std;
constexpr int MAXN = 2000001;
constexpr int logN = 21;
int f[MAXN][logN + 1], Logn[MAXN + 1];

void pre() {  // 准备工作，初始化
  Logn[1] = 0;
  Logn[2] = 1;
  for (int i = 3; i < MAXN; i++) {
    Logn[i] = Logn[i / 2] + 1;
  }
}

int main() {
  cin.tie(nullptr)->sync_with_stdio(false);
  int n, m;
  cin >> n >> m;
  for (int i = 1; i <= n; i++) cin >> f[i][0];
  pre();
  for (int j = 1; j <= logN; j++)
    for (int i = 1; i + (1 << j) - 1 <= n; i++)
      f[i][j] = max(f[i][j - 1], f[i + (1 << (j - 1))][j - 1]);  // ST表具体实现
  for (int i = 1; i <= m; i++) {
    int x, y;
    cin >> x >> y;
    int s = Logn[y - x + 1];
    cout << max(f[x][s], f[y - (1 << s) + 1][s]) << '\n';
  }
  return 0;
}
```