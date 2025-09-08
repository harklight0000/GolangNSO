package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type TasksEntity struct {
	_ID      primitive.ObjectID `json:"_id" bson:"_id" gorm:"-"`
	ID       int                `json:"id" gorm:"primary_key;" bson:"id"`
	Tasks    []int8             `json:"tasks" bson:"tasks" gorm:"column:tasks;type:varchar(500);default:'[]';serializer:my_json"`
	MapTasks []int8             `json:"maptasks" bson:"maptasks" gorm:"column:maptasks;type:varchar(5000);default:'[]';serializer:my_json"`
}

func (t TasksEntity) TableName() string {
	return "tasks"
}
