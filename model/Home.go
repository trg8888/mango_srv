package model

type HomeParameters struct {
	BaseModel
	UserID    int32 `gorm:"type:int;index:idx_category_brand"`
	User      User
	SubTitle  string  `gorm:"type:varchar(100);not null"`
	SubValue  float32 `gorm:"type:real;not null;default:1.00"`
	Title     string  `gorm:"type:varchar(50);not null;unique"`
	Uint      string  `gorm:"type:varchar(50);not null;default:'0'"`
	UintColor string  `gorm:"type:varchar(50);not null;default:'success'"`
	Value     string  `gorm:"type:varchar(50);not null;default:'0'"`
	IsTab     bool    `gorm:"default:false;not null" json:"is_tab"`
}

type HomeUpdateLog struct {
	BaseModel
	UserID int32 `gorm:"type:int;unique"`
	User   User
	Amount int32  `gorm:"type:int"`
	Data   string `gorm:"type:varchar(10);not null;default:'2023-01-01'"`
}
