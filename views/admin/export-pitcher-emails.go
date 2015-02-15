package admin

import (
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/ungerik/go-start/model"
	. "github.com/ungerik/go-start/view"
	//	"github.com/ungerik/go-start/debug"
)

func init() {
	Admin_ExportPitcherEmails = NewViewURLWrapper(
		RenderView(
			func(ctx *Context) (err error) {
				ctx.Response.ContentDispositionAttachment("pitcher-emails.csv")

				i := pitcherIterator()
				var pitcher models.EventParticipant
				// "Email Address","First Name","Last Name"
				for i.Next(&pitcher) {
					person := pitcher.GetPerson()
					if len(person.Email) > 0 {
						ctx.Response.Printf("%s, %s, %s \n", person.Email[0].Address, person.Name.First.String(), person.Name.Last.String())
					}
				}
				return i.Err()
			},
		),
	)
}

func pitcherIterator() model.Iterator {
	return models.EventParticipants.Filter("PresentsIdea", true).Iterator()
}
