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
