// 로그인없이 이용시, Project / Resume 수정불가
// 로그인하여 이용시, 수정 가능한 로직이 필요함.
// 세션을 만들어 사용자의 로그인여부를 파악하게끔...
// JS에서 사용자가 로그인하였는지 요청에 섞어 보내는것도 필요함.
// 현재 서버의 메모리 한켠에서 사용자가 관리자로 로그인 하였다면, 확인후 해당 정보를 Channel로 해당정보를 전달하여
// 사용자가 관리자인지 파악 후, Edit페이지 전달해줌.
package main

// add login func
func Login(ID string, PW string) {

}

/*
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
*/
