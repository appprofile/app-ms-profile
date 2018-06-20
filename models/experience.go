package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// ExperienceWrapper Experience wrapper.
type ExperienceWrapper struct {
	Experiences []*Experience `json:"experiences"`
}

// Experience struct.
type Experience struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Title       string        `bson:"title" json:"title" validate:"required"`
	Company     string        `bson:"company" json:"company" validate:"required"`
	From        time.Time     `bson:"from" json:"from" validate:"required"`
	To          time.Time     `bson:"to" json:"to"`
	Current     bool          `bson:"current" json:"current"`
	Description string        `bson:"description" json:"description"`
	Created     time.Time     `bson:"created" json:"created"`
	Updated     time.Time     `bson:"updated" json:"updated"`
}
