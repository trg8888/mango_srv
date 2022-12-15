package model

type Image struct {
	BaseModel
	Name     string `gorm:"type:varchar(200)"`
	URL      string `gorm:"type:varchar(200)"` // 图床链接
	UserID   int32  `gorm:"type:int;not null,unique,index"`
	User     User
	ImageID  int32    `json:"image"`
	Image    *Image   `json:"-"`
	SubImage []*Image `gorm:"foreignKey:ImageID;references:ID" json:"sub_category"`
	Level    int32    `gorm:"type:int;not null;default:1" json:"level"` // 三级才出链接
	IsTab    bool     `gorm:"default:false;not null" json:"is_tab"`
}
