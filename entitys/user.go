package entitys

type User struct {
	Id   int    `gorm:"primaryKey;autoIncrement"`
	Cpf  string `gorm:"unique"`
	Name string

	LoginUserId int
	LoginUser   LoginUser
}

func (User) TableName() string {
	return "users"
}
