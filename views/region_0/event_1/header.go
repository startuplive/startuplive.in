package event_1

import (
	. "github.com/STARTeurope/startuplive.in/views"
	. "github.com/ungerik/go-start/view"
)

func EventHeaderLogoAndTitle() View {
	return &Tag{
		Tag: "h1",
		Content: DynamicView(func(ctx *Context) (View, error) {
			data := ctx.Data.(*PageData)
			from := data.Event.Start.Format("2. - ")
			until := data.Event.End.Format("2. January 2006")
			return Views{
				&Image{Class: "logo", Src: data.Region.HeaderLogoURL.Get()},
				HTML(from),
				HTML(until),
			}, nil
		}),
	}
}

var AdditionalTopNav = DynamicView(
	func(ctx *Context) (View, error) {
		var views Views
		if ctx.Data.(*PageData).Event.IsHappeningNow() {
			views = append(
				views,
				HTML("&nbsp; | &nbsp;"),
				&Link{Model: &PageLink{Page: &Region0_Event1_Dashboard, Title: "Dashboard"}},
			)
		}
		if ok, _ := Region0_Event1_Admin_Auth.Authenticate(ctx); ok {
			views = append(
				views,
				HTML("&nbsp; | &nbsp;"),
				&Link{Model: &PageLink{Page: &Region0_Event1_Admin, Title: "Admin"}},
			)
		}
		return views, nil
	},
)

func headerEventMenu() View {
	return DynamicView(
		func(ctx *Context) (view View, err error) {
			pageData := ctx.Data.(*PageData)
			event := pageData.Event

			menuItems := []LinkModel{&PageLink{Page: &Region0_Event1, Title: "About"}}
			if event.Show.Location.Get() {
				menuItems = append(menuItems, &PageLink{Page: &Region0_Event1_Location, Title: "Location"})
			}
			if event.Show.Schedule.Get() {
				menuItems = append(menuItems, &PageLink{Page: &Region0_Event1_Schedule, Title: "Schedule"})
			}
			if event.Show.MentorsJudges.Get() {
				title := event.MentorsJudgesTab_Title.GetOrDefault("Mentors/Judges")
				menuItems = append(menuItems, &PageLink{Page: &Region0_Event1_Judges, Title: title})
			}
			if event.Show.Organisers.Get() {
				menuItems = append(menuItems, &PageLink{Page: &Region0_Event1_Organisers, Title: "Organisers"})
			}
			if event.Show.FAQ.Get() {
				menuItems = append(menuItems, &PageLink{Page: &Region0_Event1_FAQ, Title: "FAQ"})
			}
			if event.Show.Registration.Get() {
				menuItems = append(menuItems, &PageLink{Page: &Region0_Event1_Registration, Title: "Register Now"})
			}
			if event.Show.ExtraTab.Get() {
				menuItems = append(menuItems, &PageLink{Page: &Region0_Event1_Extra, Title: event.ExtraTab_Title.Get()})
			}

			view = Views{
				&Menu{
					Class:           "menu",
					ItemClass:       "menu-item",
					ActiveItemClass: "active",
					Items:           menuItems,
				},
			}
			return view, nil
		},
	)
}
