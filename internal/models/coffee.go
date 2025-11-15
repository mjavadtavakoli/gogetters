package models


type Coffee struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Late string `json:"late"`
    Amount int  `json:"amount"`
}