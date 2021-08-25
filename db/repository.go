package db

type Querys interface{
	// Read() bool
	Show() bool
	Create() bool
	Update() bool
	Delete() bool
}