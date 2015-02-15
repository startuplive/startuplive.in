package admin

import (
	"fmt"
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/STARTeurope/startuplive.in/views/admin"
	"github.com/ungerik/go-start/model"
	"github.com/ungerik/go-start/mongo"
	. "github.com/ungerik/go-start/view"
)

func init() {
	Region0_Event1_Admin_Organisers = &Page{
		OnPreRender: SetEventPageData,
		Title:       EventAdminTitle("Organisers"),
		CSS:         IndirectURL(&Region0_DashboardCSS),
		PostCSS:     StylesheetLink("/css/ui-lightness/jquery-ui-1.8.17.custom.css"),
		Scripts: Renderers{
			admin.PageScripts,
			JQueryUI,
			JQueryUIAutocompleteFromURL(".add-organiser", IndirectURL(&API_People), 2),
		},
		Content: Views{
			eventadminHeader(),
			DynamicView(
				func(ctx *Context) (View, error) {
					event := ctx.Data.(*PageData).Event

					var region models.EventRegion
					err := event.Region.Get(&region)
					if err != nil {
						return nil, err
					}

					list := region.Slug.Get() + "" + event.Number.String()
					return Escape("Your Mailing List: " + list + "@startuplive.in"), nil
				},
			),
			HR(),
			&ModelIteratorTableView{
				Class:            "visual-table",
				GetModelIterator: EventOrganiserIterator,
				GetRowModel: func(ctx *Context) (interface{}, error) {
					return new(models.Person), nil
				},
				GetHeaderRowViews: TableHeaderRowEscape("Nr", "Name", "Email", "Company", "Position", "Edit", "Make Lead", "Remove"),
				GetRowViews: func(row int, rowModel interface{}, ctx *Context) (views Views, err error) {
					person := rowModel.(*models.Person)
					editURL := Region0_Event1_Admin_Organiser2.URL(ctx.ForURLArgs(ctx.URLArgs[0], ctx.URLArgs[1], person.ID.Hex()))
					views = Views{
						Printf("%d", row+1),
						Escape(person.Name.String()),
						Escape(person.OrganiserEmail.String()),
						Escape(person.Company.Get()),
						Escape(person.Position.Get()),
						A(editURL, "Edit"),
						&If{
							Condition: row > 0,
							Content: &Form{
								SubmitButtonText:  "Make Lead",
								SubmitButtonClass: "",
								FormID:            "makelead" + person.ID.Hex(),
								OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
									event := ctx.Data.(*PageData).Event
									err := event.MakeLeadOrganiser(person)
									if err != nil {
										return "", nil, err
									}
									return "", StringURL("."), event.Save()
								},
							},
						},
						&Form{
							SubmitButtonText:  "Remove",
							SubmitButtonClass: "",
							FormID:            "remove" + person.ID.Hex(),
							OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
								event := ctx.Data.(*PageData).Event
								event.Organisers, _ = mongo.RemoveRefWithIDFromSlice(event.Organisers, person.ID)
								return "", StringURL("."), event.Save()
							},
						},
					}
					return views, nil
				},
			},
			HR(),
			&Form{
				SubmitButtonText:  "Add existing person as organiser",
				SubmitButtonClass: "button",
				FormID:            "addOrganiser",
				GetModel: func(form *Form, ctx *Context) (interface{}, error) {
					return &addOrganiserModel{}, nil
				},
				OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
					name := formModel.(*addOrganiserModel).Name.Get()
					i := models.People.Iterator()
					var person models.Person
					for i.Next(&person) {
						if name == person.Name.String() {
							person.EventOrganiser = true
							err := person.Save()
							if err != nil {
								return "", nil, err
							}
							event := ctx.Data.(*PageData).Event
							event.AddOrganiser(&person)
							return "", StringURL("."), event.Save()
						}
					}
					if i.Err() != nil {
						return "", nil, i.Err()
					}
					return "", nil, fmt.Errorf("Person '%s' not found", name)
				},
			},
			HR(),
			&Form{
				SubmitButtonText:  "Add new person as organiser",
				SubmitButtonClass: "button",
				FormID:            "createOrganiser",
				OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
					event := ctx.Data.(*PageData).Event

					person := models.NewPerson()
					person.Name.First = "[First]"
					person.Name.Last = "[Last]"
					person.EventOrganiser = true
					err := person.Save()
					if err != nil {
						return "", nil, err
					}

					event.Organisers = append(event.Organisers, person.Ref())
					err = event.Save()
					if err != nil {
						person.Delete()
						return "", nil, err
					}

					return "", StringURL(Region0_Event1_Admin_Organiser2.URL(ctx.ForURLArgs(ctx.URLArgs[0], ctx.URLArgs[1], person.ID.Hex()))), nil
				},
			},
		},
	}
}

type addOrganiserModel struct {
	Name model.String `view:"class=add-organiser"`
}
