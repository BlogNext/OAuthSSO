package model

type UserModel struct {
	BaseModel
	Nickname string `gorm:"cloumn:nickname"`
	Password string `gorm:"cloumn:password"`
}

func (UserModel) TableName() string {
	return "user"
}
