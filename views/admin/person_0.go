package admin

import (
	"labix.org/v2/mgo/bson"

	"github.com/ungerik/go-start/debug"
	"github.com/ungerik/go-start/model"
	"github.com/ungerik/go-start/user"
	. "github.com/ungerik/go-start/view"

	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
)

func init() {
	Admin_Person0 = &Page{
		Title: Render(
			func(ctx *Context) error {
				person, err := getPerson(ctx)
				if err != nil {
					return err
				}
				ctx.Response.WriteString(person.Name.String())
				ctx.Response.WriteString(" | Admin")
				return nil
			},
		),
		CSS:     IndirectURL(&Admin_CSS),
		Scripts: PageScripts,
		Content: Views{
			adminHeader(),
			DynamicView(
				func(ctx *Context) (view View, err error) {
					person, err := getPerson(ctx)
					if err != nil {
						return nil, err
					}
					debug.Nop()

					var admin *models.Person
					_, err = user.OfSession(ctx.Session, &admin)
					if err != nil {
						return nil, err
					}

					personForm := PersonForm(person, Admin_People, nil, nil, nil)
					if admin.SuperAdmin == false {
						personForm.ExcludedFields = []string{"Admin", "SuperAdmin"}
					}

					views := Views{
						H2(person.Name.String()),
						Printf("Email confirmed: %v", person.EmailPasswordConfirmed()),
						HR(),
						&Form{
							SubmitButtonText:  "Save password and mark email as confirmed",
							SubmitButtonClass: "button",
							FormID:            "password",
							DisabledFields:    []string{"Current_password"},
							GetModel: func(form *Form, ctx *Context) (interface{}, error) {
								person, err := getPerson(ctx)
								if err != nil {
									return nil, err
								}
								return &passwordFormModel{Current_password: person.Password}, nil
							},
							OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
								m := formModel.(*passwordFormModel)
								person, err := getPerson(ctx)
								if err != nil {
									return "", nil, err
								}
								person.Password.SetHashed(m.New_password.Get())
								person.ConfirmEmailPassword()
								err = person.Save()
								if err != nil {
									return "", nil, err
								}
								// if m.Send_notification_email.Get() {
								// }
								return "", StringURL("."), nil
							},
						},
						HR(),
						personForm,
					}
					return views, nil
				},
			),
		},
	}
}

/// todo Erik use person closure instead?
func getPerson(ctx *Context) (person *models.Person, err error) {
	id := bson.ObjectIdHex(ctx.URLArgs[0])
	found, err := models.People.TryDocumentWithID(id, &person)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, NotFound("404: Person not found")
	}
	return person, nil
}

type passwordFormModel struct {
	Current_password model.Password
	New_password     model.String `model:"minlen=6"`
	//Send_notification_email model.Bool
}
