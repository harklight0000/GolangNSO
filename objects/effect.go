package objects

import (
	. "nso/ainterfaces"
	"nso/entity"
	"time"
)

type Template struct {
	IconID     int    `json:"iconId"`
	Name       string `json:"name"`
	TemplateID int    `json:"templateId"`
	Type       int    `json:"type"`
}

func NewTemplate(entity *entity.EffectEntity) *Template {
	this := &Template{}
	this.IconID = int(entity.IConId)
	this.Name = entity.Name
	this.TemplateID = entity.ID
	this.Type = entity.Type
	return this
}

type Effect struct {
	ID                   int       `json:"id" bson:"id"`
	Template             *Template `json:"template" bson:"template"`
	*entity.EffectEntity `json:"-" bson:"-"`
	TimeStart            int   `json:"timeStart" bson:"timeStart"`
	TimeLength           int   `json:"timeLength" bson:"timeLength"`
	Param                int   `json:"param" bson:"param"`
	TimeRemove           int64 `json:"timeRemove" bson:"timeRemove"`
}

func (this *Effect) GetTemplate() *entity.EffectEntity {
	return this.EffectEntity
}

func (this *Effect) IsPermanentEffect() bool {
	t := this.Template.Type
	return t == 0 || t == 18 || t == 25 || t == 26 || t == 27 || t == 28
}

func (this *Effect) IsExpired() bool {
	return this.TimeRemove-CurrentTimeMillis()/int64(time.Second) <= 0
}

func NewEffectFactory(effects []*entity.EffectEntity) *EffectFactory {
	return &EffectFactory{effects: effects}
}

type EffectFactory struct {
	effects []*entity.EffectEntity
}

func (e *EffectFactory) CreateEffectWithID(id int, param int) IEffect {
	template := e.effects[id]
	eff := &Effect{EffectEntity: template, Param: param}
	eff.ID = id
	eff.Template = &Template{IconID: int(template.IConId), Name: template.Name, TemplateID: template.ID, Type: template.Type}
	return eff
}

func (e *EffectFactory) CreateEffectFull(id int, timeStart int, timeLength int, param int) IEffect {
	eff := e.CreateEffectWithID(id, param).(*Effect)
	eff.TimeStart = timeStart
	eff.TimeLength = timeLength
	eff.TimeRemove = CurrentTimeMillis()/int64(time.Second) + int64(timeLength)
	return eff
}

func (e *EffectFactory) GetTemplateByID(id int) *entity.EffectEntity {
	return e.effects[id]
}
