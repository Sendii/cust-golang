package structs

import (
	"strings"
	// _help "custgolang/helper"
	"os"
	"path/filepath"
	"fmt"
	"log"
)

type Siswa struct{
	Id string
	Nama string
	Umur int
	Jurusan string
}

var fields []string
func init(){
	fields = []string{"id", "nama", "umur", "jurusan"}
}

// return column name at table
func ColumnName() string {
	return strings.Join(fields, ",")
}

func TableName(){
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
}

// return count column name at table
func ColumnTota() int {
	return len(fields)
}

func(s Siswa) Show() bool {
	return sqlShow(s.Id)
}

func(s Siswa) Create() bool {
	return SqlInsert(s.Id, s.Nama, s.Umur , s.Jurusan)
}

func(s Siswa) Update() bool {
	return sqlUpdate(s.Id, s.Nama, s.Umur , s.Jurusan)
}

func(s Siswa) Delete() bool {
	return sqlDelete(s.Id)
}