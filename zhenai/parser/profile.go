package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
	"strconv"
)
const ageRe=`<div class="m-btn purple" data-v-bff6f798="">([\d]+)岁</div>`
const workRe  =`<div class="m-btn purple" data-v-bff6f798="">工作地:([^<]+)</div>`

func ParserProfile(content []byte)engine.ParseRequest  {
	profile := model.Profile{}

	re:=regexp.MustCompile(ageRe)
	match:=re.FindSubmatch(content)
	if match !=nil{
		age,err:=strconv.Atoi(string(match[1]))
		if err !=nil{
			profile.Age=age
		}
	}

	re=regexp.MustCompile(ageRe)
	match=re.FindSubmatch(content)
	if match !=nil{
		profile.Work=string(match[1])
	}
}
