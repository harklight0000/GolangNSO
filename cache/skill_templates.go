package cache

type SkillTemplates struct {
	SkillId  int16    `json:"skillId" bson:"skillId"`
	Point    int8     `json:"point" bson:"point"`
	Level    int      `json:"level" bson:"level"`
	ManaUse  int16    `json:"manaUse" bson:"manaUse"`
	CoolDown int64    `json:"coolDown" bson:"coolDown"`
	Dx       int16    `json:"dx" bson:"dx"`
	Dy       int16    `json:"dy" bson:"dy"`
	MaxFight int8     `json:"maxFight" bson:"maxFight"`
	Options  []Option `json:"options" bson:"options"`
}
