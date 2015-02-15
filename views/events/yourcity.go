package events

import (
	"fmt"
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/ungerik/go-mail"
	"github.com/ungerik/go-start/model"
	. "github.com/ungerik/go-start/view"
)

func init() {
	Events_YourCity = NewPublicPage("Your City", DynamicView(
		func(ctx *Context) (view View, err error) {
			 var hostContactPerson *models.Person
			 _, err = models.People.Filter("Name.First", "Florina").Filter("Name.Last", "Dumitrache").TryOneDocument(&hostContactPerson)
			 if err != nil {
			 	return nil, err
			  }

			var partnerContactPerson *models.Person
			_, err = models.People.Filter("Email.Address", "guney.kose@startuplive.in").TryOneDocument(&partnerContactPerson)
			if err != nil {
				return nil, err
			}

			view = DIV("public-content",
				H1("How to get Startup Live to your city"),
				DIV("row first",
					DIV("cell right-border",
						TitleBar("Host an event"),
						DIV("main-content",
							P("It is important to understand that Startup Live is more than just a single event. It's an attitude. It means to actively push entrepreneurship in your community and continuously strive for entrepreneurial change. In the end, it depends on you and what you want to get out of this experience. We really hope you like our vision and will choose to be a vital part of it."),
							BR(),
							 EventContact(hostContactPerson),
						),
					),
					DIV("cell right",
						TitleBarRight("Become a partner"),
						DIV("main-content",
							P("Want to be popular? You're not only partnering with a concept Europe needs, you're also being introduced to an audience that love/will love Startup Live. We offer access to a European-wide community of passionate entrepreneurs, startup geeks and people who are actively building a better ecosystem for the entrepreneurial community in Europe. The platform already has worldwide contacts and outstanding reviews and the plans are to double them in the next term. Feel free to approach!"),
							EventContact(partnerContactPerson),
						),
					),
					DivClearBoth(),
				),
				DIV("row",
					DIV("cell",
						TitleBar("Suggest to host a city"),
						DIV("main-content",
							P("Our vision is to connect all local startup communities in Europe. Your city may be missing at the moment. Just give us a hint and we we'll get back to you as soon as possible."),
							&Form{
								SubmitButtonText:  "Notify me",
								SuccessMessage:    "Thank you for your suggestion. We will check our network and will hopefully be in your city soon.",
								SubmitButtonClass: "button",
								FormID:            "suggestcity",
								GetModel: func(form *Form, ctx *Context) (interface{}, error) {
									var suggestion models.CitySuggestion
									models.CitySuggestions.InitDocument(&suggestion)
									return &suggestion, nil
								},
								OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
									city := formModel.(*models.CitySuggestion)
									city.Date.SetNowUTC()
									city.Type.Set("Startup Live Event")

									err := city.Save()
									if err != nil {
										return "", nil, err
									}

									subject := fmt.Sprintf("%s suggests city: %s [EOM]", city.Email, city.Name)
									err = email.NewBriefMessage(subject, subject, ContactEmail).Send()

									return "", nil, err
								},
							},
						),
					),
					DIV("cell right",
						DIV("main-content",
							IMG("/images/logo-startup-live-your-city.png"),
						),
					),
					DivClearBoth(),
				),
			)
			return view, nil
		},
	))
}

type suggestCityFormModel struct {
	City  model.String `view:"size=20|label=Suggested city" model:"minlen=3|maxlen=40"`
	Email model.Email  `view:"size=20|label=Your email address" model:"required|maxlen=40"`
}
