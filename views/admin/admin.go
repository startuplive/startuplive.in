package admin

import (
	"github.com/ungerik/go-start/mongoadmin"
	"github.com/ungerik/go-start/mongomedia"
	. "github.com/ungerik/go-start/view"

	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
)

var PageScripts = Renderers{
	JQuery,
	RenderUserVoice,
}

func init() {
	Admin = &Page{
		Title:   HTML("Admin | startuplive.in"),
		CSS:     IndirectURL(&Admin_CSS),
		Scripts: PageScripts,
		Content: Views{
			adminHeader(),
			&Form{
				FormID:            "clearcaches",
				SubmitButtonText:  "Clear Page Caches",
				SubmitButtonClass: "button",
				OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
					ClearAllCaches()
					return "All caches cleared", nil, nil
				},
			},
			mongoadmin.RemoveInvalidRefsButton("button"),
			DynamicView(
				func(ctx *Context) (View, error) {
					form := mongoadmin.RemoveInvalidRefsButton("button", models.Events)
					form.SubmitButtonText += " - events only"
					form.FormID += "_events_only"
					return form, nil
				},
			),
			DynamicView(
				func(ctx *Context) (View, error) {
					form := mongoadmin.RemoveInvalidRefsButton("button", mongomedia.Config.Backend.Images)
					form.SubmitButtonText += " - images only"
					form.FormID += "_images_only"
					return form, nil
				},
			),
		},
	}
}
