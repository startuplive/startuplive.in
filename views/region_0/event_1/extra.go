package event_1

import (
	. "github.com/STARTeurope/startuplive.in/views"
	. "github.com/ungerik/go-start/view"
)

func init() {
	Region0_Event1_Extra = newPublicEventPage(
		Render(
			func(ctx *Context) (err error) {
				event := ctx.Data.(*PageData).Event
				return EventTitle(event.ExtraTab_Title.GetOrDefault("Info")).Render(ctx)
			},
		),
		nil,
		DynamicView(
			func(ctx *Context) (view View, err error) {
				event := ctx.Data.(*PageData).Event
				view = &Div{
					Class: "main",
					Content: Views{
						//TitleBar(event.ExtraTab_Title.Get()),
						HTML(event.ExtraTab_HTML.Get()),
					},
				}
				return view, nil
			},
		),
	)
}
