package event_1

import (
	"github.com/STARTeurope/startuplive.in/models"
	// "errors"
	. "github.com/STARTeurope/startuplive.in/views"
	. "github.com/ungerik/go-start/view"
	// "io"
	// "github.com/ungerik/go-start/debug"
)

func init() {
	Region0_Event1_MentorJudgeFeedback = newPublicEventPage(
		EventTitle("Mentors Judges Feedback"),
		nil,
		DynamicView(eventFeedbackMentorsJudges),
	)
}

func eventFeedbackMentorsJudges(ctx *Context) (view View, err error) {
	event := ctx.Data.(*PageData).Event

	views := Views{
		H2("Feedback Form"),
		&Form{
			SubmitButtonText:  "Submit",
			SubmitButtonClass: "button",
			FormID:            "mjfeedback",
			GetModel: func(form *Form, ctx *Context) (interface{}, error) {
				var feedback models.FeedbackMentorJudge
				models.FeedbackMentorsJudges.InitDocument(&feedback)
				feedback.Event.Set(event)
				return &feedback, nil
			},
			OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {

				feedback := formModel.(*models.FeedbackMentorJudge)

				return "", Region0_Event1_Registration_Success, feedback.Save()
			},
		},
		// PersonForm(person, Region0_Event1_Registration_Success, []string{ /*"Judge"*/}, hideFields, requireFields),
	}
	return views, nil
}
