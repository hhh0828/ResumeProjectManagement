package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

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
	Cachemem.data[keyvalueset.MyKey] = &Cachedvalue{
		CacheValue: keyvalueset.MyValue,
		Createdat:  time.Now(),
	}
	DeleteCache()
	Cachemem.Unlock()
}

func DeleteCache() {
	for Mykey, MyValue := range Cachemem.data {
		Cachemem.Lock()
		if time.Since(MyValue.Createdat) > 15*time.Minute {
			delete(Cachemem.data, Mykey)
		}
		Cachemem.Unlock()
	}
}

func RetriveCache(w http.ResponseWriter, r *http.Request) {

	var cachedkey CachedKey
	err := json.NewDecoder(r.Body).Decode(&cachedkey)
	if err != nil {
		fmt.Println("error with decoding the Cachedata request", err)
	}

	var response string
	//check cached data
	for savedkey, values := range Cachemem.data {
		if savedkey == cachedkey.MyKey {
			response = values.CacheValue
		}
	}
	JsonCreator("cachedvalue", response)
	json.NewEncoder(w).Encode(response)

}
