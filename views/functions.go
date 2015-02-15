package views

import (
	"bytes"
	"labix.org/v2/mgo/bson"
	"strconv"

	"github.com/ungerik/go-start/config"
	"github.com/ungerik/go-start/debug"
	"github.com/ungerik/go-start/model"
	"github.com/ungerik/go-start/user"
	"github.com/ungerik/go-start/utils"
	. "github.com/ungerik/go-start/view"

	"github.com/STARTeurope/startuplive.in/models"
)

func init() {
	debug.Nop()
}

///////////////////////////////////////////////////////////////////////////////
// Views

func TitleBar(title string) View {
	return &Div{Class: "title-bar", Content: Escape(title)}
}

func TitleBarRight(title string) View {
	return &Div{Class: "title-bar right", Content: Escape(title)}
}

///////////////////////////////////////////////////////////////////////////////
// Page Titles

func EventTitle(mainTitle string) Renderer {
	return Render(
		func(ctx *Context) error {
			event :=
				ctx.Data.(*PageData).Event

			ctx.Response.Printf("%s | %s", mainTitle, event.Name)
			return nil
		},
	)
}

func EventAdminTitle(mainTitle string) Renderer {
	return Render(
		func(ctx *Context) error {
			event :=
				ctx.Data.(*PageData).Event

			ctx.Response.Printf("%s | %s | Admin", mainTitle, event.Name)
			return nil
		},
	)
}

func EventDashboardTitle(mainTitle string) Renderer {
	return Render(
		func(ctx *Context) error {
			event :=
				ctx.Data.(*PageData).Event

			ctx.Response.Printf("%s | %s | Dashboard", mainTitle, event.Name)
			return nil
		},
	)
}

///////////////////////////////////////////////////////////////////////////////
// models.EventRegion

func EventRegion(urlArgs []string) (region *models.EventRegion, err error) {
	region, found, err := models.GetEventRegion(urlArgs[0])
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, NotFound("404: Region not found")
	}
	return region, nil
}

func RegionTitle(what string) Renderer {
	return Render(
		func(ctx *Context) error {
			region := ctx.Data.(*PageData).Region

			ctx.Response.Printf("%s | %s", region.Name, what)
			return nil
		},
	)
}

func RegionCSSContext(ctx *Context) (context interface{}, err error) {
	eventRegion, err := EventRegion(ctx.URLArgs)
	if err != nil {
		return nil, err
	}
	return eventRegion.ColorScheme(), nil
}

func regionOnPreRender_SetFavicon(page *Page, urlArgs []string) (err error) {
	region, err := EventRegion(urlArgs)
	if err != nil {
		return err
	}
	page.Favicon16x16URL = region.Favicon16x16URL.Get()
	page.Favicon57x57URL = region.Favicon57x57URL.Get()
	page.Favicon72x72URL = region.Favicon72x72URL.Get()
	page.Favicon114x114URL = region.Favicon114x114URL.Get()
	page.Favicon129x129URL = region.Favicon129x129URL.Get()
	return nil
}

// RegionsIterator iterates models.EventRegion, use Next(*models.EventRegion)
func RegionsIterator() model.Iterator {
	return models.EventRegions.SortFunc(
		func(a, b *models.EventRegion) bool {
			return utils.CompareCaseInsensitive(a.Name.Get(), b.Name.Get())
		},
	)
}

///////////////////////////////////////////////////////////////////////////////
// models.CitySuggestion

// CitySuggestionIterator iterates models.CitySuggestion, use Next(*models.CitySuggestion)
func CitySuggestionIterator() model.Iterator {
	return models.CitySuggestions.SortFunc(
		func(a, b *models.CitySuggestion) bool {
			return utils.CompareCaseInsensitive(a.Date.Get(), b.Date.Get())
		},
	)
}

///////////////////////////////////////////////////////////////////////////////
// models.Event

func RegionAndEvent(urlArgs []string) (*models.EventRegion, *models.Event, error) {
	slug := urlArgs[0]
	number, err := strconv.ParseInt(urlArgs[1], 10, 64)
	if err != nil {
		return nil, nil, NotFound("404: Event Number not parseable")
	}
	region, event, found, err := models.GetRegionAndEvent(models.StartupLive, slug, number)
	if found && err == nil {
		return region, event, err
	}
	if err != nil {
		return nil, nil, err
	}
	region, event, found, err = models.GetRegionAndEvent(models.LiveAcademy, slug, number)
	if found && err == nil {
		return region, event, err
	}
	if err != nil {
		return nil, nil, err
	}
	region, event, found, err = models.GetRegionAndEvent(models.PioneersFestival, slug, number)
	if found && err == nil {
		return region, event, err
	}
	if err != nil {
		return nil, nil, err
	}
	region, event, found, err = models.GetRegionAndEvent(models.StartupLounge, slug, number)
	if found && err == nil {
		return region, event, err
	}
	if err != nil {
		return nil, nil, err
	}
	return nil, nil, NotFound("404: Event not found")
}

