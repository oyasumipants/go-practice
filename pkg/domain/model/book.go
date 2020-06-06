package model

type Book struct {
	ID int `gorm:"primary_key"`
	Title string
	Price int
	Author string
}
