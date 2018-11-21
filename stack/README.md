# Stack

스택은 객체와 그 객체가 저장되는 순서를 기억하는 방법에 대한 추상 자료형이다.

## 스택 추상 자료형

- 스택 객체: 0개 이상의 원소를 갖는 유한 순서 리스트
- 스택 객체 생성
  - `New(size int) Stack` ::= 스택의 크기가 maxStackSize인 빈스택을 생성하고 반환
- 인터페이스
  - `IsFull() bool` ::= stack element 수 == maxStackSize
  - `IsEmpty() bool` ::= stack element 수 == 0
  - `Push(item interface{}) Stack` ::= (IsFull) ? Stack Full 출력 : 가장 위에 item 삽입 후 stack 반환
  - `Pop() interface{}` ::= (IsEmpty) ? Stack Empty 출력 : 스택 가장 위의 원소 삭제 후 반환
  - `Peek() interface{}` ::= (IsEmpty) ? Stack Empty 출력 : 스택 가장 위의 원소 반환

## 스택 사용

- 시스템 스택: 프로그램에서 사용되는 변수들의 생명주기를 관리하기 위해 사용
- 서브루틴 호출 관리를 위해 사용
- 수식계산에 사용
- 프로그램의 수행 도중 발생되는 인터럽트 처리와 인터럽트 처리가 끝난 후 되돌아갈 명령 수행 지점 저장
- 컴파일러에서 문법 검사를 위해 사용
- 재귀호출시 실행이 끝나고 되돌아갈 실행 주소를 저장
