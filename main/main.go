// 컴파일러는 자동적으로 main 패키지를 찾아 main 함수를 찾는다.
// main은 오로지 컴파일을 위해 필요한 함수이다.
package main

import (
	"fmt"
	"go_study/accounts"
	"go_study/dict"
)

func main() {
	account := accounts.NewAccount("ran")
	fmt.Println(account)
	account.Deposit(10)
	fmt.Println(account.Balance())
	account.Withdraw(2)
	fmt.Println(account.Balance())
	// 아무일도 발생하지 않는다.(0 미만이기 때문에)
	//account.Withdraw(12)
	err := account.Withdraw(12) // 에러 반환
	if err != nil {
		// log.Fatalln(err) // 에러가 발생하면 프로그램을 종료한다.
		fmt.Println(err)
	}
	fmt.Println(account.Balance())
	account.ChangeOwner("CHOI")
	fmt.Println(account)
	fmt.Println(account.Owner())

	// String() 메서드를 호출한다.(String을 정의한 경우 내부적으로 찾아서 호출)
	fmt.Println(account)

	dictionary := dict.Dictionary{}
	dictionary["hello"] = "Greeting"
	fmt.Println(dictionary)

	definition, err := dictionary.Search("hello")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

	err = dictionary.Add("bye", "byebye")
	fmt.Println(err)

	dictionary.Update("hello", "hi")
	fmt.Println(dictionary)

	dictionary.Delete("bye")
	fmt.Println(dictionary)

}
