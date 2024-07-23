// 컴파일러는 자동적으로 main 패키지를 찾아 main 함수를 찾는다.
// main은 오로지 컴파일을 위해 필요한 함수이다.
package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

var errRequestFail = errors.New("Request failed")

// url checker를 만든다.
func main() {

	// map 선언 방식
	//var results = map[string]string{}
	var results = make(map[string]string) // make는 map을 비어있는 상태로

	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	// 1. loop를 이용한다. (속도가 느리다.)
	for _, url := range urls {
		res := hitURL(url)

		if res == nil {
			results[url] = "OK"
		} else {
			results[url] = "FAILED"
		}
		fmt.Println("Checking: ", url)
	}

	for url, results := range results {
		fmt.Println(url, results)
	}

	// 2. goroutine을 이용하여 순차적인 실행을 동시에 처리 한다.
	// goroutine은 병렬 처리를 위해 사용한다.
	// goroutine은 main 함수가 종료되면, 끝난다.
	// 지금 아래와 같이 go를 선언하면 해당 함수가 반환되기전에 main 함수가 종료되어 버려서 gorutine은 소명된다.
	for _, url := range urls {
		go hitURL(url)
	}
	// 만약 여기서 sleep을 건다면 sleep을 건 시간동안은 goroutine이 실행되어 결과를 확인할 수 있다.
	time.Sleep(time.Second * 2)
}

func hitURL(url string) error {
	response, err := http.Get(url)

	// 에러가 발생하거나, 상태코드가 400 이상이면 에러를 반환한다.
	if err != nil || response.StatusCode >= 400 {
		return errRequestFail
	}
	return nil
}
