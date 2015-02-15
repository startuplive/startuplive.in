package event_1

import (
	. "github.com/STARTeurope/startuplive.in/views"
	. "github.com/ungerik/go-start/view"
)

func init() {
	Region0_Event1_Registration = newPublicEventPage(EventTitle("Registration"), nil, DynamicView(
		func(ctx *Context) (view View, err error) {
			event := ctx.Data.(*PageData).Event
			//			tagline := event.Tagline.GetOrDefault("[Insert Event Tagline!]")
			view = &Div{
				Class: "main registration",
				Content: Views{
					TitleBar("Event registration"),
					HTML(event.AmiandoIframeCode),
				},
			}
			return view, nil
		},
	))
}
