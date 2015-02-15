package dashboard

import (
	"fmt"
	"html"

	"github.com/ungerik/go-start/model"
	. "github.com/ungerik/go-start/view"
	"labix.org/v2/mgo/bson"

	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
)

func init() {
	Region0_Event1_Dashboard_Judge2 = &Page{
		OnPreRender: SetEventPersonPageData,
		Title: Render(
			func(ctx *Context) (err error) {
				person := ctx.Data.(*PageData).Person
				return EventDashboardTitle("Judge " + person.Name.String()).Render(ctx)
			},
		),
		CSS: Region0_DashboardSubmodalCSS,
		Content: Views{
			&ModelIteratorView{
				GetModel: func(ctx *Context) (interface{}, error) {
					return new(models.Person), nil
				},
				GetModelIterator: func(ctx *Context) model.Iterator {
					return models.People.DocumentWithIDIterator(bson.ObjectIdHex(ctx.URLArgs[2]))
				},
				GetModelView: func(ctx *Context, model interface{}) (view View, err error) {
					person := model.(*models.Person)
					company := ""
					if person.Company.Get() != "" && person.Position.Get() != "" {
						company = fmt.Sprintf("<h2>%s @ %s</h2>", html.EscapeString(person.Position.Get()), person.Company.Get())
					} else {
						company = fmt.Sprintf("<h2>%s%s&nbsp;</h2>", html.EscapeString(person.Position.Get()), person.Company.Get())
					}
					//email := person.Contact.Email.Get()
					view = &Div{
						Class: "person-details",
						Content: Views{
							H1(person.Name),
							Escape(company),
							HTML("<div class='spacer-top'></div>"),
							ViewOrError(person.Image_320x320("framed-image")),
							HTML("<div class='buttons'>"),
							//HTML("<a href='#' class='button'>Start conversation</a> <a href='#' class='button'>Social Media Profiles</a>"),
							HTML("</div>"),
							Printf("<div class='tags'>%s</div>", person.Tags),
						},
					}
					return view, nil
				},
			},
		},
	}
}
