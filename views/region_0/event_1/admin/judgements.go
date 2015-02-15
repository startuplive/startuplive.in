package admin

import (
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/STARTeurope/startuplive.in/views/admin"
	"github.com/ungerik/go-start/utils"
	. "github.com/ungerik/go-start/view"
)

func init() {
	Region0_Event1_Admin_Judgements = &Page{
		OnPreRender: SetEventPageData,
		Title:       EventAdminTitle("Valuation"),
		CSS:         IndirectURL(&Region0_DashboardCSS),
		Scripts:     admin.PageScripts,
		Content: Views{
			eventadminHeader(),
			H2("Pitching Teams:"),
			&ModelIteratorTableView{
				Class:            "judgement-visual-table",
				GetModelIterator: EventPitchingTeamScoreIterator,
				GetRowModel: func(ctx *Context) (interface{}, error) {
					return new(models.EventTeam), nil
				},
				GetHeaderRowViews: TableHeaderRowEscape("Nr", "Name", "Average Score", "valuate"),
				GetRowViews:       eventValuationTeamRowViews,
			},
			BR(),
			HR(),
			BR(),
			H2("Not Pitching:"),
			&ModelIteratorTableView{
				Class:            "judgement-visual-table",
				GetModelIterator: EventNotPitchingTeamScoreIterator,
				GetRowModel: func(ctx *Context) (interface{}, error) {
					return new(models.EventTeam), nil
				},
				GetHeaderRowViews: TableHeaderRowEscape("Nr", "Name", "Average Score", "valuate"),
				GetRowViews:       eventValuationTeamRowViews,
			},
		},
	}
}

func eventValuationTeamRowViews(row int, rowModel interface{}, ctx *Context) (views Views, err error) {
	team := rowModel.(*models.EventTeam)

	judgeURL := Region0_Event1_Admin_Judgements_Team2.URL(
		ctx.ForURLArgs(ctx.URLArgs[0], ctx.URLArgs[1], team.ID.Hex()))

	views = Views{
		Printf("%d", row+1),
		Escape(team.Name.Get()),
		&If{
			Condition:   len(team.Judgements) > 0,
			Content:     Printf("%v", utils.Round(team.ComputeAverageScoreByEvent(), 2)),
			ElseContent: Printf("not judged"),
		},
		A(judgeURL, "judge"),
	}
	return views, nil
}
