package root

import (
	"errors"
	"strings"

	"github.com/ungerik/go-rss"
	"github.com/ungerik/go-start/utils"
	. "github.com/ungerik/go-start/view"

	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
)

func init() {
	Homepage = CacheView(PublicPageCacheDuration, &Page{
		Title:          HTML("Startup Live"),
		AdditionalHead: HTML("<link rel='canonical' href='http://startuplive.in/'>"),
		Content: Views{
			&Div{
				Class: "header",
				Content: Views{
					HeaderTopNav(nil),
					&Div{
						Class: "green-area",
						Content: &Div{
							Class: "highlight",
							Content: &Div{
								Class: "center",
								Content: Views{
									&Div{
										Class: "logo-container",
										Content: Views{
											IMG("/images/logo-green.png", 146, 162),
											&Image{Class: "logo-hover", Src: "/images/logo-hover.png"},
											MegaDropdown(),
										},
									},
									// &Image{
									// 	Class: "logo-starteurope",
									// 	Src:   "/images/starteuropeatsul.png",
									// },
									&Image{
										Class: "select-your-city",
										Src:   "/images/select-your-city.png",
									},
									HomepageHeaderMenu(),
									&Div{
										Class:   "description",
										Content: HTML("Where fantastic ideas are launched, creative teams are formed and valuable connections made."),
									},
									&Div{
										Class: "info-bar",
										Content: Views{
											HTML("<div class='pitches'>PITCHES</div><div class='team'>TEAMBUILDING</div><div class='work'>WORK</div><div class='feedback'>GET FEEDBACK</div><div class='present'>PRESENT</div>"),
											DivClearBoth(),
										},
									},
									&Image{
										Class: "our-next-events",
										Src:   "/images/our-next-events.png",
									},
								},
							},
						},
					},
				},
			},
			homepageUpcomingBoxes(),
			&Div{
				Class: "content",
				Content: Views{
					&Div{
						Class: "center",
						Content: Views{
							&Div{
								Class: "to-your-city",
								Content: Views{
									H3("Your city is missing? Contact us!"),
									// todo debug, why endless loop:
									//A(&Events_YourCity, HTML("Get <b>Startup Live</b> to <b>your</b> city")),
									&Link{Model: &PageLink{Page: &Events_YourCity, Content: HTML("Get <b>Startup Live</b> to <b>your</b> city")}},
								},
								},
								&Div{
									Class: "lower-content",
									Content: Views{
										homepageBoxFromTheBlog(),
										homepageBoxOurPartners(),
										homepageBoxMainEvent(),

										//homepageBoxFeaturedMentor(),
										//DivClearBoth(),
										//homepageBoxStartupOfTheMonth(),
										DivClearBoth(),
									},
								},
								&Div{
									Class: "lower-content",
									Content: Views{
										//homepageBoxStrategicPartner(),
										//homepageBoxSubsidisedBy(),
										DivClearBoth(),
									},
								},
							},
						},
					},
				},
				Footer(),
			},
		})
	}

	func homepageUpcomingBoxes() View {
	return &Div{
		Class: "upcoming",
		Content: &Div{
			Class: "center",
			Content: Views{
				DynamicView(
					func(ctx *Context) (View, error) {
						var views Views
						iter := models.UpcomingPublicStartupLiveEventIterator()
						for i := 0; i < 3; i++ {
							var content View = HTML(" ")
							var event models.Event
							if iter.Next(&event) {
								var region models.EventRegion
								found, err := event.Region.TryGet(&region)
								if err != nil {
									return nil, err
								}
								if found {
									url := Region0_Event1.URL(ctx.ForURLArgs(region.Slug.Get(), event.Number.String()))
									from := event.Start.Format("02.01. - ")
									until := event.End.Format("02.01. 2006")
									content = A(url, Views{
										IMG(region.InitialURL.Get(), 0, 100),
										H3(strings.ToUpper(region.Name.Get())),
										HTML(from + until),
									})
								} else {
									return nil, errors.New("Upcoming Event not found")
								}
							}
							views = append(views, DIV("upcoming-box-container", DIV("upcoming-box", content)))
						}
						if iter.Err() != nil {
							return nil, iter.Err()
						}
						return views, iter.Err()
					},
				),
				DIV("upcoming-box-container",
					DIV("upcoming-box main-event-box",
						A_blank("http://startuplive.in/academy/1",
							&Image{
								Src:    "/images/logo-academy-200x150.png",
								Width:  200,
								Height: 150,
								Title:  "Startup Live Academy",
							},
						),
					),
				),
				&Image{Class: "main-event", Src: "/images/accelerator.png", Width: 120, Height: 60},
				DivClearBoth(),
			},
		},
	}
}

