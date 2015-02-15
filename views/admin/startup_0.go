package admin

import (
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	"labix.org/v2/mgo/bson"
	// "github.com/ungerik/go-start/model"
	. "github.com/ungerik/go-start/view"
)

func init() {
	Admin_Startup0 = &Page{
		Title: Render(
			func(ctx *Context) error {
				startup, err := getStartup(ctx)
				if err != nil {
					return err
				}
				ctx.Response.WriteString(startup.Name.String())
				ctx.Response.WriteString(" | Admin")
				return nil
			},
		),
		CSS:     IndirectURL(&Admin_CSS),
		Scripts: PageScripts,
		Content: Views{
			adminHeader(),
			DynamicView(
				func(ctx *Context) (view View, err error) {
					startup, err := getStartup(ctx)
					if err != nil {
						return nil, err
					}

					views := Views{
						H2(startup.Name.String()),
						HR(),
						&Form{
							SubmitButtonText:  "Update",
							SubmitButtonClass: "button",
							FormID:            "startup",
							GetModel:          FormModel(startup),
							OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
								return "", StringURL("."), nil
							},
						},
					}
					return views, nil
				},
			),
		},
	}
}

func getStartup(ctx *Context) (startup *models.Startup, err error) {
	id := bson.ObjectIdHex(ctx.URLArgs[0])
	found, err := models.Startups.TryDocumentWithID(id, &startup)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, NotFound("404: Startup not found")
	}
	return startup, nil
}