func EventAllVotesIterator(ctx *Context) model.Iterator {
	return ctx.Data.(*PageData).Event.AllVotesIterator()
}

///////////////////////////////////////////////////////////////////////////////
// models.EventParticipant

func EventParticipantIterator(ctx *Context) model.Iterator {
	return ctx.Data.(*PageData).Event.ParticipantIterator()
}

func EventParticipant(ctx *Context) (participant *models.EventParticipant, err error) {
	id := bson.ObjectIdHex(ctx.URLArgs[2])
	found, err := models.EventParticipants.TryDocumentWithID(id, &participant)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, NotFound("404: Participant not found")
	}
	return participant, nil
}

///////////////////////////////////////////////////////////////////////////////
// models.EventTeam

func EventTeamIterator(ctx *Context) model.Iterator {
	return ctx.Data.(*PageData).Event.TeamIterator()
}

func EventPitchingTeamIterator(ctx *Context) model.Iterator {
	return ctx.Data.(*PageData).Event.PitchingTeamIterator()
}

func EventNotPitchingTeamIterator(ctx *Context) model.Iterator {
	return ctx.Data.(*PageData).Event.NotPitchingTeamIterator()
}

func EventPitchingTeamScoreIterator(ctx *Context) model.Iterator {
	return ctx.Data.(*PageData).Event.SortPitchingTeamsByScoreIterator()
}

func EventNotPitchingTeamScoreIterator(ctx *Context) model.Iterator {
	return ctx.Data.(*PageData).Event.SortNotPitchingTeamsByScoreIterator()
}

func EventTeam(ctx *Context) (team *models.EventTeam, err error) {
	id := bson.ObjectIdHex(ctx.URLArgs[2])
	found, err := models.EventTeams.TryDocumentWithID(id, &team)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, NotFound("404: Team not found")
	}
	return team, nil
}

///////////////////////////////////////////////////////////////////////////////
// models.Person

func EventPerson(ctx *Context) (person *models.Person, err error) {
	id := bson.ObjectIdHex(ctx.URLArgs[2])
	found, err := models.People.TryDocumentWithID(id, &person)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, NotFound("404: Person not found")
	}
	return person, nil
}

func EventMentorIterator(ctx *Context) model.Iterator {
	return ctx.Data.(*PageData).Event.MentorIterator()
}

func EventJudgeIterator(ctx *Context) model.Iterator {
	return ctx.Data.(*PageData).Event.JudgeIterator()
}

func EventOrganiserIterator(ctx *Context) model.Iterator {
	return ctx.Data.(*PageData).Event.OrganiserIterator()
}

func EventParticipantFeedbackIterator(ctx *Context) model.Iterator {
	return ctx.Data.(*PageData).Event.ParticipantFeedbackIterator()
}

func EventHostFeedbackIterator(ctx *Context) model.Iterator {
	return ctx.Data.(*PageData).Event.HostFeedbackIterator()
}

func EventMentorJudgeFeedbackIterator(ctx *Context) model.Iterator {
	return ctx.Data.(*PageData).Event.MentorJudgeFeedbackIterator()
}

var AlwaysExcludedPersonFormFields = []string{
	"Name.Organization",
	"BirthYear",
	"Password",
	"Languages",
	"Email.$.Confirmed",
	"Email.$.ConfirmationCode",
	"Facebook.$.Confirmed",
	"Facebook.$.AccessToken",
	"Twitter.$.Confirmed",
	"Twitter.$.AccessToken",
	"LinkedIn.$.Confirmed",
	"LinkedIn.$.AccessToken",
	"GitHub.$.Confirmed",
	"GitHub.$.AccessToken",
	"Skype.$.Confirmed",
	"TermsAndConditions",
	"TaCDate",
}

func PersonForm(person *models.Person, redirect URL, disableFields, excludeFields []string, requiredFields []string) *Form {
	return &Form{
		SubmitButtonText:         "Save Person Data",
		SubmitButtonClass:        "button",
		FormID:                   "person" + person.ID.Hex(),
		GetModel:                 FormModel(person),
		ExcludedFields:           append(AlwaysExcludedPersonFormFields, excludeFields...),
		DisabledFields:           disableFields,
		RequiredFields:           requiredFields,
		GeneralErrorOnFieldError: true,
		OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
			m, r, e := OnFormSubmitSaveModel(form, formModel, ctx)
			if r == nil {
				r = redirect
			}
			return m, r, e
		},
	}
}

