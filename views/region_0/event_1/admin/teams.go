package admin

import (
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/STARTeurope/startuplive.in/views/admin"
	// "github.com/ungerik/go-start/debug"
	. "github.com/ungerik/go-start/view"
)

func init() {
	Region0_Event1_Admin_Teams = &Page{
		OnPreRender: SetEventPageData,
		Title:       EventAdminTitle("Teams"),
		CSS:         IndirectURL(&Region0_DashboardCSS),
		Scripts:     admin.PageScripts,
		Content: Views{
			eventadminHeader(),
			H2("Pitching Teams:"),
			&ModelIteratorTableView{
				Class:            "visual-table",
				GetModelIterator: EventPitchingTeamIterator,
				GetRowModel: func(ctx *Context) (interface{}, error) {
					return new(models.EventTeam), nil
				},
				GetHeaderRowViews: TableHeaderRowEscape("Nr", "Name", "Tagline", "Leader", "Cancelled", "Edit", "Delete"),
				GetRowViews: func(row int, rowModel interface{}, ctx *Context) (views Views, err error) {
					team := rowModel.(*models.EventTeam)

					var cancelled string
					if team.Cancelled() {
						cancelled = "cancelled"
					}

					editURL := Region0_Event1_Admin_Team2.URL(
						ctx.ForURLArgs(ctx.URLArgs[0], ctx.URLArgs[1], team.ID.Hex()))

					views = Views{
						Printf("%d", row+1),
						Escape(team.Name.Get()),
						Escape(team.Tagline.Get()),
						Escape(team.LeaderName()),
						Escape(cancelled),
						A(editURL, "Edit"),
						&Form{
							SubmitButtonText:    "Delete",
							SubmitButtonConfirm: "Are you sure you want to delete team " + team.Name.Get() + "?",
							FormID:              "delete" + team.ID.Hex(),
							OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
								return "", StringURL("."), team.Delete()
							},
						},
					}
					return views, nil
				},
			},
			BR(),
			HR(),
			BR(),
			H2("Not Pitching:"),
			&ModelIteratorTableView{
				Class:            "visual-table",
				GetModelIterator: EventNotPitchingTeamIterator,
				GetRowModel: func(ctx *Context) (interface{}, error) {
					return new(models.EventTeam), nil
				},
				GetHeaderRowViews: TableHeaderRowEscape("Nr", "Name", "Tagline", "Leader", "Cancelled", "Delete", "Edit"),
				GetRowViews: func(row int, rowModel interface{}, ctx *Context) (views Views, err error) {
					team := rowModel.(*models.EventTeam)

					var cancelled string
					if team.Cancelled() {
						cancelled = "cancelled"
					}

					editURL := Region0_Event1_Admin_Team2.URL(
						ctx.ForURLArgs(ctx.URLArgs[0], ctx.URLArgs[1], team.ID.Hex()))

					views = Views{
						Printf("%d", row+1),
						Escape(team.Name.Get()),
						Escape(team.Tagline.Get()),
						Escape(team.LeaderName()),
						Escape(cancelled),
						A(editURL, "Edit"),
						&Form{
							SubmitButtonText:    "Delete",
							SubmitButtonConfirm: "Are you sure you want to delete team " + team.Name.Get() + "?",
							FormID:              "delete" + team.ID.Hex(),
							OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
								return "", StringURL("."), team.Delete()
							},
						},
					}
					return views, nil
				},
			},
			BR(),
			HR(),
			BR(),
			&Form{
				SubmitButtonText:  "Create Team",
				SubmitButtonClass: "button",
				FormID:            "createteam",
				OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
					event := ctx.Data.(*PageData).Event

					team := models.NewEventTeam()
					team.Event.Set(event)
					//team.Leader.Set()
					team.Name = "[Name]"
					team.Tagline = "[Tagline]"
					team.Pitching = true
					err := team.Save()
					if err != nil {
						return "", nil, err
					}

					return "", StringURL(Region0_Event1_Admin_Team2.URL(
						ctx.ForURLArgs(ctx.URLArgs[0], ctx.URLArgs[1], team.ID.Hex()))), nil
				},
			},
		},
	}
}
