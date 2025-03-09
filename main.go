package main

import (
	"fmt"
	"net/http"
	
)

// result struct
type requestResult struct {
	url string
	status string
}


//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	results := make(map[string]string)
	c := make(chan requestResult)
	urls := []string{
		"https://www.naver.com/", 
		"https://www.google.com/",
		"https://www.daum.net/",
		"https://www.nate.com/",
		"https://www.yahoo.com/",
		"https://www.bing.com/",
		"https://www.duckduckgo.com/",
		"https://www.dogpile.com/",
		"https://www.yippy.com/",
		"https://www.ask.com/",
		"https://www.aol.com/",
		"https://www.lycos.com/",
		"https://www.altavista.com/",
	}

	for _, url := range urls {
		go hitURL(url, c)
	}

	for i:=0; i<len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL(url string, c chan<- requestResult) {
	resp, err := http.Get(url)
	status := "OK"

	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- requestResult{url: url, status: status}
 }


 