package admin

import (
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	. "github.com/ungerik/go-start/view"
)

func init() {
	Region0_Event1_Admin_ExportEmails = NewViewURLWrapper(
		RenderView(
			func(ctx *Context) (err error) {
				ctx.Response.ContentDispositionAttachment("emails.txt")
				_, event, err := RegionAndEvent(ctx.URLArgs)
				if err != nil {
					return err
				}
				i := event.ParticipantPersonIterator(false)
				var person models.Person
				for i.Next(&person) {
					if len(person.Email) > 0 {
						ctx.Response.Printf("\"%s\" <%s>, ", person.Name.String(), person.Email[0].Address)
					}
				}
				return i.Err()
			},
		),
	)
}
