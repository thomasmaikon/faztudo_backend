package entitys

import "projeto/FazTudo/consts"

type Like struct {
	Id            int `gorm:"autoincrement; primarykey"`
	LoginId       int `gorm:"column:fk_login;foreignKey:Id"`
	ServicePageId int `gorm:"column:fk_service_page;foreignKey:Id"`
	Like          consts.EnumLikeType
}

func (Like) TableName() string {
	return "like"
}
