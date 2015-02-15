package dashboard

import (
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/STARTeurope/startuplive.in/views/region_0/event_1"
	. "github.com/ungerik/go-start/view"
)

func DashboardHeader() View {
	return &Div{
		Class: "header",
		Content: Views{
			&Link{
				Class: "title",
				Model: &PageLink{
					Page:    &Region0_Event1_Dashboard,
					Content: event_1.EventHeaderLogoAndTitle(),
				},
			},
			HeaderUserNav(event_1.AdditionalTopNav),
			//			Text("<table class='next-session'>"),
			//			Text("<tr><td>Next Session <span class='dash'>&nbsp;/</span></td><td align='right'><span class='next-session-name'>Team Formation</span></td></tr>"),
			//			Text("<tr><td>Time Left <span class='dash'>&nbsp;/</span></td><td align='right'>01:15:53</td></tr>"),
			//			Text("</table>"),
			&Div{
				Class: "menu-frame",
				Content: DynamicView(
					func(ctx *Context) (view View, err error) {
						pageData := ctx.Data.(*PageData)
						event := pageData.Event

						menuItems := []LinkModel{
							NewPageLink(&Region0_Event1_Dashboard, "Timetable"),
						}
						if event.Show.Info.Get() {
							menuItems = append(menuItems, NewPageLink(&Region0_Event1_Dashboard_Info, "Info"))
						}

						m := event.MentorsJudgesTab_RenameMentors.GetOrDefault("Mentors")

						menuItems = append(
							menuItems,
							NewPageLink(&Region0_Event1_Dashboard_Participants, "Participants"),
							NewPageLink(&Region0_Event1_Dashboard_Teams, "Teams"),
							NewPageLink(&Region0_Event1_Dashboard_Mentors, m),
							NewPageLink(&Region0_Event1_Dashboard_Judges, "Judges"),
							//NewPageLink(&Region0_Event1_Dashboard_Organisers, "Organisers"),
						)

						if event.Show.Voting.Get() {
							menuItems = append(
								menuItems,
								NewPageLink(&Region0_Event1_Voting, "Voting"),
							)
						}
						if event.Show.VotingResult.Get() {
							menuItems = append(
								menuItems,
								NewPageLink(&Region0_Event1_Dashboard_VotingResult, "Voting-Result"),
							)
						}
						return Views{
							&Menu{
								Class:           "menu",
								ItemClass:       "menu-item",
								ActiveItemClass: "active",
								BetweenItems:    " &nbsp;/&nbsp; ",
								Items:           menuItems,
							},
							//					&Menu{
							//						Class:           "menu profile-menu",
							//						ItemClass:       "menu-item",
							//						ActiveItemClass: "active",
							//						BetweenItems:    " &nbsp;/&nbsp; ",
							//						Items: []Linker{
							//							NewPageLink(&Region0_Event1_Dashboard_Organisers, "Profile Settings"),
							//							NewPageLink(&Region0_Event1_Dashboard_Organisers, "Logout"),
							//						},
							//					},
							DivClearBoth(),
						}, nil
					},
				),
			},
		},
	}
}
