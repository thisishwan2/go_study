// 컴파일러는 자동적으로 main 패키지를 찾아 main 함수를 찾는다.
// main은 오로지 컴파일을 위해 필요한 함수이다.
package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFail = errors.New("Request failed")

type res struct {
	url    string
	status string
}

// url checker를 만든다.
func main() {

	// map 선언 방식
	var results = map[string]string{}
	// var results = make(map[string]string) // make는 map을 비어있는 상태로

	// channel 선언
	c := make(chan res)

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

	//// 1. loop를 이용한다. (속도가 느리다.)
	//for _, url := range urls {
	//	res := hitURL(url)
	//
	//	if res == nil {
	//		results[url] = "OK"
	//	} else {
	//		results[url] = "FAILED"
	//	}
	//	fmt.Println("Checking: ", url)
	//}
	//
	//for url, results := range results {
	//	fmt.Println(url, results)
	//}
	//
	//// 2. goroutine을 이용하여 순차적인 실행을 동시에 처리 한다.
	//// goroutine은 병렬 처리를 위해 사용한다.
	//// goroutine은 main 함수가 종료되면, 끝난다.
	//// 지금 아래와 같이 go를 선언하면 해당 함수가 반환되기전에 main 함수가 종료되어 버려서 gorutine은 소명된다.
	//for _, url := range urls {
	//	go hitURL(url)
	//}
	//// 만약 여기서 sleep을 건다면 sleep을 건 시간동안은 goroutine이 실행되어 결과를 확인할 수 있다.
	//time.Sleep(time.Second * 2)
	//
	//// channel은 gorutine이 데이터를 주고 받는 방법이다.
	//cha := make(chan string) // channel을 만들어 준다.
	//for _, url := range urls {
	//	go hitURL2(url, cha) // urls의 갯수만큼 goroutine을 만들어준다.(병렬작업)
	//}
	//
	//// 메세지를 받는 작업은 blocking 작업
	//// channel에서 값을 받는다.(이 경우에는 channel로 부터 값이 올때까지 main함수를 대기한다.)
	//// channel에 모아둔 메세지중 2개의 메세지만 받는다.(<-c의 개수만큼 채널로부터 메세지를 받아오게 됨)(goroutine 보다 더 많은 개수를 받으면 deadlock이 발생한다.)
	//fmt.Println(<-c)
	//fmt.Println(<-c)

	// 3. goroutine과 channerl로 url checker를 만든다.
	for _, url := range urls {
		go hitURLWithChannel(url, c)
	}

	// 이 코드를 통해 main 함수의 종료를 막는다.
	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	// 결과를 출력한다.
	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURLWithChannel(url string, c chan<- res) { //c chan<- res 이런식으로 선언하면 해당 채널은 메세지를 수신할 수 없고, 전송만 할 수 있음을 선언하는 것임
	fmt.Println("Checking: ", url)
	response, err := http.Get(url)
	if err != nil || response.StatusCode >= 400 {
		c <- res{url: url, status: "FAILED"}
	}
	c <- res{url: url, status: "OK"}

}

func hitURL2(url string, c chan string) { // 어떤 타입의 데이터를 채널로 주고받을지를 정의한다.
	response, err := http.Get(url)

	if err != nil || response.StatusCode >= 400 {
		c <- "err" + url // channel을 통해 메세지를 보내는 방법
	}
	c <- url
}

func hitURL(url string) error {
	response, err := http.Get(url)

	// 에러가 발생하거나, 상태코드가 400 이상이면 에러를 반환한다.
	if err != nil || response.StatusCode >= 400 {
		return errRequestFail
	}
	return nil
}
