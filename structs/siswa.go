package structs

import (
	"strings"
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