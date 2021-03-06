# Queue

큐는 한쪽에서는 삽입 연산이, 반대쪽에서는 삭제 연산이 가능한 자료구조이다.

- 데이터의 삭제(remove) 연산은 큐의 `front`에서 가능하다
- 데이터는 큐의 `front`에서만 읽을 수(peek) 있다
- 데이터의 삽입(add) 연산은 큐의 `rear`에서 가능하다
- 이를 FIFO(First In, First Out)이라 한다

## 큐의 추상자료형

- 큐 객체의 정의: 0개 이상의 원소를 갖는 유한 순서 리스트
- 큐 객체생성
  - `New(size int) Queue` ::= 크기가 maxQueueSize인 빈 큐을 생성하고 반환
- 인터페이스
  - `IsFull() bool` ::= queue element 수 == maxQueueSize
  - `IsEmpty() bool` ::= queue element 수 == 0
  - `Add(item interface{}) Queue` ::= (IsFull) ? Queue Full 출력 : rear에 item 삽입 후 queue 반환
  - `Remove() interface{}` ::= (IsEmpty) ? Queue Empty 출력 : front 원소 삭제 후 반환
  - `Peek() interface{}` ::= (IsEmpty) ? Queue Empty 출력 : front 원소 반환

## 큐의 구현

- bounded queue는 고정 크기의 자료만 입력할 수 있다
- 동적배열을 사용하여 변동 크기의 자료를 입력할 수 있다
- 연결 리스트를 이용해서 구현할 수 있다
- 구현방식의 선택은 자료크기와 자료크기의 변동성에 따라 다르다 (amortized cost vs memory usage)

## 큐의 응용

- 대기열, 탐색
- First-Come First-Served 스케줄링
- Round Robin 스케줄링
- ...

## 원형 큐

큐에 원소 하나를 삽입하면 rear 값이 증가하고, 원소 하나를 삭제하면 front 값이 증가한다.
그러다 rear 값과 front 값이 같아지면 큐는 빈 상태가 되고 추가적인 원소 삽입이 불가능해질 수 있다.

이런 문제를 해결하기 위해 원형 큐를 구현할 수 있다.
원형 큐는 큐의 입구와 출구 부분을 연결시킨 형태이다.

따라서 길이가 n인 큐에서 rear의 위치가 `n-1`이 되면 다음 위치가 `0`이 되도록 구현해야 한다.

원형 큐에서 원소를 편집하는 경우에는 다음과 같이 위치를 할당하면 된다.

- `rear := (rear + 1) % n`
- `front := (front + 1) % n`

처음은 rear, front 모두 0으로 시작하고 첫 원소는 1번 주소부터 삽입한다.

- IsEmpty: `front == rear`
- IsFull: `front == (rear + 1) % n`

큐가 비어있는지/가득차있는지 확인하기 위해 원형 큐의 공간 하나는 비워둔 상태로 유지한다.

### 원형 큐 구현

- [circular_queue.go](./circular_queue.go)
- [circular_queue_test.go](./circular_queue_test.go)

## Deque

> Double-ended queue

양쪽 끝에서 삽입과 삭제가 모두 가능한 자료 구조이다.

### 인터페이스

- insert element at rear
- insert element at front
- remove last element
- remove first element
- examine last element
- examine first element

### 구현

- 큐와 동일하게 동적배열이나 이중 연결리스트를 이용하여 구현하며 구현방식 선택은 상황에 따라 다르다
