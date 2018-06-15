package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	ProfileCollectionName = "profile"
)

type Profile struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	Address     string        `bson:"address" json:"address"`
	Phone       uint64        `bson:"phone" json:"phone"`
	Experiences []*Experience `bson:"experiences" json:"experiences"`
	Education   []*Education  `bson:"education" json:"education"`
	Abilities   []string      `bson:"abilities" json:"abilities"`
	Created     time.Time     `bson:"created" json:"created"`
	Updated     time.Time     `bson:"updated" json:"updated"`
}

func (p *Profile) CollectionName() string {
	return ProfileCollectionName
}
