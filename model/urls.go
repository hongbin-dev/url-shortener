package model

import "gopkg.in/mgo.v2/bson"

type (
	Urls struct {
		Id         bson.ObjectId `json:"id" bson:"_id,omitempty"`
		OriginLink string        `json:"originLink" bson:"origin_link"`
		ShortLink  string        `json:"shortLink" bson:"short_link"`
	}
)
