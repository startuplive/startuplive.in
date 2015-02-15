package dashboard

import (
	. "github.com/ungerik/go-start/view"
	. "github.com/STARTeurope/startuplive.in/views"
)

func init() {
	Region0_Event1_Dashboard_Organisers = &Page{
		OnPreRender: SetEventPageData,
		Title:  EventDashboardTitle("Organisers"),
		CSS:         IndirectURL(&Region0_DashboardCSS),
		Content: Views{
			DashboardHeader(),
			&Div{
				Class: "content",
				Content: Views{
					&Div{
						Class:   "main organisers",
						Content: IMG("/images/organisers.jpg"),
					},
					eventDashboardSidebar(),
				},
			},
			eventDashboardFooter(),
		},
	}
}
