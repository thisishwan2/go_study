package something

// 컴파일러는 자동적으로 main 패키지를 찾아 main 함수를 찾는다.
// main은 오로지 컴파일을 위해 필요한 함수이다.

import (
	"fmt"
	"strings"
)

// import "fmt" // fmt는 Go의 패키지중 한가지로 formatting을 담당한다.

// go에서는 자바와 다르게 매개변수의 타입과 함수의 리턴 타입을 뒤에 나타낸다.
func multiply(a int, b int) int { // == (a, b int)
	return a * b
}

// go는 여러개의 리턴값을 반환할 수 있다.
func returnManyType(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// 매개변수를 여러개 받는 함수 선어
func repeatMe(words ...string) { // ... 을 통해서 복수의 매개변수를 받는다는 것을 선어
	fmt.Println(words)
}

// naked return
func lenAndUpper(name string) (length int, uppercase string) { // 리턴 변수를 미리 정의
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

// defer : 함수가 끝난 후에 실행되는 함수
func lenAndUpperDefer(name string) (length int, uppercase string) { // 리턴 변수를 미리 정의
	defer fmt.Println("I'm done") // 함수가 끝나고 실행되는 함수
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

// for 이용
func superAdd(numbers ...int) int {
	total := 0
	for index, number := range numbers {
		fmt.Println(index, number)
		total += number
	}
	return total
}

// if-else
func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 { // if else문 내부에서만 사용하기 위해서 내부에 변수 선언을 할 수 있다.
		return false
	}
	return true
}

func useSwitch(age int) bool {
	switch age { // if-else와 같이 사용
	case 10:
		return true
	case 20:
		return false
	}
	return false
}

type person struct {
	name    string
	age     int
	favFood []string
}

func exam() {

	// 외부 함수 사용
	fmt.Println("Hello, World!") // 외부에서 import 한 패키지의 함수를 사용하려면 항상 함수명의 시작은 대문자이다.
	SayHello()

	// 변수 선언,
	var tmp string = "a"
	tmp2 := "b" // 이러한 형태로 변수를 선언하면 편하게 선언할 수 있다.(type은 자동으로 추론된다.) (이러한 축약형은 func 안에서만 사용 가능)
	tmp = "aaa"
	fmt.Println(tmp)
	fmt.Println(tmp2)

	// 상수 선언
	const name string = "ran" // 상수는 변하지 않는 값이다.

	// 함수 사용
	fmt.Println(multiply(2, 2))

	// 여러개의 리턴값을 반환하는 함수 사용
	totalLength, upperName := returnManyType("ran")
	fmt.Println(totalLength, upperName)

	//totalLength, _ := returnManyType("ran")
	//fmt.Println(totalLength)

	// 매개변수를 여러개 받는 함수 사용
	repeatMe("a", "b", "c", "d", "e")

	lenght, upp := lenAndUpper("ran")
	fmt.Println(lenght, upp)

	// defer 함수의 출력문이 함수 종료후에 나온다.
	lenght, upp = lenAndUpperDefer("ran")
	fmt.Println(lenght, upp)

	// for 이용
	res := superAdd(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(res)

	fmt.Println(canIDrink(16))

	// a=10, b=2
	a := 2
	b := a
	a = 10
	fmt.Println(a, b)
	fmt.Println(&a, &b) // &a는 a의 메모리 주소를 나타낸다.

	a_ := 1
	b_ := &a_
	fmt.Println(&a_, b_) // b_는 a_의 메모리 주소를 나타낸다.
	fmt.Println(*b_)     // *b_는 b_의 메모리 주소에 있는 값을 나타낸다. a_의 값이다.

	// 배열 선언(크기가 고정)
	names := [5]string{"a", "b", "c"}
	fmt.Println(names)
	fmt.Println(len(names))
	names[3] = "d"
	fmt.Println(names)

	// slice 선언(크기가 가변)
	names2 := []string{"a", "b", "c"}
	fmt.Println(names2)
	names2 = append(names2, "d") // append 함수는 슬라이스 명과, 추가할 요소를 매개변수로 받는다.
	fmt.Println(names2)

	// map 선언
	mapEx := map[string]string{"name": "ran", "age": "25"} // map[key의 데이터타입]value의 데이터타입
	fmt.Println(mapEx)
	for key, value := range mapEx {
		fmt.Println(key, value)
	}

	// struct 선언
	favFood := []string{"kimchi", "ramen"}
	// peopleInfo := person{"ran", 25, favFood} // 이런 방식의 코드 스타일은 직관적이지 않음
	peopleInfo := person{name: "ran", age: 25, favFood: favFood} // 이 방식이 직관적
	fmt.Println(peopleInfo)
	fmt.Println(peopleInfo.name)
	// go에는 클래스가 없고, 생성자도 없다.

}