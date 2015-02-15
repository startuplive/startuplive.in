package admin

import (
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/ungerik/go-start/model"
	"github.com/ungerik/go-start/utils"
	. "github.com/ungerik/go-start/view"
	//	"github.com/ungerik/go-start/debug"
)

func init() {
	Admin_Startups = &Page{
		Title:   Escape("Startups | Admin"),
		CSS:     IndirectURL(&Admin_CSS),
		Scripts: PageScripts,
		Content: Views{
			adminHeader(),
			&ModelIteratorTableView{
				Class: "visual-table",
				GetModelIterator: func(ctx *Context) model.Iterator {
					return models.Startups.SortFunc(
						func(a, b *models.Startup) bool {
							return utils.CompareCaseInsensitive(a.Name.String(), b.Name.String())
						},
					)
				},
				GetRowModel: func(ctx *Context) (interface{}, error) {
					return new(models.Startup), nil
				},
				GetHeaderRowViews: func(ctx *Context) (views Views, err error) {
					views = Views{
						HTML("Nr"),
						HTML("Name"),
						HTML("Website"),
						HTML("Founder"),
						HTML("View"),
					}

					return views, nil
				},
				GetRowViews: func(row int, rowModel interface{}, ctx *Context) (views Views, err error) {
					startup := rowModel.(*models.Startup)
					editURL := Admin_Startup0.URL(ctx.ForURLArgs(startup.ID.Hex()))

					var person models.Person
					found, err := startup.Founder.TryGet(&person)
					if err != nil {
						return nil, err
					}
					founderName := ""
					if !found {
						founderName = "not set"
					}
					founderName = person.Name.String()

					views = Views{
						Printf("%d", row+1),
						Escape(startup.Name.String()),
						Escape(startup.Website.String()),
						Escape(founderName),
						A(editURL, "View"),
					}

					return views, nil
				},
			},
		},
	}
}
