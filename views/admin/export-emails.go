package admin

import (
	"github.com/ungerik/go-start/model"
	. "github.com/ungerik/go-start/view"
	//	"github.com/ungerik/go-start/debug"

	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
)

func init() {
	Admin_ExportEmails = NewViewURLWrapper(
		RenderView(
			func(ctx *Context) (err error) {
				ctx.Response.ContentDispositionAttachment("emails.csv")

				i := peopleIterator()
				var person models.Person
				// "Email Address","First Name","Last Name"
				for i.Next(&person) {
					if len(person.Email) > 0 {
						ctx.Response.Printf("%s, %s, %s \n", person.Email[0].Address, person.Name.First.String(), person.Name.Last.String())
					}
				}
				return i.Err()
			},
		),
	)
}

func peopleIterator() model.Iterator {
	return models.People.Iterator()
}
