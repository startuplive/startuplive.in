package events

import (
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	. "github.com/ungerik/go-start/view"
	"time"
)

func init() {
	Events_Where = NewPublicPage("Where? | Startup Live", DynamicView(
		func(ctx *Context) (view View, err error) {
			var years Views

			thisYear := time.Now().UTC().Year()
			for year := thisYear; year >= 2010; year-- {
				logos, err := LogoEventView(ctx, models.YearPublicStartupLiveEventIterator(year))
				if err != nil {
					return nil, err
				}
				if len(logos) > 0 {
					years = append(years, &Div{
						Class: "events",
						Content: Views{
							H2(Printf("%d", year)),
							logos,
							DivClearBoth(),
						},
					})
				}
			}

			view = DIV("public-content",
				DIV("main where",
					TitleBar("All the places we've been"),
					DIV("main-content",
						P(HTML("Our goal is to create one European startup community.<br/>Therefore we organise Startup Live events all over Europe.")),
						years,
					),
					&Div{
						Class: "to-your-city",
						Content: Views{
							H3("Your city is missing? Contact us!"),
							&Link{Model: NewLinkModel(&Events_YourCity, HTML("Get <b>Startup Live</b> to <b>your</b> city"))},
						},
					},
				),
			)
			return view, nil
		},
	))
}
