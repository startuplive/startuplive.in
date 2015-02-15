package admin

import (
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/ungerik/go-start/utils"
	. "github.com/ungerik/go-start/view"
)

func init() {
	Admin_ExportMentorJudgeEmails = NewViewURLWrapper(
		RenderView(
			func(ctx *Context) (err error) {
				ctx.Response.ContentDispositionAttachment("mentor-judge-emails.csv")

				ctx.Response.Printf(`"Name","Email","Phone","Web","Company","Position","Citizenship","Mentor","Judge"` + "\n")

				var person models.Person
				i := models.People.Filter("Mentor", true).Iterator()
				for i.Next(&person) {
					ctx.Response.Printf(
						`"%s","%s","%s","%s","%s","%s","%s","%t","%t"`+"\n",
						person.Name.String(),
						person.PrimaryEmail(),
						person.PrimaryPhone(),
						person.PrimaryWeb(),
						utils.StripHTMLTags(person.Company.String()),
						person.Position,
						person.Citizenship,
						person.Mentor,
						person.Judge,
					)
				}
				if i.Err() != nil {
					return i.Err()
				}
				i = models.People.Filter("Judge", true).Iterator()
				for i.Next(&person) {
					ctx.Response.Printf(
						`"%s","%s","%s","%s","%s","%s","%s","%t","%t"`+"\n",
						person.Name.String(),
						person.PrimaryEmail(),
						person.PrimaryPhone(),
						person.PrimaryWeb(),
						utils.StripHTMLTags(person.Company.String()),
						person.Position,
						person.Citizenship,
						person.Mentor,
						person.Judge,
					)
				}
				return i.Err()
			},
		),
	)
}
