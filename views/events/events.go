package events

import (

	//	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	. "github.com/ungerik/go-start/view"
)

func init() {
	Events = NewPublicPage("What? | Startup Live", DynamicView(
		func(ctx *Context) (view View, err error) {
			view = DIV("public-content",
				DIV("main what",
					TitleBar("What?"),
					DIV("main-content",
						H2(A(Events_YourCity, IMG("/images/logo-startup-live-your-city-190x90.png"))),
						P("Startup Live is a three-day startup event that provides networking, resources and incentives for individuals and teams to go from idea to launch in just one weekend. The event brings together the local startup community with other European entrepreneurs and enables the participants to connect with experienced entrepreneurs, experts and investors. As such It doesn’t serve as a singular happening but is the beginning of a great journey to build up a local startup ecosystem that helps and supports existing and future entrepreneurs."),
						P("We organize Startup Live events all over Europe and nourish regional communities as well as connecting entrepreneurs all over Europe. Those who stand out from this community - the best startups, organizers, community members - are invited to the Pioneers Festival and take part in a special track where we connect them as a group, but also get them in touch with the top-notch speakers, investors and partners. This way, we create a highly interconnected community that supports the individual members and ensure entrepreneurial success in Europe."),
						DIV("to-your-city", A(Events_YourCity, HTML("Get <b>Startup Live</b> to <b>your</b> city"))),

						H2(A_blank("http://pioneersfestival.com", IMG("/images/logo-pioneers-200x50.png"))),
						P("From 30 – 31 October ", A_blank("http://pioneersfestival.com", "Pioneers Festival"), " gathers the tech scene – Startups, Founders, Hackers, Bloggers, VCs and Innovators, in fact the Troublemakers – from all over to inspire and be inspired. Set in the heart of Vienna, at Hofburg Imperial Palace, Pioneers Festival educates, entertains and inspires a new generation of Pioneers. Furthermore the Top 50 startups of the Pioneers Challenge - a pitching competition for tech startups in the fields of web/mobile, health, aerospace, AI, robotics, energy and hardware – will be invited to pitch their ventures on 29 October, the exclusive Investorsday. What are you waiting for? Stop dreaming. Be a Pioneer."),
						P("The goal of this conference is to connect the European startup community and together pushing it to the next level."),

					),
				),
			)
			return view, nil
		},
	))
}
