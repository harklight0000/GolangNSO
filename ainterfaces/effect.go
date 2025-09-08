package ainterfaces

import "nso/entity"

type IEffectFactory interface {
	CreateEffectWithID(id int, param int) IEffect
	CreateEffectFull(id int, timeStart int, timeLength int, param int) IEffect
	GetTemplateByID(id int) *entity.EffectEntity
}

type IEffect interface {
	GetTemplate() *entity.EffectEntity
	IsPermanentEffect() bool
	IsExpired() bool
}
