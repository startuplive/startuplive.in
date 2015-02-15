package dashboard

import (
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	. "github.com/ungerik/go-start/view"
)

func init() {
	Region0_Event1_Dashboard_Participant2 = &Page{
		OnPreRender: SetEventParticipantPageData,
		Title: Render(
			func(ctx *Context) (err error) {
				participant := ctx.Data.(*PageData).Participant
				return EventDashboardTitle(participant.Name()).Render(ctx)
			},
		),
		CSS:     Region0_DashboardSubmodalCSS,
		Content: DynamicView(eventParticipantDetailsView),
	}
	//	Region0_Event1_Dashboard_Participant2.Template.GetContext = func(requestContext *Context) (context interface{}, err os.Error) {
	//		return map[string]bool{"IsModal": true}, nil
	//	}
}

func eventParticipantDetailsView(ctx *Context) (view View, err error) {
	participant := ctx.Data.(*PageData).Participant
	person := participant.GetPerson()

	var eventTeam models.EventTeam
	hasTeam, err := participant.Team.TryGet(&eventTeam)
	if err != nil {
		return nil, err
	}

	var teamName string
	var teamURL string
	var teamleader bool

	if hasTeam {
		teamName = eventTeam.Name.Get()
		teamURL = Region0_Event1_Dashboard_Team2.URL(ctx.ForURLArgs(ctx.URLArgs[0], ctx.URLArgs[1], eventTeam.ID.Hex()))
		teamleader = eventTeam.Leader.ID == participant.ID
	}

	view = &Div{
		Class: "person-details",
		Content: Views{
			Printf("<h1>%s</h1>", person.Name.String()),
			H2(
				&If{Condition: teamleader, Content: HTML("Teamleader: ")},
				A(teamURL, teamName),
			),
			HTML("<div class='spacer-top'></div>"),
			ViewOrError(person.Image_320x320("framed-image")),
			HTML("<div class='buttons'>"),
			//HTML("<a href='#' class='button'>Start conversation</a> <a href='#' class='button'>Social Media Profiles</a>"),
			HTML("</div>"),
			Printf("<div class='tags'>%s</div>", participant.Background),
			//			Printf("<p>Email: <a href='mailto:%s'>%s</a></p>", email, email),
		},
	}
	return view, nil
}