///////////////////////////////////////////////////////////////////////////////
// models.Person

func EventScheduleItem(ctx *Context) (item *models.EventScheduleItem, err error) {
	id := bson.ObjectIdHex(ctx.URLArgs[2])
	found, err := models.EventScheduleItems.TryDocumentWithID(id, &item)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, NotFound("404: Schedule item not found")
	}
	return item, nil
}

///////////////////////////////////////////////////////////////////////////////
// PageData

type PageData struct {
	Region       *models.EventRegion
	Event        *models.Event
	Location     *models.EventLocation
	Person       *models.Person
	Participant  *models.EventParticipant
	Team         *models.EventTeam
	ScheduleItem *models.EventScheduleItem
	WikiEntry    *models.WikiEntry
}

func SetRegionPageData(page *Page, ctx *Context) (err error) {
	data := new(PageData)

	data.Region, err = EventRegion(ctx.URLArgs)
	if err != nil {
		return err
	}

	ctx.Data = data

	return regionOnPreRender_SetFavicon(page, ctx.URLArgs)
}

func SetEventPageData(page *Page, ctx *Context) (err error) {
	data := new(PageData)

	data.Region, data.Event, err = RegionAndEvent(ctx.URLArgs)
	if err != nil {
		return err
	}

	err = data.Event.Location.Get(&data.Location)
	if err != nil {
		return err
	}

	ctx.Data = data

	InitRichTextEditor(data.Region)

	return regionOnPreRender_SetFavicon(page, ctx.URLArgs)
}

func SetEventPersonPageData(page *Page, ctx *Context) (err error) {
	person, err := EventPerson(ctx)
	if err != nil {
		return err
	}
	err = SetEventPageData(page, ctx)
	if err == nil {

		ctx.Data.(*PageData).Person = person
	}
	return err
}

func SetEventParticipantPageData(page *Page, ctx *Context) (err error) {
	participant, err := EventParticipant(ctx)
	if err != nil {
		return err
	}
	err = SetEventPageData(page, ctx)
	if err == nil {

		ctx.Data.(*PageData).Participant = participant
	}
	return err
}

func SetEventTeamPageData(page *Page, ctx *Context) (err error) {
	team, err := EventTeam(ctx)
	if err != nil {
		return err
	}
	err = SetEventPageData(page, ctx)
	if err == nil {

		ctx.Data.(*PageData).Team = team
	}
	return err
}

func SetEventSchedulePageData(page *Page, ctx *Context) (err error) {
	scheduleitem, err := EventScheduleItem(ctx)
	if err != nil {
		return err
	}
	err = SetEventPageData(page, ctx)
	if err == nil {

		ctx.Data.(*PageData).ScheduleItem = scheduleitem
	}
	return err
}

func SetWikiEntryPageData(page *Page, ctx *Context) (err error) {
	entry, err := WikiEntry(ctx)
	if err != nil {
		return err
	}
	err = SetEventPageData(page, ctx)
	if err == nil {

		ctx.Data.(*PageData).WikiEntry = entry
	}
	return err
}

///////////////////////////////////////////////////////////////////////////////
// Authentification

func SessionUserIsSuperAdmin(ctx *Context) bool {
	var person models.Person
	found, err := user.OfSession(ctx.Session, &person)
	if err != nil {
		config.Logger.Println(err.Error())
	}
	if !found {
		return false
	}
	return person.SuperAdmin.Get()
}

///////////////////////////////////////////////////////////////////////////////
// Wiki

func WikiEntry(ctx *Context) (item *models.WikiEntry, err error) {
	id := bson.ObjectIdHex(ctx.URLArgs[2])
	found, err := models.Wiki.TryDocumentWithID(id, &item)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, NotFound("404: Wiki Entry item not found")
	}
	return item, nil
}

///////////////////////////////////////////////////////////////////////////////
// Init Rich Text Editor

func InitRichTextEditor(region *models.EventRegion) {
	var Ctx struct {
		PrimaryColor   string
		SecondaryColor string
	}

	Ctx.PrimaryColor = models.Colors[DefaultPrimaryColorIndex]
	Ctx.SecondaryColor = models.Colors[DefaultSecondaryColorIndex]

	var toolbar bytes.Buffer
	RenderTemplate("wysihtml5-toolbar.html", &toolbar, Ctx)
	Config.RichText.DefaultToolbar = toolbar.String()
}
