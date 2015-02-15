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
	Region0_Event1_Admin_FAQ = &Page{
		OnPreRender: SetEventPageData,
		Title:       EventAdminTitle("FAQ"),
		CSS:         IndirectURL(&Region0_DashboardCSS),
		Scripts:     admin.PageScripts,
		Content: Views{
			eventadminHeader(),
			DynamicView(
				func(ctx *Context) (view View, err error) {
					debug.Nop()
					view = Views{
						H3("FAQ Insert"),
					}
					return view, nil
				},
			),
			DynamicView(
				func(ctx *Context) (view View, err error) {
					event := ctx.Data.(*PageData).Event
					return &Form{
						FormID:                   "faq",
						SubmitButtonText:         "Update FAQ",
						SubmitButtonClass:        "button",
						GeneralErrorOnFieldError: false,
						GetModel: func(form *Form, ctx *Context) (interface{}, error) {
							return &FaqFormModel{
								ShowFaq:  event.Show.FAQ,
								FAQ_HTML: event.FAQ_HTML,
							}, nil
						},
						OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
							m := formModel.(*FaqFormModel)
							event.Show.FAQ = m.ShowFaq
							event.FAQ_HTML = m.FAQ_HTML
							return "", StringURL("."), event.Save()
						},
					}, nil
				}),
		},
	}
}

type FaqFormModel struct {
	ShowFaq  model.Bool `view:"label=show faq on public site"`
	FAQ_HTML model.Text `view:"rows=30"`
}
