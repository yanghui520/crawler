package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
	"strconv"
)
var ageRe=regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">([\d]+)岁</div>`)
var workRe  =regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">工作地:([^<]+)</div>`)

func ParserProfile(content []byte,Name string)engine.ParseRequest  {
	profile := model.Profile{}
	age,err:=strconv.Atoi(extString(content,ageRe))
	if err !=nil{
		profile.Age=age
	}
	profile.Name=Name
	profile.Work=extString(content,workRe)
	res:= engine.ParseRequest{
		Items:[]interface{}{profile},
	}
	return res
}
func extString(content []byte,re *regexp.Regexp) string {
	match:=re.FindSubmatch(content)
	if len(match)>=2{
		return string(match[1])
	}else {
		return ""
	}
}
