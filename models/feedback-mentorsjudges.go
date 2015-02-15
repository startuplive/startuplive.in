package models

import (
	"github.com/ungerik/go-start/model"
	"github.com/ungerik/go-start/mongo"
)

var FeedbackMentorsJudges = mongo.NewCollection("feedbackmentorsjudges")

func NewFeedbackMentorJudge() *FeedbackMentorJudge {
	var doc FeedbackMentorJudge

	FeedbackMentorsJudges.InitDocument(&doc)

	return &doc
}

type FeedbackMentorJudge struct {
	mongo.DocumentBase `bson:",inline"`
	Event              mongo.Ref    `model:"to=events|required"`
	Communication      model.Choice `model:"options=1(awesome), 2, 3, 4, 5(bad)" view:"label=Communication leading up to the event"`
	Organisation       model.Choice `model:"options=1(awesome), 2, 3, 4, 5(bad)" view:"label=Organisation at the event"`
	OrganiserTeam      model.Choice `model:"options=1(awesome), 2, 3, 4, 5(bad)" view:"label=Working together with the Startup Live Team"`
	Sessions           model.Choice `model:"options=1(awesome), 2, 3, 4, 5(bad)" view:"label=Working / mentoring sessions with the teams"`
	QualityOfTeams     model.Choice `model:"options=1(awesome), 2, 3, 4, 5(bad)" view:"label=Quality of Teams"`
	ExpectationsMet    model.Choice `model:"options=1(awesome), 2, 3, 4, 5(bad)" view:"label=Have your expectations been met?`
	Networking         model.Choice `model:"options=1(awesome), 2, 3, 4, 5(bad)" view:"label=Networking opportunities"`
	LikesDislikes      model.Text   `model:"required" view:"label=What did you like / dislike?"`
	Comments           model.Text   `model:"required" view:"label=Comments, Hints, Suggestions, ..."`
	Overall            model.Choice `model:"options=1(awesome), 2, 3, 4, 5(bad)" view:"label=Overall rating for the event"`
}
