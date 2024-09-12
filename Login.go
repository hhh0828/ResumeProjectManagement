package main

import (
	"net/http"
)

// add login func
func Login(ID string, PW string) {

}

// 너무헷갈림 ㅜㅡㅡㅡㅡㅡㅡㅡ
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_token")
		if err != nil || cookie == nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		token := cookie.Value
		if !validateToken(token) { // 토큰 검증 로직
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

