package objects

import "nso/entity"

type SkillList []*entity.SkillEntity
type ItemList []*entity.ItemEntity

func (this ItemList) Len() int {
	return len(this)
}

func (this ItemList) Less(i, j int) bool {
	return this[i].ID < this[j].ID
}

func (this ItemList) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (s SkillList) Len() int {
	return len(s)
}

func (s SkillList) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}

func (s SkillList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
