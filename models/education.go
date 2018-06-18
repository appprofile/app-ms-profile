package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// EducationWrapper Education wrapper.
type EducationWrapper struct {
	Educations []*Education `json:"educations"`
}

// Education struct.
type Education struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Course      string        `bson:"course" json:"course" validate:"required"`
	Institute   string        `bson:"institute" json:"institute" validate:"required"`
	From        time.Time     `bson:"from" json:"from" validate:"required"`
	To          time.Time     `bson:"to" json:"to" validate:"required"`
	Description string        `bson:"description" json:"description"`
	Created     time.Time     `bson:"created" json:"created"`
	Updated     time.Time     `bson:"updated" json:"updated"`
}
