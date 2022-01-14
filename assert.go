package goreqpost

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/goreq/goreq"
	"github.com/spyzhov/ajson"
)

type JSONChecker func(node []*ajson.Node)

func AssertJSON(jsonPath string, checker JSONChecker) goreq.AfterResponseHandler {
	return func(resp *http.Response) {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			return
		}

		nodes, err := ajson.JSONPath(body, jsonPath)
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			return
		}

		if len(nodes) <= 0 {
			fmt.Println("Error: path not found")
			return
		}

		checker(nodes)

	}
}
