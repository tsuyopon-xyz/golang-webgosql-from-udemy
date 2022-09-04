package models

type Todo struct {
	BaseModel
	Title string `gorm:"type:varchar(100);not null" json:"title"`
	Body  string `gorm:"type:text;not null" json:"body"`
}
