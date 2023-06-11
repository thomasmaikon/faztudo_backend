package entitys

type LoginUser struct {
	Id       int    `gorm:primarykey;autoincrement`
	Login    string `gorm:unique; json:"login"`
	Password string `json:"password"`
}

func (LoginUser) TableName() string {
	return "credentials"
}
