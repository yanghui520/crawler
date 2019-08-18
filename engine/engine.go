package engine

import (
	"crawler/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _,r:=range seeds{
		requests=append(requests,r)
	}
	for len(requests)>0 {
		r:=requests[0]
		requests=requests[1:]
		log.Printf("url %s",r.Url)
		body,err:=fetcher.Fetch(r.Url)
		if err!=nil {
			log.Printf("fet error %s:%v",r.Url,err)
			continue
		}
		parseResult:= r.ParserFunc(body)
		requests=append(requests,parseResult.Requests...)
		for _,item:=range parseResult.Items{
			log.Printf("got item %s",item)
		}
	}

	
}
