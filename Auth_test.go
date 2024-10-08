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

// std test run
func Test(t *testing.T) {
	user := &User{
		GivenPermission: "WebMaster",
		Userid:          "hhhcjswo",
		Userpw:          "d7349801",
		Useremail:       "hhhcjswo@naver.com",
	}
	data, _ := json.Marshal(user)
	newreq, _ := http.NewRequest("POST", "https://www.hyunhoworld.site/joinus", strings.NewReader(string(data)))
	newreq.Header.Set("Content-Type", "application/json")
	getres, _ := http.DefaultClient.Do(newreq)
	fmt.Println(getres.Status)
	var p []byte
	getres.Body.Read(p)
	println(string(p))

}

// type User struct {
// 	gorm.Model
// 	GivenPermission string
// 	Userid          string
// 	Userpw          string
// 	Useremail       string
// }
