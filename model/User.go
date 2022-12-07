package model

type User struct {
	BaseModel

	Name     string `gorm:"type:varchar(100);not null;unique"`
	Mobile   string `gorm:"type:varchar(11);comment:手机号码;not null;unique"`
	Password string `gorm:"type:varchar(100);not null"`
	Position int32  `gorm:"type:int;default:1;not null;comment:0代表封禁 1代表未激活 2 代表运营 3代表管理员"`
}
