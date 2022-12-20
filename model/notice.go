package model

type Notice struct {
	BaseModel
	Title    string `gorm:"type:varchar(100);not null"`
	Describe string `gorm:"column:describe;type:text;not null"`
}
