package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	EducationCollectionName = "education"
)

type Education struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Course      string        `bson:"course" json:"course"`
	Institute   string        `bson:"institute" json:"institute"`
	From        time.Time     `bson:"from" json:"from"`
	To          time.Time     `bson:"to" json:"to"`
	Description string        `bson:"description" json:"description"`
	Created     time.Time     `bson:"created" json:"created"`
	Updated     time.Time     `bson:"updated" json:"updated"`
}

func (p *Education) CollectionName() string {
	return EducationCollectionName
}
