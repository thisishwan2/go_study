// 컴파일러는 자동적으로 main 패키지를 찾아 main 함수를 찾는다.
// main은 오로지 컴파일을 위해 필요한 함수이다.
package main

import (
	"fmt"
	"go_study/something"
)

// import "fmt" // fmt는 Go의 패키지중 한가지로 formatting을 담당한다.

func main() {

	// 외부 함수 사용
	fmt.Println("Hello, World!") // 외부에서 import 한 패키지의 함수를 사용하려면 항상 함수명의 시작은 대문자이다.
	something.SayHello()

	// 변수 선언,
	var tmp string = "a"
	tmp2 := "b" // 이러한 형태로 변수를 선언하면 편하게 선언할 수 있다.(type은 자동으로 추론된다.) (이러한 축약형은 func 안에서만 사용 가능)
	tmp = "aaa"
	fmt.Println(tmp)
	fmt.Println(tmp2)

	// 상수 선언
	const name string = "ran" // 상수는 변하지 않는 값이다.

}
