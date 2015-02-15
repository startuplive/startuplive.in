package models

import (
	//	"os"
	"github.com/ungerik/go-start/model"
	"github.com/ungerik/go-start/mongo"
)

var EventScheduleItems = mongo.NewCollection("eventscheduleitems", "Title")

func NewEventScheduleItem() *EventScheduleItem {
	var doc EventScheduleItem
	EventScheduleItems.InitDocument(&doc)
	return &doc
}

type EventScheduleItem struct {
	mongo.DocumentBase `bson:",inline"`
	Event              mongo.Ref    `model:"required|to=events"`
	Title              model.String `model:"minlen=3"`
	Location           model.String
	From               model.DateTime `model:"required"`
	Until              model.DateTime `model:"required"`
}

func (self *EventScheduleItem) String() string {
	return self.Title.Get()
}
