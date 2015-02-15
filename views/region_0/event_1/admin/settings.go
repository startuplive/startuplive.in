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
	Region0_Event1_Admin_Settings = &Page{
		OnPreRender: SetEventPageData,
		Title:       EventAdminTitle("About"),
		CSS:         IndirectURL(&Region0_DashboardCSS),
		Scripts:     admin.PageScripts,
		Content: Views{
			eventadminHeader(),
			DynamicView(
				func(ctx *Context) (view View, err error) {
					debug.Nop()
					view = Views{
						H3("Settings Site"),
					}

					return view, nil
				},
			),
			H4("Google Analytics Code"),
			DynamicView(
				func(ctx *Context) (view View, err error) {
					event := ctx.Data.(*PageData).Event

					return &Form{
						SubmitButtonText:  "Submit",
						SubmitButtonClass: "button",
						FormID:            "googleanalytics",
						GetModel: func(form *Form, ctx *Context) (interface{}, error) {
							var code model.String
							if event.GoogleAnalyticsHostAccount.Get() != "" {
								code = event.GoogleAnalyticsHostAccount
							}

							return &googleAnalyticsFormModel{
								Code: code,
							}, nil
						},
						OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
							m := formModel.(*googleAnalyticsFormModel)
							event.GoogleAnalyticsHostAccount = m.Code
							return "", StringURL("."), event.Save()
						},
					}, nil
				},
			),
		},
	}
}

type googleAnalyticsFormModel struct {
	Code model.String `view:"label=Google Analytics ID"`
}
