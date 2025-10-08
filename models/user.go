package models


type Customer struct {
	Id		  uint
    Firstname string `gorm:"not null"         json:"firstname"`
    Lastname  string `gorm:"not null"         json:"lastname"`
    Email     string `gorm:"not null;unique" json:"email"`
    Password  string `gorm:"not null"         json:"password"`
}


type LoginInput struct{
    Email     string   `json:"email"`
    Password  string    `json:"password"`
}