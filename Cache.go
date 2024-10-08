package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

//feedback

//목표 일차적으로 모든 요청을 캐시에 담도록 요청한다.
//api에서 복잡한 로직이 구성되어있을경우 서버내 부하를 줄이기위해, key와 value를 모두 저장한다.
//인터페이스를 사용하여 모든 요청에 대한 key와 value를 저장한다. 대량의 데이터 또는 배열형태로 된 형태도 목표로...

// 케시 데이터 저장
type storeCache struct {
	sync.RWMutex
	data map[string]*Cachedvalue
}

// init globally
var Cachemem = storeCache{
	data: make(map[string]*Cachedvalue),
}

type Cachedvalue struct {
	CacheValue string
	Createdat  time.Time
}

// client /cache request
// {"key" : "mykey"} , {"value" : "myvalue"}
type KeyValueSet struct {
	MyKey   string `json:"key"`
	MyValue string `json:"value"`
}

type CachedKey struct {
	MyKey string `json:"key"`
}

// {userinfo {{"name" : "hyunho"} , {"email" : "hhhcjswo"}, {"address" : "unknown"}}}

func SaveCache(w http.ResponseWriter, r *http.Request) {
	var keyvalueset KeyValueSet
	if err := json.NewDecoder(r.Body).Decode(&keyvalueset); err != nil {
		fmt.Println("the error occured with decoding the data", err)
	}
	//current object inner set below
	/*
		keyvalueset {
		Mykey = r.body.key
		Myvalue = r.body.MyValue
		}
	*/
	Cachemem.Lock()
	defer Cachemem.Unlock()
	Cachemem.data[keyvalueset.MyKey] = &Cachedvalue{
		CacheValue: keyvalueset.MyValue,
		Createdat:  time.Now(),
	}
	DeleteCache()

}

func DeleteCache() {
	Cachemem.Lock()
	defer Cachemem.Unlock()
	for Mykey, MyValue := range Cachemem.data {

		if time.Since(MyValue.Createdat) > 15*time.Minute {
			delete(Cachemem.data, Mykey)
		}

	}
}

func RetriveCache(w http.ResponseWriter, r *http.Request) {
	// get으로 받을거면 key := r.URL.Path[len(https://hyunhoworld.site/cache):]

	var cachedkey CachedKey
	err := json.NewDecoder(r.Body).Decode(&cachedkey)
	if err != nil {
		fmt.Println("error with decoding the Cachedata request", err)
	}

	//var response *Cachedvalue
	//check cached data
	/*
		for savedkey, values := range Cachemem.data {
			if savedkey == cachedkey.MyKey {
				response = values.CacheValue
			}
		}
	*/

	response, exists := Cachemem.data[cachedkey.MyKey]
	if !exists {
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}

	jsondata := JsonCreator("cachedvalue", response.CacheValue)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jsondata)

}
