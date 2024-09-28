package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	mypackage "welcome/hyunho"
)

func main() {
	a := mypackage.Thefunc(1, 2)
	fmt.Println(a)

	f, _ := os.Create("pro.prof")
	defer f.Close()

	if err := pprof.StartCPUProfile(f); err != nil {
		log.Println("working?")
	}
	defer pprof.StopCPUProfile()

	ab := mypackage.Fibt(35)
	fmt.Println(ab)
	fmt.Scan()
}
