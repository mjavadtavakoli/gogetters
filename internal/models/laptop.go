package models

type Laptop struct{

	ID uint `gorm:"primaryKey" json:"id"`
	Cpu string  `json:"cpu"`
}