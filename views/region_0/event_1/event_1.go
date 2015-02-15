package event_1

import (
	// "github.com/ungerik/go-start/debug"
	. "github.com/ungerik/go-start/view"

	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
)

func init() {
	Region0_Event1 = newPublicEventPage(EventTitle("About"), nil, DynamicView(
		func(ctx *Context) (view View, err error) {
			event := ctx.Data.(*PageData).Event
			//tagline := event.Tagline.GetOrDefault("[Enter event tagline!]")
			lead := event.DescriptionLead.GetOrDefault("[Enter event description lead!]")
			description := HTML(event.Description.GetOrDefault("[Enter event description!]"))
			howItsDone := HTML(event.HowItsDone.GetOrDefault("&nbsp;"))
			prizes := HTML(event.Prizes.GetOrDefault("&nbsp;"))
			imageURL := event.ImageURL_604x0.GetOrDefault("/images/event-default-604x233.jpg")

			var eventContactPerson *models.Person
			if len(event.Organisers) > 0 {
				err := event.Organisers[0].Get(&eventContactPerson)
				if err != nil {
					return nil, err
				}
			}

			if event.Type == models.StartupLive {
				view = &Div{
					Class: "main",
					Content: Views{
						TitleBar("About the event"),
						&If{
							Condition: event.Name == "Startup Live Jonkoping #1",
							Content: &Views{
								BR(),
								HTML("<iframe width='604' height='340' src='http://www.youtube.com/embed/MYox5PP1deg' frameborder='0' allowfullscreen></iframe>"),
							},
							ElseContent: &Image{Class: "image-box", Width: 604, Src: imageURL},
						},
						//&Image{Class: "image-box", Width: 604, Src: imageURL},
						//H2(tagline),
						&Paragraph{Class: "lead", Content: Escape(lead)},
						P(description),
						H3("How it's done"),
						IMG("/images/how-its-done.png"),
						&Div{
							Class: "how-its-done",
							Content: Views{
								&Div{Class: "pitches", Content: HTML("PITCHES")},
								&Div{Content: HTML("&rarr;")},
								&Div{Class: "teambuilding", Content: HTML("TEAMBUILDING")},
								&Div{Content: HTML("&rarr;")},
								&Div{Class: "work", Content: HTML("WORK")},
								&Div{Content: HTML("&rarr;")},
								&Div{Class: "feedback", Content: HTML("GET FEEDBACK")},
								&Div{Content: HTML("&rarr;")},
								&Div{Class: "present", Content: HTML("PRESENT")},
							},
						},
						P(howItsDone),
						H3("Prizes"),
						P(prizes),
						EventContact(eventContactPerson),
					},
				}
			} else if event.Type == models.StartupLounge {
				view = &Div{
					Class: "main",
					Content: Views{
						TitleBar("About the event"),
						&Image{Class: "image-box", Width: 604, Src: imageURL},
						//H2(tagline),
						&Paragraph{Class: "lead", Content: Escape(lead)},
						P(description),
						EventContact(eventContactPerson),
					},
				}
			} else if event.Type == models.LiveAcademy {
				view = &Div{
					Class: "main",
					Content: Views{
						TitleBar("About the Academy"),
						&Image{Class: "image-box", Width: 604, Src: imageURL},
						//H2(tagline),
						&Paragraph{Class: "lead", Content: Escape(lead)},
						P(description),
						H3("How it's done"),
						P(howItsDone),
						H3("Prizes"),
						P(prizes),
						EventContact(eventContactPerson),
					},
				}
			
			}
			return view, nil
		},
	))
}

func newPublicEventPage(title, scripts Renderer, main View) *Page {
	logo := DynamicView(
		func(ctx *Context) (View, error) {
			region := ctx.Data.(*PageData).Region
			return Views{
				&Image{Class: "logo", Src: region.PublicHeaderLogoURL.Get()},
				&Image{Class: "logo-hover", Src: region.HoverLogoURL.Get()},
			}, nil
		},
	)

	menu := headerEventMenu()

	content := Views{
		eventInfoView(),
		DIV("event-content",
			eventSidebar(),
			main,
			DivClearBoth(),
			localEventSponsors(),
		),
	}

	return &Page{
		OnPreRender: func(page *Page, ctx *Context) (err error) {
			err = SetEventPageData(page, ctx)
			if err != nil {
				return err
			}
			event := ctx.Data.(*PageData).Event
			if roundupURL := event.RoundupURL.Get(); roundupURL != "" {
				return PermanentRedirect(roundupURL)
			}
			if event.Status != models.EventPublic {
				admin, err := Region0_Event1_Admin_Auth.Authenticate(ctx)
				if err != nil {
					return err
				}
				if !admin {
					return Forbidden("403 Forbidden: event not public")
				}
			}
			return nil
		},
		Title: title,
		CSS:   IndirectURL(&Region0_CSS),
		HeadScripts: Renderers{
			renderGoogleAnalyticsForHosts,
		},
		Scripts: Renderers{
			IndirectRenderer(&Config.Page.DefaultScripts),
			scripts,
		},
		Content: PublicPageStructure("menu-area event-menu", logo, menu, content, AdditionalTopNav),
	}
}

