package main

import (
	"encoding/json"
	"fmt"
	"github.com/oelmekki/chromessr"
	"log"
	"net/http"
)

type PageRetriever struct {
	Url         string `json:"url"`
	Placeholder string `json:"placeholder"`
}

func (handler PageRetriever) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	encoded := request.FormValue("payload")
	fmt.Println(encoded)
	err := json.Unmarshal([]byte(encoded), &handler)

	if err != nil {
		log.Printf("Wrong payload for retriever: %v\n", err)
		return
	}

	content, err := chromessr.Retrieve(handler.Url, handler.Placeholder)
	if err != nil {
		log.Printf("Error while retrieving content: %v\n", err)
		return
	}

	fmt.Println(content)

	response.Write([]byte(content))
}
