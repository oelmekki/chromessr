package chromessr

import (
	"github.com/raff/godet"
	"log"
	"os/exec"
	"time"
)

var remote *godet.RemoteDebugger
var enqueue chan Page

func Init() (err error) {
	err = startBrowser()
	if err != nil {
		log.Println(err)
		return
	}

	time.Sleep(time.Second * 10) // allow chrome to start
	enqueue = make(chan Page)

	remote, err = godet.Connect("127.0.0.1:9222", false)
	if err != nil {
		return
	}

	err = remote.SetBlockedURLs("*.jpg", "*.png", "*.gif")
	if err != nil {
		return
	}

	go processQueue()
	return err
}

func startBrowser() (err error) {
	err = exec.Command("google-chrome-stable", "--headless", "--disable-gpu", "--remote-debugging-address=127.0.0.1", "--remote-debugging-port=9222").Start()
	return
}
