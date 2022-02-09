package main

import (
	"encoding/json"
	"fmt"
	"jasonblanchard/depfootprint/pkg/npmjs"
	"log"
)

func main() {
	npmjsFetcher := &npmjs.NpmJS{}

	// pkg, err := npmjsFetcher.Get("express")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	deps, err := npmjsFetcher.Tree("remix")
	if err != nil {
		log.Fatal(err)
	}

	resultJson, err := json.Marshal(deps)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(resultJson))
}
