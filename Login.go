// 사용자가 관리자인지 파악 후, Edit페이지 전달해줌.
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"gorm.io/gorm"
)

// add login func
func Login(ID string, PW string) {

}

type User struct {
	gorm.Model
	GivenPermission string `json:"gp"`
	Userid          string `json:"id"`
	Userpw          string `json:"pw"`
	Useremail       string
}
type Message struct {
	Status            uint16
	MessagefromMaster string
}

// Set and transforming to Json []byte arra
func (mess *Message) Messagesetter(a uint16, m string) ([]byte, error) {
	mess.Status = a
	mess.MessagefromMaster = m
	data, err := json.Marshal(mess)
	return data, err
}

func (user *User) Encryption() [32]byte {
	bytepassword := []byte(user.Userpw)
	encpw := sha256.Sum256(bytepassword)
	return encpw
}

func (user *User) ChecksinDB(encpwstr string) bool {
	comparinguser := new(User)
	db := ConnectDB()
	db.First(&comparinguser, "userid = ?", user.Userid) // UserId로 찾기
	fmt.Println(comparinguser, user, "여긴 user check sin 함수 내")
	if comparinguser.Userpw == encpwstr {
		return true
	} else {
		return false
	}
}
func LoginRequest(w http.ResponseWriter, r *http.Request) {
	//send user data from Client.
	loguser := new(User)
	err := json.NewDecoder(r.Body).Decode(&loguser)
	if err != nil {
		log.Println("the error occured with decoding", err)
	}
	encpw := (loguser.Encryption())
	loguser.Userpw = hex.EncodeToString(encpw[:])
	//Matching user
	fmt.Println(loguser, "여기는 Request")
	if loguser.ChecksinDB(loguser.Userpw) {
		//쿠키 추가하여 던져주고 확인하는.
		w.Header().Set("Content-Type", "application/json")
		a := NewCookie(GenerateToken(Jheader{Alg: "HS256", Typ: "JWT"}, JPayload{Userid: loguser.Userid, LoggedinAs: loguser.GivenPermission, Exp: time.Now().Add(15 * time.Minute)}))
		http.SetCookie(w, a)
		response, err := json.Marshal(&Message{Status: 200, MessagefromMaster: loguser.GivenPermission})
		if err != nil {
			log.Println("marshaling error", err)
		}
		w.Write(response)
	} else {
		fmt.Println("test if it works ?")
		a, err := json.Marshal(&Message{Status: 401, MessagefromMaster: "A Yo You need to input correct password"})
		if err != nil {
			log.Println("marshaling failed", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(a)
	}
}
func JoinasMember(w http.ResponseWriter, r *http.Request) {
	//User 객체 포인터 생성
	requesttobeuser := new(User)
	// 사용자 입력 값 토대로
	err := json.NewDecoder(r.Body).Decode(&requesttobeuser)
	if err != nil {
		log.Println("fail to decode the User join data", err)
	}
	//사용자 엔크립션 해시256
	a := requesttobeuser.Encryption()
	//패스워드 헥사로
	requesttobeuser.Userpw = hex.EncodeToString(a[:])
	db := ConnectDB()
	db.Create(&requesttobeuser)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&Message{Status: 200, MessagefromMaster: "the request has been sent to the server please try to login. if you have issue with the logon send us a feedback on the contact page"})
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./home/Login.html")
}

/*
Feedbacks from GPT4
토큰 발급: 로그인 성공 시 토큰을 발급하는 로직이 필요하다고 하셨는데, 이를 통해 사용자의 세션을 관리하는 것이 좋습니다. 앞서 제안드린 JWT를 사용하는 방법을 참고하시면 도움이 될 것입니다. Done.

에러 처리: 에러 처리가 적절히 이루어져 있지만, 사용자가 어떤 문제로 인해 실패했는지 명확하게 전달하는 것이 좋습니다. 예를 들어, 비밀번호가 틀린 경우와 사용자가 존재하지 않는 경우를 구분하여 사용자에게 메시지를 제공하면 더 친절한 인터페이스가 될 것입니다.Done

인증 미들웨어: 사용자가 로그인한 상태인지 확인하는 미들웨어가 필요합니다. 이를 통해 보호해야 하는 엔드포인트에 대해 인증을 쉽게 관리할 수 있습니다. need to do authmiddelware함수 작성.

비밀번호 저장: 비밀번호를 해시한 후 DB에 저장할 때는 16진수로 변환한 후 저장하는 것이 좋습니다. 현재는 해시 값을 바로 비교하는 형태인데, 보안적인 측면에서 해시 값을 저장하는 것이 일반적입니다. Done.
*/

// JWT Token =
