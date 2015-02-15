package admin

import (
	// "github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/STARTeurope/startuplive.in/views/admin"
	// "github.com/ungerik/go-start/model"
	// "github.com/ungerik/go-start/user"
	// "github.com/ungerik/go-start/utils"
	. "github.com/ungerik/go-start/view"
)

func init() {
	Region0_Event1_Admin_WikiEntry0 = &Page{
		OnPreRender: SetWikiEntryPageData,
		Title:       Escape("Compendium"),
		CSS:         IndirectURL(&Admin_CSS),
		Scripts:     admin.PageScripts,
		Content: Views{
			eventadminHeader(),
			BR(),
			A("../../", HTML("&larr; Back to the compendium")),
			HR(),
			DynamicView(
				func(ctx *Context) (view View, err error) {
					entry := ctx.Data.(*PageData).WikiEntry
					if err != nil {
						return nil, err
					}

					return Views{
						DIV("", H1(entry.Title.String())),
						DIV("", HTML(entry.Content.String())),
						HR(),
						DynamicView(
							func(ctx *Context) (view View, err error) {
								return admin.ShowAnswers(entry)
							},
						),
					}, nil
				},
			),
		},
	}
}
