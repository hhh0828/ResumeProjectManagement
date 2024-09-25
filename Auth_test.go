package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestMain(t *testing.T) {

	header := Jheader{
		Alg: "HS256",
		Typ: "JWT",
	}
	payload := JPayload{
		Userid:     "testname",
		LoggedinAs: "webmaster",
		Exp:        time.Now().Add(15 * time.Minute),
		Sub:        "test",
	}

	tk := GenerateToken(header, payload)
	fmt.Println(tk)

	ValidateToken(tk)
}

func Test(t *testing.T) {
	feedback := &Feedback{
		Name:    "hyun",
		Email:   "mess@hyunho.com",
		Message: "test",
	}
	data, _ := json.Marshal(feedback)
	newreq, _ := http.NewRequest("POST", "https://www.hyunhoworld.site/submit", strings.NewReader(string(data)))
	newreq.Header.Set("Content-Type", "application/json")
	getres, _ := http.DefaultClient.Do(newreq)
	fmt.Println(getres.Status)
	var p []byte
	getres.Body.Read(p)
	println(string(p))

}
