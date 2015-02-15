package event_1

import (

	//	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	. "github.com/ungerik/go-start/view"
)

func init() {
	Region0_Event1_FAQ = newPublicEventPage(EventTitle("FAQ"), nil, DynamicView(
		func(ctx *Context) (view View, err error) {
			view = &Div{
				Class: "main",
				Content: Views{
					TitleBar("Fequently Asked Questions"),
					&Div{
						Class:   "faq",
						Content: HTML(ctx.Data.(*PageData).Event.FAQ_HTML.Get()),
					},
				},
			}
			return view, nil
		},
	))
}
