package models

import (
	"github.com/ungerik/go-start/model"
	"github.com/ungerik/go-start/mongo"
)

var Surveys = mongo.NewCollection("surveys")

///////////////////////////////////////////////////////////////////////////////
// Survey

type Survey struct {
	mongo.DocumentBase `bson:",inline"`
	CreatedBy          mongo.Ref `model:"to=users"`
	CreatedAt          model.DateTime
	ModifiedBy         mongo.Ref `model:"to=users"`
	ModfiedAt          model.DateTime
	Questions          []SurveyQuestion
}

type SurveyQuestion struct {
	Type        model.Choice `model:"options=Single Choice,Multiple Choice,Points,Slider,Text Answer"`
	Question    model.String
	Description model.String
	Options     []model.String
	Answers     []struct {
		User mongo.Ref `model:"to=users"`
		// todo
	}
}
