# Binary Tree

트리에 속한 모든 노드의 차수(degree)가 2 이하인 트리를 이진트리라고 한다.

이진트리에서 각 레벨(level)이 허용되는 최대 개수 노드를 가질 때 **포화 이진트리**라고 한다.
따라서 높이(height)가 `k`인 포화 이진트리의 노드 개수는 다음과 같다: `2^k -1`

왼쪽부터 노드가 순서대로 채워진 트리를 **완전 이진트리**라고 한다.
(가운데 빈 곳이 없어야 하지만 해당 레벨(level))을 다 채울필요는 없음)

- 노드(node): 트리를 구성하는 항목
- 차수(degree): 특정 노드의 자식 노드의 수(진출 차수)
- 루트(root): 부모가 존재하지 않는 노드(진입 차수가 0인 노드)
- 리프(leaf): 자식이 존재하지 않는 노드(진출 차수가 0인 노드)
- 형제(sibling): 같은 부모 노드를 갖는 노드
- 높이(height): 루트 노드에서 단말(terminal) 노드까지의 깊이
- 레벨(level): 루트로부터 해당 노드까지 이어진 경로의 길이(루트의 레벨: 0 또는 1)

## 구분

- Binary Tree
  - Threaded binary tree
  - **Heap**
  - **Binary Search Tree**
  - Self-balancing Binary Search Tree
    - Splay Tree
    - AVL Tree
    - BB(Weight-balanced) Tree
    - **Red-Black Tree**

## 구현

이진트리는 연속된 배열에 저장할 수 있다.
이 방법은 포화 이진트리나 완전 이진트리인 경우 낭비되는 공간이 없거나 적어 효율적이다.
하지만 노드의 차수가 1씩만 존재하는 경우 기억장소의 낭비가 커진다.

이러한 이유로 이진 트리는 보통 연결리스트로 구현한다.

이진트리를 구현하기 위한 연결리스트의 노드는 데이터, 왼쪽 서브트리 포인터, 오른쪽 서브트리 포인터로 구성된다.

## 연산

### 순회(traverse)

트리의 각 노드를 빠지지 않고 중복없이 한 번씩 방문하는 것을 순회라고 한다.

순회 순서에 따라 다음으로 구분하여 부른다

- 전위 순회(PLR)
  - 루트 방문
  - 왼쪽 서브트리를 전위 순회로 순회
  - 오른쪽 서브트리를 전위 순회로 순회

- 중위 순회(LPR)
  - 왼쪽 서브트리를 중위 순회로 순회
  - 루트를 방문
  - 오른쪽 서브트리를 중위 순회로 순회

- 후위 순회(LRP)
  - 왼쪽 서브트리를 후위 순회로 순회
  - 오른쪽 서브트리를 후위 순회로 순회
  - 루트를 방문

트리의 순회법은 수식의 표기법과 밀접하게 괸련있다.

순회는 순환 호출로 구현하거나 스택을 통해 비순환으로 구현한다.

```go
func preorder(node *node) {
    if node == nil {
        return
    }
    fmt.Println(node.data)
    preorder(node.left)
    preorder(node.right)
}

func inorder(node *node) {
    if node == nil {
        return
    }
    inorder(node.left)
    fmt.Println(node.data)
    inorder(node.right)
}

func postorder(node *node) {
    if node == nil {
        return
    }
    postorder(node.left)
    postorder(node.right)
    fmt.Println(node.data)
}
```

### 삽입/삭제

- 노드를 삽입하는 경우 부모 노드의 링크가 자식노드를 가리키도록 하면 된다
- 노드를 삭제하려는 경우 해당 노드를 가리키는 부모의 포인터를 null 처리하고 노드의 메모리를 반환한다
- 삭제하려는 노드가 leaf가 아닌 경우(자식 노드가 있는 경우) 자식 노드들을 모두 삭제해야 한다
- 정점을 삭제하는경우 자식노드 중 어느 노드를 정점노드위치로 이동시킬지는 정해야 한다

### 노드 개수

- 노드가 null이 아닌 동안 반복하며 자식 노드의 개수를 반환하는 순환 호출로 구현할 수 있다

  ```go
  func getNodeCount(node *node) int {
      if node == nil {
          return 0
      }
      result := 1
      result += getNodeCount(node.left) + getNodeCount(node.right)
      return result
  }
  ```

- leaf 노드의 개수를 세려면 leaf 노드에 도달했을 때 1을 반환하는 순환 호출로 구현할 수 있다

  ```go
  func getLeafCount(node *node) int {
      result := 0
      if node == nil {
          return 0
      } else if node.left == nil && node.right == nil {
          return 1
      }
      result += getLeafCount(node.left) + getLeafCount(node.right)
      return result
  }
  ```

## 일반 트리를 이진 트리로 변환

- 주어진 트리에 대해 각 노드의 형제들을 연결한다 (수평연결)
- 각 노드에 대하여 가장 왼쪽 링크만 남기고 모두 제거한다 (수직연결)
- 루트 노드는 반드시 왼쪽 자식 하나만 가져야 한다
