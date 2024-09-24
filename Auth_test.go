package main

import (
	"fmt"
	"testing"
	"time"
)

func TestMain(t *testing.T) {

	header := Jheader{
		Alg: "HS256",
		Typ: "JWT",
	}
	payload := JPayload{
		Userid:     "testname",
		LoggedinAs: "webmaster",
		Exp:        time.Now().Add(15 * time.Minute),
		Sub:        "test",
	}

	tk := GenerateToken(header, payload)
	fmt.Println(tk)

	ValidateToken(tk)
}
