package admin

import (
	"errors"
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/STARTeurope/startuplive.in/views/admin"
	"github.com/ungerik/go-start/model"
	// "github.com/ungerik/go-start/debug"
	. "github.com/ungerik/go-start/view"
)

func init() {
	Region0_Event1_Admin_Schedule = &Page{
		OnPreRender: SetEventPageData,
		Title:       EventAdminTitle("Schedule"),
		CSS:         IndirectURL(&Region0_DashboardCSS),
		Scripts:     admin.PageScripts,
		Content: Views{
			eventadminHeader(),
			&ModelIteratorTableView{
				Class: "visual-table",
				GetModelIterator: func(ctx *Context) model.Iterator {
					return ctx.Data.(*PageData).Event.ScheduleItemIterator()
				},
				GetRowModel: func(ctx *Context) (interface{}, error) {
					return new(models.EventScheduleItem), nil
				},
				GetHeaderRowViews: TableHeaderRowEscape("From", "Until", "Title", "Location", "Remove"),
				GetRowViews: func(row int, rowModel interface{}, ctx *Context) (views Views, err error) {
					item := rowModel.(*models.EventScheduleItem)
					editURL := Region0_Event1_Admin_Schedule2.URL(
						ctx.ForURLArgs(ctx.URLArgs[0], ctx.URLArgs[1], item.ID.Hex()))
					views = Views{
						HTML(item.From.Format(model.ShortDateTimeFormat)),
						HTML(item.Until.Format("&nbsp;- 15:04")),
						Escape(item.Title.Get()),
						Escape(item.Location.Get()),
						A(editURL, "edit"),
						&Form{
							SubmitButtonText:  "Remove",
							SubmitButtonClass: "",
							FormID:            "remove" + item.ID.Hex(),
							OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
								return "", StringURL("."), item.Delete()
							},
						},
					}
					return views, nil
				},
			},
			HR(),
			&Form{
				SubmitButtonText:  "Add",
				SubmitButtonClass: "button",
				FormID:            "addscheduleitem",
				GetModel: func(form *Form, ctx *Context) (interface{}, error) {
					item := models.NewEventScheduleItem()
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
					return "", StringURL("."), item.Save()
				},
			},
		},
	}
}
