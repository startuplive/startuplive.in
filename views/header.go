package views

import (
	"github.com/ungerik/go-start/model"
	"github.com/ungerik/go-start/user"
	. "github.com/ungerik/go-start/view"

	"github.com/STARTeurope/startuplive.in/models"
)

func HeaderTopNav(additional View) View {
	return &Div{
		Class: "top-nav",
		Content: &Div{
			Class: "center",
			Content: Views{
				HeaderUserNav(additional),
				DynamicView(
					func(ctx *Context) (View, error) {
						if ctx.Request.RequestURI == "/" {
							return nil, nil
						}
						return A("/", HTML("&larr; Back to startuplive.in homepage")), nil
					},
				),
			},
		},
	}
}

func HeaderMenu() *Menu {
	return &Menu{
		Class:           "menu",
		ItemClass:       "menu-item",
		ActiveItemClass: "active",
		Items: []LinkModel{
			NewLinkModel(&Homepage, "Home"),
			NewLinkModel(&Events, "What?"),
			NewLinkModel(&Events_Where, "Where?"),
			// //&PageLink{Page: &Investors},
			// //&PageLink{Page: &Mentors},
			// //&PageLink{Page: &Startups},
			NewLinkModel(&Organisers, "Organisers"),
			NewLinkModel(&Blog, "Blog"),
		},
	}
}

func HomepageHeaderMenu() *Menu {
	menu := HeaderMenu()
	//menu.Items = menu.Items[1:]
	return menu
}

func HeaderUserNav(additional View) View {
	// sessionUser := ""
	// doc := user.OfSession(response)
	// if doc != nil {
	// 	sessionUser = doc.(*models.Person).Name.String()
	// }

	return DIV("login-nav",
		user.Nav(
			A(&LoginSignup, "Login / Sign up"),
			nil,
			A(&Logout, "Logout"),
			A(&Profile, "My profile"),
			HTML("&nbsp; | &nbsp;"),
		),
		additional,
	)
}

func LogoEventView(ctx *Context, i model.Iterator) (views Views, err error) {
	var event models.Event
	for i.Next(&event) {
		var region models.EventRegion
		err = event.Region.Get(&region)
		if err != nil {
			return nil, err
		}
		eventURL := Region0_Event1.URL(ctx.ForURLArgs(region.Slug.Get(), event.Number.String()))
		//logoURL := region.InitialURL_60x0.GetOrDefault(region.InitialURL.Get())
		logoURL := region.InitialURL.Get()
		start := event.Start.Format("02/01 - ")
		end := event.End.Format("02/01 2006")
		view := &Div{
			Class: "logo-event",
			Content: A(eventURL, Views{
				IMG(logoURL, 0, 60),
				H3(HTML("Startup Live")),
				H3(region.Name.Get()),
				HTML(start + end),
			}),
		}
		views = append(views, view)
	}
	return views, i.Err()
}

func MegaDropdown() View {
	return DynamicView(
		func(ctx *Context) (view View, err error) {
			upcomingEvents, err := LogoEventView(ctx, models.UpcomingPublicStartupLiveEventIterator())
			if err != nil {
				return nil, err
			}
			currentEvents, err := LogoEventView(ctx, models.CurrentPublicStartupLiveEventIterator())
			if err != nil {
				return nil, err
			}
			var pastEvents Views
			i := models.PastPublicStartupLiveEventIterator()
			var event models.Event
			for i.Next(&event) {
				var region models.EventRegion
				err = event.Region.Get(&region)
				if err != nil {
					return nil, err
				}
				eventURL := Region0_Event1.URL(ctx.ForURLArgs(region.Slug.Get(), event.Number.String()))
				start := event.Start.Format("02/01 - ")
				end := event.End.Format("02/01 2006")
				view = &Div{
					Class: "past-event",
					Content: A(eventURL, Views{
						H3(region.Name.Get()),
						HTML(start + end),
					}),
				}
				pastEvents = append(pastEvents, view)
			}
			if i.Err() != nil {
				return nil, i.Err()
			}

			logos := &Div{
				Class: "upcoming-events",
				Content: Views{
					H2(Printf("%d UPCOMING EVENTS", len(upcomingEvents))),
					upcomingEvents,
					DivClearBoth(),
				},
			}

			if len(currentEvents) > 0 {
				logos = &Div{
					Class: "upcoming-and-current-events",
					Content: Views{
						logos,
						&Div{
							Class: "current-events",
							Content: Views{
								H2(Printf("%d HAPPENING RIGHT NOW", len(currentEvents))),
								currentEvents,
								DivClearBoth(),
							},
						},
						DivClearBoth(),
					},
				}
			}

			view = &Div{
				Class: "megadropdown",
				Content: Views{
					logos,
					&Div{
						Class: "past-events",
						Content: Views{
							H2("PAST EVENTS"),
							pastEvents,
							DivClearBoth(),
						},
					},
				},
			}
			return view, nil
		},
	)
}
