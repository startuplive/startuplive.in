package admin

import (
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/STARTeurope/startuplive.in/views/admin"
	"github.com/ungerik/go-start/debug"
	"github.com/ungerik/go-start/model"
	"github.com/ungerik/go-start/mongo"
	"github.com/ungerik/go-start/utils"
	. "github.com/ungerik/go-start/view"
)

func init() {
	debug.Nop()
	Region0_Event1_Admin_Judgements_Team2 = &Page{
		OnPreRender: SetEventTeamPageData,
		Title: Render(
			func(ctx *Context) (err error) {
				team := ctx.Data.(*PageData).Team
				return EventAdminTitle("Team " + team.Name.Get()).Render(ctx)
			},
		),
		CSS:     IndirectURL(&Region0_DashboardCSS),
		Scripts: admin.PageScripts,
		Content: Views{
			eventadminHeader(),
			DynamicView(teamJudgementHeadEditorView),
			DynamicView(teamJudgementEditorView),
			DivClearBoth(),
		},
	}
}

func eventValuationJudgesRowViews(rowModel interface{}, response *Response) (views Views, err error) {
	//team := ctx.Data.(*PageData).Team
	person := rowModel.(*models.Person)

	views = Views{
		Escape(person.Name.String()),
		Printf("Valuate"),
		HR(),
	}
	return views, nil
}

func teamJudgementHeadEditorView(ctx *Context) (view View, err error) {
	team := ctx.Data.(*PageData).Team
	event := ctx.Data.(*PageData).Event

	teamName := team.Name.Get()

	views := Views{
		H2(A(Region0_Event1_Admin_Judgements, HTML("&larr;Back"))),
		HR(),
		H2(teamName),
		HR(),
		&If{
			Condition: len(team.Judgements) != 0,
			Content: &Views{
				H3(Printf("Average Score: %v", utils.Round(team.ComputeAverageScoreByEvent(), 2))),
				H3(Printf("Average Score / Category: %v", utils.Round(team.ComputeAverageCatScoreByEvent(), 2))),
				HTML("<table>"),
				HTML(Printf("<tr><td>presentation, value proposition clarity (average):</td><td>%v</td></tr>", utils.Round(team.ComputeAverageJudgementCategory(models.JudgementPresentation), 2))),
				HTML(Printf("<tr><td>innovation, non-immitability (average):</td><td>%v</td></tr>", utils.Round(team.ComputeAverageJudgementCategory(models.JudgementInnovation), 2))),
				HTML(Printf("<tr><td>traction, proof of market (average):</td><td>%v</td></tr>", utils.Round(team.ComputeAverageJudgementCategory(models.JudgementTraction), 2))),
				HTML(Printf("<tr><td>founder impression, team impression (average):</td><td>%v</td></tr>", utils.Round(team.ComputeAverageJudgementCategory(models.JudgementTeamImpression), 2))),
				HTML(Printf("<tr><td>market, attractiveness, competition (average):</td><td>%v</td></tr>", utils.Round(team.ComputeAverageJudgementCategory(models.JudgementMarket), 2))),
				HTML(Printf("<tr><td>monetizabilty, scalability (average):</td><td>%v</td></tr>", utils.Round(team.ComputeAverageJudgementCategory(models.JudgementScalability), 2))),
				HTML(Printf("<tr><td>feasibility, Vredible going to market (average):</td><td>%v</td></tr>", utils.Round(team.ComputeAverageJudgementCategory(models.JudgementFeasability), 2))),
				&If{ // For Pioneersfestival
					Condition: event.Type.String() == "PioneersFestival",
					Content:   HTML(Printf("<tr><td>image voting (average):</td><td>%v</td></tr>", utils.Round(team.ComputeAverageJudgementCategory(models.JudgementImageVoting), 2))),
				},
				HTML("</table>"),
				/*H4(
					&If{
						Condition:   team.ComputeAverageJudgementCategory(models.JudgementPresentation) > 0,
						Content:     Printf("presentation, value proposition clarity (average): %v", team.ComputeAverageJudgementCategory(models.JudgementPresentation)),
						ElseContent: Printf("presentation, value proposition clarity (average): k.A"),
					}),
				H4(
					&If{
						Condition:   team.ComputeAverageJudgementCategory(models.JudgementInnovation) > 0,
						Content:     Printf("presentation, value proposition clarity (average): %v", team.ComputeAverageJudgementCategory(models.JudgementInnovation)),
						ElseContent: Printf("presentation, value proposition clarity (average): k.A"),
					}),
				H4(Printf("traction, proof of market (average): %v", team.ComputeAverageJudgementCategory(models.JudgementTraction))),
				H4(Printf("founder impression, team impression (average): %v", team.ComputeAverageJudgementCategory(models.JudgementTeamImpression))),
				H4(Printf("market, attractiveness, competition (average): %v", team.ComputeAverageJudgementCategory(models.JudgementMarket))),
				H4(Printf("monetizabilty, scalability (average): %v", team.ComputeAverageJudgementCategory(models.JudgementScalability))),
				H4(Printf("Feasibility, Vredible going to market (average): %v", team.ComputeAverageJudgementCategory(models.JudgementFeasability))),
				*/
			},
			ElseContent: &Views{
				H3("Average Score: - "),
			},
		},
		HR(),
	}

	return views, nil
}

