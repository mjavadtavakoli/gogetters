package models

type Car struct{
    ID     uint   `gorm:"primaryKey" json:"id"`
	Name   string `json:"name"` 

}