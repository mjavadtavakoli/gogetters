package models

type Book struct {
    ID     uint   `gorm:"primaryKey" json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
    Year   int    `json:"year"`
    Lan    string `json:"lan"`
}

/*type Coffee struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Late string `json:"late"`
    Amount int  `json:"amount"`
}*/


/*type motorcycle struct{
    ID uint `gorm:"primaryKey" json:"id"`
    Brand string `json:"brand"`
    Totalspeed int `json:"totalsped"`
    Fueltype  string  `json:"feultype"`
    Price float64 `json:"price"`
}*/