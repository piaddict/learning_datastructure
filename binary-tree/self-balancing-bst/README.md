# 자가 균형 이진 탐색 트리

자가 균형 이진 탐색 트리는 삽입과 삭제가 일어나는 경우에 자동으로 균형의 차이를 작게 유지하는 노드 기반 이진 탐색 트리이다

이진 탐색트리는 동적 자료구조이므로 삽입/삭제 연산과 함께 구조가 계속 변하게 되고 이는 성능에 영향을 준다.
균형이 맞지않는 트리는 최악의 경우 linked list 수준의 탐색 경로를 갖게된다.

따라서 트리의 균형을 유지하도록 만드는데, 만약 완전 균형을 유지하게 하려면 `O(n)`의 시간복잡도가 필요하다.

효율성을 위해 (거의) 완전한 균형 트리를 구성하는 여러 구현방식이 존재한다.
아래 방식의 경우 탐색/삽입/삭제 연산의 시간복잡도가 `O(logn)`이다.

## 공통연산: Rotate(x)

대상 `x`, 부모 `p`, 회전방향과 같은쪽 `x`의 자식 `T2`, 조부모 `g`

```text
    p                              x
   / \       Right Rotation       /  \
  x   T3     – – – – – – – >     T1   p
 / \                                 / \
T1  T2                              T2  T3
```

- Right Rotation
  - p.left = T2
  - x.right = p
  - g.child = x || p가 루트라면 tree.root = x

```text
    p                              x
   / \       Left Rotation        /  \
  T3  x     – – – – – – – >      p    T2
     / \                        / \
    T1  T2                     T3  T1
```

- Left Rotation
  - p.right = T2
  - x.left = p
  - g.child = x || p가 루트라면 tree.root = x

## Splay 트리

| Algorithm | Average | Worst case |
| --- | --- | --- |
| Space | O(n) | O(n) |
| Search | O(log n) | amortized O(log n) |
| Insert | O(log n) | amortized O(log n) |
| Delete | O(log n) | amortized O(log n) |

Splay 트리는 자주 탐색하는 키를 가진 노드를 루트에 가깝게 위치하도록 구성한 BS 트리이다.
이 트리는 Splay 연산을 적용하여 최근에 접근한 노드를 루트에 위치시켜 트리를 재구성한다.

### Splay 연산

Splay(x) 연산은 임의의 노드 x를 루트로 만드는 연산이다.

Rotate(x) 연산은 x를 x의 부모위치로 올리는 연산이다.

> 최근에 접근한 노드(x), x의 부모노드(p), x의 조부모노드(g)

- x가 루트이면 종료한다
- x의 부모 p가 루트이면 Rotate(x) 실행후 종료: **Zig**
- x의 조부모가 g인경우
  - g -> p의 방향과 p -> x의 방향이 같으면, Rotate(p) -> Rotate(x) 실행: **Zig-Zig**
  - g -> p의 방향과 p -> x의 방향이 다르면, Rotate(x)를 두변 실행: **Zig-Zag**
- x가 루트가 될때까지 반복한다

```go
func Splay(x *node) {
    for x.p != nil {
        p := x.p
        g := p.p
        if g != nil {
          if (x == p.left) == (p == g.left) {
            Rotate(p)
          } else {
            Rotate(x)
          }
        }
        Rotate(x);
    }
}
```

### Zig / Zig-Zig / Zig-Zag

- Zig
  - p가 트리의 루트이면 p와 x의 간선연결(edge joining)을 회전시킨다
- Zig-Zig
  - p가 루트가 아니고 x와 p가 동일한 방향의 왼쪽(오른쪽) 자식인경우
  - p와 조부모 g와의 간선연결을 회전시키고(`Rotate(p)`)
  - x와 p의 간선연결을 회전시킨다(`Rotate(x)`)
- Zig-Zag
  - p가 루트가 아니고 x가 왼쪽(오른쪽)자식, p가 오른쪽(왼쪽) 자식인경우
  - x와 p의 간선 연결을 회전시키고(`Rotate(x)`)
  - x와 x의 새로운 부모 p와의 간선 연결을 회전시킨다(`Rotate(x)`)

