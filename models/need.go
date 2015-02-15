package models

import (
	"github.com/ungerik/go-start/model"
	"github.com/ungerik/go-start/mongo"
)

var Needs = mongo.NewCollection("needs", "Description")

///////////////////////////////////////////////////////////////////////////////
// Need

type Need struct {
	mongo.DocumentBase `bson:",inline"`
	Created            model.DateTime
	Type               model.MultipleChoice `model:"options=Co-Founder;Employee;Freelancer;Advisor;Funding;Coaching;Business Plan;Legal Advice;Opinion;Publicity Blitz;Help"`
	Description        model.String
	Modified           model.DateTime
	Closed             model.DateTime
	CloseReason        model.String
}
