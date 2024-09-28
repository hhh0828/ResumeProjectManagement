package mypackage

func Thefunc(a, b int) int {
	return a + b
}

func Fibt(a int) int {
	if a <= 1 {
		return a
	}
	return Fibt(a-2) + Fibt(a-1)

}
