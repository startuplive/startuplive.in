package models

import (
	//"os"
	//"time"
	"github.com/ungerik/go-start/model"
	"github.com/ungerik/go-start/mongo"
	//"gostart/debug"
)

var Votes = mongo.NewCollection("votes")

func NewVote() *Vote {
	var doc Vote
	Votes.InitDocument(&doc)
	return &doc
}

type Vote struct {
	mongo.DocumentBase `bson:",inline"`
	Created            model.DateTime
	IP                 model.String
	Cookie             model.String
	Team               mongo.Ref `model:"to=eventteams"`
	Event              mongo.Ref `model:"to=events"`
}
