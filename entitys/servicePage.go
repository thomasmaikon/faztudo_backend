package entitys

type ServicePage struct {
	Id                  uint64 `gorm:"primarykey;autoincrement"`
	UserId              uint64 `gorm:"column:fk_login;foreignKey:Id"`
	Name                string
	Description         string
	Image               string
	Value               float64
	PositiveEvaluations uint64
	NegativeEvaluations uint64
}

func (ServicePage) TableName() string {
	return "service"
}
