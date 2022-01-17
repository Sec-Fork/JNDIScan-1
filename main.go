package main

import (
	"flag"
	"fmt"
	"github.com/EmYiQing/JNDIScan/config"
	"github.com/EmYiQing/JNDIScan/core"
	"github.com/EmYiQing/JNDIScan/log"
	"github.com/EmYiQing/JNDIScan/model"
	"github.com/EmYiQing/JNDIScan/util"
	"os"
	"os/signal"
	"syscall"
)

var (
	ResultChan chan *model.Result
	RenderChan chan *model.Result
)

func main() {
	initApp()
	go core.StartFakeServer(&ResultChan)
	go core.StartOutput(&RenderChan)
	go startApp()
	core.PrintPayload()
	wait()
}

func initApp() {
	core.PrintLogo(config.GetAuthors())
	parserInput()
	ResultChan = make(chan *model.Result, config.DefaultChannelSize)
	RenderChan = make(chan *model.Result, config.DefaultChannelSize)
}

func startApp() {
	for {
		select {
		case res := <-ResultChan:
			info := fmt.Sprintf("%s->%s", res.Name, res.Host)
			log.Info(info)
			data := &model.Result{
				Host:   res.Host,
				Name:   res.Name,
				Finger: res.Finger,
				Path:   res.Path,
			}
			RenderChan <- data
		}
	}
}

func parserInput() {
	var (
		port int
		help bool
	)
	flag.IntVar(&port, "p", 8001, "server port")
	flag.BoolVar(&help, "help", false, "help info")
	flag.Parse()
	if help {
		flag.PrintDefaults()
		return
	}
	if !util.CheckPort(port) {
		os.Exit(-1)
	}
	log.Info("use port: %d", port)
	config.Port = port
}

func wait() {
	sign := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sign
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()
	<-done
}
