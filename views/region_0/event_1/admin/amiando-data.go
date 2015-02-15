package admin

import (
	// "github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/STARTeurope/startuplive.in/views/admin"
	// 	"github.com/ungerik/go-start/debug"
	"github.com/ungerik/go-start/model"
	// 	"github.com/AlexTi/go-amiando"
	. "github.com/ungerik/go-start/view"
)

func init() {
	Region0_Event1_Admin_AmiandoData = &Page{
		OnPreRender: SetEventPageData,
		Title:       EventAdminTitle("Amiando Data"),
		CSS:         IndirectURL(&Region0_DashboardCSS),
		Scripts:     admin.PageScripts,
		Content: Views{
			eventadminHeader(),
			DynamicView(amiandoDataView),
		},
	}
}

func amiandoDataView(ctx *Context) (view View, err error) {
	event := ctx.Data.(*PageData).Event
	// region := ctx.Data.(*PageData).Region
	var views Views
	views = Views{
		H2("Information we need in order to setup Amiando"),
		HTML(`<ol><li>Choose your Ticket Categories for the Event</li>
			<li>Add payment information</li>
			<li>Add billing information</li>
			<li>Add Welcome Text for Pitchers</li></ol>`),
		HR(),
	}

	//=== Add Ticket Categories for Amiando Event ===
	view = &ShortTag{Tag: "h3", Content: HTML("1. Choose your Ticket Categories for the Event")}
	views = append(views, view)
	views = append(views, &Form{
		SubmitButtonText:  "Set Ticket Categories",
		SubmitButtonClass: "button",
		FormID:            "createamiandoevent",
		GetModel: func(form *Form, ctx *Context) (interface{}, error) {
			// if len(event.AmiandoTicketCategories) == 0 {
			// 	return &SetupAmiandoTicketCategories{
			// 		TicketCategories: []models.AmiandoTicketCategory{
			// 			{Name: "student (early bird)"},
			// 			{Name: "normal (early bird)"},
			// 			{Name: "student"},
			// 			{Name: "normal"},
			// 			{Name: "student (late bird)"},
			// 			{Name: "normal (late bird)"},
			// 		},
			// 	}, nil
			// }
			return &SetupAmiandoTicketCategories{
				TicketCategories: event.AmiandoTicketCategories,
			}, nil

		},
		OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
			m := formModel.(*SetupAmiandoTicketCategories)
			// theevent := ctx.Data.(*PageData).Event
			// debug.Print("ticketCategories created")
			// var cats []models.AmiandoTicketCategory
			// cats = m.TicketCategories
			event.AmiandoTicketCategories = m.TicketCategories

			return "", StringURL("."), event.Save()
		},
	})

	//=== Add Payment Information for Amiando Event ===
	view = &ShortTag{Tag: "h3", Content: HTML("2. Add payment information")}
	views = append(views, HR())
	views = append(views, view)
	views = append(views, &Form{
		SubmitButtonText:  "Add Payment Information",
		SubmitButtonClass: "button",
		RequiredFields: []string{
			"AccountHolder",
			"BankName",
			"Swift",
			"Iban",
			"Country",
		},
		FormID: "createpaymentinfo",
		GetModel: func(form *Form, ctx *Context) (interface{}, error) {
			return &event.AmiandoPaymentData, nil
		},
		OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
			return "", StringURL("."), event.Save()
		},
	})

	//=== Add Billing Information for Amiando Event ===
	view = &ShortTag{Tag: "h3", Content: HTML("3. Add billing information")}
	views = append(views, HR())
	views = append(views, view)
	views = append(views, &Form{
		SubmitButtonText:  "Add Billing Information",
		SubmitButtonClass: "button",
		RequiredFields: []string{
			"FirstName",
			"LastName",
			"Company",
			"Street",
			"Zip",
			"City",
			"Country",
		},
		FormID: "createbillinginfo",
		GetModel: func(form *Form, ctx *Context) (interface{}, error) {
			return &event.AmiandoBillingData, nil
		},
		OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
			return "", StringURL("."), event.Save()
		},
	})

	//=== Add Welcome Text for Pitcher ===
	views = append(views, HR())
	views = append(views, &ShortTag{Tag: "h3", Content: HTML("4. Add welcome text for pitcher")})
	views = append(views, &Form{
		Style:             "float:left; width:600px",
		SubmitButtonText:  "Update",
		SubmitButtonClass: "button",
		FormID:            "updatewelcometext",
		GetModel: func(form *Form, ctx *Context) (interface{}, error) {
			var txt model.Text

			if event.PitcherRegistrationWelcomeText.String() == "" {
				txt.Set(event.GetDefaultPitcherRegistrationText())
			} else {
				txt = event.PitcherRegistrationWelcomeText
			}
			return &PitcherRegistrationWelcomeForm{
				Text: txt,
			}, nil
		},
		OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
			m := formModel.(*PitcherRegistrationWelcomeForm)
			event.PitcherRegistrationWelcomeText = m.Text
			return "", StringURL("."), event.Save()
		},
	})
	views = append(views, &Div{Content: Views{
		H3("Description:"),
		P("On the left you see an email template which will be sent to pitchers after they bought a ticket."),
		P("The names in brackets i.e.: {{EventName}} gets automatically added by the system. So for Startup Live Vienna i.e.: {{EventName}} will be changed to 'Startup Live Vienna' "),
		P("Please do not remove the login data!"),
		H4("Template Markup:"),
		DIV("", HTML("<strong>{{EventName}}</strong> - inserts the name of you event")),
		DIV("", HTML("<strong>{{Date}}</strong> - inserts the date of the event")),
		DIV("", HTML("<strong>{{ParticipantFirstName}}</strong> - inserts the first name of the participant")),
		DIV("", HTML("<strong>{{ParticipantLastName}}</strong> - inserts the last name of the participant")),
		DIV("", HTML("<strong>{{ParticipantEmail}}</strong> - inserts the email of the participant")),
		DIV("", HTML("<strong>{{ParticipantPassword}}</strong> - inserts the temporary password of the participant")),
		DIV("", HTML("<strong>{{OrganiserName}}</strong> - inserts the name of the lead organiser")),
		DIV("", HTML("<strong>{{OrganiserEmail}}</strong> - inserts the email of the lead organiser")),
		DIV("", HTML("<strong>{{PitcherForm}}</strong> - inserts the 'pitcher form' url, every pitcher has to fill out")),
	},
	})
	views = append(views, DivClearBoth())

	return views, err
}

type PitcherRegistrationWelcomeForm struct {
	Text model.Text `view:"label=Text sent to pitcher after ticket purchase|rows=30"`
}
