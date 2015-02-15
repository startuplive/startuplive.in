package dashboard

import (
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/ungerik/go-start/debug"
	// "github.com/ungerik/go-start/utils"
	. "github.com/ungerik/go-start/view"
)

func init() {
	debug.Nop()

	Region0_Event1_Dashboard_Teams = &Page{
		OnPreRender: SetEventPageData,
		Title:       EventDashboardTitle("Teams"),
		CSS:         IndirectURL(&Region0_DashboardCSS),
		Content: Views{
			DashboardHeader(),
			&Div{
				Class: "content",
				Content: Views{
					&Div{
						Class: "main",
						Content: Views{
							&Div{
								Class: "main-header",
								Content: Views{
									DynamicView(
										func(ctx *Context) (view View, err error) {
											region := ctx.Data.(*PageData).Region
											return H1("Teams in " + region.Name.Get()), nil
										},
									),
									//							&Form{
									//								Class: "search-mentor",
									//								Content: &TextField{},
									//							},
								},
							},
							&ModelIteratorView{
								GetModel: func(ctx *Context) (interface{}, error) {
									return new(models.EventTeam), nil
								},
								GetModelIterator: EventTeamIterator,
								GetModelView:     eventTeamPreview,
							},
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

func eventTeamPreview(ctx *Context, model interface{}) (view View, err error) {
	eventTeam := model.(*models.EventTeam)

	if eventTeam.Cancelled() {
		return nil, nil
	}

	teamURL := Region0_Event1_Dashboard_Team2.URL(
		ctx.ForURLArgs(ctx.URLArgs[0], ctx.URLArgs[1], eventTeam.ID.Hex()))

	logo, err := eventTeam.LogoImage("framed-image", 160)
	if err != nil {
		return nil, err
	}

	//	i := eventTeam.PersonIterator()
	//	memberNames := make([]string, 0, 16)
	//	for doc := i.Next(); doc != nil; doc = i.Next() {
	//		participant := doc.(*models.EventParticipant)
	//		memberNames = append(memberNames, participant.Name())
	//	}
	//	if i.Err() != nil {
	//		return nil, i.Err()
	//	}

	//needOther := eventTeam.NeedOther.Get()

	view = &Div{
		Class: "preview",
		Content: Views{
			A(teamURL, logo),
			&Div{
				Class: "info",
				Content: Views{
					&Div{
						Class: "info-header",
						Content: Views{
							Printf("<div class='name'>%s</div>", eventTeam.Name),
						},
					},
					&Paragraph{
						Content: &TextPreview{
							PlainText: eventTeam.Abstract.String(),
							MaxLength: 400,
							MoreLink:  NewLinkModel(teamURL, ReadMoreArrow),
						},
					},
					//Text("<table class='buttons'><tr><td>"),
					Printf("<a class='button' href='%s'>Show Details</a>", teamURL),
					//					Text("</td><td>"),
					//					Text("<ul>"),
					//					Text("<li><div class='button'>Get in Touch <div class='dropdown-arrow'></div></div>"),
					//					Text("<ul class='dropdown'>"),
					//					Text("<li><a target='_blank' href='http://twitter.com/moritzplassnig'><img src='/images/icons/twitter-mini.png'/> @moritzplassnig</a></li>"),
					//					Text("<li><a target='_blank' href='http://facebook.com/moritzplassnig'><img src='/images/icons/facebook-mini.png'/>/moritzplassnig</a></li>"),
					//					Text("<li><a target='_blank' href='http://twitter.com/moritzplassnig'><img src='/images/icons/xing-mini.png'/>Moritz Plassnig</a></li>"),
					//					Text("<li><a target='_blank' href='http://twitter.com/moritzplassnig'><img src='/images/icons/linkedin-mini.png'/>Moritz Plassnig</a></li>"),
					//					Text("<li><a target='_blank' href='mailto:moritz.plassnig@starteurope.at'><img src='/images/icons/mail-mini.png'/>moritz.plassnig@start...</a></li>"),
					//					Text("</ul></li></ul>"),
					//Text("</td></tr></table>"),
				},
			},
		},
	}
	return view, nil
}
