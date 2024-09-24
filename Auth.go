package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

//JWT 완벽히 이해할떄까지.

type Jheader struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type JPayload struct {
	Userid     string    `json:"userid"`
	LoggedinAs string    `json:"LoggedinAs"`
	Exp        time.Time `json:"exp"`
	Sub        string    `json:"sub"`
}

const Secretkey = "This is my Secret Key"

/*
토큰 제너레이션
헤더와 페이로드를 받고 시크릿키와 함께 서명을 만들어
구분을 . 으로하며 토큰을 생성 header.payload.signature
*/
func GenerateToken(header Jheader, payload JPayload) string {

	he := Encodetobase64url(header)
	pa := Encodetobase64url(payload)
	sign := EncryptSigature(header, payload, Secretkey)

	data := he + "." + pa

	Token := data + "." + sign
	return Token
}

// 서명 생성 및 암호화 //GO에서 Json 바이트배열로 인코딩해주고, 해당 데이터를 합쳐서 headpay만들어줌
// 만들어진 데이터를 hmac-sha256 해쉬함수에 write 데이터 넣어줌.
// 그렇게만들어진 바이트 서명을 []바이트로 SUM을 통해 반환 하고 엔코딩하여 패딩을 제거해줌.
// 그리고 암호화된 서명 반환해줌.
func EncryptSigature(header Jheader, payload JPayload, secretkey string) string {

	//get header, Payload and change the type to Json from GO
	headerdata, err := json.Marshal(header)
	if err != nil {
		log.Println("the error occured with transforming the header to Json type", err)
	}
	payloaddata, err := json.Marshal(payload)
	if err != nil {
		log.Println("the error occured with transforming thepayload to Json type", err)
	}
	//Encoding the data to base64URL
	encodedHeader := strings.TrimRight(base64.URLEncoding.EncodeToString(headerdata), "=")
	encodedPayload := strings.TrimRight(base64.URLEncoding.EncodeToString(payloaddata), "=")
	dataToSign := encodedHeader + "." + encodedPayload

	//Create Hmac-sha256 signature
	hm := hmac.New(sha256.New, []byte(secretkey))
	hm.Write([]byte(dataToSign))
	bhmsignature := strings.TrimRight(base64.URLEncoding.EncodeToString(hm.Sum(nil)), "=")
	return bhmsignature
}

// Marshal and Encode to 64URL do not use for different type.
func Encodetobase64url(object interface{}) string {
	//여기서 json 바이트 슬라이스변환
	data, err := json.Marshal(object)
	if err != nil {
		log.Println(err)
	}
	//바뀐 배열을 base64로 인코딩 패딩빼고
	return strings.TrimRight(base64.URLEncoding.EncodeToString(data), "=")
}

func CreateHmac(data string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	a := strings.TrimRight(base64.URLEncoding.EncodeToString(h.Sum(nil)), "=")
	return a
}

func ValidateToken(receivedjwt string) {
	separatedjwt := strings.Split(receivedjwt, ".")

	header := separatedjwt[0]
	payload := separatedjwt[1]
	receivedsignature := separatedjwt[2]

	expectedsignature := CreateHmac(header+"."+payload, Secretkey)

	fmt.Println(expectedsignature, "this is expected")
	fmt.Println(receivedsignature, "this is received one")

	if receivedsignature == expectedsignature {
		fmt.Println("the token validated")
	} else {
		fmt.Println("dead token")
	}

	//check exp time
	payloadbyte, _ := base64.RawURLEncoding.DecodeString(payload)
	var payloadi JPayload
	json.Unmarshal(payloadbyte, &payloadi)

	if int64(payloadi.Exp.Unix()) > time.Now().Unix() {
		fmt.Println("time ok")
	} else {
		fmt.Println("time over")
	}
}

//Unix ?

//시간 비교
