package entitys

type LoginUser struct {
	Id       uint64 `gorm: primarykey; autoincrement`
	Login    string
	Password string
}

func (LoginUser) TableName() string {
	return "credentials"
}
