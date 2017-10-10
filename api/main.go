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

func main() {
	InitSSR()
	defer chromessr.CloseBrowser()

	http.Handle("/retrieve", PageRetriever{})

	log.Fatal(http.ListenAndServe(":3001", nil))
}
