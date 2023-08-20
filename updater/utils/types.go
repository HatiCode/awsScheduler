package utils

import "go.mongodb.org/mongo-driver/bson/primitive"

type Player struct {
	ID     primitive.ObjectID `bson:"_id"`
	name   string             `bson:"name"`
	region string             `bson:"region"`
	soloq  []Soloq            `bson:"soloq"`
}

type Soloq struct {
	name string `bson:"name"`
	rank string `bson:"rank"`
	lp   int    `bson:"lp"`
}
