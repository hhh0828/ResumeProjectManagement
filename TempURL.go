package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

// URL 저장
type storeURL struct {
	sync.RWMutex
	urls map[string]*Urldata
}

type Urldata struct {
	sync.RWMutex
	originalurl string
	Createdat   time.Time
}

var store = storeURL{
	urls: make(map[string]*Urldata),
}

type RequestBody struct {
	URL string `json:"url"`
}

/*
var store = storerUR
func initstore() {
	store.urls = make(map[string]string)
}
*/

// 긴 URL을 짧은 URL로 변환 6글자 문자열
func GenerateURL() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	//6 칸짜리 []byte 배열
	b := make([]byte, 6)

	//b 슬 에 랜덤값 주입 랜덤 숫자를 넣어줌
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// 짧은 URL을 사용자에게 반환해줌
func ShortenURL(w http.ResponseWriter, r *http.Request) {
	var body RequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
	}

	shortURL := GenerateURL()

	//URL정보를 담아서 저장한다 shortURL키..
	store.Lock()

	//리터럴 - URLdata 초기화
	store.urls[shortURL] = &Urldata{
		originalurl: body.URL,
		Createdat:   time.Now(),
	}
	//메모리 공간절약을 위하여 shortURL 점검. 15분 넘어가는게 있는지 확인
	// for key, value := range map {}
	/*
		need to know :

		map 순회
		객체의 필드 중 embeded mutex가 있을 경우 복사 문제가 일어날 수있음.
		속도 및 메모리 최적화를 위해 포인터를 사용해 메모리 주소값을 가르키면서 Mutex copies 해결
		time값에 대한 연산 더하기빼기 등 시간차 계산 확인

	*/
	for shortURL, data := range store.urls {
		store.Lock()
		if time.Since(data.Createdat) > 15*time.Minute {
			delete(store.urls, shortURL)
		}
		store.Unlock()
	}
	store.Unlock()

	//json 형식으로 ..
	response := JsonCreator("short_url", "http://localhost:8080/"+shortURL)
	w.Header().Set("Content-Type", "application/json")
	//data, _ := json.Marshal(response)
	json.NewEncoder(w).Encode(response)

}

// Map creator. only for strings
func JsonCreator(key string, value string) *map[string]string {
	jsondata := map[string]string{key: value}
	return &jsondata
}

// 짧은 URL요청 시 긴 URL로 Redirect 해줌
func RedirectURL(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[1:] // ?

	store.RLock()
	originalURL := store.urls[shortURL].originalurl
	store.RUnlock()

	//responsedata := JsonCreator("originalurl", "https://www.hyunhoworld.site" + originalURL)

	http.Redirect(w, r, originalURL, http.StatusFound)
}

//저장소에 시간을 두어서 타임 만료 시 해당 저장소의 URL 파기
/*
저장소의 URL의 시간객체를 심어두고 파기할수있는가 ? NO
Go 루틴을 생성하여 각 저장소를 순회하며 저장소의 URL의 만들어진 시간을 체크하여 +15분일 시 제거.
보통은 어떤 방법을 ?
*/
