package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestXxx(t *testing.T) {

	req, _ := http.NewRequest("GET", "http://172.30.1.19:8700/metrics", nil)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}
