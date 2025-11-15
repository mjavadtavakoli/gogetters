package models

type Motorcycle struct{
    ID uint `gorm:"primaryKey" json:"id"`
    Brand string `json:"brand"`
    Totalspeed int `json:"totalspeed"`
    Fueltype  string  `json:"fueltype"`
    Price float64 `json:"price"`
}