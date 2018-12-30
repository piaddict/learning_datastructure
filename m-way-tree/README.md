# m-way Tree

트리의 노드가 m개 이하의 가지를 가질 수 있는 트리를 m-way 트리라고 한다.

## 구분

- m-way Tree
  - **Trie**
  - m-way Search Tree
    - **B tree**
      - B* Tree
      - **B+ Tree**
      - 2-3 Tree
        - 2-3-4 Tree

## m-way Search Tree

m-way 탐색트리는 이진 트리를 확장한 것이다.
이진 탐색트리의 제한을 따르되 두 개 이상 m개 이하의 자식을 가질 수 있다.

조건은 다음과 같다

- Pi는 서브트리에 대한 포인터이고 Ki는 키값이다, n <= m-1
- i = 0 ... n-2 인 i에 대해 Ki < Ki+1이 성립한다
- i = 0 ... n-1 인 i에 대해 Pi가 가리키는 서브트리의 모든 키값은 Ki의 키값 보다작다
- Pn이 가리키는 서브트리의 모든 키값은 Kn-1의 키값보다 크다
- i = 0 ... n 인 i에 대해 Pi가 가리키는 서브트리는 m-way 탐색트리이다

```go
type (
    Entry struct {
        key   string
        value int
    }

    Node struct {
        parent *Node
        e      []Entry
        child  []*Node
    }
)
```

### m-way Search Tree: 탐색

- m-way 트리를 위한 노드는 중첩 구조체로 정의한다
- m-way 트리의 노드에는 서브트리를 위한 키값과 포인터가 반복해서 나타난다
- 트리의 노드에는 각 노드가 갖고있는 서브트리의 개수를 n으로 표현한다
- 키 값을 찾는과정은 이진 탐색 트리에서 키값을 찾는 과정을 확장하면 된다
