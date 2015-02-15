package admin

import (
	"github.com/STARTeurope/startuplive.in/models"
	// "bytes"
	// "encoding/json"
	. "github.com/STARTeurope/startuplive.in/views"
	// "github.com/ungerik/go-start/debug"
	// "github.com/ungerik/go-start/utils"
	// "errors"
	. "github.com/ungerik/go-start/view"
	"strconv"
	"strings"
)

func init() {
	API_PartnerOrder = NewViewURLWrapper(
		RenderView(
			func(ctx *Context) (err error) {
				urlArgs := ctx.URLArgs
				slug := urlArgs[0]
				number, err := strconv.ParseInt(urlArgs[1], 10, 64)
				if err != nil {
					err = NotFound("404: Event Number not parseable")
				}
				_, event, found, err := models.GetRegionAndEvent(models.StartupLive, slug, number)
				if err != nil {
					return err
				}
				if found {

					// cat := ctx.Request.Params["Category"]
					// order := ctx.Request.Params["Order"]

					// type Result struct {
					// 	Order []string
					// }

					// var result Result
					// err = json.Unmarshal([]byte(order), &result)
					// if err != nil {
					// 	return err
					// }
					// debug.Print(result)

					// debug.Print("partner order")

					params := ctx.Request.Params

					// debug.Print(params)

					cat := ""

					var newpartnersOrder []string

					for k, v := range params {
						if k == "Category" {
							cat = v
						} else if k == "Order" {
							newpartnersOrder = strings.Split(v, ",")
						}

					}
					// debug.Print("cat: ", cat)
					// debug.Print("newpartnersOrder: ", newpartnersOrder)
					// oldpartners, err := event.GetPartnersByCategory(cat)
					// if err != nil {
					// 	return err
					// }
					err = event.OrderEventPartnersInCategory(cat, newpartnersOrder)

				} else {
					err = NotFound("404: Event Number not parseable")
				}

				ctx.Response.SetContentTypeByExt(".json")
				ctx.Response.WriteByte('[')
				// first := true
				// var person models.Person
				// for i.Next(&person) {
				// 	name := person.Name.String()
				// 	// company := ""
				// 	// if doc.(*models.Person).Company.String() != "" {
				// 	// 	company = " @" + doc.(*models.Person).Company.String()
				// 	// }
				// 	// email := ""
				// 	// if len(doc.(*models.Person).User.Email) > 0 {
				// 	// 	email = " #" + doc.(*models.Person).User.Email[0].Address.String()
				// 	// }
				// 	if first {
				// 		first = false
				// 	} else {
				// 		ctx.Response.WriteByte(',')
				// 	}
				ctx.Response.WriteByte('"')
				if err != nil {
					ctx.Response.WriteString("success: false, ")
					ctx.Response.WriteString("error: " + err.Error())
				} else {
					ctx.Response.WriteString("success: true, ")
					ctx.Response.WriteString("error: null")
				}

				ctx.Response.WriteByte('"')
				// 	//writer.Printf(`{"id":"%s","label":"%s","value":"%s"}`, name, name, name)
				// }
				// // if first == true {
				// // 	ctx.Response.WriteByte('"')
				// // 	ctx.Response.WriteString("Nothing found.")
				// // 	ctx.Response.WriteByte('"')
				// // }
				ctx.Response.WriteByte(']')

				return nil
			},
		),
	)
}
