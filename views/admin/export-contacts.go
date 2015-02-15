package admin

import (
	"bytes"
	"encoding/csv"

	// "github.com/ungerik/go-start/debug"
	// "github.com/ungerik/go-start/model"
	. "github.com/ungerik/go-start/view"

	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
)

func StringCSV(records [][]string) string {
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)
	err := writer.WriteAll(records)
	if err != nil {
		return ""
	}
	return buf.String()
}

func init() {
	Admin_ExportContacts = NewViewURLWrapper(
		RenderView(
			func(ctx *Context) (err error) {
				ctx.Response.ContentDispositionAttachment("live-contacts.csv")

				records := [][]string{
					{"Name", "Email", "Tel", "Web", "Facebook", "Twitter", "LinkedIn", "Xing"},
				}
				i := models.People.Iterator()
				var person models.Person
				for i.Next(&person) {
					records = append(records, []string{
						person.Name.String(),
						person.PrimaryEmail(),
						person.PrimaryPhone(),
						person.PrimaryWeb(),
						person.PrimaryFacebookIdentity().ProfileURL(),
						person.PrimaryTwitterIdentity().ProfileURL(),
						person.PrimaryLinkedInIdentity().ProfileURL(),
						person.PrimaryXingIdentity().ProfileURL(),
					})
				}
				if i.Err() != nil {
					return i.Err()
				}
				_, err = ctx.Response.WriteString(StringCSV(records))
				return err
			},
		),
	)
}
