package engine

import (
	"log"
)

type ConcurrentEengine struct {
	Scheduler Scheduler
	WorkerCount int
	ItemChan chan interface{}
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkChan(chan Request)
}

func (e *ConcurrentEengine) Run(seeds ...Request)  {

	in:= make(chan Request)
	out:=make(chan ParseRequest)
	e.Scheduler.ConfigureMasterWorkChan(in)
	for i:=0;i<=e.WorkerCount ;i++  {
		createWorker(in,out)
	}
	for _,r:=range seeds{
		e.Scheduler.Submit(r)
	}
	for  {
		requset:=<-out
		for _,item:=range requset.Items{
			log.Printf("got itme %v",item)
			go func() {e.ItemChan<-item}()
		}
		for _,request:= range requset.Requests{
			e.Scheduler.Submit(request)
		}
	}
}
func createWorker(in chan Request,out chan ParseRequest)  {
	go func() {
		for  {
			request:=<-in
			result,err:= worker(request)
			if err!=nil {
				continue
			}
			out<-result
		}
	}()
}