func eventInfoView() View {
	return DynamicView(
		func(ctx *Context) (view View, err error) {
			data := ctx.Data.(*PageData)
			event := data.Event
			location := data.Location

			from := event.Start.Format("02. January - ")
			until := event.End.Format("02. January 2006")

			registrationURL := "#"
			if event.Show.Registration.Get() {
				registrationURL = Region0_Event1_Registration.URL(ctx)
			}
			registrationButtonText := event.RegistrationButton.GetOrDefault("Register Now")

			var locationView View
			if event.Show.Location.Get() {
				locationURL := Region0_Event1_Location.URL(ctx)
				locationView = A(locationURL, location.Name.Get())
			} else {
				locationView = Escape(location.Name.Get())
			}

			twitterURL := event.TwitterURL.GetOrDefault(STARTeuropeTwitterURL)
			facebookURL := event.FacebookURL.GetOrDefault(STARTeuropeFacebookURL)
			linkedInURL := event.LinkedInURL.GetOrDefault(STARTeuropeLinkedInURL)
			flickrURL := event.FlickrURL.GetOrDefault(STARTeuropeFlickrURL)
			spotieURL := event.SpotieURL.Get()

			var stampImage View
			if !event.StampImageURL.IsEmpty() {
				stampImage = &Image{Class: "stamp", Src: event.StampImageURL.Get(), Height: 72}
			}

			view = &Div{
				Class: "event-info",
				Content: Views{
					Printf("<span class='event-location'><i>startuplive.in/</i> <b>%s</b></span>", data.Region.Name),
					stampImage,
					&Table{
						HeaderRow: true,
						Model: ViewsTableModel{
							{HTML("Topic"), Escape(event.GetTopic())},
							{HTML("When:"), HTML(from + until)},
							{HTML("Where:"), locationView},
							{HTML("Language:"), Escape(event.Language.String())}, //data.Event.Language.EnglishName()},
						},
					},
					&If{
						Condition: ctx.Request.URLString() != registrationURL,
						Content: &Div{
							Class: "register-now",
							Content: Views{
								&Link{Class: "button", Model: NewLinkModel(registrationURL, registrationButtonText)},
								Escape(event.RegistrationTagline.Get()),
							},
						},
					},
					&Div{
						Class: "socialmedia-bar",
						Content: Views{
							A_blank(
								twitterURL,
								&Image{Class: "hover-visible", Src: "/images/icons/twitter-circle.png"},
								&Image{Class: "hover-invisible", Src: "/images/icons/twitter-circle-gray.png"},
							),
							A_blank(
								facebookURL,
								&Image{Class: "hover-visible", Src: "/images/icons/facebook-circle.png"},
								&Image{Class: "hover-invisible", Src: "/images/icons/facebook-circle-gray.png"},
							),
							A_blank(
								linkedInURL,
								&Image{Class: "hover-visible", Src: "/images/icons/linkedin-circle.png"},
								&Image{Class: "hover-invisible", Src: "/images/icons/linkedin-circle-gray.png"},
							),
							A_blank(
								flickrURL,
								&Image{Class: "hover-visible", Src: "/images/icons/flickr-circle.png"},
								&Image{Class: "hover-invisible", Src: "/images/icons/flickr-circle-gray.png"},
							),
							&If{
								Condition: spotieURL != "",
								Content: A_blank(
									spotieURL,
									&Image{Class: "hover-visible", Src: "/images/icons/spotie-circle.png"},
									&Image{Class: "hover-invisible", Src: "/images/icons/spotie-circle-gray.png"},
								),
							},
						},
					},
				},
			}
			return view, nil
		},
	)
}

func featuredBox(title string, content ...View) *Div {
	return DIV("featured-box",
		DIV("box-title", title),
		Views(content),
		DivClearBoth(),
	)
}

func featuredBoxSmall(title string, content ...View) *Div {
	return DIV("featured-box small",
		DIV("box-title", title),
		Views(content),
		DivClearBoth(),
	)
}

var renderGoogleAnalyticsForHosts = Render(
	func(ctx *Context) (err error) {
		data := ctx.Data.(*PageData)
		event := data.Event
		// if !Config.IsProductionServer {
		// 	return nil
		// }

		if event.GoogleAnalyticsHostAccount.Get() != "" {
			analyticsID := event.GoogleAnalyticsHostAccount.String()
			ctx.Response.Write([]byte(`<script type="text/javascript">

  var _gaq = _gaq || [];
  _gaq.push(['_setAccount', '` + analyticsID + `']);
  _gaq.push(['_trackPageview']);

  (function() {
    var ga = document.createElement('script'); ga.type = 'text/javascript'; ga.async = true;
    ga.src = ('https:' == document.location.protocol ? 'https://ssl' : 'http://www') + '.google-analytics.com/ga.js';
    var s = document.getElementsByTagName('script')[0]; s.parentNode.insertBefore(ga, s);
  })();

</script>`))

		}

		return nil
	},
)
