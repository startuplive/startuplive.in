package models

import (
	"github.com/ungerik/go-start/model"
	"github.com/ungerik/go-start/mongo"
)

var CitySuggestions = mongo.NewCollection("citysuggestions", "Name")

func NewCitySuggestion() *CitySuggestion {
	var doc CitySuggestion
	CitySuggestions.InitDocument(&doc)
	return &doc
}

type CitySuggestion struct {
	mongo.DocumentBase `bson:",inline"`
	Date               model.DateTime       `view:"hidden"`
	Type               model.MultipleChoice `model:"options=Startup Live Event, Startup Lounge Event" view:"hidden"`
	Name               model.String         `view:"size=20|label=Suggested city" model:"required|minlen=3|maxlen=40"`
	Email              model.Email          `view:"size=20|label=Your email address" model:"required|maxlen=40"`
}
