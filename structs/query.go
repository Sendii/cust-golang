package structs
import (
	"fmt"
	db "custgolang/db"
)

func sqlShow(id string)bool{
	db, err := db.PConenctDB()
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer db.Close()

	var res = Siswa{}
	err = db.
	QueryRow("select "+ColumnName()+" from siswa where id = ?", id).
	Scan(&res.Id, &res.Nama, &res.Umur, &res.Jurusan) // yg di select, kudu dipanggil semua biar ga error

	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

func SqlInsert(id string, nama string, umur int, jurusan string)bool{
	db, err := db.PConenctDB()
	if err != nil{
		fmt.Println(err.Error())
		return false
	}
	defer db.Close()

	_, err = db.Exec("insert into siswa values(?, ?, ?, ?)", id, nama, umur, jurusan)
	if err != nil{
		fmt.Println(err.Error())
		return false
	}
	return true
}

func sqlUpdate(id string, nama string, umur int, jurusan string)bool{
	db, err := db.PConenctDB()
	if err != nil{
		fmt.Println(err.Error())
		return false
	}
	defer db.Close()

	_, err = db.Exec("update siswa set nama = ?, umur = ?, jurusan = ? where id = ?", nama, umur, jurusan, id)
	if err != nil{
		fmt.Println(err.Error())
		return false
	}
	return true
}

func sqlDelete(id string)bool{
	db, err := db.PConenctDB()
	if err != nil{
		fmt.Println(err.Error())
		return false
	}
	defer db.Close()

	_, err = db.Exec("delete from siswa where id = ?", id)
	if err != nil{
		fmt.Println(err.Error())
		return false
	}
	return true
}