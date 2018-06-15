package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	ExperienceCollectionName = "experience"
)

type Experience struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Title       string        `bson:"title" json:"title"`
	Company     string        `bson:"company" json:"company"`
	From        time.Time     `bson:"from" json:"from"`
	To          time.Time     `bson:"to" json:"to"`
	Description string        `bson:"description" json:"description"`
	Created     time.Time     `bson:"created" json:"created"`
	Updated     time.Time     `bson:"updated" json:"updated"`
}

func (p *Experience) CollectionName() string {
	return ExperienceCollectionName
}
