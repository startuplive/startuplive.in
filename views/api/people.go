package api

import (
	"strings"

	"github.com/ungerik/go-start/utils"
	. "github.com/ungerik/go-start/view"

	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
)

func init() {
	API_People = SearchPeopleJSONView(
		func(person *models.Person, searchTerm string) (ok bool) {
			name := strings.ToLower(person.Name.String())
			return person.Blocked == false && name != "admin" && (searchTerm == "" || strings.Contains(name, searchTerm))
		},
	)
}

type SearchPersonFilterFunc func(person *models.Person, searchTerm string) (ok bool)

func SearchPeopleJSONView(f SearchPersonFilterFunc) ViewWithURL {
	return RenderViewWithURL(
		func(ctx *Context) (err error) {
			searchTerm, _ := ctx.Request.Params["term"]
			searchTerm = strings.ToLower(searchTerm)
			i := models.People.FilterFunc(
				func(doc interface{}) (ok bool) {
					return f(doc.(*models.Person), searchTerm)
				},
			)
			ctx.Response.Header().Set("Content-Type", "application/json")
			ctx.Response.WriteByte('[')
			first := true
			var person models.Person
			for i.Next(&person) {
				name := person.Name.String()
				// company := ""
				// if doc.(*models.Person).Company.String() != "" {
				// 	company = " @" + doc.(*models.Person).Company.String()
				// }
				// email := ""
				// if len(doc.(*models.Person).User.Email) > 0 {
				// 	email = " #" + doc.(*models.Person).User.Email[0].Address.String()
				// }
				if first {
					first = false
				} else {
					ctx.Response.WriteByte(',')
				}
				ctx.Response.WriteByte('"')
				ctx.Response.WriteString(utils.EscapeJSON(name))
				// ctx.Response.WriteString(utils.EscapeJSON(email))
				// ctx.Response.WriteString(utils.EscapeJSON(company))
				ctx.Response.WriteByte('"')
				//writer.Printf(`{"id":"%s","label":"%s","value":"%s"}`, name, name, name)
			}
			// if first == true {
			// 	ctx.Response.WriteByte('"')
			// 	ctx.Response.WriteString("Nothing found.")
			// 	ctx.Response.WriteByte('"')
			// }
			ctx.Response.WriteByte(']')
			return i.Err()
		},
	)
}
