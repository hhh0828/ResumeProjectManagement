package main

import (
	"encoding/base64"
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

//여기는 request part

// 인증요청시...
type NaverLoginAuth struct {
	//인증과정에 대한 내부 구분값으로 code를 전송해야함.
	Response_type string `json:"response_type"`
	//dev center 등록시 발급받은 Client ID값
	Client_id string `json:"clinet_id"`
	//Call back URI
	Redirect_Uri string `json:"redirect_uri"`
	//Crossorigin - blocking
	State string `json:"state"`
	//No need to send out this,, the scope of access
	Scope string `json:"scope"`
}

type Tokens struct {
	//인증과정 구분값 1 발급, 2 갱신, 3 삭제
	Grant_type string `json:"grante_type"`
	//dev center 등록시 발급받은 Client ID값
	Client_id string `json:"clinet_id"`
	//the code which you get once you get registered your app to naver application
	Client_secret string `json:"client_secret"`
	//로그인 인증후 성공하고 리턴받은 인증코드값
	Code string `json:"code"`
	//크로스사이트공격방지
	State string `json:"state"`
	//인증에성공하고 발급받은 갱신토큰
	Refresh_token string `json:"refresh_token"`
	//발급받은 접근토큰이고 URL 인코딩 적용
	Access_token string `json:"access_token"`
	//인증제공자 이름 NAVER.
	Service_provider string `json:"service_provider"`
}

/*
// Resource owner > Client(API server)//

1. 인증이 필요한 리소스 접근 요청
Access request for authorization.

// Client to Resource Owner
2. 로그인페이지를 응답,
response the page to RO, for Log into Resource server

RO > RS
3.RO need to sign in with login page that provided by Client

//response part



//
*/

type ResponseAuth struct {
	Code              string `json:"code"`
	State             string `json:"state"`
	Error             string `json:"error"`
	Error_description string `json:"error_description"`
}

type ResponseReqToken struct {
	Access_token      string
	Refresh_token     string
	Token_type        string
	Expires_in        int //Sec
	Error             string
	Error_description string
}

////   https://nid.naver.com/oauth2.0/token?grant_type=authorization_code&client_id=jyvqXeaVOVmV&client_secret=527300A0_COq1_XV33cf&code=EIc5bFrl4RibFls1&state=9kgsGTfH4j7IyAkg

func GenerateOauthstate(w http.ResponseWriter) string {
	expiration := time.Now().Add(3600 * time.Second)

	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	//6 칸짜리 []byte 배열
	b := make([]byte, 16)
	rand.Seed(time.Now().Unix())
	//b 슬 에 랜덤값 주입 랜덤 숫자를 넣어줌
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	state := base64.URLEncoding.EncodeToString(b)
	cook := &http.Cookie{
		Name:    "ostate",
		Value:   state,
		Expires: expiration,
	}

	http.SetCookie(w, cook)
	return state
}

func OauthSignin(w http.ResponseWriter, r *http.Request) {
	//response a login page for final client
	naverLoginURL := "https://nid.naver.com/oauth2.0/authorize"
	clientID := "FfJDLNxLwC5I_H3NV7z6"
	redirectURI := "https://www.hyunhoworld.site/index"
	responseType := "code"
	state := GenerateOauthstate(w)

	// 쿼리 파라미터를 구성
	queryParams := url.Values{}
	queryParams.Add("client_id", clientID)
	queryParams.Add("redirect_uri", redirectURI)
	queryParams.Add("response_type", responseType)
	queryParams.Add("state", state)

	// GET 요청 URL 생성
	fullURL := fmt.Sprintf("%s?%s", naverLoginURL, queryParams.Encode())

	// 사용자를 네이버 로그인 페이지로 리다이렉트
	// req, err := http.NewRequest("GET", fullURL, nil)
	// if err != nil {
	// 	log.Println("failed to create request", err)
	// 	return
	// }
	//fmt.Println(req)
	// 네이버로 리다이렉트
	http.Redirect(w, r, fullURL, http.StatusFound)
}

////   https://nid.naver.com/oauth2.0/token?grant_type=authorization_code&client_id=jyvqXeaVOVmV&client_secret=527300A0_COq1_XV33cf&code=EIc5bFrl4RibFls1&state=9kgsGTfH4j7IyAkg

func OauthCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")
	cookie, _ := r.Cookie("ostate")
	if cookie.Value != state {
		log.Println("invalid cookie")
		http.Redirect(w, r, "/index", http.StatusTemporaryRedirect)
		return
	}

	//돌아온 콜백 요청이 괜찮으면, 그길로 aCCESS 토큰 요청
	data := Tokenrequest(code, state)

	req, err := http.NewRequest("POST", "https://nid.naver.com/oauth2.0/token", strings.NewReader(data.Encode()))
	if err != nil {
		log.Println("the error occured with request", err)
	}
	res, _ := http.DefaultClient.Do(req)
	responedtoken := &ResponseReqToken{}
	json.NewDecoder(res.Body).Decode(responedtoken)

	// and send it back to us the Auth code

}

// to Resource server
func Tokenrequest(code string, state string) *url.Values {

	data := &url.Values{}
	data.Set("client_id", "FfJDLNxLwC5I_H3NV7z6")
	data.Set("client_secret", "FJzqN73rzl")
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("state", state)

	return data

}

func OauthTokenValidation() {

}
