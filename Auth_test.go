package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
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

	Nreq := &NaverLoginAuth{}

	Nreq.Client_id = "FfJDLNxLwC5I_H3NV7z6"
	Nreq.Redirect_Uri = "https://wwww.hyunhoworld.site/index"
	Nreq.Response_type = "code"
	state := "test crossss"
	EncState := base64.URLEncoding.EncodeToString([]byte(state))
	Nreq.State = EncState
	data, _ := json.Marshal(Nreq)
	req, err := http.NewRequest("POST", "https://nid.naver.com/oauth2.0/authorize", strings.NewReader(string(data)))
	if err != nil {
		log.Println("failed to create request", err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("err occrured during get response from logon server", err)
	}
	authcode := new(ResponseAuth)
	json.NewDecoder(res.Body).Decode(&authcode)
	if authcode.State != EncState {
		fmt.Println(authcode)
	}
	fmt.Println(authcode)

}

// std test run
// func Test(t *testing.T) {
// 	user := &User{
// 		GivenPermission: "WebMaster",
// 		Userid:          "soyeon",
// 		Userpw:          "lovelove",
// 		Useremail:       "hhhcjswo@naver.com",
// 	}
// 	data, _ := json.Marshal(user)
// 	newreq, _ := http.NewRequest("POST", "https://www.hyunhoworld.site/joinus", strings.NewReader(string(data)))
// 	newreq.Header.Set("Content-Type", "application/json")
// 	getres, _ := http.DefaultClient.Do(newreq)
// 	fmt.Println(getres.Status)
// 	var p []byte
// 	getres.Body.Read(p)
// 	println(string(p))

// }

// type User struct {
// 	gorm.Model
// 	GivenPermission string
// 	Userid          string
// 	Userpw          string
// 	Useremail       string
// }
