package models

import (
	"github.com/ungerik/go-start/model"
	"github.com/ungerik/go-start/mongo"
)

var FeedbackParticipants = mongo.NewCollection("feedbackparticipants")

func NewFeedbackParticipant() *FeedbackParticipant {
	var doc FeedbackParticipant
	FeedbackParticipants.InitDocument(&doc)
	return &doc
}

type FeedbackParticipant struct {
	mongo.DocumentBase `bson:",inline"`
	Event              mongo.Ref           `model:"to=events|required"`
	Team               mongo.Ref           `model:"to=eventteams"`
	TeamChoice         model.DynamicChoice `view:"label=Team"`
	Communication      model.Choice        `model:"options=1(awesome), 2, 3, 4, 5(bad)" view:"label=Communication leading up to the event"`
	Organisation       model.Choice        `model:"options=1(awesome), 2, 3, 4, 5(bad)" view:"label=Organisation at the event"`
	Motivation         model.Choice        `model:"options=1(awesome), 2, 3, 4, 5(bad)" view:"label=Motivation through Startup Live Team"`
	Ideas              model.Choice        `model:"options=1(awesome), 2, 3, 4, 5(bad)" view:"label=Ideas and Topics"`
	Workshops          model.Choice        `model:"options=1(awesome), 2, 3, 4, 5(bad)" view:"label=Workshops & Talks"`
	WorkshopsBest      model.String        `view:"label=Which Workshop / Talk did you like best"`
	Mentors            model.Choice        `model:"options=1(awesome), 2, 3, 4, 5(bad)" view:"label=Mentors overall rating"`
	MentorsBest        model.String        `view:"label=Your favorite mentors|placeholder=name,name,..."`
	// // EachMentor         []MentorValuation `model:"options=1(awesome), 2, 3, 4, 5(bad)" view:"label=Workshops & Talks"`
	Teambuilding model.Choice `model:"options=,1(awesome), 2, 3, 4, 5(bad)" view:"label=Teambuilding activity"`
	Location     model.Choice `model:"options=1(awesome), 2, 3, 4, 5(bad)" view:"label=Location"`
	Food         model.Choice `model:"options=1(awesome), 2, 3, 4, 5(bad)" view:"label=Food & Beverages"`
	Comments     model.Text   `model:"required" view:"label=Comments, Hints, Suggestions, ..."`
	Overall      model.Choice `model:"options=1(awesome), 2, 3, 4, 5(bad)" view:"label=Overall rating for the event"`
}

type MentorValuation struct {
}
