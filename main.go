package main

import (
	"crawler/engine"
	"crawler/persist"
	"crawler/scheduler"
	"crawler/zhenai/parser"
)

func main() {
	e:=engine.ConcurrentEengine{
		Scheduler:&scheduler.SimpleScheduler{},
		WorkerCount:10,
		ItemChan:persist.ItemSaver(),
	}
	e.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenhun",
		ParserFunc:parser.ParseCityList,
	})
}
