package admin

import (
	"errors"

	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/STARTeurope/startuplive.in/views/admin"
	// "github.com/ungerik/go-start/debug"
	. "github.com/ungerik/go-start/view"
)

func init() {
	Region0_Event1_Admin_Schedule2 = &Page{
		OnPreRender: SetEventSchedulePageData,
		Title:       EventAdminTitle("Schedule"),
		CSS:         IndirectURL(&Region0_DashboardCSS),
		Scripts:     admin.PageScripts,
		Content: Views{
			eventadminHeader(),
			&Form{
				SubmitButtonText:  "Save",
				SubmitButtonClass: "button",
				FormID:            "editscheduleitem",
				GetModel: func(form *Form, ctx *Context) (interface{}, error) {
					item := ctx.Data.(*PageData).ScheduleItem

					// item := models.EventScheduleItems.NewDocument().(*models.EventScheduleItem)
					item.Event.Set(ctx.Data.(*PageData).Event)
					return item, nil
				},
				OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
					item := formModel.(*models.EventScheduleItem)
					from := item.From.Time()
					until := item.Until.Time()
					if !from.Before(until) {
						return "", nil, errors.New("From must be before Until")
					}
					return "", StringURL("../"), item.Save()
				},
			},
		},
	}
}
