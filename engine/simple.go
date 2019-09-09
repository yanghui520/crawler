package engine

import (
	"crawler/fetcher"
	"log"
)

type SimpleEngine struct {}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _,r:=range seeds{
		requests=append(requests,r)
	}
	for len(requests)>0 {
		r:=requests[0]
		requests=requests[1:]
		parseResult,err:=worker(r)
		if err!=nil {
			continue
		}
		requests=append(requests,parseResult.Requests...)
		for _,item:=range parseResult.Items{
			log.Printf("got item %s",item)
		}
	}
}
func worker(r Request) (ParseRequest,error)  {
	log.Printf("url %s",r.Url)
	body,err:=fetcher.Fetch(r.Url)
	if err!=nil {
		log.Printf("fet error %s:%v",r.Url,err)
		return ParseRequest{},err
	}
	return r.ParserFunc(body),nil
}
