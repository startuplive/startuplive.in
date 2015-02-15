package region_0

import (
	"errors"
	"fmt"
	"path"
	"strings"

	"github.com/AlexTi/go-plesk"
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/STARTeurope/startuplive.in/views/admin"
	"github.com/ungerik/go-start/debug"
	"github.com/ungerik/go-start/model"
	. "github.com/ungerik/go-start/view"
)

func init() {
	Admin_Region0 = &Page{
		OnPreRender: SetRegionPageData,
		Title:       RegionTitle("Admin"),
		CSS:         IndirectURL(&Region0_DashboardCSS),
		Scripts:     admin.PageScripts,
		Content: Views{
			Header(),
			H2("General Settings"),
			&Form{
				FormID: "editregionname",
				GetModel: func(form *Form, ctx *Context) (interface{}, error) {
					return &editRegionnFormModel{}, nil
				},
				SubmitButtonText:  "Edit Region Name",
				SubmitButtonClass: "button",

				OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
					// oldslug := ctx.URLArgs[0]
					// region := models.EventRegions.Filter("Slug", oldslug)
					region, err := EventRegion(ctx.URLArgs)
					if err != nil {
						return "", nil, err
					}

					m := formModel.(*editRegionnFormModel)
					if m.Name == "" {
						return "", nil, errors.New("Empty Name")
					}

					slug := strings.ToLower(m.Name.Get())
					count, err := models.EventRegions.Filter("Slug", slug).Count()
					if err != nil {
						return "", nil, err
					}
					if count > 0 {
						return "", nil, fmt.Errorf("Event region with slug '%s' already exists", slug)
					}

					region.Name = m.Name
					region.Slug.Set(slug)
					// region.ColorScheme.Init(0, 0)
					return "", StringURL("../regions"), region.Save()
				},
			},
			DynamicView(deleteRegion),
			H2("Events in Region"),
			&ModelIteratorView{
				GetModel: func(ctx *Context) (interface{}, error) {
					return new(models.Event), nil
				},
				GetModelIterator: func(ctx *Context) model.Iterator {
					// return ctx.Data.(*PageData).Region.EventIterator(models.StartupLive)
					return ctx.Data.(*PageData).Region.RegionEventIterator()
				},
				GetModelView: eventRegionAdminView,
			},
			H2("Import Event From Amiando:"),
			&Form{
				FormID: "importevent",
				GetModel: func(form *Form, ctx *Context) (interface{}, error) {
					return &ImportEventFormModel{}, nil
				},
				SubmitButtonText:  "Import Event",
				SubmitButtonClass: "button",
				OnSubmit:          ImportEvent,
			},
			H2("Create Event Without Amiando:"),
			&Form{
				FormID: "createevent",
				GetModel: func(form *Form, ctx *Context) (interface{}, error) {
					return &CreateEventFormModel{}, nil
				},
				SubmitButtonText:  "Create Event",
				SubmitButtonClass: "button",
				OnSubmit:          CreateEvent,
			},
		},
	}
}

func eventRegionAdminView(ctx *Context, model interface{}) (view View, err error) {
	event := model.(*models.Event)

	slug := ctx.URLArgs[0]
	number := event.Number.String()

	eventPublicURL := Region0_Event1.URL(ctx.ForURLArgs(slug, number))
	eventDashboardURL := Region0_Event1_Dashboard.URL(ctx.ForURLArgs(slug, number))
	eventAdminURL := Region0_Event1_Admin.URL(ctx.ForURLArgs(slug, number))
	syncURL := Admin_Region0_SyncEvent1.URL(ctx.ForURLArgs(slug, event.ID.Hex()))

	eventdeleteform := &Form{
		SubmitButtonText:    "Delete Event",
		SubmitButtonConfirm: "Do you really really want to delete this event " + event.Name.Get() + "?",
		SubmitButtonClass:   "delete-button",
		FormID:              "delete" + event.ID.Hex(),
		OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
			return "", StringURL("."), event.Delete()
		},
	}

	from := event.Start.Format("2.")
	until := event.End.Format("2. Jan 2006")

	view = Views{
		&Tag{
			Tag: "h3",
			Content: Views{
				Printf("%s: %s - %s, ", event.Name, from, until),
				Printf("<a href='%s'>Public Site</a>, <a href='%s'>Dashboard</a>, <a href='%s'>Admin</a>", eventPublicURL, eventDashboardURL, eventAdminURL),
				// NonProductionServerView(eventdeleteform),
				&If{
					Condition: !event.AmiandoEventIdentifier.IsEmpty(),
					Content:   Printf(", <a href='https://www.amiando.com/mycenter/event/%s/overview' target='_blank'>Amiando</a>, <a href='%s'>Sync</a>", event.AmiandoEventIdentifier, syncURL),
				},
				&If{
					Condition: event.Type.String() != "StartupLive" || event.Status.String() == "Canceled",
					Content:   eventdeleteform,
				},
			},
		},
		BR(),
	}
	return view, nil
}

