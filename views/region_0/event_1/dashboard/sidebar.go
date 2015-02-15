package dashboard

import (
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/ungerik/go-rss"
	"github.com/ungerik/go-start/model"
	. "github.com/ungerik/go-start/view"
	"strings"
)

func eventDashboardSidebar() View {
	return Views{
		DIV("sidebar",
			DIV("box-header",
				DIV("box-title", "Upcoming Events"),
				A(IndirectURL(&Events_Where), HTML("show all &raquo;")),
			),
			DIV("box-body",
				&ModelIteratorView{
					GetModel: func(ctx *Context) (interface{}, error) {
						return new(models.Event), nil
					},
					GetModelIterator: func(ctx *Context) model.Iterator {
						i := models.UpcomingPublicStartupLiveEventIterator()
						return model.NewLimitedIterator(i, 5)
					},
					GetModelView: eventSidebarView,
				},
			),
			DIV("box-header",
				DIV("box-title", "From the Blog"),
				A(IndirectURL(&Blog), HTML("show all &raquo;")),
			),
			DIV("box-body",
				DynamicView(
					func(ctx *Context) (View, error) {
						blogURL := Blog.URL(ctx)
						feed, err := rss.Read(StartupLiveBlogFeedURL)
						if err != nil {
							return nil, err
						}
						var views Views
						for i := 0; i < 5 && i < len(feed.Item); i++ {
							item := &feed.Item[i]
							itemURL := strings.Replace(item.Link, HiddenStartupLiveBlogURL, blogURL, 1)
							views = append(views, &Div{
								Class: "entry",
								Content: Views{
									H5(A_blank(itemURL, item.Title)),
									Escape(item.PubDate.MustFormat("January 2, 2006")),
								},
							})
						}
						return views, nil
					},
				),
			),
		),
		DivClearBoth(),
	}
}

//func eventSidebarEntry(url, title, description string) View {
//	return &Div{
//		Class: "entry",
//		Content: Views{
//			H5(&Link{Model: NewLinkModel(url, title), NewWindow: true}),
//			Escape(description),
//		},
//	}
//}

func eventSidebarView(ctx *Context, model interface{}) (view View, err error) {
	event := model.(*models.Event)

	var region models.EventRegion
	err = event.Region.Get(&region)
	if err != nil {
		return nil, err
	}

	var location models.EventLocation
	err = event.Location.Get(&location)
	if err != nil {
		return nil, err
	}

	day := event.Start.Format("2")
	month := event.Start.Format("Jan")

	regionURL := Region0.URL(ctx.ForURLArgs(region.Slug.String()))

	locationName := location.Name.Get()
	if !location.ShortName.IsEmpty() {
		locationName = location.ShortName.Get()
	}

	view = &Div{
		Class: "event",
		Content: Views{
			&Div{
				Class: "date-box",
				Content: Views{
					&Div{Class: "day", Content: HTML(day)},
					&Div{Class: "month", Content: HTML(month)},
				},
			},
			&Div{
				Class: "info",
				Content: Views{
					Printf("<a href='%s' target='_blank'>startuplive.in/<span class='region'>%s</span></a>", regionURL, region.Name),
					Printf("<div class='location'>%s</div>", locationName),
				},
			},
			DivClearBoth(),
		},
	}

	return view, nil
}
