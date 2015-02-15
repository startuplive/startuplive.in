package admin

import (
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/STARTeurope/startuplive.in/views/admin"
	"github.com/ungerik/go-start/debug"
	"github.com/ungerik/go-start/model"
	. "github.com/ungerik/go-start/view"
)

func init() {
	Region0_Event1_Admin_Participant2 = &Page{
		OnPreRender: SetEventParticipantPageData,
		// Title: Render(
		// 	func(ctx *Context) (err error) {
		// 		participant := ctx.Data.(*PageData).Participant
		// 		return EventAdminTitle("Participant " + participant.Name()).Render(ctx)
		// 	},
		// ),
		CSS:     IndirectURL(&Region0_DashboardCSS),
		Scripts: admin.PageScripts,
		Content: Views{
			eventadminHeader(),
			DynamicView(
				func(ctx *Context) (view View, err error) {
					participant := ctx.Data.(*PageData).Participant
					person := participant.GetPerson()

					var cancelView View
					if participant.Cancelled() {
						cancelView = &Form{
							SubmitButtonText:  "Revoke Cancelation",
							SubmitButtonClass: "button",
							FormID:            "uncancel" + participant.ID.Hex(),
							OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
								participant.Uncancel()
								return "", StringURL("."), participant.Save()
							},
						}
					} else {
						cancelView = &Form{
							SubmitButtonText:  "Cancel Participant",
							SubmitButtonClass: "button",
							FormID:            "cancel" + participant.ID.Hex(),
							OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
								participant.Cancel(nil) // todo canceling admin user
								return "", StringURL("."), participant.Save()
							},
						}
					}

					//TODO extract functionality:
					teams := make([]*models.EventTeam, 0, 32)
					teamNames := make([]string, 1, 32) // First name will be an empty string
					selectedTeamName := ""
					i := EventTeamIterator(ctx)
					var doc models.EventTeam
					for i.Next(&doc) {
						team := doc
						teams = append(teams, &team)
						teamName := team.Name.Get()
						teamNames = append(teamNames, teamName)
						if team.ID == participant.Team.ID {
							selectedTeamName = teamName
						}
					}
					if i.Err() != nil {
						return nil, i.Err()
					}

					var teamNamesModel struct {
						Team model.DynamicChoice
					}
					teamNamesModel.Team.SetOptions(teamNames)
					teamNamesModel.Team.SetString(selectedTeamName)

					participantForm := &Form{
						SubmitButtonText:  "Save Event Specific Data",
						SubmitButtonClass: "button",
						FormID:            "participant" + participant.ID.Hex(),
						GetModel:          FormModel(participant),
						OnSubmit:          OnFormSubmitSaveModelAndRedirect(StringURL(".")),
					}
					participantForm.AddRestrictedMongoRefController(EventTeamIterator(ctx), "Team")

					excludedFields, err := ExcludedPersonFormFields(ctx)
					if err != nil {
						return nil, err
					}
					personForm := &Form{
						SubmitButtonText:  "Save Person Specific Data",
						SubmitButtonClass: "button",
						FormID:            "person" + person.ID.Hex(),
						GetModel:          FormModel(person),
						ExcludedFields:    append(AlwaysExcludedPersonFormFields, excludedFields...),
						RequiredFields: []string{
							"Name.First",
							"Name.Last",
						},
						OnSubmit: OnFormSubmitSaveModelAndRedirect(StringURL(".")),
					}

					participantsURL := Region0_Event1_Admin_Participants.URL(ctx.ForURLArgs(ctx.URLArgs[0], ctx.URLArgs[1]))

					name := participant.Name()
					if participant.Cancelled() {
						name += " (cancelled)"
					} else if !participant.CheckedIn.Get() {
						name += " (absent)"
					}

					debug.Nop(personForm, participantForm)

					views := Views{
						Printf("<h2><a href='%s'>&larr;Back to Participants</a></h2>", participantsURL),
						Printf("<h2>%s</h2>", name),
						cancelView,
						HTML("<h3>Event Specific Data</h2>"),
						participantForm,
						HR(),
						HTML("<h3>Person Specific Data</h2>"),
						personForm,
						HR(),
						Printf("<h2><a href='%s'>&larr;Back to Participants</a></h2>", participantsURL),
					}
					return views, nil
				},
			),
		},
	}
}

func ExcludedPersonFormFields(ctx *Context) (excludeFields []string, err error) {
	orga, err := OnlyRegionAdminOrEventOrga.Authenticate(ctx)
	if err != nil {
		return nil, err
	}
	if orga {
		excludeFields = []string{
			"Username",
			"Blocked",
			"Admin",
			"SuperAdmin",
			"BirthYear",
			"University",
			"PostalAddress.Country",
		}
	}
	return excludeFields, nil
}
