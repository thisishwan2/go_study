package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct { // 외부에서 사용해야하기 때문에 대문자
	// 필드명도 외부에서 접근하기 위해서는 대문자.
	// 소문자는 private이다.
	owner   string
	balance int
}

var errNoMoney = errors.New("can't withdraw")

// struct 생성자 함수
func NewAccount(owner string) *Account { // 포인터를 리턴한다.
	account := Account{owner: owner, balance: 0}
	return &account // 주소 반환
}

// receiver를 설정한다. 만약 struct에 포인터를 설정하지 않으면,
// 새로운 Account 객체가 생성되는 것이기 때문에 수정사항이 반영되지 않음
// 입금 함수
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

func (a Account) Balance() int {
	return a.balance
}

// 출금 함수. 0이하면 출금이 안된다.(예외처리)
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil //null과 같이 예외처리를 하면 반환해줘야한다.
}

// owner 변경
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.owner, " ", a.balance)

}
