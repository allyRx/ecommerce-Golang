package dto

type RegisterDTO struct {
	Id		  uint
    Firstname string `json:"first_name" form:"first_name" xml:"first_name"`
    Lastname  string `json:"lastname"  form:"lastname"  xml:"lastname"`
    Email     string `json:"email"     form:"email"     xml:"email"     gorm:"unique"` 
    Password  string `json:"password"  form:"password"  xml:"password"`
}

