# Hash Table

해시 테이블은 key에 대한 해시 값으로 value를 탐색/삽입/삭제하며,
key-value 쌍에 따라 동적으로 크기가 변하는 연관 배열(associate array)이다.

해시값을 사용한 해시함수로 색인을 찾으므로 탐색/삽입/삭제의 시간복잡도가 낮지만,
해시 범위에 해당되는 메모리를 미리 할당해야 하므로 공간 효율적이지는 않다.

## Direct Address Table

키의 전체 개수와 동일한 크기의 bucket을 사용하고, key값을 직접 색인으로 사용하는 테이블이다.

최선의 탐색/삽입/삭제 연산을 할 수 있으나, 메모리 효율이 크게 떨어진다.

## 해시 테이블의 해시 함수 구현

### 제산 잔여(mod function)

값을 나눈 나머지를 해시값으로 출력한다.

이 경우 어떤 값이 주어져도 제수(하위 비트)에만 전적으로 의존하고 해시 결과도 고르게 분포되지 않을 가능성이 많다.

즉, 제수는 해시함수 테이블의 성능을 크게 좌우하게 되는데 보통 키의 수 3배가 적당하다.
(적재율 30%정도 까지는 충돌이 많이 발생하지 않는다)

### Binning

키 값의 범위를 해시 테이블 크기로 나눈다.

이 경우 상위 비트의 분포가 고르지 못하다면 해시값도 고르게 분포되지 않는 문제가 발생한다.

### 중간 제곱(mid-squre)

정수 키 값을 사용하는 경우 사용하기 좋은 방식이다.

키 값을 제곱한 후, 결과의 중간에 있는 n비트를 취하여 `0` 부터 `2^n - 1` 범위를 해시 결과로 사용한다.

## 해시 충돌 회피

### Separate Chaining

충돌하는 해시값이 있다면 같은 해시값을 가지는 value를 포인터를 통해 연속적으로 연결하여 보관한다.

이 경우 수행시간은 최초 해시값을 이용하여 key-value를 찾는 경우 상수시간이 걸리고,
만약 충돌하는 key가 있어 여러개의 value가 존재하는 경우 추가적으로 선형시간의 수행시간이 필요하다.

### Open Addressing

#### Linear Probing

충돌하는 해시값이 있다면 바로 정해진 값의 인덱스를 이동한 뒤 데이터를 저장한다.
이동한 이후에도 데이터가 있다면 빈 공간이 있을 때 까지 이동하여 해당 위치에 저장한다.

Linear Probing은 데이터들이 특정 위치에 밀집하는 primary clustering 문제가 발생한다.

#### Quadratic Probing

Linear probing의 primary clustering을 회피하기 충돌시 이동거리를 2차식의 형태로 만든다.

Quadratic probing도 시작 해시값이 같은경우 연속한 탐사위치는 같으므로 secondary clustering 문제가 발생한다.

#### Double Hasing

Double Hashing은 탐사할 해시값의 규칙성을 제거하여 clustering을 방지하는 기법이다.

두 종류의 해시함수를 준비해서 해시값을 얻을 때와 충돌시 이동거리를 얻을 때 사용할 방식을 구분한다.

단, 두 해시함수는 연관성이 없어야 한다. (ex: mod function의 경우 서로소이어야 한다)

## Java의 HashMap

<https://d2.naver.com/helloworld/831311>

Java에서는 기본적으로 각 객체의 `hashCode()` 메소드가 반환하는 값(32bit 정수)을 해시값으로 사용한다.

HashMap에서 해시 테이블 해시충돌을 해결하기 위해 사용하는 방식은 Separate Chaining이다.

Open Addressing은 데이터 삭제가 효율적이지 않고,
데이터의 수가 일정이상이라면 Open Addressing이 더 느리기 때문이다.

Separate Chaining을 위해 해시 테이블의 해시충돌을 줄이기 위한 보조 해시함수를 사용한다.

### HashMap의 Separate Chaining

Java8 이후로 Separate Chaining에서 충돌값이 일정이상이라면 Linked List 대신 Tree를 사용한다.

### Hash Bucket 동적확장

해시 버킷 크기가 작다면 해시충돌로 인한 성능상 손실이 발생한다.
따라서 key-value 쌍이 증가하면 해시 버킷의 크기를 두배로 늘린다.

해시 버킷의 기본값은 16이고 임계점이 될 때마다 두배씩 늘리며 최대 2^30까지 늘어난다.

그러나 동적확장이 발생할 때마다 새로운 Separate Chaining을 구성해야 하는 문제가 있으므로,
데이터의 개수를 예측할 수 있는 경우 초기 해시 버킷 개수를 생성자에서 지정하여 문제를 회피할 수 있다.

해시 버킷 크기를 늘리는 임계점은 `현재 데이터 개수 = load factor(기본값 0.75) * 현재의 해시 버킷수` 이다.

load factor역시 생성자에서 지정 가능하다.

### 보조 해시함수

`index = X.hashCOde() % M`의 M 값이 소수인경우 index 분포가 가장 균등할 수 있다.

그러나 M값이 소수가 아니므로 별도의 보조 해시 함수를 통해 균등화 해야한다.

Java8에서는 상위 16비트 값을 XOR연산하는 보조 해시 함수를 사용한다

```java
static final int hash(Object key) {
    int h;
    return (key == null) ? 0 : (h = key.hashCode()) ^ (h >>> 16);
}
```

### String 객체 해시함수

String 객체에 대한 해시 함수 수행시간은 문자열 길이에 비례한다.

문자열은 앞 글자의 분포가 균등하지 않아 해시 충돌이 일어날 가능성이 크므로
(URL의 경우 https, 일반 단어의 경우 e로 시작하는 단어 ...)
Java8의 경우 다음 방식을 사용한다.

```java
public int hashCode() {  
    int h = hash;
    if (h == 0 && value.length > 0) {
        char val[] = value;

        for (int i = 0; i < value.length; i++) {
            h = 31 * h + val[i];
        }
        hash = h;
    }
    return h;
}
```

31을 사용하는 이유는, 31이 소수이며 `31=32N-N`으로 어떤수에 32를 곱한값은 shift연산으로 쉽게 구현가능하기 때문이다.
