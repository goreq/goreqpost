package main

import (
	"fmt"
	"net/http"

	"github.com/goreq/goreq"
	"github.com/goreq/goreqpost"
)

func main() {
	goreq.Get("https://my-json-server.typicode.com/hadihammurabi/flutter-webservice/contacts", nil,
		goreq.WithAfterResponseHandler(
			goreqpost.AssertStatus(http.StatusOK),
			goreqpost.AssertBody(".*"),
			goreqpost.AssertJSON("$..id", jsonChecker),
		),
	)
}

func jsonChecker(nodes []*goreqpost.JSON) {
	ids := []float64{1, 2, 3}
	for i, node := range nodes {
		if node.MustNumeric() != ids[i] {
			fmt.Printf("%f is not %f\n", node.MustNumeric(), ids[i])
			break
		}
	}
}
