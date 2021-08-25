package main

import (
	"fmt"
	// mat "belajar/materi"
	// "belajar/db"
	"custgolang/db"
	_str "custgolang/structs"
)


func main(){
	var query db.Querys

	query = _str.Siswa{"A001", "Rizki", 19, "Inggris"}
	fmt.Println("show : ", query.Show())

	if query.Show() {
		query = _str.Siswa{"A001", "Rizki", 19, "Inggris"}
		fmt.Println("delete : ", query.Delete())
	}	

	query = _str.Siswa{"A001", "Rizki", 19, "Inggris"}
	fmt.Println("create : ", query.Create())

	if query.Show() {
		query = _str.Siswa{"A001", "Rizki Update", 20, "Inggris Update"}
		fmt.Println("update : ", query.Update())
	}	
}