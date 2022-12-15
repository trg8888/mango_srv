package model

type User struct {
	BaseModel

	Name     string `gorm:"type:varchar(100);not null;unique"`
	Mobile   string `gorm:"type:varchar(11);comment:手机号码;not null;unique"`
	Password string `gorm:"type:varchar(100);not null"`
	Position int32  `gorm:"type:int;default:1;not null;comment:0代表封禁 1代表未激活 2 代表运营 3代表管理员"`
}

type HomeDirectory struct {
	BaseModel
	Name  string `gorm:"type:varchar(20);not null"`
	Path  string `gorm:"type:varchar(20);not null;unique"`
	Icon  string `gorm:"type:varchar(200);default:'MoreFilled';not null"`
	IsTab bool   `gorm:"default:false;not null" json:"is_tab"`
}

type Directory struct {
	BaseModel
	Name  string `gorm:"type:varchar(20);not null"`
	Path  string `gorm:"type:varchar(20);not null;unique"`
	Icon  string `gorm:"type:varchar(200);default:'MoreFilled';not null"`
	IsTab bool   `gorm:"default:false;not null" json:"is_tab"`
}

type UserCategoryDirectory struct {
	BaseModel

	UserID          int32 `gorm:"type:int;index:idx_category_brand,unique"`
	User            User
	HomeDirectoryID int32 `gorm:"type:int;index:idx_category_brand,unique"`
	HomeDirectory   HomeDirectory
	DirectoryID     int32 `gorm:"type:int;index:idx_category_brand,unique"`
	Directory       Directory
}

type HomeDirectoryCategoryDirectory struct {
	BaseModel

	UserID          int32 `gorm:"type:int;index:idx_category_brand,unique"`
	User            User
	HomeDirectoryID int32 `gorm:"type:int;index:idx_category_brand,unique"`
	HomeDirectory   HomeDirectory
}
