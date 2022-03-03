package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Create Struct
type Comment struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Chat string             `json:"chat" bson:"chat,omitempty"`
}

type Reply struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CmtId     string             `json:"cmtid,omitempty" bson:"cmtid,omitempty"`
	ChatReply string             `json:"chatreply" bson:"chatreply,omitempty"`
}