type ImportEventFormModel struct {
	AmiandoEventIdentifier model.String `model:"minlen=3" view:"label=Amiando Event Identifier or URL"`
	AmiandoApiKey          model.String `view:"label=Amiando API Key"`
	EventType              model.Choice `model:"options=StartupLive,StartupLounge,PioneersFestival" view:"label=Event Type"`
}

func ImportEvent(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
	region := ctx.Data.(*PageData).Region

	m := formModel.(*ImportEventFormModel)
	identifier := m.AmiandoEventIdentifier.String()
	apiKey := m.AmiandoApiKey.String()
	eventType := m.EventType

	if identifier == "" {
		return "", nil, errors.New("Empty Event Identifier")
	}
	if strings.HasPrefix(identifier, "http://") {
		identifier = path.Base(identifier)
	}

	count, err := models.Events.Filter("AmiandoEventIdentifier", identifier).Count()
	if err != nil {
		return "", nil, err
	}

	if count > 0 {
		return "", nil, fmt.Errorf("Event with identifier '%s' already added", identifier)
	}

	if apiKey == "" {
		apiKey = models.DefaultAmiandoApiKey
	}

	event, err := region.CreateEvent(eventType)
	if err != nil {
		return "", nil, err
	}
	event.AmiandoEventIdentifier.Set(identifier)
	event.AmiandoEventApiKey.Set(apiKey)

	err = event.SyncBasicDataWithAmiando()
	if err != nil {
		return "", nil, err
	}

	return "", StringURL("."), event.Save()
}

type CreateEventFormModel struct {
	EventType model.Choice `model:"options=StartupLive,StartupLounge,PioneersFestival" view:"label=Event Type"`
	Start     model.DateTime
	End       model.DateTime
}

func CreateEvent(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
	region := ctx.Data.(*PageData).Region
	m := formModel.(*CreateEventFormModel)
	eventType := m.EventType

	debug.Print("create Event form get eventType: ", eventType)

	event, err := region.CreateEvent(eventType)
	if err != nil {
		return "", nil, err
	}
	event.Start = m.Start
	event.End = m.End
	event.FAQ_HTML.Set(event.GetDefaultFAQ())
	event.PitcherRegistrationWelcomeText.Set(event.GetDefaultPitcherRegistrationText())
	event.Show.FAQ = true
	event.Show.Schedule = true
	event.Show.MentorsJudges = true
	event.Show.Organisers = true

	if eventType == models.StartupLive {
		//create Email
		amiandomail := region.Slug.String() + "" + event.Number.String() + "admin"
		debug.Print("amiando admin email: " + amiandomail)
		createAdminEmail(amiandomail)
		amiandomail = amiandomail + "@startuplive.in"
		event.AmiandoAccEmail.Set(amiandomail)
	}

	return "", StringURL("."), event.Save()
}

type editRegionnFormModel struct {
	Name model.String `model:"minlen=4"`
}

func deleteRegion(ctx *Context) (view View, err error) {
	if SessionUserIsSuperAdmin(ctx) {
		region, err := EventRegion(ctx.URLArgs)
		if err != nil {
			return nil, err
		}
		eventcount, err := region.EventCount(models.StartupLive)
		if err != nil {
			return nil, err
		}
		if eventcount == 0 {
			var views Views
			views = append(views,
				HR(),
				&Form{
					SubmitButtonText:    "Delete Region",
					SubmitButtonConfirm: "Are you sure you want to delete the region " + region.Name.Get() + "?",
					SubmitButtonClass:   "button",
					FormID:              "delete" + region.ID.Hex(),
					OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
						return "", StringURL("../regions"), region.Delete()
					},
				},
				HR(),
			)
			return views, nil
		}
	}

	return nil, nil
}

func createAdminEmail(name string) (err error) {

	plesk := plesk.NewPleskApi(PleskUrl, PleskUser, PleskPW, false)

	// POST requets
	debug.Nop()

	forward := []string{AmiandoAdminEmail}

	mailacc, err := plesk.EmailExists(name)
	if err != nil {
		return err
	}

	if mailacc.Status != "error" {
		// update amiando admin email
		plesk.UpdateEmail(name, forward)

	} else {
		// create amiando admin email
		plesk.CreateEmail(name, forward)

	}

	return nil
}
