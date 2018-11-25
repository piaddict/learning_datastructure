# Queue

큐는 한쪽에서는 삽입 연산이, 반대쪽에서는 삭제 연산이 가능한 자료구조이다.

원소의 삭제 연산이 이루어지는 곳을 `front`라 하고, 삽입 연산이 이루어지는 곳을 `rear`라 한다.

## 큐의 추상자료형

- 큐 객체의 정의: 0개 이상의 원소를 갖는 유한 순서 리스트
- 큐 객체생성
  - `New(size int) Queue` ::= 크기가 maxQueueSize인 빈큐을 생성하고 반환
- 인터페이스
  - `IsFull() bool` ::= queue element 수 == maxQueueSize
  - `IsEmpty() bool` ::= queue element 수 == 0
  - `Add(item interface{}) Queue` ::= (IsFull) ? Queue Full 출력 : 가장 위에 item 삽입 후 queue 반환
  - `Remove() interface{}` ::= (IsEmpty) ? Queue Empty 출력 : 스택 가장 위의 원소 삭제 후 반환
  - `Peek() interface{}` ::= (IsEmpty) ? Queue Empty 출력 : 스택 가장 위의 원소 반환
