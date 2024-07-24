package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type extractedJob struct {
	title       string
	location    string
	companyName string
}

var baseUrl string = "https://www.saramin.co.kr/zf_user/search/recruit?&searchword=java"

func main() {
	var jobs []extractedJob
	totalPages := getPages()

	for i := 0; i < totalPages; i++ {
		extract := getPage(i)
		jobs = append(jobs, extract...)
	}
	fmt.Println(jobs)
}

func getPage(number int) []extractedJob {
	// pageUrl := baseUrl + "&recruitPage=" + string(number+1) 이 방식은 정수를 문자열로 변환하지 못함
	var jobs []extractedJob
	pageUrl := baseUrl + "&recruitPage=" + strconv.Itoa(number+1)
	fmt.Println("Requesting: ", pageUrl)
	res, err := http.Get(pageUrl)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".item_recruit")
	searchCards.Each(func(i int, card *goquery.Selection) {
		job := extractJob(card)
		jobs = append(jobs, job)

	})
	return jobs
}

func extractJob(card *goquery.Selection) extractedJob {
	id, _ := card.Attr("value")
	title := card.Find(".area_job").Find(".job_tit>a").Text()
	location := cleanString(card.Find(".area_job").Find(".job_condition>a").Text())
	companyName := cleanString(card.Find(".area_job.job_tit > a > span").Text())
	fmt.Println(id)
	return extractedJob{title: title, location: location, companyName: companyName}
}

func getPages() int {
	pages := 0
	res, err := http.Get(baseUrl)

	checkErr(err)
	checkCode(res)

	// defer는 main 마지막에 실행한다.(finally와 비슷)
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status: ", res.StatusCode)
	}
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
