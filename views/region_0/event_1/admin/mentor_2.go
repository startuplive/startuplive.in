package admin

import (
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/STARTeurope/startuplive.in/views/admin"
	. "github.com/ungerik/go-start/view"
)

func init() {
	Region0_Event1_Admin_Mentor2 = &Page{
		OnPreRender: SetEventPersonPageData,
		Title: Render(
			func(ctx *Context) (err error) {
				person := ctx.Data.(*PageData).Person
				return EventAdminTitle("Mentor " + person.Name.String()).Render(ctx)
			},
		),
		CSS:     IndirectURL(&Region0_DashboardCSS),
		Scripts: admin.PageScripts,
		Content: Views{
			eventadminHeader(),
			DynamicView(
				func(ctx *Context) (view View, err error) {
					person := ctx.Data.(*PageData).Person
					//mentorsURL := Region0_Event1_Admin_Mentors.URL(context, ctx.URLArgs[0], ctx.URLArgs[1])
					excludeFields, err := ExcludedPersonFormFields(ctx)
					if err != nil {
						return nil, err
					}
					person.Mentor = true

					requireFields := []string{
						"Name.First",
						"Name.Last",
						"Company",
						"Position",
					}

					views := Views{
						H2("Mentor " + person.Name.String()),
						PersonForm(person, Region0_Event1_Admin_Mentors, []string{ /*"Mentor"*/}, excludeFields, requireFields),
					}
					return views, nil
				},
			),
		},
	}
}
