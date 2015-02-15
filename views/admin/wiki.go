package admin

import (
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/ungerik/go-start/model"
	// "github.com/ungerik/go-start/user"
	// "github.com/ungerik/go-start/utils"
	. "github.com/ungerik/go-start/view"
)

func init() {
	Admin_Wiki = &Page{
		Title:   Escape("Wiki | Admin"),
		CSS:     IndirectURL(&Admin_CSS),
		Scripts: PageScripts,
		Content: Views{
			adminHeader(),
			DynamicView(
				func(ctx *Context) (view View, err error) {
					wikientries := models.GetAllWikiEntries()
					return renderWiki(ctx, wikientries)
				},
			),
			&Form{
				SubmitButtonText:  "Create New Wiki Entry",
				SubmitButtonClass: "button",
				FormID:            "createWikiEntry",
				OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
					return "", NewURLWithArgs(Admin_WikiEntry0, "new"), nil
				},
			},
		},
	}
}

func renderWiki(ctx *Context, wiki model.Iterator) (view View, err error) {
	var views Views

	var e models.WikiEntry
	for wiki.Next(&e) {
		entry := e

		showUrl := Admin_WikiEntry0.URL(ctx.ForURLArgs(entry.ID.Hex()))

		author := "no author set"
		var person models.Person
		found, err := entry.CreatedBy.TryGet(&person)
		if err != nil {
			return nil, err
		}
		if found {
			author = person.Name.String()
		}

		views = append(views, DIV("",
			HTML("Title: "+entry.Title.String()),
			BR(),
			HTML("Author: "+author),
			BR(),
			A(showUrl, "show"),
			&Form{
				SubmitButtonText:    "Delete",
				SubmitButtonClass:   "delete",
				FormID:              "deleteWikiEntry",
				SubmitButtonConfirm: "Are you sure you want to delete this entry: " + entry.Title.Get() + "?",
				OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
					return "", StringURL("."), entry.Delete()
				},
			},
			HR(),
		),
		)
	}
	if wiki.Err() != nil {
		return nil, wiki.Err()
	}

	return views, nil
}
