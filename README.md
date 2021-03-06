# learn_go

간단한 문법 및 이론에 대해서 정리

## Packages and Imports

```go
// somethings/somethings.go
package sometings

import "fmt"

func SayHello() {
	fmt.Println("Hello")
}

func sayBye() {
	fmt.Println("Bye")
}

// main.go
package main

import (
	"fmt"
	"github.com/JaeHong94/learngo/sometings"
)

func main() {
	sometings.SayHello()
	fmt.Println("Hello world!")
}
```

export처럼 바깥에서 사용하려면 함수명 첫번째 글자를 대문자로 시작해야한다.

## Variables and Constants

```go
package main

import (
	"fmt"
)

const boolean bool = false

func main() {
	name := "jaehong"
	name = "min"
	fmt.Println(name, boolean)
}
```

func안에서는 타입을 축약으로 사용가능, 자동으로 감지, var로 지정된다.  
단, func 밖에서는 축약 사용 불가능  
const = 상수

## Functions Part.1

```go
package main

import (
	"fmt"
	"strings"
)

func multiply(a int, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func main() {
	fmt.Println(multiply(2, 3))
	totalLength, upperName := lenAndUpper("jaehong")
	fmt.Println(totalLength, upperName)
}

```

func의 parameter를 사용하기 위해서는 변수마다 type을 지정해줘야 합니다.  
단, parameter가 모두 같은 타입이라면 multiply(a, b int)로 축약이 가능  
또한 return type도 지정해 주어야합니다.

### 여러가지 return하기

lenAndUpper 함수는 name의 길이와 모든글자를 대문자로 출력해주는 함수입니다.
len은 name의 길이, strings.ToUpper는 name의 글자를 모두 대문자로 변경해주는 기능입니다.  
반환값이 두가지이기 때문에 2개의 변수를 지정합니다. totalLength(name의 길이), upperName(name글자에 대문자)  
go에서는 변수를 만들어두고 사용하지 않으면 오류가 생기므로 주의해야합니다.
2가지를 return받는다고 꼭 변수 두개를 만들어야할 필요는 없습니다.  
totalLength, * := lenAndUpper("jaehong")  
필요하지 않는곳에 *를 사용하면 컴파일할때 무시하고 실행합니다.

## Functions Part.2

```go
package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (length int, uppercase string) {
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	fmt.Println(lenAndUpper("jaehong"))
	defer fmt.Println("World")
	fmt.Println("Hello")
}
```

function에서 return을 명시해주지 않고도 반환하는 방법이 있습니다.  
그것은 function을 작성할때 미리 return할 변수를 만들어주고 코드안에서 데이터를 넣어주는 것입니다.

defer - 연기하다  
defer는 실행을 나중에 할때 사용합니다.  
World와 Hello가 있는데 World에 defer를 추가하니 Hello가 먼저 출력되는것을 볼 수 있습니다.

## for, range, ...args

```go
package main

import "fmt"

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}
```

go에서 for를 사용하는 방법
for 뒤에 *는 index를 의미하지만 출력할 필요가 없기때문에 *를 사용함으로써 무시하고 for문을 사용한다.

```go
func superAdd(numbers ...int) int {
	total := 0
	for i := 0; i < len(numbers); i++ {
		total += i + 1
	}
	return total
}
```

일반적은 for문 사용법처럼 코딩을 만들수도 있지만 각자 편한방법을 사용하면 될것같다.

## If with a Twist

```go
package main

import "fmt"

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge > 19 {
		return true
	}
	return false
}

func main() {
	fmt.Println(canIDrink(18))
}
```

if문안에 변수를 생성해서 사용했는데 이는 바깥에서 만들어 사용할 수도 있습니다.

```go
func canIDrink(age int) bool {
	koreanAge := age + 2
	if koreanAge > 19 {
		return true
	}
	return false
}
```

하지만 이렇게 사용한다면 koreanAge를 다른 곳에서도 사용할 수 있다는 생각을 들게 할 수도 있기 때문에 if문에서만 사용한다는 인식을 주기 위해서 전자에 방식을 사용합니다.

## Switch

```go
func canIDrink(age int) bool {
	switch age {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

func canIDrink(age int) bool {
	switch  {
	case age < 10:
		return false
	case age == 18:
		return true
	case age > 50:
		return false
	}
	return false
}

func canIDrink(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 20:
		return true
	}
	return false
}
```

switch는 다른 언어들과 같이 같은 방식으로 만들 수도 있고 if {} else if {} else {}의 방식, if를 배울 때 활용한 switch안에서만 사용하는 변수를 만들어 사용하는 방법이 있다.  
3가지 방식이 있으니 상황에 맞게 사용하면 될 것 같다.

## Pointers

```go
package main

import "fmt"

func main() {
	a := 2
	b := 4
	fmt.Println(a, &a)
	fmt.Println(b, &b)
}
```

변수 앞에 &를 붙이면 데이터의 pointer 값을 출력합니다.  
하지만 pointer 값을 b라는 데이터에 넣은것이므로 b의 pointer 값은 다릅니다.

```go
package main

import "fmt"

func main() {
	a := 2
	b := &a
	*b = 20
	fmt.Println(a)
}
```

*은 메모리 주소에서 해당 주소의 현재값으로 포인터를 역참조(dereference) 합니다.  
*b = 20을 하게되면 a의 데이터값도 같이 바뀌게 됩니다.
