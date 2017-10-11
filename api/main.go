package main

import (
	"fmt"
	"github.com/oelmekki/chromessr"
	"log"
	"net/http"
	"os"
)

func InitSSR() {
	err := chromessr.Init()
	if err != nil {
		fmt.Println("Can't initiate SSR")
		fmt.Println(err)
		os.Exit(1)
	}
}

func getPort() (port string) {
	port = os.Getenv("CHROMESSR_PORT")
	if len(port) > 0 {
		port = ":" + port
	} else {
		port = ":3001"
	}

	return
}

func main() {
	InitSSR()
	defer chromessr.CloseBrowser()

	http.Handle("/retrieve", PageRetriever{})

	log.Fatal(http.ListenAndServe(getPort(), nil))
}
