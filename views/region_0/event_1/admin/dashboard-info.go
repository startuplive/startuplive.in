package admin

import (
	// "github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/STARTeurope/startuplive.in/views/admin"
	"github.com/ungerik/go-start/debug"
	"github.com/ungerik/go-start/model"
	// "github.com/AlexTi/go-amiando"
	. "github.com/ungerik/go-start/view"
)

func init() {
	Region0_Event1_Admin_DashboardInfo = &Page{
		OnPreRender: SetEventPageData,
		Title:       EventAdminTitle("Dashboard Info Page"),
		CSS:         IndirectURL(&Region0_DashboardCSS),
		Scripts:     admin.PageScripts,
		Content: Views{
			eventadminHeader(),
			DynamicView(
				func(ctx *Context) (view View, err error) {
					debug.Nop()
					view = Views{
						H3("Dashboard Info Insert"),
					}
					return view, nil
				},
			),
			DynamicView(
				func(ctx *Context) (view View, err error) {
					event := ctx.Data.(*PageData).Event
					return &Form{
						FormID:                   "info",
						SubmitButtonText:         "Update Dashboard Info",
						SubmitButtonClass:        "button",
						GeneralErrorOnFieldError: false,
						GetModel: func(form *Form, ctx *Context) (interface{}, error) {
							return &InfoFormModel{
								ShowInfo:  event.Show.Info,
								Info_HTML: event.DashboardInfo_HTML,
							}, nil
						},
						OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
							m := formModel.(*InfoFormModel)
							event.Show.Info = m.ShowInfo
							event.DashboardInfo_HTML = m.Info_HTML
							return "", StringURL("."), event.Save()
						},
					}, nil
				}),
		},
	}
}

type InfoFormModel struct {
	ShowInfo  model.Bool `view:"label=show info on dashboard site"`
	Info_HTML model.Text `view:"rows=30"`
}
