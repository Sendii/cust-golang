package structs

type Siswa struct{
	Id string
	Nama string
	Umur int
	Jurusan string
}

func ColumnName() [4]string {
	return [4]string{"id", "nama", "umur", "jurusan"}
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