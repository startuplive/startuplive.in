package admin

import (
	"github.com/ungerik/go-start/config"
	"github.com/ungerik/go-start/debug"
	"github.com/ungerik/go-start/model"
	"github.com/ungerik/go-start/utils"
	. "github.com/ungerik/go-start/view"

	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
)

func init() {
	Admin_People = &Page{
		Title:   Escape("People | Admin"),
		CSS:     IndirectURL(&Admin_CSS),
		Scripts: PageScripts,
		Content: Views{
			adminHeader(),
			DynamicView(
				func(ctx *Context) (view View, err error) {
					exportEmailsURL := Admin_ExportEmails.URL(ctx)
					exportPitcherEmailsURL := Admin_ExportPitcherEmails.URL(ctx)
					exportPeople := Admin_ExportPeople.URL(ctx)
					exportMentorsJudges := Admin_ExportMentorJudgeEmails.URL(ctx)
					exportContacts := Admin_ExportContacts.URL(ctx)
					return Views{
						Printf("<a href='%s'>download pitcher emails as csv</a>", exportPitcherEmailsURL),
						BR(),
						Printf("<a href='%s'>download all emails as csv</a>", exportEmailsURL),
						BR(),
						Printf("<a href='%s'>download quality people data as csv</a>", exportPeople),
						BR(),
						Printf("<a href='%s'>download mentors and judges as csv</a>", exportMentorsJudges),
						BR(),
						Printf("<a href='%s'>download all user contact data as csv</a>", exportContacts),
					}, nil
				},
			),
			&ModelIteratorTableView{
				Class: "visual-table",
				GetModelIterator: func(ctx *Context) model.Iterator {
					return models.People.SortFunc(
						func(a, b *models.Person) bool {
							return utils.CompareCaseInsensitive(a.Name.String(), b.Name.String())
						},
					)
					return models.People.Iterator()
				},
				GetRowModel: func(ctx *Context) (interface{}, error) {
					return new(models.Person), nil
				},
				GetHeaderRowViews: func(ctx *Context) (views Views, err error) {
					views = Views{
						HTML("Nr"),
						HTML("Name"),
						HTML("Email"),
						HTML("Position"),
						HTML("Company"),
						HTML("Edit"),
					}
					if SessionUserIsSuperAdmin(ctx) {
						views = append(views, HTML("Delete"))
					}
					return views, nil
				},
				GetRowViews: func(row int, rowModel interface{}, ctx *Context) (views Views, err error) {
					person := rowModel.(*models.Person)
					editURL := Admin_Person0.URL(ctx.ForURLArgs(person.ID.Hex()))
					views = Views{
						Printf("%d", row+1),
						Escape(person.Name.String()),
						Escape(person.PrimaryEmail()),
						Escape(person.Position.Get()),
						HTML(person.Company.Get()),
						A(editURL, "Edit"),
					}
					if SessionUserIsSuperAdmin(ctx) {
						views = append(views,
							&Form{
								SubmitButtonText:    "Delete",
								SubmitButtonConfirm: "Are you sure you want to delete " + person.Name.String() + "?",
								FormID:              "delete" + person.ID.Hex(),
								OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
									config.Logger.Printf("Deleting user '%s' with ID %s", person.Name.String(), person.ID.Hex())
									config.Logger.Printf("FormID: " + form.FormID)
									debug.Dump(person)
									return "", StringURL("."), person.Delete()
								},
							},
						)
					}
					return views, nil
				},
			},
			&Form{
				SubmitButtonText:  "Add Person",
				SubmitButtonClass: "button",
				FormID:            "addperson",
				OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
					var person models.Person
					models.People.InitDocument(&person)
					person.Name.First = "[First]"
					person.Name.Last = "[Last]"
					person.AddEmail("", "")
					err := person.Save()
					if err != nil {
						return "", nil, err
					}
					return "", NewURLWithArgs(Admin_Person0, person.ID.Hex()), nil
				},
			},
		},
	}
}
