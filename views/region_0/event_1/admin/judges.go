package admin

import (
	"fmt"

	"github.com/ungerik/go-start/model"
	"github.com/ungerik/go-start/mongo"
	. "github.com/ungerik/go-start/view"

	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/STARTeurope/startuplive.in/views/admin"
)

func init() {
	Region0_Event1_Admin_Judges = &Page{
		OnPreRender: SetEventPageData,
		Title:       EventAdminTitle("Judges"),
		CSS:         IndirectURL(&Region0_DashboardCSS),
		PostCSS:     StylesheetLink("/css/ui-lightness/jquery-ui-1.8.17.custom.css"),
		Scripts: Renderers{
			admin.PageScripts,
			JQueryUI,
			JQueryUIAutocompleteFromURL(".add-judge", IndirectURL(&API_People), 2),
		},
		Content: Views{
			eventadminHeader(),
			&ModelIteratorTableView{
				Class:            "visual-table",
				GetModelIterator: EventJudgeIterator,
				GetRowModel: func(ctx *Context) (interface{}, error) {
					return new(models.Person), nil
				},
				GetHeaderRowViews: TableHeaderRowEscape("Nr", "Name", "Company", "Position", "Edit", "Remove"),
				GetRowViews:       eventJudgeRowViews,
			},
			HR(),
			&Form{
				SubmitButtonText:  "Add existing person as judge",
				SubmitButtonClass: "button",
				FormID:            "addJudge",
				GetModel: func(form *Form, ctx *Context) (interface{}, error) {
					return &addJudgeModel{}, nil
				},
				OnSubmit: addJudge,
			},
			HR(),
			&Form{
				SubmitButtonText:  "Add new person as judge",
				SubmitButtonClass: "button",
				FormID:            "createJudge",
				OnSubmit:          createJudge,
			},
		},
	}
}

type addJudgeModel struct {
	Name model.String `view:"class=add-judge"`
}

func addJudge(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
	name := formModel.(*addJudgeModel).Name.Get()
	i := models.People.Iterator()
	var person models.Person
	for i.Next(&person) {
		if name == person.Name.String() {
			person.Judge = true
			err := person.Save()
			if err != nil {
				return "", nil, err
			}
			event := ctx.Data.(*PageData).Event
			event.Judges = append(event.Judges, person.Ref())
			return "", StringURL("."), event.Save()
		}
	}
	if i.Err() != nil {
		return "", nil, i.Err()
	}
	return "", nil, fmt.Errorf("Person '%s' not found", name)
}

func eventJudgeRowViews(row int, rowModel interface{}, ctx *Context) (views Views, err error) {
	person := rowModel.(*models.Person)

	editURL := Region0_Event1_Admin_Judge2.URL(ctx.ForURLArgs(ctx.URLArgs[0], ctx.URLArgs[1], person.ID.Hex()))

	views = Views{
		Printf("%d", row+1),
		Escape(person.Name.String()),
		Escape(person.Company.Get()),
		Escape(person.Position.Get()),
		A(editURL, "Edit"),
		&Form{
			SubmitButtonText:    "Remove",
			SubmitButtonConfirm: "Are you sure you want to remove " + person.Name.String() + "?",
			FormID:              "remove" + person.ID.Hex(),
			OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
				event := ctx.Data.(*PageData).Event
				event.Judges, _ = mongo.RemoveRefWithIDFromSlice(event.Judges, person.ID)
				return "", StringURL("."), event.Save()
			},
		},
	}
	return views, nil
}

func createJudge(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
	event := ctx.Data.(*PageData).Event

	person := models.NewPerson()
	person.Name.First = "[First]"
	person.Name.Last = "[Last]"
	person.Judge = true
	err := person.Save()
	if err != nil {
		return "", nil, err
	}

	event.Judges = append(event.Judges, person.Ref())
	err = event.Save()
	if err != nil {
		person.Delete()
		return "", nil, err
	}

	return "", StringURL(Region0_Event1_Admin_Judge2.URL(
		ctx.ForURLArgs(ctx.URLArgs[0], ctx.URLArgs[1], person.ID.Hex()))), nil
}
