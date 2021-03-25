package model

//授信的客户端
type OAuthClientModel struct {
	BaseModel
	//客户id
	ClientId     string `gorm:"cloumn:client_id"`
	ClientSecret string `gorm:"cloumn:client_secret"`
	ClientName   string `gorm:"cloumn:client_name"`
}

func (OAuthClientModel) TableName() string {
	return "oauth_client"
}
