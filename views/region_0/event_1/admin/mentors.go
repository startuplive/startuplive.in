package admin

import (
	"fmt"
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/STARTeurope/startuplive.in/views/admin"
	"github.com/ungerik/go-start/model"
	"github.com/ungerik/go-start/mongo"
	. "github.com/ungerik/go-start/view"
	// "strings"
)

func init() {
	Region0_Event1_Admin_Mentors = &Page{
		OnPreRender: SetEventPageData,
		Title:       EventAdminTitle("Mentors"),
		CSS:         IndirectURL(&Region0_DashboardCSS),
		PostCSS:     StylesheetLink("/css/ui-lightness/jquery-ui-1.8.17.custom.css"),
		Scripts: Renderers{
			admin.PageScripts,
			JQueryUI,
			JQueryUIAutocompleteFromURL(".add-mentor", IndirectURL(&API_People), 2),
		},
		Content: Views{
			eventadminHeader(),
			&ModelIteratorTableView{
				Class:            "visual-table",
				GetModelIterator: EventMentorIterator,
				GetRowModel: func(ctx *Context) (interface{}, error) {
					return new(models.Person), nil
				},
				GetHeaderRowViews: TableHeaderRowEscape("Nr", "Name", "Company", "Position", "Edit", "Remove"),
				GetRowViews:       eventMentorRowViews,
			},
			HR(),
			&Form{
				SubmitButtonText:  "Add existing personn as mentor",
				SubmitButtonClass: "button",
				FormID:            "addMentor",
				GetModel: func(form *Form, ctx *Context) (interface{}, error) {
					return &addMentorModel{}, nil
				},
				OnSubmit: addMentor,
			},
			HR(),
			&Form{
				SubmitButtonText:  "Add new person as mentor",
				SubmitButtonClass: "button",
				FormID:            "createMentor",
				OnSubmit:          createMentor,
			},
			// HR(),
			// H1("Mentor Bookings"),
			// &ModelIteratorView{
			// 	GetModel: func(ctx *Context) (interface{}, error) {
			// 		return new(models.Person), nil
			// 	},
			// 	GetModelIterator: EventMentorIterator,
			// 	GetModelView:     eventAdminMentorBookingView,
			// },
		},
	}
}

func eventMentorRowViews(row int, rowModel interface{}, ctx *Context) (views Views, err error) {
	person := rowModel.(*models.Person)

	editURL := Region0_Event1_Admin_Mentor2.URL(ctx.ForURLArgs(ctx.URLArgs[0], ctx.URLArgs[1], person.ID.Hex()))

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
				event.Mentors, _ = mongo.RemoveRefWithIDFromSlice(event.Mentors, person.ID)
				return "", StringURL("."), event.Save()
			},
		},
	}
	return views, nil
}

type addMentorModel struct {
	Name model.String `view:"class=add-mentor|placeholder=firstname lastname"`
}

func addMentor(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
	event := ctx.Data.(*PageData).Event
	name := formModel.(*addMentorModel).Name.Get()
	i := models.People.Iterator()
	var person models.Person
	for i.Next(&person) {
		if name == person.Name.String() {
			person.Mentor = true
			err := person.Save()
			if err != nil {
				return "", nil, err
			}
			event.Mentors = append(event.Mentors, person.Ref())
			return "", StringURL("."), event.Save()
		}
	}
	if i.Err() != nil {
		return "", nil, i.Err()
	}

	return "", nil, fmt.Errorf("Person '%s' not found", name)

	// person := models.People.NewDocument().(*models.Person)

	// namesplit := strings.SplitAfter(name, " ")
	// if len(namesplit) == 1 {
	// 	person.Name.First.Set(namesplit[0])
	// } else if len(namesplit) > 1 {
	// 	person.Name.First.Set(namesplit[0])
	// 	person.Name.Last.Set(namesplit[1])
	// }

	// person.Mentor = true
	// err := person.Save()
	// if err != nil {
	// 	return "", nil, err
	// }

	// event.Mentors = append(event.Mentors, person.Ref())
	// err = event.Save()
	// if err != nil {
	// 	person.Remove()
	// 	return "", nil, err
	// }

	// return "", StringURL(Region0_Event1_Admin_Mentor2.URL(
	// 	ctx.ForURLArgs(ctx.URLArgs[0], ctx.URLArgs[1], person.ID.Hex()))), nil
}

func createMentor(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
	event := ctx.Data.(*PageData).Event

	person := models.NewPerson()
	person.Name.First = "[First]"
	person.Name.Last = "[Last]"
	person.Mentor = true
	err := person.Save()
	if err != nil {
		return "", nil, err
	}

	event.Mentors = append(event.Mentors, person.Ref())
	err = event.Save()
	if err != nil {
		person.Delete()
		return "", nil, err
	}

	return "", StringURL(Region0_Event1_Admin_Mentor2.URL(ctx.ForURLArgs(ctx.URLArgs[0], ctx.URLArgs[1], person.ID.Hex()))), nil
}

func eventAdminMentorBookingView(ctx *Context, model interface{}) (view View, err error) {
	person := model.(*models.Person)

	i := models.EventTeams.Filter("BookedMentors", person.ID).Iterator()

	teamNames := make([]string, 0, 32)
	var team models.EventTeam
	for i.Next(&team) {
		teamNames = append(teamNames, team.Name.Get())
	}
	if i.Err() != nil {
		return nil, i.Err()
	}

	result := Views{
		H3(person.Name.String()),
		&List{Model: EscapeStringsListModel(teamNames)},
	}
	return result, nil
}
