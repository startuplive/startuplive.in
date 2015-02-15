package dashboard

import (
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/STARTeurope/startuplive.in/views/admin"
	eventAdmin "github.com/STARTeurope/startuplive.in/views/region_0/event_1/admin"
	. "github.com/ungerik/go-start/view"
)

func init() {
	Region0_Event1_Dashboard_VotingResult = &Page{
		OnPreRender:  SetEventPageData,
		Title:   EventAdminTitle("Voting-Result"),
		CSS:          IndirectURL(&Region0_DashboardCSS),
		Scripts: admin.PageScripts,
		Content: Views{
			DashboardHeader(),
			eventAdmin.VotingEvalView(),
		},
	}
}
