package admin

import (
	"github.com/AlexTi/go-amiando"
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/STARTeurope/startuplive.in/views/admin"
	"github.com/ungerik/go-start/debug"
	"github.com/ungerik/go-start/model"
	. "github.com/ungerik/go-start/view"
)

func init() {
	Region0_Event1_Admin_Amiando = &Page{
		OnPreRender: SetEventPageData,
		Title:       EventAdminTitle("Amiando Setup"),
		CSS:         IndirectURL(&Region0_DashboardCSS),
		Scripts:     admin.PageScripts,
		Content: Views{
			eventadminHeader(),
			DynamicView(amiandoSyncView),
		},
	}
}

func amiandoSyncView(ctx *Context) (view View, err error) {
	event := ctx.Data.(*PageData).Event
	region := ctx.Data.(*PageData).Region
	var views Views
	views = Views{
		H2("Amiando Account and Event Setup"),
	}

	nextTodo := false

	//=== Step 1: Create the event - we assume it is otherwise we would not see that page ===
	view = &ShortTag{Tag: "h4", Content: HTML("1. Create Event on the platform")}
	view.(*ShortTag).Class = "todo-done"
	views = append(views, view)

	//=== Step 2: Add a Lead Organiser ===
	//TODO: on organiser creation and on "make lead" - add organiser forwarding email address to cityNrAdmin email
	view = &ShortTag{Tag: "h4", Content: HTML("2. Add Lead Organiser")}
	views = append(views, view)
	if event.HasLeadOrganiser() {
		view.(*ShortTag).Class = "todo-done"
	} else if !nextTodo {
		view.(*ShortTag).Class = "todo-next"
		nextTodo = true
	}

	//=== Step 3: Create an API Key ===
	view = &ShortTag{Tag: "h4", Content: HTML("3. Create API Key")}
	views = append(views, view)
	if event.AmiandoEventApiKey != "" {
		view.(*ShortTag).Class = "todo-done"
	} else {
		if !nextTodo {
			view.(*ShortTag).Class = "todo-next"
			nextTodo = true

			views = append(views, &Form{
				SubmitButtonText:  "Create API Key",
				SubmitButtonClass: "button",
				FormID:            "createapikey",
				OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
					keyname := region.Slug.String() + string(event.Number)
					key, err := amiando.CreateAPIKey(keyname)
					if err != nil {
						return "", nil, err
					}

					event.AmiandoEventApiKey.Set(key)

					return "", StringURL("."), event.Save()
				},
			})
		}
	}

	//=== Step 4: Create an amiando account ===
	view = &ShortTag{Tag: "h4", Content: HTML("4. Create an amiando account")}
	views = append(views, view)
	if event.AmiandoHostId != "" {
		view.(*ShortTag).Class = "todo-done"
	} else {
		if !nextTodo {
			view.(*ShortTag).Class = "todo-next"
			nextTodo = true

			views = append(views, &Form{
				SubmitButtonText:  "Create Amiando Account",
				SubmitButtonClass: "button",
				FormID:            "createamiandoaccount",
				GetModel: func(form *Form, ctx *Context) (interface{}, error) {
					var firstname model.String
					var lastname model.String

					debug.Print(event.GetLeadOrganiser())
					lead, err := event.GetLeadOrganiser()
					if err != nil {
						firstname = ""
						lastname = ""
					} else {
						firstname = lead.Name.First
						lastname = lead.Name.Last
					}

					pw := region.Slug.String() + event.Number.String() + "3gy"
					var password model.String
					password.Set(pw)

					return &amiandoAccount{
						Username:  event.AmiandoAccEmail,
						FirstName: firstname,
						LastName:  lastname,
						Password:  password,
					}, nil
				},
				OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {

					m := formModel.(*amiandoAccount)

					api := amiando.NewApi(event.AmiandoEventApiKey.Get())

					hostIds, err := api.HostId(m.Username.String())
					if err != nil {
						debug.Print("err ", err)
						return "", nil, err
					}
					var userid string
					if len(hostIds) != 0 {
						userid = hostIds[0].String()
						debug.Print("found user: ", userid)
					} else {
						userid, err = api.CreateAmiandoUser(m.FirstName.String(), m.LastName.String(),
							m.Username.String(), m.Password.String(), "en")

						if err != nil {
							debug.Print("err ", err)
							return "", nil, err
						}
						debug.Print("created new user: ", userid)
					}

					event.AmiandoHostId.Set(userid)
					event.AmiandoHostPassword = m.Password
					err = event.Save()
					if err != nil {
						return "", nil, err
					}
					return "", StringURL("."), nil

				},
			})
		}
	}

	//=== Step 5: Create Event on Amiando ===
	view = &ShortTag{Tag: "h4", Content: HTML("5. Create Event on Amiando")}
	views = append(views, view)
	if event.AmiandoEventId != "" {
		view.(*ShortTag).Class = "todo-done"
	} else {
		if !nextTodo {
			view.(*ShortTag).Class = "todo-next"
			nextTodo = true

			views = append(views, &Form{
				SubmitButtonText:  "Create Amiando Event",
				SubmitButtonClass: "button",
				FormID:            "createamiandoevent",

				GetModel: func(form *Form, ctx *Context) (interface{}, error) {
					var identifier model.String
					identifier.Set("startuplive" + region.Slug.String() + "" + event.Number.String())
					return &amiandoCreateEventFormModel{
						Username:   event.AmiandoAccEmail,
						Title:      event.Name,
						Identifier: identifier,
						HostId:     event.AmiandoHostId,
						Date:       event.Start,
					}, nil
				},
				OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
					m := formModel.(*amiandoCreateEventFormModel)

					title := m.Title.String()
					id := m.Identifier.String()
					hostId := m.HostId.String()
					country := m.Country.String()
					date := m.Date.String()

					// create event on amiando
					// func (self *Api) CreateEvent(hostId string, title string, country string, date string) (err error) {
					api := amiando.NewApi(event.AmiandoEventApiKey.Get())
					eventid, err := api.CreateEvent(hostId, title, country, date, id)
					if err != nil {
						return "", nil, err
					}

					event.AmiandoEventId.Set(eventid)
					event.AmiandoEventIdentifier.Set(id)

					return "", StringURL("."), event.Save()
				},
			})
		}
	}

	//=== Step 6: Create Ticket Categories for Amiando Event ===
	view = &ShortTag{Tag: "h4", Content: HTML("6. Create Ticket Categories for the Event")}
	views = append(views, view)
	if event.AmiandoTicketCategoriesAdded == true {
		view.(*ShortTag).Class = "todo-done"
	} else {
		if !nextTodo {
			view.(*ShortTag).Class = "todo-next"
			nextTodo = true

			views = append(views, &Form{
				SubmitButtonText:  "Create Ticket Categories",
				SubmitButtonClass: "button",
				FormID:            "createamiandoevent",

				GetModel: func(form *Form, ctx *Context) (interface{}, error) {
					if len(event.AmiandoTicketCategories) == 0 {
						return &SetupAmiandoTicketCategories{
							TicketCategories: []models.AmiandoTicketCategory{
								{Name: "student (early bird)"},
								{Name: "normal (early bird)"},
								{Name: "student"},
								{Name: "normal"},
								{Name: "student (late bird)"},
								{Name: "normal (late bird)"},
							},
						}, nil
					}
					return &SetupAmiandoTicketCategories{
						TicketCategories: event.AmiandoTicketCategories,
					}, nil
				},
				OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
					m := formModel.(*SetupAmiandoTicketCategories)

					for _, v := range m.TicketCategories {

						price := v.Price.String() + "00"
						api := amiando.NewApi(event.AmiandoEventApiKey.Get())
						_, err := api.CreateTicketCategory(event.AmiandoEventId.String(), v.Name.String(), price, v.Available.String(), v.SaleStart.String(), v.SaleEnd.String())
						if err != nil {
							return "", nil, err
						}

					}
					// debug.Print("ticketCategories created")
					event.AmiandoTicketCategories = m.TicketCategories
					event.AmiandoTicketCategoriesAdded = true

					return "", StringURL("."), event.Save()
				},
			})
		}
	}

	//=== Step 7: Add Payment information to the event ===
	view = &ShortTag{Tag: "h4", Content: HTML("7. Add payment information directly on amiando by hand (no API call possible).")}
	views = append(views, view)
	if event.AmiandoPaymentAdded == true {
		view.(*ShortTag).Class = "todo-done"
	} else {
		if !nextTodo {
			view.(*ShortTag).Class = "todo-next"
			nextTodo = true

			view = &TextField{Text: event.AmiandoPaymentData.AccountHolder.String(), Readonly: true}
			views = append(views, &Label{For: view, Content: Printf("Account Holder: "), Class: "amiando-label"})
			views = append(views, view)
			views = append(views, BR())

			view = &TextField{Text: event.AmiandoPaymentData.BankName.String(), Readonly: true}
			views = append(views, &Label{For: view, Content: Printf("Bank Name: "), Class: "amiando-label"})
			views = append(views, view)
			views = append(views, BR())

			view = &TextField{Text: event.AmiandoPaymentData.Swift.String(), Readonly: true}
			views = append(views, &Label{For: view, Content: Printf("SWIFT code: "), Class: "amiando-label"})
			views = append(views, view)
			views = append(views, BR())

			view = &TextField{Text: event.AmiandoPaymentData.Iban.String(), Readonly: true}
			views = append(views, &Label{For: view, Content: Printf("IBAN code: "), Class: "amiando-label"})
			views = append(views, view)
			views = append(views, BR())

			view = &TextField{Text: event.AmiandoPaymentData.Country.String(), Readonly: true}
			views = append(views, &Label{For: view, Content: Printf("Country: "), Class: "amiando-label"})
			views = append(views, view)
			views = append(views, BR())
			views = append(views, BR())

			views = append(views, &Form{

				SubmitButtonText:    "Yes! I swear I have added *PAYMENT DATA* on Amiando by hand.",
				SubmitButtonClass:   "button",
				SubmitButtonConfirm: "Seriously, are you sure?",
				FormID:              "createpaymentdata",
				OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
					event.AmiandoPaymentAdded = true
					return "", StringURL("."), event.Save()
				},
			})

		}
	}

	//=== Step 8: Add Billing information to the event ===
	view = &ShortTag{Tag: "h4", Content: HTML("8. Add billing information directly on amiando by hand (no API call possible).")}
	views = append(views, view)
	if event.AmiandoBillingAdded == true {
		view.(*ShortTag).Class = "todo-done"
	} else {
		if !nextTodo {
			view.(*ShortTag).Class = "todo-next"
			nextTodo = true

			view = &TextField{Text: event.AmiandoBillingData.FirstName.String(), Readonly: true}
			views = append(views, &Label{For: view, Content: Printf("First Name: "), Class: "amiando-label"})
			views = append(views, view)
			views = append(views, BR())

			view = &TextField{Text: event.AmiandoBillingData.LastName.String(), Readonly: true}
			views = append(views, &Label{For: view, Content: Printf("Last Name: "), Class: "amiando-label"})
			views = append(views, view)
			views = append(views, BR())

			view = &TextField{Text: event.AmiandoBillingData.Company.String(), Readonly: true}
			views = append(views, &Label{For: view, Content: Printf("Comapny: "), Class: "amiando-label"})
			views = append(views, view)
			views = append(views, BR())

			view = &TextField{Text: event.AmiandoBillingData.Street.String(), Readonly: true}
			views = append(views, &Label{For: view, Content: Printf("Street: "), Class: "amiando-label"})
			views = append(views, view)
			views = append(views, BR())

			view = &TextField{Text: event.AmiandoBillingData.Zip.String(), Readonly: true}
			views = append(views, &Label{For: view, Content: Printf("Zip: "), Class: "amiando-label"})
			views = append(views, view)
			views = append(views, BR())

			view = &TextField{Text: event.AmiandoBillingData.City.String(), Readonly: true}
			views = append(views, &Label{For: view, Content: Printf("City: "), Class: "amiando-label"})
			views = append(views, view)
			views = append(views, BR())

			view = &TextField{Text: event.AmiandoBillingData.Country.String(), Readonly: true}
			views = append(views, &Label{For: view, Content: Printf("Country: "), Class: "amiando-label"})
			views = append(views, view)
			views = append(views, BR())
			views = append(views, BR())

			views = append(views, &Form{
				SubmitButtonText:    "Yes! I swear I have added *BILLING DATA* on Amiando by hand.",
				SubmitButtonClass:   "button",
				SubmitButtonConfirm: "Seriously, are you sure?",

				FormID: "createbillingdata",
				OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
					event.AmiandoBillingAdded = true
					return "", StringURL("."), event.Save()
				},
			})

		}
	}

	//=== Step 9: Add Promo Codes to the event ===
	view = &ShortTag{Tag: "h4", Content: HTML("9. Add Promo Code 'kredito49191' (20% discount) directly on amiando by hand (no API call possible).")}
	views = append(views, view)
	if event.AmiandoPromoCodeAdded == true {
		view.(*ShortTag).Class = "todo-done"
	} else {
		if !nextTodo {
			view.(*ShortTag).Class = "todo-next"
			nextTodo = true

			views = append(views, HTML("Add Promo Code 'kredito49191' to the event -> Price Reduction 20% -> valid for every category and the whole presale period."))
			views = append(views, BR())

			views = append(views, &Form{
				SubmitButtonText:    "Oh boy, added the promo code. Nice one!",
				SubmitButtonClass:   "button",
				SubmitButtonConfirm: "Keep up the fantastic work. Nearly done. Promocode really added?",
				FormID:              "promocode",
				OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
					event.AmiandoPromoCodeAdded = true
					return "", StringURL("."), event.Save()
				},
			})

		}
	}

	//=== Step 9: Add ticket attendee data form fields to the event ===
	view = &ShortTag{Tag: "h4", Content: HTML("10. Add ticket attendee data form fields to the event.")}
	views = append(views, view)
	if event.AmiandoAttendeeDataAdded == true {
		view.(*ShortTag).Class = "todo-done"
	} else {
		if !nextTodo {
			view.(*ShortTag).Class = "todo-next"
			nextTodo = true

			views = append(views, &Span{Content: Printf("Type"), Class: "amiando-label"})
			views = append(views, &Span{Content: Printf("Name of Fields")})
			views = append(views, BR())
			view = &TextField{Text: "Date of birth", Readonly: true}
			views = append(views, &Label{For: view, Content: Printf("Date of birth"), Class: "amiando-label"})
			views = append(views, view)
			views = append(views, BR())
			view = &TextField{Text: "Gender", Readonly: true}
			views = append(views, &Label{For: view, Content: Printf("Gender"), Class: "amiando-label"})
			views = append(views, view)
			views = append(views, BR())
			view = &TextField{Text: "Citizenship", Readonly: true}
			views = append(views, &Label{For: view, Content: Printf("Country"), Class: "amiando-label"})
			views = append(views, view)
			views = append(views, BR())
			view = &TextField{Text: "Email", Readonly: true}
			views = append(views, &Label{For: view, Content: Printf("Email field"), Class: "amiando-label"})
			views = append(views, view)
			views = append(views, BR())
			view = &TextField{Text: "Phone number", Readonly: true}
			views = append(views, &Label{For: view, Content: Printf("Telephone field"), Class: "amiando-label"})
			views = append(views, view)
			views = append(views, BR())
			view = &TextField{Text: "Address", Readonly: true}
			views = append(views, &Label{For: view, Content: Printf("Address field"), Class: "amiando-label"})
			views = append(views, view)
			views = append(views, BR())
			view = &TextField{Text: "University", Readonly: true}
			views = append(views, &Label{For: view, Content: Printf("Text field"), Class: "amiando-label"})
			views = append(views, view)
			views = append(views, BR())
			view = &TextField{Text: "Background", Readonly: true}
			views = append(views, &Label{For: view, Content: Printf("Dropdown"), Class: "amiando-label"})
			views = append(views, view)
			views = append(views, BR())
			view = &TextField{Text: "Business", Readonly: true}
			views = append(views, &Label{For: view, Content: Printf(""), Class: "amiando-label"})
			views = append(views, view)
			views = append(views, BR())
			view = &TextField{Text: "Design", Readonly: true}
			views = append(views, &Label{For: view, Content: Printf(""), Class: "amiando-label"})
			views = append(views, view)
			views = append(views, BR())
			view = &TextField{Text: "Tech", Readonly: true}
			views = append(views, &Label{For: view, Content: Printf(""), Class: "amiando-label"})
			views = append(views, view)
			views = append(views, BR())
			view = &TextField{Text: "Other", Readonly: true}
			views = append(views, &Label{For: view, Content: Printf(""), Class: "amiando-label"})
			views = append(views, view)
			views = append(views, BR())
			view = &TextField{Text: "Other", Readonly: true}
			views = append(views, &Label{For: view, Content: Printf("Subfield of Background & Textfield"), Class: "amiando-label"})
			views = append(views, view)
			views = append(views, BR())
			view = &TextField{Text: "Will you present an idea?", Readonly: true}
			views = append(views, &Label{For: view, Content: Printf("Radio Buttons"), Class: "amiando-label"})
			views = append(views, view)
			views = append(views, BR())
			view = &TextField{Text: "Yes", Readonly: true}
			views = append(views, &Label{For: view, Content: Printf(""), Class: "amiando-label"})
			views = append(views, view)
			views = append(views, BR())
			view = &TextField{Text: "No", Readonly: true}
			views = append(views, &Label{For: view, Content: Printf(""), Class: "amiando-label"})
			views = append(views, view)
			views = append(views, BR())
			view = &TextField{Text: "Do you need an accommodation?", Readonly: true}
			views = append(views, &Label{For: view, Content: Printf("Radio Buttons"), Class: "amiando-label"})
			views = append(views, view)
			views = append(views, BR())
			view = &TextField{Text: "Yes", Readonly: true}
			views = append(views, &Label{For: view, Content: Printf(""), Class: "amiando-label"})
			views = append(views, view)
			views = append(views, BR())
			view = &TextField{Text: "No", Readonly: true}
			views = append(views, &Label{For: view, Content: Printf(""), Class: "amiando-label"})
			views = append(views, view)
			views = append(views, BR())

			views = append(views, &Form{
				SubmitButtonText:    "Puh, added all the attendee data form fields!",
				SubmitButtonClass:   "button",
				SubmitButtonConfirm: "Next step - infinity. Ticket Data added?",
				FormID:              "ticketdate",
				OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
					event.AmiandoAttendeeDataAdded = true
					return "", StringURL("."), event.Save()
				},
			})

		}
	}

	view = &ShortTag{Tag: "h4", Content: HTML("11. Activate Event")}
	views = append(views, view)
	if event.AmiandoEventActivated == true {
		view.(*ShortTag).Class = "todo-done"
	} else {
		if !nextTodo {
			view.(*ShortTag).Class = "todo-next"
			nextTodo = true

			views = append(views, &Form{
				SubmitButtonText:    "Go Go Gadgeto EVENT ACTIVATION!",
				SubmitButtonClass:   "button",
				SubmitButtonConfirm: "Yes, do it!",
				FormID:              "eventactivation",
				OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {

					api := amiando.NewApi(event.AmiandoEventApiKey.Get())
					err := api.ActivateEvent(event.AmiandoEventId.String())
					if err != nil {
						return "", nil, err
					}
					event.AmiandoEventActivated = true

					// iframecode := `<script type="text/javascript" src="https://www.amiando.com/resources/js/amiandoExport.js"></script>
					// 					<iframe src="https://www.amiando.com/` + event.AmiandoEventIdentifier.String() + `.html?viewType=iframe&distributionChannel=CHANNEL_IFRAME&panelId=1656776&useDefaults=false&resizeIFrame=true" frameborder="0" width="650px" height="450px" name="_amiandoIFrame1656776u1Nsxmf1" id="_amiandoIFrame1656776u1Nsxmf1">
					// 					<p>This page requires frame support. Please use a frame compatible browser to see the amiando ticket sales widget.</p>
					// 					<p> Try out the <a href="http://www.amiando.com/">online event registration system</a> from amiando.</p>
					// 					</iframe><p style="text-align: left; font-size:10px;">
					// 					<a href="http://www.amiando.com" target="_blank" alt="Networking event - Online Event Management" title="Networking event - Online Event Management" >Networking event - Online Event Management</a> with the ticketing solution from amiando</p>`

					// event.AmiandoIframeCode.Set(iframecode)

					return "", StringURL("."), event.Save()
				},
			})
		}
	}

	view = &ShortTag{Tag: "h4", Content: &Views{
		HTML("You are nearly done - insert the iframe code from amiando and then enjoy!"),
		HTML("Enter Server Callback on Amiando > Integration > Ticket shop callback > Server Callback with this url: http://startuplive.in/" + region.Slug.String() + "/" + event.Number.String() + "/amiando-tracking"),
	}}
	if !nextTodo {
		view = &ShortTag{Tag: "h3", Content: &Views{
			HTML("You are nearly done - insert the iframe code from amiando and then enjoy!"),
			HTML("Enter Server Callback on Amiando > Integration > Ticket shop callback > Server Callback with this url: http://startuplive.in/" + region.Slug.String() + "/" + event.Number.String() + "/amiando-tracking"),
		}}
		view.(*ShortTag).Class = "todo-next"
	}
	views = append(views, view)

	return views, err
}

// Amiando User Account
// we set the language to english
type amiandoAccount struct {
	Username  model.Email  `model:"required"`
	FirstName model.String `model:"required"`
	LastName  model.String `model:"required"`
	Password  model.String `model:"required"`
	// Language	model.Choice `model:"required|options=en,de,fr,es"`
}

type amiandoCreateEventFormModel struct {
	Username   model.Email
	Title      model.String   `model:"required"`
	Identifier model.String   `model:"required" view:"label=Event Identifier"`
	HostId     model.String   `model:"required"`
	Country    model.String   `model:"required"`
	Date       model.DateTime `model:"required"`
}

type SetupAmiandoTicketCategories struct {
	TicketCategories []models.AmiandoTicketCategory
}
