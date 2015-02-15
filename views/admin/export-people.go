package admin

import (
	"github.com/ungerik/go-start/debug"
	"github.com/ungerik/go-start/model"
	. "github.com/ungerik/go-start/view"

	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
)

func init() {
	Admin_ExportPeople = NewViewURLWrapper(
		RenderView(
			func(ctx *Context) (err error) {
				ctx.Response.ContentDispositionAttachment("people.csv")

				i := participantIterator()
				var p models.EventParticipant
				// "Email Address","First Name","Last Name"
				count := 0
				for i.Next(&p) {
					person := p.GetPerson()
					if &p != nil && person != nil {

						// ctx.Response.Printf("%s, %s, %s \n", person.Email[0].Address, person.Name.First.String(), person.Name.Last.String())
						if len(person.Email) > 0 && person.Name.Last.String() != "" && person.Gender.Get() != "" && person.Citizenship.Get() != "" && p.Background.Get() != "" {
							ctx.Response.Printf("%s, %s, %s, %s, %s, %s\n", person.Email[0].Address, person.Name.First.String(), person.Name.Last.String(), person.Citizenship.Get(), person.Gender.Get(), p.Background.Get())
							count++
						}
					}

				}
				debug.Print(count)
				return i.Err()
			},
		),
	)
}

func participantIterator() model.Iterator {
	return models.EventParticipants.FilterExists("person", true).Iterator()
}
