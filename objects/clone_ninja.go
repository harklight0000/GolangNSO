package objects

import "nso/entity"

type CloneNinja struct {
	*entity.BodyEntity
	ChuThan *Ninja
	*Ninja
}
