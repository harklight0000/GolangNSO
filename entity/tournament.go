package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type TournamentEntity struct {
	_ID         primitive.ObjectID `json:"_id" bson:"_id" gorm:"-"`
	ID          int                `json:"id" gorm:"primary_key;" bson:"id"`
	Tournaments string             `json:"tournaments" bson:"tournaments" gorm:"column:tournaments;type:varchar(500);default:'[]'"`
}
