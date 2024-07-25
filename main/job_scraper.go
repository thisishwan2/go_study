package main

import (
	"encoding/csv"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type extractedJob struct {
	title       string
	location    string
	companyName string
}

var baseUrl string

func scrape(term string) {
	var jobs []extractedJob
	baseUrl = "https://www.saramin.co.kr/zf_user/search/recruit?&searchword=" + term
	c := make(chan []extractedJob)
	totalPages := getPages()

	for i := 0; i < totalPages; i++ {
		go getPage(i, c)
	}

	for i := 0; i < totalPages; i++ {
		extractedJob := <-c
		jobs = append(jobs, extractedJob...)
	}

	fmt.Println(jobs)
	writeJobs(jobs)
	fmt.Println("done extracting")
}

func getPage(number int, mainChannel chan<- []extractedJob) {
	// pageUrl := baseUrl + "&recruitPage=" + string(number+1) 이 방식은 정수를 문자열로 변환하지 못함
	var jobs []extractedJob
	c := make(chan extractedJob)
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
		go extractJob(card, c)
	})

	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}
	mainChannel <- jobs
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

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("value")
	title := card.Find(".area_job").Find(".job_tit>a").Text()
	location := cleanString(card.Find(".area_job").Find(".job_condition>a").Text())
	companyName := cleanString(card.Find(".area_job.job_tit > a > span").Text())
	fmt.Println(id)
	c <- extractedJob{title: title, location: location, companyName: companyName}
}

// csv로 저장
func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv") // 파일 생성
	checkErr(err)

	w := csv.NewWriter(file) // csv writer 생성
	defer w.Flush()          // main 함수가 끝나기 직전에 실행 (파일에 데이터 입력)

	headers := []string{"title", "location", "companyName"}

	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		w.Write([]string{job.title, job.location, job.companyName})
	}
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
