package main

import "net/http"

func main() {
	//go routines and server is alivinig in hereee!!
	AdjustingScale("./home/assets/test1.png")
	http.ListenAndServe("0.0.0.0:8700", NewHandlers())

}
