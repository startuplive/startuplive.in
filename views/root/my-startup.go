package root

import (
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	"labix.org/v2/mgo/bson"
	// "github.com/ungerik/go-start/model"
	// "github.com/ungerik/go-start/modelext"
	// "github.com/ungerik/go-start/debug"
	"github.com/ungerik/go-start/user"
	. "github.com/ungerik/go-start/view"
)

func init() {
	MyStartup = NewPublicPage("My Startup | Startup Live",
		DIV("public-content",
			DIV("main",
				DIV("main-content",
					DynamicView(
						func(ctx *Context) (view View, err error) {

							if !user.LoggedIn(ctx.Session) {
								return H1("You have to be logged in to edit your startup"), nil
							}

							var startup *models.Startup

							startup, err = getStartup(ctx)
							if err != nil {
								return nil, err
							}

							views := Views{
								A(Profile.URL(ctx), "< Back to Profile"),
								HR(),
								H2(startup.Name.String()),
								HR(),
								&Form{
									SubmitButtonText:  "Update",
									SubmitButtonClass: "button",
									FormID:            "startup",
									Class:             "public-form",
									ExcludedFields:    []string{},
									GetModel:          FormModel(startup),
									OnSubmit:          OnFormSubmitSaveModelAndRedirect(StringURL(".")),
								},
							}
							return views, nil
						},
					),
				),
			),
		),
	)
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
