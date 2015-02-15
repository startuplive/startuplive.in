package root

import (
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	// "github.com/ungerik/go-start/debug"
	// "github.com/ungerik/go-start/model"
	. "github.com/ungerik/go-start/view"
)

func init() {
	Organisers = NewPublicPage("Organisers | Startup Live", DynamicView(
		func(ctx *Context) (view View, err error) {
			var images Views
			i := models.People.Filter("EventOrganiser", true).Sort("Name.First").Sort("Name.Last").Iterator()
			// i = model.NewRandomIterator(i)
			var person models.Person
			for i.Next(&person) {
				if person.OrganiserInfo != "" && person.HasImage_160x160_and_284x144() {
					images = append(images, DIV("organiser-container",
						ViewOrError(person.Image_160x160("")),
						EventOrganiserView(&person),
					))
				}
			}
			if i.Err() != nil {
				return nil, i.Err()
			}

			view = DIV("public-content",
				H1("Meet the amazing Startup Live organisers!"),
				DIV("organisers",
					//P("Volutpat claram option luptatum nihil claritatem. Saepius vel futurum zzril lorem typi. Humanitatis enim dolor nunc eodem vulputate. Saepius adipiscing in quam aliquam duis."),
					images,
				),
			)
			return view, nil
		},
	))
}
