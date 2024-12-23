package main

import (
	"encoding/base64"
	"encoding/json"

	"fmt"
	"io"
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
	Access_token      string `json:"access_token"`
	Refresh_token     string `json:"refresh_token"`
	Token_type        string `json:"token_type"`
	Expires_in        int    `json:"expires_in"`
	Error             string `json:"error"`
	Error_description string `json:"error_description"`
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
	redirectURI := "https://www.hyunhoworld.site/navercallback"
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

type Responses struct {
	ResultCode string `json:"resultcode"`
	Message    string `json:"message"`
	Data       struct {
		ID        string `json:"id"`
		Age       string `json:"age"`
		Gender    string `json:"gender"`
		Name      string `json:"name"`
		Birthday  string `json:"birthday"`
		BirthYear string `json:"birthyear"`
		Mobile    string `json:"mobile"`
		Mobile2   string `json:"mobile_e164"`
	} `json:"response"`
}

func OauthCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")
	//크로스 공격 방지.
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
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, _ := http.DefaultClient.Do(req)

	responedtoken := new(ResponseReqToken)
	json.NewDecoder(res.Body).Decode(responedtoken)
	fmt.Println(responedtoken)
	fmt.Println(responedtoken.Access_token, "해당 시간뒤에 만료 됨 : ", responedtoken.Expires_in)

	getreq, _ := http.NewRequest("GET", "https://openapi.naver.com/v1/nid/me", nil)

	getreq.Header.Set("Authorization", " Bearer "+responedtoken.Access_token)

	response, _ := http.DefaultClient.Do(getreq)

	resp := &Responses{}
	datas, _ := io.ReadAll(response.Body)
	defer response.Body.Close()
	//fmt.Println(string(datas))
	json.Unmarshal(datas, resp)

	fmt.Println(resp)
	fmt.Println(resp.Data.Name)
	//별도 회원가입 없이 가능하게.. 자동회원가입
	//네이버 회원정보 토대로 DB에 기록한다.
	//데이터 가공해서 만들기!
	//DB 조회 하고 ID가 이미있으면

	if !CheckUser(resp.Data.ID) {
		CreateUser(resp)

		uid := &User{}
		db := ConnectDB()
		db.First(uid, "userid = ?", resp.Data.ID)
		forwardedip := r.Header.Get("X-Forwarded-For")
		cookie := NewCookie(NewToken(uid.Userid, uid.GivenPermission, r.UserAgent()+forwardedip))
		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/index", http.StatusPermanentRedirect)
		return
	} else {
		// 사용자가 이미 DB에 있는 경우 index핸들러로가면, 알아서 사이트 토큰 체크함, 60분 지난경우 Oauth 받더라도 다시 로그인해야함.
		// 문제 ? = 내 서버에서 60분짜리 토큰 쿠키 받고 토큰이 유효하지않아 리다이렉트 되었지만 네이버에서 발급한 토큰은 유효한경우 네이버로 로그인은 되지만 내서버에서는
		uid := &User{}
		db := ConnectDB()
		db.First(uid, "userid = ?", resp.Data.ID)
		forwardedip := r.Header.Get("X-Forwarded-For")
		//해당 사용자의 permission을 확인하기위해 DB에 접속해야함.
		cookie := NewCookie(NewToken(uid.Userid, uid.GivenPermission, r.UserAgent()+forwardedip))
		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/index", http.StatusMovedPermanently)
	}

	//로그인 기록 체크 해야함. AccessToken 체크해서 Cache되어있는 사용자인지 체크 한시간 지났는지 체크하는 로직 만들기- 준비물 MAP
	//로그인 후, 데이터베이스에 사용자 고유 식별정보와 이름을 매칭하여 저장함. 없는 회원의 겨우 회원 가입 유도.
	//

	// and send it back to us the Auth code
	//test
}

// a := NewCookie(GenerateToken(Jheader{Alg: "HS256", Typ: "JWT"}, JPayload{Userid: loguser.Userid, LoggedinAs: loguser.GivenPermission, Exp: time.Now().Add(15 * time.Minute)}))
func NewToken(userid, paycl2, payclsession string) string {
	Jh := &Jheader{
		Alg: "HS256",
		Typ: "JWT",
	}
	Jp := &JPayload{
		Userid:     userid,
		LoggedinAs: paycl2,
		Exp:        time.Now().Add(60 * time.Minute),
		SessionID:  payclsession,
	}
	NewToken := GenerateToken(*Jh, *Jp)
	return NewToken
}

// 고유 ID 받고 있는지 없는지 먼저 체크
func CheckUser(oid string) bool {
	db := ConnectDB()
	model := new(User)
	db.First(&model, "userid = ?", oid)
	if model.Userid == oid {
		return true
	} else {
		return false
	}

}

// redirected from callback().
func OauthLogin(w http.ResponseWriter, r *http.Request) {
	//로그인 정보가 등록되면,
}

// Naver response CreateUser()
func (r *Responses) CreateUser() {
	Ouser := &User{
		GivenPermission: "0",
		Userid:          r.Data.ID,
		Mobile:          r.Data.Mobile2,
		Oauth:           1,
		Name:            r.Data.Name,
	}
	db := ConnectDB()
	db.Create(&Ouser)
	fmt.Println("new user added")
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
