# 2-3 Tree

| Algorithm | Average | Worst case |
| --- | --- | --- |
| Space | O(n) | O(n) |
| Search | O(log n) | O(log n) |
| Insert | O(log n) | O(log n) |
| Delete | O(log n) | O(log n) |

차수가 2 또는 3인 내부 노드를 갖는 트리를 2-3 트리라고 한다.
이때 차수가 2인 노드를 2-노드, 3인 노드를 3-노드라고 부른다.

즉, 2-3 트리는 2-노드와 3-노드만으로 구성된 특수한 형태의 B 트리이다.

- 각각의 내부 노드는 2-노드이거나 3-노드이다
  - 2-노드는 한 개의 키값을, 3-노드는 두 개의 키값을 갖는다

- l-child, m-child를 각각 2-노드의 좌측 및 중간 자식이라고 하고, l-key가 노드의 키값이라고 한다면
  - 루트가 l-child인 모든 2-3 서브트리 키값은 l-key 보다 작고
  - 루트가 m-child인 모든 2-3 서브트리 키값은 l-key 보다 크다

- l-child, m-child, 및 r-child를 각각 3-노드의 왼쪽, 중간 및 오른쪽 자식이라 하고, l-key및 r-key를 이 노드의 두 키 값이라 하면
  - l-key < r-key
  - 루트가 l-child인 모든 2-3 서브트리 키값은 l-key 보다 작다
  - 루트가 m-child인 모든 2-3 서브트리 키값은 r-key 보다 작고, l-key 보다 크다
  - 루트가 r-child인 모든 2-3 서브트리 데이터는 r-key 보다 크다

- 모든 leaf 노드는 같은 레벨에 있다

높이가 h인 2-3 트리의 키 개수는 `2^h - 1`개에서 `3^h - 1` 개 사이이다

## 탐색

2-3 트리의 탐색은 이진트리 탐색 알고리즘을 확장하여 구현할 수 있다

```go
type TwoThree struct {
  lkey   int
  rkey   int
  lchild *TwoThree
  mchild *TwoThree
  rchild *TwoThree
}

func Search(t *TwoThree, int x) *TwoThree {
    for t != nil {
        if x가 왼쪽 키값보다 작은 경우 {
            t = t.lchild
        } else if x가 왼쪽 키값보다 크고 오른쪽 키값보다 작은 경우 {
            t = t.mchild
        } else if x가 오른쪽 키값보다 큰 경우 {
            t = t.rchild
        } else if x가 이들 중 어느 것과 같은 경우 {
            return t
        }
    }
    return nil
}
```

## 2-3-4 Tree

2-3 트리를 확장하여 네 개의 자식을 가진 4-노드를 허용하는 것이 2-3-4 트리이다

2-3-4 탐색트리는 다음 조건을 만족한다

- 각각의 내부 노드는 2-노드, 3-노드 및 4-노드이다

- lchild, mchild는 각각 2-노드의 자식이라 하고 lkey를 키값이라 하면
  - 루트가 lchild인 모든 2-3-4 서브트리 키값은 lkey 보다 작고
  - 루트가 mchild인 모든 2-3-4 서브트리 키값은 lkey 보다 크다

- lchild, mchild, rchild를 각각 3-노드의 자식이라 하고 lkey, rkey를 노드의 키값이라 하면 다음이 성립한다
  - lkey < mkey
  - 루트가 lchild인 모든 2-3-4 서브트리 키값은 lkey보다 작다
  - 루트가 lmchild인 모든 2-3-4 서브트리 키값은 lkey보다 크고 mkey보다 작다
  - 루트가 rmchild인 모든 2-3-4 서브트리 키값은 mkey보다 크다

- lchild, lmchild, rmchild 및 rchild를 각각 4-노드의 자식이라 하고 lkey, mkey, rkey를 이 노드의 키값이라 하면
  - lkey < mkey < rkey
  - 루트가 lchild인 모든 2-3-4 서브트리 키값은 lkey보다 작다
  - 루트가 lmchild인 모든 2-3-4 서브트리 키값은 mkey보다 작고 lkey보다 크다
  - 루트가 rmchild인 모든 2-3-4 서브트리 키값은 rkey보다 작고 mkey보다 크다
  - 루트가 rchild인 모든 2-3-4 서브트리 키값은 rkey보다 크다

- 모든 leaf 노드들은 같은 레벨에 있다
