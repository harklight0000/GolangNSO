package objects

import (
	"nso/ainterfaces"
	"nso/cache"
	"nso/entity"
)

type Skill struct {
	ID       byte
	Point    byte
	CoolDown int64
	Type     byte
	entity   *entity.SkillEntity
	template *cache.SkillTemplates
}

func NewSkill(skillEntity *entity.SkillEntity, point byte) *Skill {
	this := &Skill{entity: skillEntity}
	this.ID = this.entity.ID
	this.template = this.Template(point)
	this.Point = point
	this.CoolDown = this.template.CoolDown
	this.Type = this.entity.Type
	return this
}

func (s *Skill) Template(point byte) *cache.SkillTemplates {
	return &s.entity.SkillTemplates[point]
}

func (s *Skill) Data() *entity.SkillEntity {
	return s.entity
}

func (s *Skill) ToJSON() ainterfaces.SkillJSON {
	return ainterfaces.SkillJSON{
		ID:    s.ID,
		Point: s.Point,
	}
}
