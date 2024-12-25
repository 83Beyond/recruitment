package spider

import (
	"github.com/83Beyond/recruitment/network"
	"github.com/83Beyond/recruitment/website"
)

var (
	ReqQueue  = make(chan network.Request, 10)
	RespQueue chan network.Response
)

type Scheduler struct{}

func (s *Scheduler) Schedule() {
	for _, r := range website.AllWebsites {
		req := r
		go func() {
			for i := 0; i < req.ScrapePage; i++ {
				req.CurrentPage = i
				ReqQueue <- req
			}
		}()
	}
}

func (s *Scheduler) Start() {
	go func() {
		for req := range ReqQueue {
			resp := StartRequest(req)
			req.Parse(resp)
		}
	}()
}