func teamJudgementEditorView(ctx *Context) (view View, err error) {
	judgeiterator := ctx.Data.(*PageData).Event.JudgeIterator()
	team := ctx.Data.(*PageData).Team
	event := ctx.Data.(*PageData).Event

	views := Views{}

	excludedFields := []string{}
	if event.Type.String() != "PioneersFestival" {
		excludedFields = append(excludedFields, "ImageVoting")
	}

	var j models.Person
	for judgeiterator.Next(&j) {
		judge := j // copy by value because it will be used in a closure later on
		key, status := team.HasJudged(&judge)

		if status == models.TeamNotJudgedByJudge {
			//TODO: create new judgement for this judge
			views = append(views,
				DIV("judgeitem",
					&Tag{
						Tag:   "h3",
						Class: "judge-notjudged",
						Content: Views{
							Escape(judge.String()),
						},
					},
					&Tag{
						Tag:     "h4",
						Content: Printf("Score: - "),
					},
					&Form{
						FormID:            judge.ID.Hex(),
						SubmitButtonClass: "button",
						ExcludedFields:    excludedFields,
						GetModel: func(form *Form, ctx *Context) (interface{}, error) {
							judgement := new(models.Judgement)
							mongo.InitRefs(judgement)
							debug.Print("****** set judge to team 1st: ", judge.ID.Hex())
							judgement.Judge.Set(&judge)
							return judgement, nil
						},
						OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
							val := formModel.(*models.Judgement)
							val.Score.Set(team.ComputeTeamScoreByJudge(val))

							val.Judge.Set(&judge)
							team.Judgements = append(team.Judgements, *val)
							return "", Region0_Event1_Admin_Judgements_Team2, team.Save()
						},
					},
				),
			)
		} else if status == models.TeamJudgedByJudge {
			// debug.Print("********* : ", team.Judgements[key].Score)
			views = append(views,
				DIV("judgeitem",
					&Tag{
						Tag:   "h3",
						Class: "judge-hasjudged",
						Content: Views{
							Escape(judge.String()),
						},
					},
					&Tag{
						Tag:     "h4",
						Content: Printf("Score: %v", utils.Round(team.Judgements[key].Score.Get(), 2)),
					},
					&Form{
						FormID:            judge.ID.Hex(),
						SubmitButtonClass: "button",
						ExcludedFields:    excludedFields,
						GetModel: func(form *Form, ctx *Context) (interface{}, error) {
							return &team.Judgements[key], nil
						},
						OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
							val := formModel.(*models.Judgement)
							val.Score.Set(team.ComputeTeamScoreByJudge(val))
							team.Judgements[key] = *val
							return "", Region0_Event1_Admin_Judgements_Team2, team.Save()
						},
					},
				),
			)
		} else { //status == models.TeamJudgingIncompleteByJudge
			views = append(views,
				DIV("judgeitem",
					&Tag{
						Tag:   "h3",
						Class: "judge-incomplete",
						Content: Views{
							Escape(judge.String()),
						},
					},
					&Tag{
						Tag:     "h4",
						Content: Printf("Score: %v*", utils.Round(team.Judgements[key].Score.Get(), 2)),
					},
					&Form{
						FormID:            judge.ID.Hex(),
						SubmitButtonClass: "button",
						ExcludedFields:    excludedFields,
						GetModel: func(form *Form, ctx *Context) (interface{}, error) {
							return &team.Judgements[key], nil
						},
						OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
							val := formModel.(*models.Judgement)
							val.Score.Set(team.ComputeTeamScoreByJudge(val))
							team.Judgements[key] = *val
							return "", Region0_Event1_Admin_Judgements_Team2, team.Save()
						},
					},
				),
			)
		}
	}
	if judgeiterator.Err() != nil {
		return nil, judgeiterator.Err()
	}

	return views, nil
}

type judged struct {
	key    model.Int //key in team.Valuations
	judge  mongo.Ref `model:"to=people"` //the judge
	status model.Int //0 - not judged at all, 1 - missing fields, 2 - done
}