func homepageBoxFromTheBlog() View {
	return &Div{
		Class: "dashed-box dashed-box-left",
		Content: Views{
			&Div{
				Class:   "box-title",
				Content: HTML("From the blog"),
			},
			&Div{
				Class: "box-content",
				Content: Views{
					DynamicView(
						func(ctx *Context) (View, error) {
							blogURL := Blog.URL(ctx)
							feed, err := rss.Read(StartupLiveBlogFeedURL)
							if err != nil {
								return nil, err
							}
							var class string
							var views Views
							for i := 0; i < HomepageNumBlogPosts && i < len(feed.Item); i++ {
								item := &feed.Item[i]
								if i%2 == 0 {
									class = "item even"
								} else {
									class = "item odd"
								}
								var image View
								if item.Enclosure.URL != "" {
									image = &Image{Width: 100, Src: item.Enclosure.URL}
								}
								itemURL := strings.Replace(item.Link, HiddenStartupLiveBlogURL, blogURL, 1)
								views = append(views, &Div{
									Class: class,
									Content: Views{
										&Link{Model: NewLinkModel(itemURL, Views{image, H4(item.Title)})},
										&TextPreview{
											PlainText:   utils.StripHTMLTags(item.Description),
											MaxLength:   150,
											ShortLength: 150 - 15,
											MoreLink:    NewLinkModel(itemURL, HTML("read more &rarr;")),
										},
									},
								})
							}
							return views, nil
						},
					),
					HTML("<a class='rss' href='" + StartupLiveBlogFeedURL + "' target='_blank' rel='me'><img src='/images/icons/rss16.png' alt='RSS' width='16' height='16' /> Get the RSS Feed</a>"),
					&Link{Class: "older", Model: NewLinkModel(Blog, HTML("older entries &rarr;"))},
				},
			},
		},
	}
}

func homepageBoxOurPartners() View {
	return DIV("dashed-box dashed-box-right our-partners",
		DIV("box-title", HTML("Our global partner")),
		DIV("box-content",
			A_blank("https://www.conda.at", IMG("http://i.imgur.com/oA5Fzsb.png")),
		),
	)
	// return Views{
	// 	A_blank("https://www.mingo.at/", IMG("http://dl.dropbox.com/u/5565424/MINGO.jpg", 210)),
	// 	A_blank("http://www.accent.at/", IMG("/images/sponsors/accent-210x46.png", 210, 46)),
	// 	A_blank("http://www.tecnet.co.at/", IMG("/images/sponsors/tecnet-210x104.png", 210, 104)),
	// 	A_blank("http://www.sektor5.at/", IMG("/images/sponsors/sektor5-210x80.png", 210, 80)),
	// }
}

func homepageBoxMainEvent() View {
	return &Div{
		Class: "dashed-box dashed-box-right",
		Content: Views{
			&Div{
				Class:   "box-title red-box-title",
				Content: HTML("Supporter"),
			},
			&Div{
				Class: "box-content pioneers",
				Content: &Link{
					NewWindow: true,
					Model: NewLinkModel("http://pioneersfestival.com",
						&Image{
							Src:    "/images/logo-pioneers-360x90.png",
							Width:  360,
							Height: 90,
							Title:  "Pioneers Festival",
						},
					),
				},
			},
		},
	}
}

func homepageBoxFeaturedMentor() View {
	return &Div{
		Class: "dashed-box dashed-box-right",
		Content: Views{
			&Div{
				Class:   "box-title",
				Content: HTML("Mentor: Morten Lund"),
			},
			&Div{
				Class: "box-content mentor",
				Content: Views{
					HTML("<a href='http://www.startupweek2011.com/speaker/morten-lund/'><img src='/images/Morten-Lund1.jpeg' width='100' height='100' /></a>"),
					&Div{
						Class: "text",
						Content: Views{
							H4("Morten Lund (DEN)"),
							P("Morten Lund, is the CEO of Everbread Ltd. a travel-search company, and Chairman of Tradeshift Ltd. the worlds first open business network."),
							P("Described variously as an ‘archangel investor’, startup ‘ideologist’ and ‘visionary’, Morten has been one of the most active seed investors in Europe, through his startup catalyst LundXY."),
							//HTML("<a href='http://www.startupweek2011.com/speaker/morten-lund/'>read more &rarr;</a>"),
						},
					},
				},
			},
		},
	}
}

func homepageBoxStartupOfTheMonth() View {
	return &Div{
		Class: "dashed-box",
		Content: Views{
			&Div{
				Class:   "box-title",
				Content: HTML("Startup of the month"),
			},
			&Div{
				Class: "box-content",
				Content: Views{
					&Div{
						Class: "startup-of-month",
						Content: Views{
							H4("Rails on Fire"),
							P("We provide Continuous Integration and Continuous Deployment for Ruby code hosted on GitHub. Follow a modern day development method with regular testing and deployment in the cloud. In less than two minutes you can go from your first login to testing and deploying to Heroku."),
							HTML("<a href='http://railsonfire.com' target='_blank'>www.railsonfire.com</a>"),
						},
					},
					HTML("<a href='http://railsonfire.com' target='_blank'><img class='startup-logo' src='/images/railsonfire-dummy.png' /></a>"),
					DivClearBoth(),
				},
			},
		},
	}
}

func homepageBoxStrategicPartner() View {
	return DIV("dashed-box dashed-box-strategic-partner our-partners",
		DIV("box-title", HTML("Strategic partner")),
		DIV("box-content",
			A_blank("http://www.bmwfj.gv.at/", IMG("http://pioneersfestival.com/wp-content/uploads/2012/07/Logo-BMWFJ-CMYK-01-e1342697673507.png")),
		),
	)
}

func homepageBoxSubsidisedBy() View {
	return DIV("dashed-box dashed-box-subsidised our-partners",
		DIV("box-title", HTML("Subsidised by")),
		DIV("box-content",
			A_blank("http://www.impulse-awsg.at/", IMG("http://pioneersfestival.com/wp-content/uploads/2012/07/impulseaws_evolvebmwfj_ohnelogo_4c-01-e1342697653827.png")),
		),
	)
}
