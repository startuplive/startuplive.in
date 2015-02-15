package event_1

import (
	// "github.com/ungerik/go-start/debug"
	. "github.com/ungerik/go-start/view"

	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
)

func init() {
	Region0_Event1_Mentor_Registration = newPublicEventPage(
		EventTitle("Mentor Registration"),
		JQuery,
		DynamicView(eventMentorRegistration),
	)
}

func eventMentorRegistration(ctx *Context) (view View, err error) {
	event := ctx.Data.(*PageData).Event

	excludedFields, err := excludedPersonFormFields(ctx)
	if err != nil {
		return nil, err
	}
	excludedFields = append(excludedFields, "JudgeInfo")

	requireFields := []string{
		"Name.First",
		"Name.Last",
		"Company",
		"Position",
		"Email",
		"MentorInfo",
		"Gender",
	}

	views := Views{
		H2("Mentor Form"),
		&Form{
			SubmitButtonText:  "Save Data",
			SubmitButtonClass: "button",
			FormID:            "newMentor",
			GetModel: func(form *Form, ctx *Context) (interface{}, error) {
				return models.NewPerson(), nil
			},
			ExcludedFields: append(AlwaysExcludedPersonFormFields, excludedFields...),
			DisabledFields: []string{ /*"Judge"*/},
			RequiredFields: requireFields,
			OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
				person := formModel.(*models.Person)
				person.Mentor = true
				err = person.Save()
				if err != nil {
					return "", nil, err
				}

				event.Mentors = append(event.Mentors, person.Ref())
				err = event.Save()
				if err != nil {
					person.Delete()
					return "", nil, err
				}
				return "", Region0_Event1_Registration_Success, nil
			},
		},
		// PersonForm(person, Region0_Event1_Registration_Success, []string{ /*"Judge"*/}, hideFields, requireFields),
	}
	return views, nil
}
