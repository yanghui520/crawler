package parser

import (
	"crawler/engine"
	"regexp"
)

const cityRe  = `<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a>`
func ParseCity(content []byte) engine.ParseRequest{
	res:=regexp.MustCompile(cityRe)
	matchs:=res.FindAllSubmatch(content,-1)
	result:= engine.ParseRequest{}
	for _,m:=range matchs{
		result.Items=append(result.Items,"user"+string(m[2]))
		result.Requests=append(result.Requests,engine.Request{
			Url:string(m[1]),
			ParserFunc:engine.NilParser,
		})
	}
	return result
}
