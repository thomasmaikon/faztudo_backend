package entitys

type Commit struct {
	Id     uint64 `gorm:"primarykey;autoincrement"`
	Commit string
	UserId uint64 `gorm:"column:fk_login;foreignKey:Id"`
	PageId uint64 `gorm:"column:fk_service_page;foreignKey:Id"`
}

func (Commit) TableName() string {
	return "commit"
}
