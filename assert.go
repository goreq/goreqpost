package goreqpost

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/goreq/goreq"
	"github.com/spyzhov/ajson"
)

type JSON struct {
	*ajson.Node
}
type JSONChecker func(node []*JSON)

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

		nodesToJson := make([]*JSON, len(nodes))
		for _, node := range nodes {
			nodesToJson = append(nodesToJson, &JSON{node})
		}
		checker(nodesToJson)

	}
}

func AssertBody(regex string) goreq.AfterResponseHandler {
	return func(resp *http.Response) {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			return
		}

		matched, err := regexp.MatchString(regex, string(body))
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			return
		}

		if !matched {
			fmt.Printf("Error: invalid body\n")
			return
		}

	}
}
