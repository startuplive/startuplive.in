package event_1

import (
	"github.com/ungerik/go-start/user"
	. "github.com/ungerik/go-start/view"

	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
)

func init() {
	Region0_Event1_Pitcher_Registration = newPublicEventPage(
		EventTitle("Pitcher Registration"),
		JQuery,
		DynamicView(eventPitcherRegistration),
	)
}

func eventPitcherRegistration(ctx *Context) (view View, err error) {
	event := ctx.Data.(*PageData).Event

	excludeFields := []string{
		"Judgements",
		"CrowdVotes",
		"CancelledDate",
		"CancelledBy",
		"Pitching",
		"PitchtrainingBooked",
		"PitchtrainingAttended",
		"BookedMentors",
		"Leader",
		"Event",
	}

	views := Views{
		H2("Pitcher Form"),
		&Form{
			SubmitButtonText:  "Save Data",
			SubmitButtonClass: "button",
			FormID:            "newIdea",
			GetModel: func(form *Form, ctx *Context) (interface{}, error) {
				team := models.NewEventTeam()
				team.Event.Set(event)
				var person models.Person
				ok, err := user.OfSession(ctx.Session, &person)
				if err != nil {
					return nil, err
				}
				if ok {
					team.Leader.Set(&person)
				}
				return team, nil
			},
			ExcludedFields: excludeFields,
			DisabledFields: []string{ /* Pitcher */},
			RequiredFields: []string{"LeaderNameInput", "EmailInput"},
			OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {

				team := formModel.(*models.EventTeam)

				team.Pitching = true

				err = team.Save()
				if err != nil {
					team.Delete()
					// debug.Print(" - removed Person")
					return "", nil, err
				}
				// debug.Print(" before return", err)
				return "", Region0_Event1_Registration_Success, nil
			},
		},
		// PersonForm(person, Region0_Event1_Registration_Success, []string{ /*"Judge"*/}, hideFields, requireFields),
	}
	return views, nil
}
