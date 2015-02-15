package admin

import (
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/STARTeurope/startuplive.in/views/admin"
	. "github.com/ungerik/go-start/view"
)

func init() {
	Region0_Event1_Admin_Partner2 = &Page{
		OnPreRender: SetEventPersonPageData,
		Title: Render(
			func(ctx *Context) (err error) {
				// person := ctx.Data.(*PageData).Person
				return EventAdminTitle("Partner ").Render(ctx)
			},
		),
		CSS:     IndirectURL(&Region0_DashboardCSS),
		Scripts: admin.PageScripts,
		Content: Views{
			eventadminHeader(),
			// DynamicView(EventAdminJudgeView),
		},
	}
}
