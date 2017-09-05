package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	BackendDB = "backend"
	UserC     = "user"
)

// Account is the storage model of account.
type Account struct {
	ID         bson.ObjectId `bson:"_id" json:"id"`
	Mobile     string        `bson:"mobile" json:"mobile"`
	Avatar     string        `bson:"avatar" json:"avatar"`
	IdentityID string        `bson:"identity_id"`
	Realname   string        `bson:"realname"`
	Password   string        `bson:"password" json:"password"`
	CreatedAt  time.Time     `json:"_" bson:"created_at"`
}
