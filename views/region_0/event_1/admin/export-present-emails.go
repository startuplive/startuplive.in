package admin

import (
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	. "github.com/ungerik/go-start/view"
)

func init() {
	Region0_Event1_Admin_ExportPresentEmails = NewViewURLWrapper(
		RenderView(
			func(ctx *Context) (err error) {
				ctx.Response.ContentDispositionAttachment("emails.txt")
				_, event, err := RegionAndEvent(ctx.URLArgs)
				if err != nil {
					return err
				}
				i := event.ParticipantIterator()
				var participant models.EventParticipant
				for i.Next(&participant) {
					person := participant.GetPerson()
					if !participant.Cancelled() && participant.PresentsIdea.Get() && len(person.Email) > 0 {
						ctx.Response.Printf("\"%s\" <%s>, ", person.Name.String(), person.Email[0].Address)
					}
				}
				return i.Err()
			},
		),
	)
}
