package main

import (
	"encoding/json"
	"fmt"
	"jasonblanchard/depfootprint/pkg/npmjs"
	"log"
)

func main() {
	npmjsFetcher := &npmjs.NpmJS{}

	deps, err := npmjsFetcher.Tree("express")
	if err != nil {
		log.Fatal(err)
	}

	resultJson, err := json.Marshal(deps)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(resultJson))
}
