package gojsonq

import (
	"fmt"
	"net/http"
	"strings"
)

func UselessFunc(s string) {
	req, _ := http.NewRequest("GET", "google.com", strings.NewReader(s))
	res, _ := http.DefaultClient.Do(req)
	fmt.Println(res.Body)
}
