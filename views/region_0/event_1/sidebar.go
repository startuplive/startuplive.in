package event_1

import (
	. "github.com/STARTeurope/startuplive.in/views"
	. "github.com/ungerik/go-start/view"
)

func eventSidebar() View {
	return &Div{
		Class: "sidebar",
		Content: Views{
			DynamicView(
				func(ctx *Context) (view View, err error) {
					event := ctx.Data.(*PageData).Event
					logoURL := event.HostLogoURL_200x0.Get()
					if logoURL != "" {
						var logo View = IMG(logoURL, 200)
						logoLinkURL := event.HostLogoLinkURL.Get()
						if logoLinkURL != "" {
							logo = A_blank(logoLinkURL, logo)
						}
						var logos View = logo

						logoURL = event.SecondHostLogoURL_200x0.Get()
						if logoURL != "" {
							var logo View = IMG(logoURL, 200)
							logoLinkURL := event.SecondHostLogoLinkURL.Get()
							if logoLinkURL != "" {
								logo = A_blank(logoLinkURL, logo)
							}
							logos = Views{logos, logo}
						}

						view = &Div{
							Class: "featured-box",
							Content: Views{
								&Div{
									Class:   "box-title",
									Content: HTML("EVENT ORGANISED BY"),
								},
								logos,
							},
						}
					}
					return view, nil
				},
			),
			DynamicView(
				func(ctx *Context) (view View, err error) {
					event := ctx.Data.(*PageData).Event
					logoURL := event.PoweredByLogoURL_200x0.Get()
					if logoURL != "" {
						var logo View = IMG(logoURL, 200)
						logoLinkURL := event.PoweredByLogoLinkURL.Get()
						if logoLinkURL != "" {
							logo = A_blank(logoLinkURL, logo)
						}
						var logo2 View
						logoURL = event.SecondPoweredByLogoURL_200x0.Get()
						if logoURL != "" {
							logo2 = IMG(logoURL, 200)
							logoLinkURL = event.SecondPoweredByLogoLinkURL.Get()
							if logoLinkURL != "" {
								logo2 = A_blank(logoLinkURL, logo2)
							}
						}
						view = &Div{
							Class: "featured-box",
							Content: Views{
								&Div{
									Class:   "box-title",
									Content: HTML("EVENT POWERED BY"),
								},
								logo,
								logo2,
							},
						}
					}
					return view, nil
				},
			),
			&Div{
				Class: "featured-box",
				Content: Views{
					&Div{
						Class:   "box-title",
						Content: HTML("GLOBAL PARTNER"),
					},
					DynamicView(
						func(ctx *Context) (view View, err error) {
							event := ctx.Data.(*PageData).Event
							return A_blank(event.GlobalPartnerLinkURL.GetOrDefault("https://www.conda.at"), IMG("http://i.imgur.com/oA5Fzsb.png", 200)), nil
						},
					),
				},
			},
		},
	}
}
