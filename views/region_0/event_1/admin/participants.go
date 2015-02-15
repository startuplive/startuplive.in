package admin

import (
	"fmt"

	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/STARTeurope/startuplive.in/views/admin"
	"github.com/ungerik/go-start/debug"
	"github.com/ungerik/go-start/model"
	. "github.com/ungerik/go-start/view"
)

func init() {
	debug.Nop()

	Region0_Event1_Admin_Participants = &Page{
		OnPreRender: SetEventPageData,
		Title:       EventAdminTitle("Participants"),
		CSS:         IndirectURL(&Region0_DashboardCSS),
		PostCSS:     StylesheetLink("/css/ui-lightness/jquery-ui-1.8.17.custom.css"),
		Scripts: Renderers{
			admin.PageScripts,
			JQueryUI,
			JQueryUIAutocompleteFromURL(".add-participant", IndirectURL(&API_People), 2),
		},
		Content: Views{
			eventadminHeader(),
			&ModelIteratorTableView{
				Class:            "visual-table",
				GetModelIterator: EventParticipantIterator,
				GetRowModel: func(ctx *Context) (interface{}, error) {
					return new(models.EventParticipant), nil
				},
				GetHeaderRowViews: TableHeaderRowEscape("Nr", "Name", "Team", "Background", "Application Date", "Checked In", "Cancelled", "Idea", "Edit", "Remove"),
				GetRowViews: func(row int, rowModel interface{}, ctx *Context) (views Views, err error) {
					p := rowModel.(*models.EventParticipant)
					participant := p

					var teamName string
					var team models.EventTeam
					ok, err := participant.Team.TryGet(&team)
					if err != nil {
						return nil, err
					}
					if ok {
						teamName = team.Name.Get()
					}

					var present string
					if participant.PresentsIdea {
						present = "presents"
					}

					var checkedIn View
					if participant.CheckedIn {
						checkedIn = Escape("checked in")
					} //else {

					// 	checkedIn = &Form{
					// 		SubmitButtonText:  "Arrived",
					// 		SubmitButtonClass: "button",
					// 		FormID:            "checkinparticipant" + participant.ID.Hex(),
					// 		OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
					// 			// event := ctx.Data.(*PageData).Event
					// 			// debug.Print("p old: ", participant.ID.Hex())
					// 			// p, err := models.GetParticipantById(participant.ID)
					// 			// debug.Print("p new: ", p.ID.Hex())
					// 			// if err != nil {
					// 			// 	return "", StringURL("."), err
					// 			// }

					// 			// debug.Print("p checkedin: ", p.CheckedIn)

					// 			debug.Print("p checkedin: ", participant.CheckedIn)
					// 			debug.Print("p id: ", participant.ID.Hex())

					// 			p := participant
					// 			// p.Init(collection, embeddingStructPtr)
					// 			p.CheckedIn.Set(true)
					// 			return "", StringURL("."), p.Save()
					// 		},
					// 	}

					// }

					var cancelled string
					if participant.Cancelled() {
						cancelled = "cancelled"
					}

					editURL := Region0_Event1_Admin_Participant2.URL(
						ctx.ForURLArgs(ctx.URLArgs[0], ctx.URLArgs[1], participant.ID.Hex()))

					views = Views{
						Printf("%d", row+1),
						Escape(participant.Name()),
						Escape(teamName),
						Escape(participant.Background.Get()),
						Escape(participant.AppliedDate.Get()),
						checkedIn,
						Escape(cancelled),
						Escape(present),
						A(editURL, "Edit"),
						&Form{
							SubmitButtonText:    "Remove",
							SubmitButtonClass:   "delete-button",
							SubmitButtonConfirm: "Are you sure you want to delete " + participant.Name() + "?",
							FormID:              "remove" + participant.ID.Hex(),
							OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
								return "", StringURL("."), participant.Delete()
							},
						},
					}
					return views, nil
				},
			},
			HR(),
			&Form{
				SubmitButtonText:  "Add existing person as participant",
				SubmitButtonClass: "button",
				FormID:            "addParticipant",
				GetModel: func(form *Form, ctx *Context) (interface{}, error) {
					return &addParticipantModel{}, nil
				},
				OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
					name := formModel.(*addParticipantModel).Name.Get()
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

							eventParticipant := models.NewEventParticipant()
							_, err = models.EventParticipants.Filter("Event", event.ID).Filter("Person", person.ID).TryOneDocument(&eventParticipant)
							if err != nil {
								return "", nil, err
							}
							eventParticipant.Event.Set(event)
							eventParticipant.Person.Set(&person)
							err = eventParticipant.Save()
							if err != nil {
								return "", nil, err
							}
							return "", StringURL(Region0_Event1_Admin_Participant2.URL(ctx.ForURLArgs(ctx.URLArgs[0], ctx.URLArgs[1], eventParticipant.ID.Hex()))), nil
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
				SubmitButtonText:  "Add new person as participant",
				SubmitButtonClass: "button",
				FormID:            "createParticipant",
				OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
					event := ctx.Data.(*PageData).Event

					person := models.NewPerson()
					person.Name.First = "[First]"
					person.Name.Last = "[Last]"
					err := person.Save()
					if err != nil {
						return "", nil, err
					}

					eventParticipant := models.NewEventParticipant()
					eventParticipant.Event.Set(event)
					eventParticipant.Person.Set(person)
					err = eventParticipant.Save()
					if err != nil {
						return "", nil, err
					}

					return "", StringURL(Region0_Event1_Admin_Participant2.URL(ctx.ForURLArgs(ctx.URLArgs[0], ctx.URLArgs[1], eventParticipant.ID.Hex()))), nil
				},
			},
		},
	}
}

type addParticipantModel struct {
	Name model.String `view:"class=add-participant"`
}
