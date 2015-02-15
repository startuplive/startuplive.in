package admin

import (
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/STARTeurope/startuplive.in/views/admin"
	. "github.com/ungerik/go-start/view"
)

func init() {
	Region0_Event1_Admin_Location = &Page{
		OnPreRender: SetEventPageData,
		Title:       EventAdminTitle("Location"),
		CSS:         IndirectURL(&Region0_DashboardCSS),
		Scripts:     admin.PageScripts,
		Content: Views{
			eventadminHeader(),
			DynamicView(
				func(ctx *Context) (view View, err error) {
					data := ctx.Data.(*PageData)
					views := Views{
						H2(data.Region.Name.Get() + " Location"),
						&Form{
							SubmitButtonClass: "button",
							FormID:            data.Location.ID.Hex(),
							ExcludedFields:    []string{"Address.Country", "GeoLocation"},
							RequiredFields:    data.Location.GetRequiredFields(),
							GetModel:          FormModel(data.Location),
							OnSubmit:          OnFormSubmitSaveModelAndRedirect(StringURL(".")),
						},
					}
					return views, nil
				},
			),
		},
	}
}
