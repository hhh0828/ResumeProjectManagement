package main

type DBupdater interface {
	Upload() //add the Data column to yout database table.
}

// func that implements the DBupdater interface. // 업로드 함수를 가진 놈이면 디비를 업로드함.
func Upload(a DBupdater) {
	a.Upload()
}
