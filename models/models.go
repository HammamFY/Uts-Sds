package models

type User struct {
	Id_user               uint `json:"id_user" gorm:"primaryKey;autoIncrement:true"`
	Nama, Email, Username string
	Password              string
}

type Admin struct {
	Id_admin              uint `json:"id_admin" gorm:"primaryKey;autoIncrement:true"`
	Nama, Email, Username string
	Password              string
}
