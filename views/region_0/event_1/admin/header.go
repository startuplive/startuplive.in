package admin

import (
	. "github.com/ungerik/go-start/view"

	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/STARTeurope/startuplive.in/views/region_0/event_1"
)

func eventadminHeader() View {
	return &Div{
		Class: "header",
		Content: Views{
			&Link{
				Class: "title",
				Model: &PageLink{
					Page:    &Region0_Event1_Admin,
					Content: event_1.EventHeaderLogoAndTitle(),
				},
			},
			HeaderUserNav(nil),
			//			Text("<table class='next-session'>"),
			//			Text("<tr><td>Next Session <span class='dash'>&nbsp;/</span></td><td align='right'>Team Formation</td></tr>"),
			//			Text("<tr><td>Time Left <span class='dash'>&nbsp;/</span></td><td align='right'>01:15:53</td></tr>"),
			//			Text("</table>"),
			DynamicView(
				func(ctx *Context) (View, error) {
					event := ctx.Data.(*PageData).Event

					menuItems := []LinkModel{
						NewPageLink(&Region0_Event1_Admin, "Overview"),
						NewPageLink(&Region0_Event1_Admin_Wiki, "Compendium"),
						NewPageLink(&Region0_Event1_Admin_About, "About"),
						NewPageLink(&Region0_Event1_Admin_Location, "Location"),
						NewPageLink(&Region0_Event1_Admin_Schedule, "Schedule"),
						NewPageLink(&Region0_Event1_Admin_Partners, "Partners"),
						NewPageLink(&Region0_Event1_Admin_Mentors, "Mentors"),
						NewPageLink(&Region0_Event1_Admin_Judges, "Judges"),
						NewPageLink(&Region0_Event1_Admin_Organisers, "Organisers"),
						NewPageLink(&Region0_Event1_Admin_FAQ, "FAQ"),
						NewPageLink(&Region0_Event1_Admin_Participants, "Participants"),
						NewPageLink(&Region0_Event1_Admin_Teams, "Teams"),
						NewPageLink(&Region0_Event1_Admin_Judgements, "Judgements"),
						NewPageLink(&Region0_Event1_Admin_Voting, "Voting"),
						NewPageLink(&Region0_Event1_Admin_Feedback, "Feedback"),
						NewPageLink(&Region0_Event1_Admin_Settings, "Settings"),
						NewPageLink(&Region0_Event1_Dashboard, "Dashboard"),
						NewPageLink(&Region0_Event1_Admin_DashboardInfo, "Dashboard Info"),
					}
					isAdmin := SessionUserIsSuperAdmin(ctx)

					if (!event.SetupAmiandoEventRequest.Get() && !event.AmiandoEventActivated.Get()) || isAdmin {
						menuItems = append(menuItems, NewPageLink(&Region0_Event1_Admin_AmiandoData, "Amiando Data"))
					}

					if isAdmin {
						menuItems = append(menuItems, NewPageLink(&Admin_Region0, "Region"))
						menuItems = append(menuItems, NewPageLink(&Region0_Event1_Admin_Amiando, "Amiando Setup"))
					}

					return &Div{
						Class: "menu-frame",
						Content: Views{
							&Menu{
								Class:           "menu",
								ItemClass:       "menu-item",
								ActiveItemClass: "active",
								BetweenItems:    " &nbsp;/&nbsp; ",
								Items:           menuItems,
							},
							DivClearBoth(),
						},
					}, nil
				}),
		},
	}
}
