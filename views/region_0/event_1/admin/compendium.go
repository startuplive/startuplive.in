package admin

import (
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/STARTeurope/startuplive.in/views/admin"
	"github.com/ungerik/go-start/model"
	// "github.com/ungerik/go-start/user"
	// "github.com/ungerik/go-start/utils"
	. "github.com/ungerik/go-start/view"
)

func init() {
	Region0_Event1_Admin_Wiki = &Page{
		Title:   Escape("Compendium"),
		CSS:     IndirectURL(&Admin_CSS),
		Scripts: admin.PageScripts,
		Content: Views{
			eventadminHeader(),
			BR(),
			DynamicView(
				func(ctx *Context) (view View, err error) {
					wikientries := models.GetAllPublicWikiEntries()
					return renderPublicWiki(ctx, wikientries)
				},
			),
		},
	}
}

func renderPublicWiki(ctx *Context, wiki model.Iterator) (view View, err error) {
	var views Views

	views = append(views, H1("Startup Live Compendium"))

	var entry models.WikiEntry
	for wiki.Next(&entry) {
		showUrl := Region0_Event1_Admin_WikiEntry0.URL(ctx.ForURLArgs(ctx.URLArgs[0], ctx.URLArgs[1], entry.ID.Hex()))

		views = append(views, DIV("",
			A(showUrl, entry.Title.String()),
			HR(),
		),
		)
	}
	if wiki.Err() != nil {
		return nil, wiki.Err()
	}

	return views, nil
}
