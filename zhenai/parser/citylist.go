package parser

import (
	"awesomeProject/crawler/engine"
	"regexp"
)

const re  = `<a target="_blank" href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a> `
func ParseCityList(content []byte) engine.ParseRequest{
	res:=regexp.MustCompile(re)
	matchs:=res.FindAllSubmatch(content,-1)
	result:= engine.ParseRequest{}
	for _,m:=range matchs{
		result.Items=append(result.Items,string(m[2]))
		result.Requests=append(result.Requests,engine.Request{
			Url:string(m[1]),
			ParserFunc:engine.NilParser,
		})
	}
	return result
}