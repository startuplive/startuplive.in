package event_1

import (
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/ungerik/go-start/view"
	//	"github.com/ungerik/go-start/debug"
	. "github.com/STARTeurope/startuplive.in/views"
)

func init() {
	Region0_Event1_Organisers = newPublicEventPage(EventTitle("Organisers"), nil, DynamicView(
		func(ctx *Context) (view View, err error) {
			data := ctx.Data.(*PageData)

			var teamMembers Views
			var rightCol *Div

			i := data.Event.OrganiserIterator()
			var person models.Person
			for i.Next(&person) {
				view := EventOrganiserView(&person)
				if rightCol == nil {
					rightCol = &Div{Class: "right-col"}
					row := &Div{
						Content: Views{
							&Div{Class: "left-col", Content: view},
							rightCol,
							DivClearBoth(),
						},
					}
					teamMembers = append(teamMembers, row)
				} else {
					rightCol.Content = view
					rightCol = nil
				}
			}
			if i.Err() != nil {
				return nil, i.Err()
			}
			teamImageURL := data.Event.TeamImageURL_604x0.GetOrDefault("/images/event-team-default-604x233.jpg")
			view = &Div{
				Class: "main",
				Content: Views{
					TitleBar("Our team in " + data.Region.Name.Get()),
					&Image{Class: "image-box", Width: 604, Src: teamImageURL},
					H2("Meet the team"),
					teamMembers,
				},
			}
			return view, nil
		},
	))
}
