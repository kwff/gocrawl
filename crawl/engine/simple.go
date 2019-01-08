package engine

import (
	"crawl/fetch"
	"log"
)

type SimpleEngine struct {

}

func (e SimpleEngine)Run(seeds ...Request) {

	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult, err := worker(r)
		if err!=nil{
			continue
		}
		//把requests放入requests队列
		requests = append(requests, parseResult.Requests...)
		//把Items打印输出
		for _, item := range parseResult.Items {
			log.Printf("got item %v", item)
		}

	}

}

func worker(r Request) (ParseResult, error) {
	log.Printf("fetching %s", r.Url)

	body, err := fetch.Fetch(r.Url)
	if err != nil {
		log.Printf("fetch error fetching url %s %v", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParseFunc(body), nil
}
