package main

type DBupdater interface {
	Upload() //add the Data column to yout database table.
	Delete()
}

// func that implements the DBupdater interface.
func Upload(a DBupdater) {
	a.Upload()
}
func Delete(a DBupdater) {
	a.Delete()
}
