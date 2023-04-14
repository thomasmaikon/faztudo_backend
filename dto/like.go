package dto

import "projeto/FazTudo/consts"

type LikeInput struct {
	LoginId       int
	ServicePageId int
	Like          consts.EnumLikeType
}