## AVL 트리

| Algorithm | Average | Worst case |
| --- | --- | --- |
| Space | O(n) | O(n) |
| Search | O(log n) | O(log n) |
| Insert | O(log n) | O(log n) |
| Delete | O(log n) | O(log n) |

노드의 삽입과 삭제가 일어날 때 노드 키 값과 서브트리 키값 사이의 관계를 유지하며 균형을 만들기는 쉽지않다.

BS 트리의 탐색 성능을 개선하면서 균형을 유지하는데 따르는 수고를 줄이려면
제한조건을 완화하여 트리가 (완전한) 균형이 아닌것을 허용해야 한다.

AVL 트리는 거의 완전한 균형 트리의 한 형태로 높이가 균형잡힌 높이 균형트리이다.

트리의 각 노드 Vi에 대해서 다음과 같은 성질을 만족할 때 그 트리는 높이가 균형 잡힌 트리이다.

> 노드 Vi의 왼쪽 서브트리 높이와 Vi의 오른쪽 서브트리 높이가 최대 1만큼 차이난다

AVL 트리는 직접 탐색 성능이 좋다. 비록 완전히 균형 잡히지는 않았지만 탐색 경로의 길이는 거의 차이가 없다.
만일 삽입이나 삭제 수행이 AVL 트리의 균형을 깨면 트리를 균형잡힌 상태로 재구성하는데 이를 rotation이라 한다.

### Height

- height: Leaf 노드로부터 거리, 즉 leaf 노드의 높이는 0이다. null 노드의 경우 -1이다
- Balance Factor: 왼쪽 서브트리의 높이 - 오른쪽 서브트리의 높이
- 삽입, 삭제 이후 해당노드로부터 루트까지 올라가며 높이를 갱신하면서 BF를 계산한다
- BF가 허용범위를 넘어서면 갱신을 중지하고 Rotation을 실행한다

### Single Rotation

불균형이 바깥쪽 서브트리에서 발생하는 경우, Single Rotation을 실행한다

Left Left Case / Right Right Case

```text
    y                              x
   / \       Right Rotation       /  \
  x   T3     – – – – – – – >     T1   y
 / \         < - - - - - - -         / \
T1  T2       Left Rotation         T2  T3
```

### Double Rotation

불균형이 안쪽 서브트리에서 발생하는 경우, 반대 방향의 Single Rotation을 연속으로 실행한다

Left Right Case

```text
     z                               z                             x
    / \                            /   \                          /  \
   y   T4  Left Rotate (y)        x    T4    Right Rotate(z)    y      z
  / \      - - - - - - - - ->    /  \        - - - - - - - ->  / \    / \
T1   x                          y    T3                      T1  T2 T3  T4
    / \                        / \
  T2   T3                    T1   T2
```

Right Left Case

```text
   z                              z                              x
  / \                            / \                            /  \
T1   y   Right Rotate (y)      T1   x        Left Rotate(z)   z      y
    / \  - - - - - - - - ->       /  \     - - - - - - - ->  / \    / \
   x   T4                        T2   y                    T1  T2  T3  T4
  / \                                /  \
T2   T3                             T3   T4
```

## BB(Weight-balanced) 트리

거의 완전히 균형 잡힌 트리의 다른 종류로 무게가 균형잡힌 트리가 있다.
이 트리를 무게 균형트리 또는 BB 트리라고 한다.

무게 균형트리는 양쪽 서브트리 무게가 균형을 유지하는 트리이다.
즉, 임의의 노드 x에 대해 다음식이 만족하도록 트리를 구성한다

`sqrt(2) - 1 < 왼쪽 서브트리 무게 / 오른쪽 서브트리 무게 < sqrt(2) + 1`

BB 트리의 경우 서브트리 높이가 아닌 서브트리 노드 개수에 제한을 둔 것이다.

BB 트리의 평균 탐색 길이 역시 균형트리와 크게 차이나지 않는다.
또한 삽입이나 삭제에 의해 균형을 잃으면 rotation 연산을 통해 균형을 유지한다.
