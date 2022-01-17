package core

import (
	"github.com/EmYiQing/JNDIScan/model"
	"io/ioutil"
	"sync"
)

var (
	resultList []*model.Result
	lock       sync.Mutex
)

func StartOutput(renderChan *chan *model.Result) {
	go listenData(renderChan)
}

func listenData(renderChan *chan *model.Result) {
	for {
		select {
		case res := <-*renderChan:
			lock.Lock()
			resultList = append(resultList, res)
			data := RenderHtml(resultList)
			_ = ioutil.WriteFile("result.html", data, 0666)
			lock.Unlock()
		}
	}
}
