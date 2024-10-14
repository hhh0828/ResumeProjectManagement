package main

import (
	_ "encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
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

func OauthSignin(w http.ResponseWriter, r *http.Request) {
	//response a login page for final client
	naverLoginURL := "https://nid.naver.com/oauth2.0/authorize"
	clientID := "FfJDLNxLwC5I_H3NV7z6"
	redirectURI := "https://www.hyunhoworld.site/index"
	responseType := "code"
	state := "test_crossss"

	// 쿼리 파라미터를 구성
	queryParams := url.Values{}
	queryParams.Add("client_id", clientID)
	queryParams.Add("redirect_uri", redirectURI)
	queryParams.Add("response_type", responseType)
	queryParams.Add("state", state)

	// GET 요청 URL 생성
	fullURL := fmt.Sprintf("%s?%s", naverLoginURL, queryParams.Encode())

	// 사용자를 네이버 로그인 페이지로 리다이렉트
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		log.Println("failed to create request", err)
		return
	}
	fmt.Println(fullURL)
	// 네이버로 리다이렉트
	http.Redirect(w, req, fullURL, http.StatusFound)
}

func OauthSigninpost(w http.ResponseWriter, r *http.Request) {
	naverLoginURL := "https://nid.naver.com/oauth2.0/authorize"
	clientID := "FfJDLNxLwC5I_H3NV7z6"
	redirectURI := "https://www.hyunhoworld.site/index"
	responseType := "code"
	state := "test_crossss"

	// POST 요청에 사용할 폼 데이터를 URL 인코딩 형식으로 설정
	formData := url.Values{}
	formData.Add("client_id", clientID)
	formData.Add("redirect_uri", redirectURI)
	formData.Add("response_type", responseType)
	formData.Add("state", state)

	// POST 요청 생성
	req, err := http.NewRequest("POST", naverLoginURL, strings.NewReader(formData.Encode()))
	if err != nil {
		log.Println("failed to create request", err)
		return
	}

	// Content-Type 설정
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// 요청 실행
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Println("error occurred during POST request", err)
		return
	}
	defer res.Body.Close()

	// 리다이렉트가 처리되는지 확인 (응답 헤더에서 Location 확인)
	if res.StatusCode == http.StatusFound || res.StatusCode == http.StatusMovedPermanently {
		fmt.Println("Redirected to:", res.Header.Get("Location"))
	} else {
		fmt.Println("Response Status:", res.Status)
		fmt.Println("Response Headers:", res.Header)
	}
}

// func (Nreq *NaverLoginAuth) SigninAuthRequest() string {

// 	Nreq.Client_id = "FfJDLNxLwC5I_H3NV7z6"
// 	Nreq.Redirect_Uri = "https://wwww.hyunhoworld.site/index"
// 	Nreq.Response_type = "code"
// 	state := "test crossss"
// 	EncState := base64.URLEncoding.EncodeToString([]byte(state))
// 	Nreq.State = EncState
// 	data, _ := json.Marshal(Nreq)
// 	req, err := http.NewRequest("POST", "https://nid.naver.com/oauth2.0/authorize", strings.NewReader(string(data)))
// 	if err != nil {
// 		log.Println("failed to create request", err)
// 	}
// 	res, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		log.Println("err occrured during get response from logon server", err)
// 	}
// 	authcode := new(ResponseAuth)
// 	json.NewDecoder(res.Body).Decode(&authcode)
// 	if authcode.State != EncState {
// 		return "0"
// 	}

// 	return authcode.Code
// }

func OauthCallback() {
	//need to wait for the response from resource server
	// and send it back to us the Auth code
}

// to Resource server
func Tokenrequest() {
	//토큰 받았으면 받은토큰으로

	//request to get an Access Token for client with Auth code that provided by Resource server.

}

func OauthTokenValidation() {
	//Validate Token during the logon state persist.
}
