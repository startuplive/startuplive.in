package dashboard

import (
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/ungerik/go-start/debug"
	. "github.com/ungerik/go-start/view"
)

func init() {
	debug.Nop()

	Region0_Event1_Dashboard_Info = &Page{
		OnPreRender: SetEventPageData,
		Title:  EventDashboardTitle("Info"),
		CSS:         IndirectURL(&Region0_DashboardCSS),
		Content: Views{
			DashboardHeader(),
			&Div{
				Class: "content",
				Content: Views{
					&Div{
						Class: "main mentors",
						Content: Views{
							&Div{
								Class: "main-header",
								Content: Views{
									H1("Event Info"),
									//							&Form{
									//								Class: "search-mentor",
									//								Content: &TextField{},
									//							},
								},
							},
							DynamicView(func(ctx *Context) (view View, err error) {
								info := ctx.Data.(*PageData).Event.DashboardInfo_HTML.Get()
								return HTML(info), nil
							}),
							DivClearBoth(),
						},
					},
					eventDashboardSidebar(),
				},
			},
			eventDashboardFooter(),
		},
	}
}
