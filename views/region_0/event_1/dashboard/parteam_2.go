package dashboard

import (
	"labix.org/v2/mgo/bson"

	"github.com/ungerik/go-start/model"
	. "github.com/ungerik/go-start/view"

	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
)

func init() {
	Region0_Event1_Dashboard_Team2 = &Page{
		OnPreRender: SetEventTeamPageData,
		Title: Render(
			func(ctx *Context) (err error) {
				team := ctx.Data.(*PageData).Team
				return EventDashboardTitle("Team " + team.Name.Get()).Render(ctx)
			},
		),
		CSS: IndirectURL(&Region0_DashboardCSS),
		Content: Views{
			DashboardHeader(),
			&Div{
				Class: "content",
				Content: Views{
					&ModelIteratorView{
						GetModel: func(ctx *Context) (interface{}, error) {
							return new(models.EventTeam), nil
						},
						GetModelIterator: func(ctx *Context) model.Iterator {
							return models.EventTeams.DocumentWithIDIterator(bson.ObjectIdHex(ctx.URLArgs[2]))
						},
						GetModelView: func(ctx *Context, m interface{}) (view View, err error) {
							eventTeam := m.(*models.EventTeam)

							logo, err := eventTeam.LogoImage("team-logo framed-image", 300)
							if err != nil {
								return nil, err
							}

							var needs Views
							if eventTeam.NeedTechies.Get() {
								needs = append(needs, HTML("<li>Techies</li>"))
							}
							if eventTeam.NeedBizPeople.Get() {
								needs = append(needs, HTML("<li>Biz People</li>"))
							}
							if eventTeam.NeedDesigners.Get() {
								needs = append(needs, HTML("<li>Designers</li>"))
							}
							if !eventTeam.NeedOther.IsEmpty() {
								needs = append(needs, Printf("<li>%s</li>", eventTeam.NeedOther))
							}
							if len(needs) > 0 {
								needs = append(needs, HTML("</ul>"))
								needs = append(Views{HTML("<h3>Need:</h3><ul>")}, needs...)
							}

							view = &Div{
								Class: "main team-details",
								Content: Views{
									&Div{
										Class: "main-header",
										Content: Views{
											H1(eventTeam.Name.Get()),
											HTML("<a href='../../teams/'>&larr;back to overview</a>"),
										},
									},
									&Div{
										Class: "short-info",
										Content: Views{
											logo,
											&If{
												Condition: !eventTeam.Tagline.IsEmpty(),
												Content:   Printf("<h2>%s:</h2>&quot;%s&quot;", eventTeam.Name, eventTeam.Tagline),
											},
											&Div{
												Class: "social-media-icons",
												Content: Views{
													&If{
														Condition: eventTeam.FacebookURL != "",
														Content:   Printf("<a href='%s' target='_blank'><img src='/images/icons/facebook30.png'/></a>", eventTeam.FacebookURL),
													},
													&If{
														Condition: eventTeam.TwitterURL != "",
														Content:   Printf("<a href='%s' target='_blank'><img src='/images/icons/twitter30.png'/></a>", eventTeam.TwitterURL),
													},
													&If{
														Condition: eventTeam.LinkedInURL != "",
														Content:   Printf("<a href='%s' target='_blank'><img src='/images/icons/linkedin30.png'/></a>", eventTeam.LinkedInURL),
													},
													&If{
														Condition: eventTeam.YoutubeURL != "",
														Content:   Printf("<a href='%s' target='_blank'><img src='/images/icons/youtube30.png'/></a>", eventTeam.YoutubeURL),
													},
												},
											},
										},
									},
									DivClearBoth(),
									&If{
										Condition: eventTeam.Abstract != "",
										Content:   Views{H3("Description:"), HTML(eventTeam.Abstract.Get())},
									},
									&If{
										Condition: eventTeam.ProblemOpportunity != "",
										Content:   Views{H3("Problem/Opportunity:"), HTML(eventTeam.ProblemOpportunity.Get())},
									},
									&If{
										Condition: eventTeam.Haves != "",
										Content:   Views{H3("Have:"), HTML(eventTeam.Haves.Get())},
									},
									needs,
									H3(PrintfEscape("%s Team Members:", eventTeam.Name)),
									&Div{
										Class: "team-members",
										Content: Views{
											&ModelIteratorView{
												GetModel: func(ctx *Context) (interface{}, error) {
													return new(models.EventParticipant), nil
												},
												GetModelIterator: func(ctx *Context) model.Iterator {
													return ctx.Data.(*PageData).Team.ParticipantIterator()
												},
												GetModelView: eventParticipantPreview,
											},
											DivClearBoth(),
										},
									},
								},
							}
							return view, nil
						},
					},
					eventDashboardSidebar(),
				},
			},
		},
	}
}
