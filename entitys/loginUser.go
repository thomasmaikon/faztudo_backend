package entitys

type LoginUser struct {
	Id       uint64 `gorm: primarykey; autoincrement`
	Login    string `gorm: unique; json:"login"`
	Password string `json:"password"`
}

func (LoginUser) TableName() string {
	return "credentials"
}
