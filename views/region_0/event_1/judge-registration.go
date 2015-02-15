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
	Region0_Event1_Judge_Registration = newPublicEventPage(
		EventTitle("Judge Registration"),
		JQuery,
		DynamicView(eventJudgeRegistration),
	)
}

func eventJudgeRegistration(ctx *Context) (view View, err error) {
	event := ctx.Data.(*PageData).Event

	excludedFields, err := excludedPersonFormFields(ctx)
	if err != nil {
		return nil, err
	}
	excludedFields = append(excludedFields, "MentorInfo")

	requireFields := []string{
		"Name.First",
		"Name.Last",
		"Company",
		"Position",
		"Email.Address",
		"JudgeInfo",
	}
	views := Views{
		H2("Judge Form"),
		&Form{
			SubmitButtonText:  "Save Data",
			SubmitButtonClass: "button",
			FormID:            "newJudge",
			GetModel: func(form *Form, ctx *Context) (interface{}, error) {
				return models.NewPerson(), nil
			},
			ExcludedFields: append(AlwaysExcludedPersonFormFields, excludedFields...),
			DisabledFields: []string{ /*"Judge"*/},
			RequiredFields: requireFields,
			OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
				person := formModel.(*models.Person)

				person.Judge = true
				// debug.Print(" saving Person")

				err = person.Save()
				// debug.Print(" saved Person", err)
				if err != nil {
					// debug.Print(" - removed Person")
					return "", nil, err
				}

				event.Judges = append(event.Judges, person.Ref())
				err = event.Save()
				if err != nil {
					person.Delete()
					return "", nil, err
				}
				// debug.Print(" before return", err)
				return "", Region0_Event1_Registration_Success, person.Save()
			},

		},
		// PersonForm(person, Region0_Event1_Registration_Success, []string{ /*"Judge"*/}, hideFields, requireFields),
	}
	return views, nil
}

func excludedPersonFormFields(ctx *Context) (excludedFields []string, err error) {

	excludedFields = []string{
		"Username",
		"Blocked",
		"Admin",
		"SuperAdmin",
		"BirthDate",
		"BirthYear",
		"University",
		"PostalAddress.FirstLine",
		"PostalAddress.SecondLine",
		"PostalAddress.ZIP",
		"PostalAddress.City",
		"PostalAddress.State",
		"PostalAddress.Country",
		"Citizenship",
		"Images.URL_50x50",
		"Images.URL_100x100",
		"Images.URL_160x160",
		"Images.URL_320x320",
		"Images.URL_284x144",
		"EventOrganiser",
		"OrganiserInfo",
		"OrganiserEmail",
		"OrganiserForwardingEmail",
		"Judge",
		"Mentor",
		"FeaturedMentor",
		"FeaturedMentorInfo",
		"Name.Prefix",
		"Name.Postfix",
		"Name.Middle",
		"GitHub",
		"Xing",
		"Twitter",
		"Facebook",
		"Web",
	}

	return excludedFields, nil
}
