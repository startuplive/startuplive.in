package models

import (
	"github.com/ungerik/go-start/model"
	"github.com/ungerik/go-start/mongo"
)

var FeedbackHosts = mongo.NewCollection("feedbackhosts", "Name")

func NewFeedbackHost() *FeedbackHost {
	var doc FeedbackHost
	FeedbackHosts.InitDocument(&doc)
	return &doc
}

type FeedbackHost struct {
	mongo.DocumentBase `bson:",inline"`
	Event              mongo.Ref    `model:"to=events|required"`
	Organisation       model.Choice `model:"options=1(awesome), 2, 3, 4, 5(bad)" view:"label=Organisation at the event (workload, stress, ...)"`
	Feeling            model.Choice `model:"options=1(awesome), 2, 3, 4, 5(bad)" view:"label=Feeling at the event"`
	Ideas              model.Choice `model:"options=1(awesome), 2, 3, 4, 5(bad)" view:"label=Ideas and Topics"`
	Workshops          model.Choice `model:"options=1(awesome), 2, 3, 4, 5(bad)" view:"label=Workshops & Talks"`
	WorkshopsBest      model.String `view:"label=Which Workshop / Talk did you like best"`
	Mentors            model.Choice `model:"options=1(awesome), 2, 3, 4, 5(bad)" view:"label=Mentors overall rating"`
	// // EachMentor         []MentorValuation `model:"options=1(awesome), 2, 3, 4, 5(bad)" view:"label=Workshops & Talks"`
	Teambuilding  model.Choice `model:"options=,1(awesome), 2, 3, 4, 5(bad)" view:"label=Teambbuilding Activity"`
	Location      model.Choice `model:"options=1(awesome), 2, 3, 4, 5(bad)" view:"label=Location"`
	Food          model.Choice `model:"options=1(awesome), 2, 3, 4, 5(bad)" view:"label=Food & Beverages"`
	LikesDislikes model.Text   `model:"required" view:"label=What did you like / dislike?"`
	Comments      model.Text   `model:"required" view:"label=Comments, Hints, Suggestions, ..."`
	Overall       model.Choice `model:"options=1(awesome), 2, 3, 4, 5(bad)" view:"label=Overall rating for the event"`
}
