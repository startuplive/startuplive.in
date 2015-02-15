package admin

import (
	. "github.com/ungerik/go-start/view"

	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/STARTeurope/startuplive.in/views/admin"
)

func init() {
	Region0_Event1_Admin_Team2 = &Page{
		OnPreRender: SetEventTeamPageData,
		Title: Render(
			func(ctx *Context) (err error) {
				team := ctx.Data.(*PageData).Team
				return EventAdminTitle("Team " + team.Name.Get()).Render(ctx)
			},
		),
		CSS:     IndirectURL(&Region0_DashboardCSS),
		Scripts: admin.PageScripts,
		Content: Views{
			eventadminHeader(),
			DynamicView(
				func(ctx *Context) (view View, err error) {
					team := ctx.Data.(*PageData).Team
					event := ctx.Data.(*PageData).Event

					teamForm := &Form{
						FormID:            team.ID.Hex(),
						SubmitButtonClass: "button",
						ExcludedFields:    []string{"Judgements"},
						GetModel:          FormModel(team),
						OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
							return "", Region0_Event1_Admin_Teams, team.Save()
						},
					}
					teamForm.AddRestrictedMongoRefController(event.ParticipantPersonIterator(false), "Leader")

					var cancelView View
					if team.Cancelled() {
						cancelView = &Form{
							SubmitButtonText:  "Revoke Cancelation",
							SubmitButtonClass: "button",
							FormID:            "uncancel" + team.ID.Hex(),
							OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
								return "", StringURL("."), team.UncancelAndSave()
							},
						}
					} else {
						cancelView = &Form{
							SubmitButtonText:  "Cancel Team",
							SubmitButtonClass: "button",
							FormID:            "cancel" + team.ID.Hex(),
							OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
								return "", StringURL("."), team.CancelAndSave(nil) // todo canceling admin user
							},
						}
					}

					teamName := team.Name.Get()
					if team.Cancelled() {
						teamName += " (cancelled)"
					}

					views := Views{
						H2(A(Region0_Event1_Admin_Teams, HTML("&larr;Back"))),
						H2(teamName),
						cancelView,
						HR(),
						// teamLeaderChooser,
						// HR(),
						H2("Team Data:"),
						teamForm,
					}
					return views, nil
				},
			),
		},
	}
}
