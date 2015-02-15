package event_1

import (
	"github.com/ungerik/go-start/debug"
	. "github.com/ungerik/go-start/view"

	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
)

func init() {
	Region0_Event1_FeedbackParticipants = newPublicEventPage(
		EventTitle("Participants Feedback"),
		nil,
		DynamicView(eventFeedbackParticipants),
	)
}

func eventFeedbackParticipants(ctx *Context) (view View, err error) {
	event := ctx.Data.(*PageData).Event

	eventTeams := make([]*models.EventTeam, 0, 128)
	eventTeamNames := make([]string, 0, 128)
	eventTeamNames = append(eventTeamNames, "")
	i := event.TeamIterator()
	var team models.EventTeam
	for i.Next(&team) {
		eventTeams = append(eventTeams, &team)
		eventTeamNames = append(eventTeamNames, team.Name.String())
	}
	if i.Err() != nil {
		return nil, i.Err()
	}

	views := Views{
		H2("Feedback Form"),
		&Form{
			SubmitButtonText:  "Submit",
			SubmitButtonClass: "button",
			FormID:            "participantfeedback",
			GetModel: func(form *Form, ctx *Context) (interface{}, error) {
				feedback := models.NewFeedbackParticipant()
				feedback.Event.Set(event)
				feedback.TeamChoice.SetOptions(eventTeamNames)
				return feedback, nil
			},
			OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
				feedback := formModel.(*models.FeedbackParticipant)
				teamName := feedback.TeamChoice.String()
				for _, team := range eventTeams {
					if team.Name.String() == teamName {
						feedback.Team.Set(team)
						feedback.TeamChoice.SetString(teamName)
					}
				}
				feedback.Save()

				// var feedback2 models.FeedbackParticipant
				// err := models.FeedbackParticipants.DocumentWithID(feedback.ID, &feedback2)
				// debug.Dump(err)
				debug.Nop()

				// i := event.TeamIterator()
				// for doc := i.Next(); doc != nil; doc = i.Next() {
				// 	team := doc.(*models.EventTeam)
				// 	if team.Name.String() == teamName {
				// 		feedback.Team.Set(team)
				// 	}
				// }
				// if i.Err() != nil {
				// 	return "", nil, i.Err()
				// }

				// debug.Print(" before return", err)
				return "", Region0_Event1_Registration_Success, nil
			},

			// Redirect: Region0_Event1_Registration_Success,
		},
		// PersonForm(person, Region0_Event1_Registration_Success, []string{ /*"Judge"*/}, hideFields, requireFields),
	}
	return views, nil
}
