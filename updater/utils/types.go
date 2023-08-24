package utils

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Player struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	name   string             `bson:"name,omitempty"`
	region string             `bson:"region,omitempty"`
	soloq  []Soloq            `bson:"soloq,omitempty"`
}

type Soloq struct {
	name string `bson:"name,omitempty"`
	rank string `bson:"rank,omitempty"`
	lp   int    `bson:"lp,omitempty"`
}
