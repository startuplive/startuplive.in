package admin

import (
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/STARTeurope/startuplive.in/views/admin"
	//"github.com/ungerik/go-start/debug"
	. "github.com/ungerik/go-start/view"
)

func init() {
	Region0_Event1_Admin_Voting = &Page{
		OnPreRender: SetEventPageData,
		Title:       EventAdminTitle("Voting"),
		CSS:         IndirectURL(&Region0_DashboardCSS),
		Scripts:     admin.PageScripts,
		Content: Views{
			eventadminHeader(),
			VotingEvalView(),
			BR(),
			&Form{
				SubmitButtonText:    "Delete votings",
				SubmitButtonConfirm: "Are you sure you want to delete all votings?",
				SubmitButtonClass:   "button",
				FormID:              "deletevotings",
				OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
					_, err := ctx.Data.(*PageData).Event.DeleteAllVotes()
					return "", StringURL("."), err
				},
			},
			BR(),
			HR(),
			H2("Voting LOG:"),
			DynamicView(
				func(ctx *Context) (view View, err error) {
					event :=
						ctx.Data.(*PageData).Event
					votes, err := event.GetVotes()
					return Views{
						Printf("#votes: %v", votes),
					}, err
				},
			),
			&ModelIteratorTableView{
				Class:            "fullwidth-table voting-visual-table",
				GetModelIterator: EventAllVotesIterator,
				GetRowModel: func(ctx *Context) (interface{}, error) {
					return new(models.Vote), nil
				},
				GetHeaderRowViews: TableHeaderRowEscape("Time", "IP", "Team"),
				GetRowViews:       votingLogRowViews,
			},
		},
	}
}

/*func eventVotesRowViews(rowModel interface{}, response *Response) (views Views, err os.Error) {
	team := rowModel.(*models.EventTeam)
	event := ctx.Data.(*PageData).Event
	allvotes, err := event.GetVotes()
	debug.Print("############### ", allvotes)
	votes, err := event.GetVotesByTeam(team)
	debug.Print("############### ", votes)
	relvotes := 0
	if votes > 0 {
	relvotes = (votes/allvotes)*100
	} 

	debug.Print("############### ", relvotes)

	views = Views{
		Escape(team.Name.Get()),
		Escape(team.Tagline.Get()),
		Printf("%v", relvotes),
	}
	return views, nil
}*/

func VotingEvalView() View {
	return DynamicView(
		func(ctx *Context) (view View, err error) {
			event :=
				ctx.Data.(*PageData).Event
			allvotes, err := event.GetVotes()

			var views Views
			i := event.PitchingTeamsSortedByVotesIterator()

			var team models.EventTeam
			for i.Next(&team) {

				//debug.Print("### allvotes ", allvotes)
				votes, err := event.GetVotesByTeam(&team)
				//debug.Print("### teamvotes ", votes)
				relvotes := 0
				if votes > 0 {
					relvotes = int((float64(votes) / float64(allvotes)) * 100)
				}

				if err != nil {
					return nil, err
				}

				//debug.Print("### relative votes ", relvotes)

				views = append(views, Views{
					Printf("<div style='width:650px; margin-left:30px; height:40px; border-bottom:dashed gray 1px'>"),
					Printf("<span style='width:100px; float:left'>%s</span>", team.Name.Get()),
					Printf("<div style='width:%vpx;height:30px;background-color:red; float:left;'></div>", (relvotes * 5)),
					Printf("<span style='float:left; margin:7px 5px'>%v%% </span>", relvotes),
					Printf("</div>"),
					BR(),
					DivClearBoth(),
				})
			}
			return views, nil
		},
	)
}

func votingLogRowViews(row int, rowModel interface{}, ctx *Context) (views Views, err error) {
	vote := rowModel.(*models.Vote)

	var team models.EventTeam
	err = vote.Team.Get(&team)
	if err != nil {
		return nil, err
	}
	views = Views{
		Escape(vote.Created.Get()),
		Escape(vote.IP.Get()),
		Escape(team.Name.Get()),
	}
	return views, nil
}
