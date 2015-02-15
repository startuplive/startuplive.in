package root

import (
	// "github.com/ungerik/go-start/model"
	// "github.com/ungerik/go-start/modelext"
	// "github.com/ungerik/go-start/debug"
	"github.com/ungerik/go-start/user"
	. "github.com/ungerik/go-start/view"

	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
)

// var disabledProfileFields []string = []string{
// 	"Blocked",
// 	"Admin",
// 	"BirthYear",
// 	"EventOrganiser",
// 	"OrganiserInfo",
// 	"Judge",
// 	"JudgeInfo",
// 	"Mentor",
// 	"MentorInfo",
// 	"FeaturedMentor",
// 	"FeaturedMentorInfo",
// }

func init() {
	Profile = NewPublicPage("My Profile | Startup Live",
		DIV("public-content",
			DynamicView(
				func(ctx *Context) (view View, err error) {
					var person models.Person
					found, err := user.OfSession(ctx.Session, &person)
					if err != nil {
						return nil, err
					}
					if !found {
						return H1("You have to be logged in to edit your startup"), nil
					}
					email := person.PrimaryEmail()

					startupsview := Views{}
					i := person.GetStartups()
					var startup models.Startup
					for i.Next(&startup) {
						viewurl := MyStartup.URL(ctx.ForURLArgs(startup.ID.Hex()))
						startupsview = append(startupsview, DIV("profile-startup",
							A(viewurl, HTML(startup.Name.String())),
							BR(),
							SPAN("", HTML("last updated: "+startup.CreationDate.Get())),
							HR(),
						))
					}
					if i.Err() != nil {
						return nil, i.Err()
					}

					eventsview := Views{}
					i = person.GetOrganisedEvents()
					hasEvents := false
					var event models.Event
					for i.Next(&event) {
						var region models.EventRegion
						found, err := event.Region.TryGet(&region)
						if err != nil {
							return nil, err
						}
						if found {
							eventURL := Region0_Event1.URL(ctx.ForURLArgs(region.Slug.Get(), event.Number.String()))
							eventAdminURL := Region0_Event1_Admin.URL(ctx.ForURLArgs(region.Slug.Get(), event.Number.String()))
							hasEvents = true
							logoURL := region.HeaderLogoURL.Get()
							start := event.Start.Format("02/01 - ")
							end := event.End.Format("02/01 2006")
							view := &Div{
								Class: "profile-event",
								Content: Views{
									IMG(logoURL, 0, 60),
									DIV("",
										H3("Startup Live"+region.Name.Get()+" #"+event.Number.String()),
										HTML(start+end),
									),
									&If{
										Condition: event.IsPublished(),
										Content: Views{
											A(eventURL, "Public Page"),
											HTML(" | "),
										},
									},
									A(eventAdminURL, "Admin Page"),
									DivClearBoth(),
								},
							}
							eventsview = append(eventsview, view)
						}

					}

					view = DIV("row",
						DIV("cell right-border",
							TitleBar("My Profile"),
							DIV("main-content",
								H3("Email: ", email),
								H3("Name:"),
								P(&Form{
									SubmitButtonText:  "Save name",
									SubmitButtonClass: "button",
									FormID:            "profile",
									ExcludedFields:    []string{"Organization"},
									Labels:            map[string]string{"First": "Given Name", "Middle": "Middle Name", "Last": "Family Name"},
									GetModel: func(form *Form, ctx *Context) (interface{}, error) {
										return &person.Name, nil
									},
									OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
										return "", StringURL("."), person.Save()
									},
								}),
								H3("Password:"),
								P(&Form{
									SubmitButtonText:  "Save password",
									SubmitButtonClass: "button",
									FormID:            "password",
									GetModel: func(form *Form, ctx *Context) (interface{}, error) {
										return new(user.PasswordFormModel), nil
									},
									OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
										m := formModel.(*user.PasswordFormModel)
										person.Password.SetHashed(m.Password1.Get())
										return "", StringURL("."), person.Save()
									},
								}),
							),
						),
						DIV("cell right",
							TitleBarRight("My Startups"),
							DIV("main-content",
								startupsview,
								DIV("profile-startup",
									A(StartupForm.URL(ctx), "Create new"),
								),
							),
						),
						&If{
							Condition: hasEvents,
							Content: DIV("cell right",
								TitleBarRight("My Events"),
								DIV("main-content",
									eventsview,
								),
							),
						},

						DivClearBoth(),
					)
					return view, nil

				},
			),
		),
	)
}
