package admin

import (
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/STARTeurope/startuplive.in/views/admin"
	. "github.com/ungerik/go-start/view"
)

func init() {
	Region0_Event1_Admin_Judge2 = &Page{
		OnPreRender: SetEventPersonPageData,
		Title: Render(
			func(ctx *Context) (err error) {
				person := ctx.Data.(*PageData).Person
				return EventAdminTitle("Judge " + person.Name.String()).Render(ctx)
			},
		),
		CSS:     IndirectURL(&Region0_DashboardCSS),
		Scripts: admin.PageScripts,
		Content: Views{
			eventadminHeader(),
			DynamicView(EventAdminJudgeView),
		},
	}
}

func EventAdminJudgeView(ctx *Context) (view View, err error) {
	person := ctx.Data.(*PageData).Person
	//judgesURL := Region0_Event1_Admin_Judges.URL(response, ctx.URLArgs[0], ctx.URLArgs[1])

	excludedFields, err := ExcludedPersonFormFields(ctx)
	if err != nil {
		return nil, err
	}
	person.Judge = true

	requireFields := []string{
		"Name.First",
		"Name.Last",
		"Company",
		"Position",
	}
	views := Views{
		H2("Judge " + person.Name.String()),
		PersonForm(person, Region0_Event1_Admin_Judges, []string{ /*"Judge"*/}, excludedFields, requireFields),
	}
	return views, nil
}
