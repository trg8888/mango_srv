package model

type User struct {
	BaseModel

	Name     string `gorm:"type:varchar(100);not null"`
	Mobile   int32  `gorm:"type:int;comment:手机号码;not null"`
	Password string `gorm:"type:varchar(100);not null"`
	Position int32  `gorm:"type:int;efault:1;not null;comment:0代表封禁 1代表未激活 2 代表运营 3代表管理员"`
}
