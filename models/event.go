package models

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/AlexTi/go-amiando"
	"github.com/ungerik/go-start/config"
	"github.com/ungerik/go-start/debug"
	"github.com/ungerik/go-start/errs"
	"github.com/ungerik/go-start/media"
	"github.com/ungerik/go-start/model"
	"github.com/ungerik/go-start/mongo"
	"github.com/ungerik/go-start/user"
	"github.com/ungerik/go-start/utils"
	. "github.com/ungerik/go-start/view"
)

const EventDefaultTopic = "Web and Mobile"
const DefaultAmiandoApiKey = "VqmuCQsRg7G1LnFaLt8lfB5xjqIT6Byc8bjaFx6Wrknkl51X6Y"

const (
	StartupLive      = "StartupLive"
	StartupLounge    = "StartupLounge"
	PioneersFestival = "PioneersFestival"
        LiveAcademy      = "LiveAcademy"
)

const (
	EventPlanned  = "Planned"
	EventApproved = "Approved"
	EventCanceled = "Canceled"
	EventPublic   = "Public"
)

var Events = mongo.NewCollection("events", "Name")

func NewEvent() *Event {
	var doc Event
	Events.InitDocument(&doc)
	return &doc
}

type Event struct {
	mongo.DocumentBase `bson:",inline"`
	Type               model.Choice `model:"options=StartupLive,LiveAcademy,StartupLounge,PioneersFestival" view:"auth=admin"`
	Status             model.Choice `model:"options=Planned,Approved,Canceled,Public" view:"auth=admin"`
	Show               ShowPages
	Name               model.String `model:"required"`
	Topic              model.String `model:"required" view:"placeholder=your events topic"`
	//Tagline                 model.String
	DescriptionLead model.Text     `model:"required" view:"label=Short Description (one sentence)"`
	Description     model.Text     `model:"required"`
	HowItsDone      model.Text     `model:"required" view:"label=How it is done"`
	Prizes          model.Text     `model:"required" view:"label=Prizes"`
	Region          mongo.Ref      `model:"required|to=eventregions" view:"auth=admin"`
	Number          model.Int      `view:"auth=admin"`
	Location        mongo.Ref      `model:"required|to=eventlocations"`
	Start           model.DateTime `model:"required" view:"auth=admin"`
	End             model.DateTime `model:"required" view:"auth=admin"`
	TimeZone        model.Int      `view:"auth=admin"`
	Language        model.String   `model:"required"`

	// Todo Erik: model.ImageRef
	StampImageURL      model.Url `view:"label=Old StampImageURL"`
	ImageURL_604x0     model.Url `view:"label=Old ImageURL 604"`
	TeamImageURL_604x0 model.Url `view:"label=Old TeamImageURL 604"`

	// Todo Erik: []model.ImageRef
	HostLogoURL_200x0       model.Url `view:"label=Old HostLogoURL"`
	HostLogoLinkURL         model.Url
	SecondHostLogoURL_200x0 model.Url `view:"label=Old SecondHostLogoURL"`
	SecondHostLogoLinkURL   model.Url

	// Todo Erik: []model.ImageRef
	PoweredByLogoURL_200x0       model.Url `view:"label=Old PoweredByLogoURL"`
	PoweredByLogoLinkURL         model.Url
	SecondPoweredByLogoURL_200x0 model.Url `view:"label=Old SecondPoweredByLogoURL"`
	SecondPoweredByLogoLinkURL   model.Url

	// EventPartners []EventPartner `view:"hidden"`
	EventPartners []EventPartner `view:"auth=admin|hidden"`

	GlobalPartnerLinkURL           model.Url `view:"label=Global Partner Link URL (if not standard)|auth=admin"`
	FacebookURL                    model.Url
	TwitterURL                     model.Url
	LinkedInURL                    model.Url
	FlickrURL                      model.Url
	SpotieURL                      model.Url
	FAQ_HTML                       model.Text   `view:"rows=20"`
	DashboardInfo_HTML             model.Text   `view:"rows=20"`
	MentorsJudgesTab_Title         model.String `view:"auth=admin"` // will be "Mentors/Judges" if empty
	MentorsJudgesTab_RenameMentors model.String `view:"auth=admin"` // will be "Mentors" if empty
	ExtraTab_Title                 model.String

	ExtraTab_HTML model.Text `model:"rows=20"`
	//TechnicalContact       mongo.Ref `gostart:"to=people"`
	//SecurityContact        mongo.Ref `gostart:"to=people"`
	RegistrationButton             model.String `model:"required|label=Registration Button Text"`
	RegistrationTagline            model.String `model:"required|label=Registration Button Subtext"`
	Organisers                     []mongo.Ref  `model:"to=people"` // first is chief organiser
	Mentors                        []mongo.Ref  `model:"to=people"`
	Judges                         []mongo.Ref  `model:"to=people"`
	Started                        model.Bool   `model:"label=Started (dashboard restricted to present participants)"` // Absent participants are hidden and can't access the dashboard anymore
	RoundupURL                     model.Url    `view:"auth=admin"`
	PitcherRegistrationWelcomeText model.Text   `view:"rows=20"`

	GoLiveRequested          model.Bool `view:"auth=admin"`
	SetupAmiandoEventRequest model.Bool `view:"auth=admin"`

	GoogleAnalyticsHostAccount model.String `view:"hidden"`

	AmiandoAccEmail              model.Email  `view:"auth=admin"`
	AmiandoHostId                model.String `view:"auth=admin"`
	AmiandoHostPassword          model.String `view:"auth=admin"`
	AmiandoEventIdentifier       model.String `view:"auth=admin"`
	AmiandoEventApiKey           model.String `view:"auth=admin"`
	AmiandoEventId               model.String `view:"auth=admin"`
	AmiandoIframeCode            model.Text   `view:"auth=admin|label=Amiando iframe code"`
	AmiandoTicketCategories      []AmiandoTicketCategory
	AmiandoTicketCategoriesAdded model.Bool `view:"auth=admin"`
	AmiandoPaymentData           AmiandoPaymentData
	AmiandoPaymentAdded          model.Bool `view:"auth=admin"`
	AmiandoBillingData           AmiandoBillingData
	AmiandoBillingAdded          model.Bool `view:"auth=admin"`
	AmiandoPromoCodeAdded        model.Bool `view:"auth=admin"`
	AmiandoAttendeeDataAdded     model.Bool `view:"auth=admin"`
	AmiandoEventActivated        model.Bool `view:"auth=admin"`

	//OpeningTimes model.Array{Of: &DateTime{}}},
	//ClosingTimes model.Array{Of: &DateTime{}}},
	//Rooms model.Array{Of: &Str{}}},     // todo
	//Equipment model.Array{Of: &Str{}}}, /models.StartupLive,/ todo Beamer, Sound, Internet
	//Sponsors model.Array{Of: &Ref{To: &Organisation}}},
	//Prizes model.Array{Of: &Str{}}}, // todo
	//WorkshopLeaders model.Array{Of: &Ref{To: &Person}}},
}

type ShowPages struct {
	Info          model.Bool `view:"label=Dashboard Info"`
	Location      model.Bool
	Schedule      model.Bool
	MentorsJudges model.Bool
	Organisers    model.Bool
	FAQ           model.Bool
	Registration  model.Bool
	Voting        model.Bool `view:"label=Dashboard Voting"`
	VotingResult  model.Bool `view:"label=Dashboard Voting-Result"`
	ExtraTab      model.Bool
}

type EventPartner struct {
	Order    model.Int
	Name     model.String
	Partners []mongo.Ref `model:"to=partners"`
}

func (self *Event) String() string {
	return self.Name.Get()
}

func (self *Event) IsPublished() bool {
	return self.Status == "Public"
}

func (self *Event) IsPublishedOrApproved() bool {
	return self.Status == "Public" || self.Status == "Approved"
}

func (self *Event) GetDate() string {
	start := self.Start.Format("02/01 - ")
	end := self.End.Format("02/01 2006")
	return start + end
}

func (self *Event) AboutDone() bool {
	aboutdone := (self.Name != "" &&
		self.Topic != "" &&
		self.DescriptionLead != "" &&
		self.Description != "" &&
		self.HowItsDone != "" &&
		self.Prizes != "" &&
		self.Language != "" &&
		self.RegistrationButton != "" &&
		self.RegistrationTagline != "")
	return aboutdone
}

func (self *Event) LocationSet() bool {
	var loc EventLocation
	err := self.Location.Get(&loc)
	if err != nil {
		return false
	}

	done := (loc.Name != "" &&
		loc.Description != "" &&
		loc.ShortName != "" &&
		loc.PublicTransport != "" &&
		loc.PrivateCar != "" &&
		loc.Address.City != "" &&
		loc.Address.State != "" &&
		loc.Address.ZIP != "" &&
		loc.Address.FirstLine != "")
	return done
}

func (self *Event) AmiandoDataSetup() bool {
	done := (self.AmiandoBillingAdded &&
		self.AmiandoPaymentAdded &&
		self.AmiandoTicketCategoriesAdded)
	return done.Get()
}

func (self *Event) GetTopic() string {
	if self.Topic.IsEmpty() {
		return EventDefaultTopic
	}
	return self.Topic.String()
}

func (self *Event) ParticipantIterator() model.Iterator {
	// return EventParticipants.Filter("Event", self.ID).Sort("Name.First").Sort("Name.Last").Iterator()
	return EventParticipants.Filter("Event", self.ID).SortFunc(
		func(a, b *EventParticipant) bool {
			return utils.CompareCaseInsensitive(a.Name(), b.Name())
		},
	)
}

func (self *Event) ParticipantsCount() (int, error) {
	return EventParticipants.Filter("Event", self.ID).Count()
}

func (self *Event) ParticipantPersonIterator(includeCancelled bool) model.Iterator {
	peopleIDs := make([]interface{}, 0, 128)
	i := EventParticipants.Filter("Event", self.ID).Iterator()

	var participant *EventParticipant
	for i.Next(&participant) {
		if includeCancelled || !participant.Cancelled() {
			peopleIDs = append(peopleIDs, participant.Person.ID)
		}
	}
	if i.Err() != nil {
		return model.NewErrorOnlyIterator(i.Err()) //i.Err()
	}

	// model.Iterate(i, func(participant *EventParticipant) {
	// 	if includeCancelled || !participant.Cancelled() {
	// 		peopleIDs = append(peopleIDs, participant.Person.ID)
	// 	}
	// })
	// if i.Err() != nil {
	// 	return model.NewErrorOnlyIterator(i.Err())
	// }
	return People.FilterIn("_id", peopleIDs...).Sort("Name.First").Sort("Name.Last").Iterator()
}

func (self *Event) TeamsCount() (int, error) {
	return EventTeams.Filter("Event", self.ID).Count()
}

func compareTeamNames(a, b *EventTeam) bool {
	return utils.CompareCaseInsensitive(a.Name.Get(), b.Name.Get())
}

// TeamIterator iterates models.EventTeam, use Next(*models.EventTeam)
func (self *Event) TeamIterator() model.Iterator {
	return EventTeams.Filter("Event", self.ID).SortFunc(compareTeamNames)
}

// PitchingTeamIterator iterates models.EventTeam, use Next(*models.EventTeam)
func (self *Event) PitchingTeamIterator() model.Iterator {
	return EventTeams.Filter("Event", self.ID).Filter("Pitching", true).SortFunc(compareTeamNames)
}

// NotPitchingTeamIterator iterates models.EventTeam, use Next(*models.EventTeam)
func (self *Event) NotPitchingTeamIterator() model.Iterator {
	return EventTeams.Filter("Event", self.ID).Filter("Pitching", false).SortFunc(compareTeamNames)
}

func compareScoreAndTeamnames(a, b *EventTeam) bool {
	scoreA := a.ComputeAverageScoreByEvent()
	scoreB := b.ComputeAverageScoreByEvent()
	nameA := a.Name.Get()
	nameB := b.Name.Get()
	if scoreA > scoreB {
		return true
	} else if scoreA < scoreB {
		return false
	}
	return utils.CompareCaseInsensitive(nameA, nameB)
}

// SortPitchingTeamsByScoreIterator iterates models.EventTeam, use Next(*models.EventTeam)
func (self *Event) SortPitchingTeamsByScoreIterator() model.Iterator {
	return EventTeams.Filter("Event", self.ID).Filter("Pitching", true).SortFunc(compareScoreAndTeamnames)
}

// SortNotPitchingTeamsByScoreIterator iterates models.EventTeam, use Next(*models.EventTeam)
func (self *Event) SortNotPitchingTeamsByScoreIterator() model.Iterator {
	return EventTeams.Filter("Event", self.ID).Filter("Pitching", false).SortFunc(compareScoreAndTeamnames)
}

// MentorIterator iterates models.Person, use Next(*models.Person)
func (self *Event) MentorIterator() model.Iterator {
	return People.FilterReferenced(self.Mentors...).Sort("Name.First").Sort("Name.Last").Iterator()
}

// JudgeIterator iterates models.Person, use Next(*models.Person)
func (self *Event) JudgeIterator() model.Iterator {
	return People.FilterReferenced(self.Judges...).Sort("Name.First").Sort("Name.Last").Iterator()
}

// ParticipantFeedbackIterator iterates models.FeedbackParticipant, use Next(*models.FeedbackParticipant)
func (self *Event) ParticipantFeedbackIterator() model.Iterator {
	return FeedbackParticipants.Filter("Event", self.ID).Iterator()
}

// HostFeedbackIterator iterates models.FeedbackHost, use Next(*models.FeedbackHost)
func (self *Event) HostFeedbackIterator() model.Iterator {
	return FeedbackHosts.Filter("Event", self.ID).Iterator()
}

// MentorJudgeFeedbackIterator iterates models.FeedbackMentorJudge, use Next(*models.FeedbackMentorJudge)
func (self *Event) MentorJudgeFeedbackIterator() model.Iterator {
	return FeedbackMentorsJudges.Filter("Event", self.ID).Iterator()
}

// OrganiserIterator iterates models.Person, use Next(*models.Person)
func (self *Event) OrganiserIterator() model.Iterator {
	refs, invalid, err := mongo.CheckRefs(self.Organisers)
	if err != nil {
		return model.NewErrorOnlyIterator(err)
	}
	for _, r := range invalid {
		config.Logger.Printf("event.OrganiserIterator(): Invalid Ref %s", r)
	}
	return mongo.NewDereferenceIterator(refs...)
}

func (self *Event) sortOrganisers() {
	// TODO fix crash

	// if len(self.Organisers) == 0 {
	// 	return
	// }
	// lead := self.Organisers[:1]
	// rest := self.Organisers[1:]
	// mongo.SortRefs(rest,
	// 	func(a, b *mongo.Ref) bool {
	// 		doc, err := a.Get()
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		nameA := doc.(*Person).Name.String()

	// 		doc, err = b.Get()
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		nameB := doc.(*Person).Name.String()

	// 		return utils.CompareCaseInsensitive(nameA, nameB)
	// 	},
	// )
	// self.Organisers = append(lead, rest...)
}

func (self *Event) MakeLeadOrganiser(person *Person) error {
	for i := range self.Organisers {
		if self.Organisers[i].ID == person.ID {
			self.Organisers[0], self.Organisers[i] = self.Organisers[i], self.Organisers[0]
			self.sortOrganisers()
			return nil
		}
	}
	return errs.Format("No organiser with name: %s", person.Name.String())
}

func (self *Event) HasLeadOrganiser() bool {
	if len(self.Organisers) > 0 {
		return true
	}
	return false
}

//TODO: get lead organiser
func (self *Event) GetLeadOrganiser() (*Person, error) {
	debug.Nop()
	// if len(self.Organisers)>0 {
	// 	debug.Print("organiser: ", self.Organisers[0].Name.First)
	// 	// self.Organisers[0] = self.Organisers[0]
	// 	// doc, err := self.Organisers[0].Get()
	// 	orga, err:= People.DocumentWithID(self.Organisers[0].ID)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return orga.(*Person), nil

	// }
	return nil, errs.Format("No organisers at this event")
}

func (self *Event) RemoveOrganiser(person *Person) error {
	if refs, ok := mongo.RemoveRefWithIDFromSlice(self.Organisers, person.ID); ok {
		self.Organisers = refs
		return nil
	}
	return errs.Format("No organiser with name: %s", person.Name.String())
}

func (self *Event) AddOrganiser(person *Person) {
	self.Organisers = append(self.Organisers, person.Ref())
	self.sortOrganisers()
}

func compareScheduleItem(a, b *EventScheduleItem) bool {
	return a.From.Time().Before(b.From.Time())
}

// ScheduleItemIterator iterates models.EventScheduleItem, use Next(*models.EventScheduleItem)
func (self *Event) ScheduleItemIterator() model.Iterator {
	// return EventScheduleItems.Filter("Event", self.ID).Sort("From").Iterator()
	return EventScheduleItems.Filter("Event", self.ID).SortFunc(compareScheduleItem)
}

// DaySchedulItemIterator iterates models.EventScheduleItem, use Next(*models.EventScheduleItem)
func (self *Event) DaySchedulItemIterator(someTimeOfTheDay time.Time) model.Iterator {
	from, until := utils.DayTimeRange(someTimeOfTheDay)
	return EventScheduleItems.Filter("Event", self.ID).Sort("From").FilterFunc(
		func(doc interface{}) (ok bool) {
			item := doc.(*EventScheduleItem)
			t := item.From.Time()
			return utils.TimeInRange(t, from, until)
		},
	)
	// iter := EventScheduleItems.Filter("Event", self.ID).Sort("From").FilterFunc(
	// 	func(doc interface{}) (ok bool) {
	// 		item := doc.(*EventScheduleItem)
	// 		t := item.From.Time()
	// 		return utils.TimeInRange(t, from, until)
	// 	},
	// )
	// return model.NewSortIterator(iter, compareScheduleItem)
}

func (self *Event) Days() ([]time.Time, error) {
	daymap := make(map[time.Time]bool)
	for t := utils.DayBeginningTime(self.Start.Time()); t.Before(self.End.Time()); t = t.Add(time.Hour * 24) {
		daymap[t] = true
	}
	i := self.ScheduleItemIterator()
	var scheduleitem EventScheduleItem
	for i.Next(&scheduleitem) {
		t := scheduleitem.From.Time()
		t = utils.DayBeginningTime(t)
		daymap[t] = true
	}
	if i.Err() != nil {
		return nil, i.Err()
	}
	var days utils.SortableTimeSlice
	for t, _ := range daymap {
		days = append(days, t)
	}
	sort.Sort(days)
	return days, nil
}

func (self *Event) IsHappeningNow() bool {
	return self.Started.Get() && self.End.Time().After(time.Now())
}

///////////////////////////////////////////////////////////////////////////////
// Voting

func (self *Event) GetVotes() (votes int, err error) {
	return Votes.Filter("Event", self.ID).Count()
}

func (self *Event) GetVotesByTeam(team *EventTeam) (votes int, err error) {
	return Votes.Filter("Event", self.ID).Filter("Team", team.ID).Count()
}

func (self *Event) PitchingTeamsSortedByVotesIterator() model.Iterator {
	return EventTeams.Filter("Event", self.ID).Filter("Pitching", true).SortFunc(
		func(a, b *EventTeam) bool {
			scoreA, err := self.GetVotesByTeam(a)
			scoreB, err := self.GetVotesByTeam(b)

			if err != nil {
				return false
			}

			if scoreA > scoreB {
				return true
			} else {
				return false
			}
			return false
		},
	)
}

func (self *Event) AllVotesIterator() model.Iterator {
	return Votes.Filter("Event", self.ID).Iterator()
}

func (self *Event) DeleteAllVotes() (int, error) {
	return Votes.Filter("Event", self.ID).RemoveAll()
}

///////////
// Partner Categories

func (self *Event) GetPartnersByCategory(cat string) ([]Partner, error) {
	var partners []Partner
	var partnerRefs []mongo.Ref

	for j := 0; j < len(self.EventPartners); j++ {
		i := j
		if self.EventPartners[i].Name.String() == cat {
			for k := 0; k < len(self.EventPartners[i].Partners); k++ {
				l := k
				var p Partner
				if !self.EventPartners[i].Partners[l].IsEmpty() {
					err := self.EventPartners[i].Partners[l].Get(&p)
					if err != nil {
						return partners, err
					}
					partners = append(partners, p)
					partnerRefs = append(partnerRefs, p.Ref())
				}

			}
			self.EventPartners[i].Partners = partnerRefs
			return partners, self.Save()
		}
	}

	// for j := 0; j < len(self.EventPartners); j++ {
	// 	i := j
	// 	if self.EventPartners[i].Name.String() == cat {
	// 		for k := 0; k < len(self.EventPartners[i].Partners); k++ {
	// 			l := k
	// 			var p Partner
	// 			err := self.EventPartners[i].Partners[l].Get(&p)
	// 			if err != nil {
	// 				return partners, err
	// 			}
	// 			partners = append(partners, p)
	// 		}
	// 	}
	// }

	return partners, nil
}

func (self *Event) AddEventPartnerToCategory(cat string, partner *Partner) error {
	foundcat := false
	for j := 0; j < len(self.EventPartners); j++ {
		i := j
		if self.EventPartners[i].Name.String() == cat {

			self.EventPartners[i].Partners = append(self.EventPartners[i].Partners, partner.Ref())

			foundcat = true
			ref := partner.Ref()
			if ref.IsEmpty() == false {
				return self.Save()
			} else {
				return errors.New("Could not add new Event Partner")
			}
		}
	}
	if !foundcat {
		debug.Print("not found - Partner Ref: ", partner.Ref())
		var eventPartner EventPartner
		eventPartner.Name.Set(cat)
		eventPartner.Partners = []mongo.Ref{partner.Ref()}
		self.EventPartners = append(self.EventPartners, eventPartner)
		ref := partner.Ref()
		if ref.IsEmpty() == false {
			return self.Save()
		} else {
			return errors.New("Could not add new Event Partner")
		}
	}

	return errors.New("Could not add new Event Partner")
}

func (self *Event) OrderEventPartner() error {
	for j := 0; j < len(self.EventPartners); j++ {
		if self.EventPartners[j].Name == "Partners" {
			self.EventPartners[j].Order = 0
		} else if self.EventPartners[j].Name == "Supporters" {
			self.EventPartners[j].Order = 1
		} else if self.EventPartners[j].Name == "Media Partners" {
			self.EventPartners[j].Order = 2
		}
	}
	return self.Save()
}

func (self *Event) RemoveEventPartner(partner *Partner) error {

	for j := 0; j < len(self.EventPartners); j++ {
		for i := 0; i < len(self.EventPartners[j].Partners); i++ {
			if self.EventPartners[j].Partners[i].ID == partner.ID {
				if i == (len(self.EventPartners[j].Partners) - 1) {
					self.EventPartners[j].Partners = self.EventPartners[j].Partners[:i]
				} else {
					self.EventPartners[j].Partners = append(self.EventPartners[j].Partners[:i], self.EventPartners[j].Partners[i+1:]...)
				}
				return self.Save()
			}
		}
	}
	return errors.New("Did not find Partner to remove")
}

func (self *Event) OrderEventPartnersInCategory(cat string, neworder []string) error {

	var partners []mongo.Ref
	for j := 0; j < len(self.EventPartners); j++ {
		i := j

		if self.EventPartners[i].Name.String() == cat {
			partners = self.EventPartners[i].Partners
			debug.Print(partners)
			var newpartners = make([]mongo.Ref, len(partners))

			for k := 0; k < len(neworder); k++ {

				l, err := strconv.Atoi(neworder[k])
				if err != nil {
					fmt.Errorf("not able to parse string value to int")
				}

				newpartners[k] = partners[l]
			}

			self.EventPartners[i].Partners = newpartners
			break
		}
	}

	return self.Save()

}

func (self *Event) HasEventPartners() bool {
	if len(self.EventPartners) == 0 {
		return false
	}
	return true
}

func (self *Event) GetEventPartnersLength() int {
	length := 0
	if len(self.EventPartners) == 0 {
		return 0
	} else {
		for j := 0; j < len(self.EventPartners); j++ {
			if self.EventPartners[j].Name != "" && len(self.EventPartners[j].Partners) != 0 {
				length++
			}
		}
	}
	return length
}

func (self *Event) RepairEventPartners() error {
	var repairedpartners []EventPartner
	if len(self.EventPartners) == 0 {
		return nil
	} else {
		for j := 0; j < len(self.EventPartners); j++ {
			if self.EventPartners[j].Name != "" && len(self.EventPartners[j].Partners) != 0 {
				repairedpartners = append(repairedpartners, self.EventPartners[j])
			}
		}
	}
	self.EventPartners = repairedpartners
	return self.Save()
}

///////////////////////////////////////////////////////////////////////////////
// Default Texts

func (self *Event) GetDefaultFAQ() string {
	var faq bytes.Buffer
	RenderTemplate("faq_template.html", &faq, nil)
	return faq.String()
}

func (self *Event) GetDefaultPitcherRegistrationText() string {
	filename := "pitcher-registration_template.html"
	filePath, found, _ := FindTemplateFile(filename)
	if !found {
		return ""
	}
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return ""
	}
	return string(b)
}

///////////////////////////////////////////////////////////////////////////////
// Amiando Data

type AmiandoTicketCategory struct {
	Name      model.String   `model:"required"`
	Price     model.String   `model:"required" view:"placeholder=only whole numbers"`
	Available model.String   `model:"required" view:"placeholder=amount of available tickets"`
	SaleStart model.DateTime `model:"required"`
	SaleEnd   model.DateTime `model:"required"`
}

type AmiandoPaymentData struct {
	AccountHolder model.String
	BankName      model.String
	Swift         model.String
	Iban          model.String
	Country       model.String
}

type AmiandoBillingData struct {
	FirstName model.String
	LastName  model.String
	Company   model.String
	Street    model.String
	Zip       model.String
	City      model.String
	Country   model.String
}

///////////////////////////////////////////////////////////////////////////////
// Amiando Sync

/*******
/ Amiando Participant go-amiando
 	PaymentID     ID            `json:"-"`
	PaymentUserID ID            `json:"buyerId"`      // payment
	PaymentStatus PaymentStatus `json:"status"`       // payment
	InvoiceNumber string        `json:"identifier"`   // payment
	CreatedDate   string        `json:"creationTime"` // payment
	ModifiedDate  string        `json:"lastModified"` // payment

	UserData []UserData `json:"userData"` // payment & ticket

	TicketID           ID         `json:"-"`
	FirstName          string     `json:"firstName"`         // ticket
	LastName           string     `json:"lastName"`          // ticket
	Email              string     `json:"email"`             // ticket
	CheckedDate        string     `json:"lastChecked"`       // ticket
	CancelledDate      string     `json:"cancelled"`         // ticket
	TicketType         TicketType `json:"ticketType"`        // ticket
	RegistrationNumber string     `json:"displayIdentifier"` // ticket
*******/

var amiandoData = map[string][]string{
	"Phone":             []string{"Telephone number", "Phone number", "Telephone", "Phone"},
	"Twitter":           []string{"Twitter handle", "Twitter?", "Twitter account"},
	"Facebook":          []string{"Facebook profile?", "Facebook Profile"},
	"LinkedIn":          []string{"LinkedIn profile"},
	"Age":               []string{"Age"},
	"Birthday":          []string{"Date of birth", "Date of Birth"},
	"Gender":            []string{"Gender"},
	"Citizenship":       []string{"Citizenship"},
	"University":        []string{"University"},
	"Company":           []string{"Company"},
	"Background":        []string{"Background"},
	"Other":             []string{"Ohter", "Other?"},
	"IsPitching":        []string{"I will present an idea?", "I will present an idea!", "Will you present an idea?"},
	"StartupName":       []string{"Which one?", "Name"},
	"StartupWebsite":    []string{"Website"},
	"StartupNrEmployee": []string{"Nr. of Employees"},
	"StartupYears":      []string{"Years active"},
	"Accommodation":     []string{"Accommodation", "Do you need an accommodation?"},

	// "TeamName" : []string{"Awesome! Name it!"},
	// "TeamAbstract" : []string{"Abstract for the masses..."},
	// "TeamLogo" : []string{"URL of logo"},
	// "TeamProblem" : []string{"Problem which is solved? Opportunity which is addressed?"},
	// "TeamSolution" : []string{"How do you solve the problem?"},
	// "TeamHaves" : []string{"Have?"},
	// "TeamNeedTechies" : []string{"Need techies?"},
	// "TeamNeedBiz" : []string{"Need biz people?"},
	// "TeamNeedDesign" : []string{"Need designers?"},
	// "TeamNeedSthElse" : []string{"Other needs?"},
	// "TeamPitchtraining" : []string{"Will you attend the pitchtraining?"},
}

func amiandoToModelDate(date string) (string, error) {
	return utils.ConvertTimeString(date, amiando.DateFormat, model.DateFormat)
}

func amiandoToModelDateTime(date string) (string, error) {
	return utils.ConvertTimeString(date, amiando.DateFormat, model.DateTimeFormat)
}

func (self *Event) SyncBasicDataWithAmiando() error {
	api := amiando.NewApi(self.AmiandoEventApiKey.Get())
	amiandoEvent, err := amiando.NewEvent(api, self.AmiandoEventIdentifier.Get())
	if err != nil {
		return err
	}
	return self.syncBasicDataWithAmiando(amiandoEvent)
	return nil
}

func (self *Event) SyncAllDataWithAmiando(logger *log.Logger) error {
	identifier := self.AmiandoEventIdentifier.Get()

	logger.Print("Begin sync Amiando event: ", identifier)
	logger.Println()

	api := amiando.NewApi(self.AmiandoEventApiKey.Get())
	amiandoEvent, err := amiando.NewEvent(api, identifier)
	if err != nil {
		return err
	}

	// deactivated event syncing
	// err = self.syncBasicDataWithAmiando(amiandoEvent)
	// if err != nil {
	// 	return err
	// }

	err = self.syncParticipantsDataWithAmiando(amiandoEvent, logger)
	if err != nil {
		return err
	}

	logger.Print("End sync Amiando event: ", identifier)
	return nil
}

// func (self *Event) SyncPaymentTicketDataWithAmiando(paymentId string, logger *log.Logger) error {
// 	identifier := self.AmiandoEventIdentifier.Get()

// 	api := amiando.NewApi(self.AmiandoEventApiKey.Get())
// 	amiandoEvent, err := amiando.NewEvent(api, identifier)
// 	if err != nil {
// 		return err
// 	}

// 	err = self.syncParticipantsDataPerPaymentWithAmiando(amiandoEvent, paymentId, logger)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func (self *Event) syncBasicDataWithAmiando(amiandoEvent *amiando.Event) error {
	//self.Language.Set(amiandoEvent.Data.Language)

	start, err := amiandoToModelDateTime(amiandoEvent.Data.StartDate)
	if err != nil {
		return err
	}
	self.Start.Set(start)

	end, err := amiandoToModelDateTime(amiandoEvent.Data.EndDate)
	if err != nil {
		return err
	}
	self.End.Set(end)

	title := amiandoEvent.Data.Title
	self.Name.Set(title)
	pos := strings.LastIndex(title, "#")
	if pos >= 0 && pos < len(title)-1 {
		number, err := strconv.ParseInt(title[pos+1:], 10, 64)
		if err == nil {
			self.Number.Set(number)
		}
	}

	var location EventLocation
	_, err = self.Location.TryGet(&location)
	if err != nil {
		return err
	}

	location.Name.Set(amiandoEvent.Data.Location)
	location.Address.FirstLine.Set(amiandoEvent.Data.Street)
	location.Address.ZIP.Set(amiandoEvent.Data.ZipCode)
	location.Address.City.Set(amiandoEvent.Data.City)
	location.GeoLocation.Longitude.Set(amiandoEvent.Data.Longitude)
	location.GeoLocation.Latitude.Set(amiandoEvent.Data.Latitude)
	err = EventLocations.InitAndSaveDocument(&location)
	if err != nil {
		return err
	}

	self.Location.Set(&location)
	return self.Save()
}

func (self *Event) syncParticipantsDataWithAmiando(amiandoEvent *amiando.Event, logger *log.Logger) error {
	numParticipants := 0
	p, e := amiandoEvent.EnumParticipants()
	for participant, ok := <-p; ok; participant, ok = <-p {
		_, err := self.UpdateAmiandoParticipant(logger, participant)
		if err != nil {
			return err
		}
		numParticipants++
	}
	if err, ok := <-e; ok {
		return err
	}

	logger.Printf("Synced %d Amiando participants", numParticipants)
	return nil
}

// func (self *Event) syncParticipantsDataPerPaymentWithAmiando(amiandoEvent *amiando.Event, paymentId string, logger *log.Logger) error {
// 	numParticipants := 0
// 	p, e := amiandoEvent.EnumParticipantsByPayment(paymentId)
// 	for participant, ok := <-p; ok; participant, ok = <-p {
// 		err := self.updateAmiandoParticipant(logger, participant)
// 		if err != nil {
// 			return err
// 		}
// 		numParticipants++
// 	}
// 	if err, ok := <-e; ok {
// 		return err
// 	}

// 	debug.Print("Synced %d Amiando participants", numParticipants)
// 	return nil
// }

func (self *Event) UpdateAmiandoParticipant(logger *log.Logger, amiandoParticipant *amiando.Participant) (*EventParticipant, error) {
	logger.Printf("Update participant: %s %s <%s>", amiandoParticipant.FirstName, amiandoParticipant.LastName, amiandoParticipant.Email)

	// Name and email of participant can already be used in
	// in User or Person database entry.
	// So try and match the data from Amiando with our existing database:

	// Find or create Person
	var person Person
	found, err := user.WithEmail(amiandoParticipant.Email, &person)
	if err != nil {
		return nil, err
	}
	if found {
		logger.Printf("models.Person with email %s already exists", amiandoParticipant.Email)
	} else {
		logger.Print("Creating models.Person")
		People.InitDocument(&person)
	}

	err = self.updatePersonFromAmiandoParticipant(logger, &person, amiandoParticipant)
	if err != nil {
		return nil, err
	}

	//debug.Print("AFTER: updatePersonFromAmiandoParticipant")

	// Find or create EventParticipant
	var eventParticipant EventParticipant
	found, err = EventParticipants.Filter("Event", self.ID).Filter("Person", person.ID).TryOneDocument(&eventParticipant)
	if err != nil {
		return nil, err
	}
	if found {
		logger.Print("models.EventParticipant already exists")
	} else {
		logger.Print("Creating models.EventParticipant")
		EventParticipants.InitDocument(&eventParticipant)
	}
	err = self.updateEventParticipantFromAmiandoParticipant(logger, &person, &eventParticipant, amiandoParticipant)
	if err != nil {
		return nil, err
	}
	err = self.setEventTeamFromAmiandoParticipant(logger, &person, &eventParticipant, amiandoParticipant)

	logger.Println()
	return &eventParticipant, err
}

// sync person participant data
func (self *Event) updatePersonFromAmiandoParticipant(logger *log.Logger, person *Person, amiandoParticipant *amiando.Participant) error {
	person.Name.SetForPerson("", amiandoParticipant.FirstName, "", amiandoParticipant.LastName, "")
	//debug.Print("Person Name: " + person.Name.String())

	if !person.HasEmail(amiandoParticipant.Email) {
		err := person.AddEmail(amiandoParticipant.Email, "via Amiando")
		if err != nil {
			return err
		}
	}

	////
	//// SYNC PERSON DATA
	////

	// PHONE
	aPhoneData := amiandoData["Phone"]
	for i := 0; i < len(aPhoneData); i++ {
		if phone, ok := amiandoParticipant.FindUserData(aPhoneData[i]); ok {
			if !person.HasPhone(phone.String()) {
				person.AddPhone(phone.String(), "via Amiando")
			}
			break
		}
	}

	// TWITTER
	aTwitterData := amiandoData["Twitter"]
	for i := 0; i < len(aTwitterData); i++ {
		if amiandotwitter, ok := amiandoParticipant.FindUserData(aTwitterData[i]); ok {
			twitter := amiandotwitter.String()
			pos := strings.LastIndex(twitter, "@")
			if pos != -1 {
				twitter = twitter[pos+1:]
			}
			twitterfound := false
			for j := 0; j < len(person.Twitter); j++ {
				if person.Twitter[j].Name.String() == twitter {
					twitterfound = true
				}
			}

			if !twitterfound {
				var twitterIdentity user.TwitterIdentity
				twitterIdentity.Name.Set(twitter)

				person.Twitter = append(person.Twitter, twitterIdentity)
			}

			break
		}
	}

	// FACEBOOK
	aFacebookData := amiandoData["Facebook"]
	for i := 0; i < len(aFacebookData); i++ {
		if facebook, ok := amiandoParticipant.FindUserData(aFacebookData[i]); ok {
			facebookfound := false
			for j := 0; j < len(person.Facebook); j++ {
				if person.Facebook[j].Name.String() == facebook.String() {
					facebookfound = true
				}
			}

			if !facebookfound {
				var facebookIdentity user.FacebookIdentity
				facebookIdentity.Name.Set(facebook.String())

				person.Facebook = append(person.Facebook, facebookIdentity)
			}

			break
		}
	}

	// Linked In
	aLinkedInData := amiandoData["LinkedIn"]
	for i := 0; i < len(aLinkedInData); i++ {
		if linkedin, ok := amiandoParticipant.FindUserData(aLinkedInData[i]); ok {
			linkedinfound := false
			for j := 0; j < len(person.LinkedIn); j++ {
				if person.LinkedIn[j].Name.String() == linkedin.String() {
					linkedinfound = true
				}
			}

			if !linkedinfound {
				var linkedinIdentity user.LinkedInIdentity
				linkedinIdentity.Name.Set(linkedin.String())

				person.LinkedIn = append(person.LinkedIn, linkedinIdentity)
			}

			break
		}
	}

	// Address
	if address, ok := amiandoParticipant.FindUserData("Address"); ok {
		a := address.Address()
		if a.Street != "" {
			person.PostalAddress.FirstLine.Set(a.Street)
		}
		if a.Street2 != "" {
			person.PostalAddress.SecondLine.Set(a.Street2)
		}
		if a.City != "" {
			person.PostalAddress.City.Set(a.City)
		}
		if a.ZipCode != "" {
			person.PostalAddress.ZIP.Set(a.ZipCode)
		}
		if a.Country != "" {
			person.PostalAddress.Country.Set(a.Country)
		}
	}

	// Citizenship
	aCitizenshipData := amiandoData["Citizenship"]
	for i := 0; i < len(aCitizenshipData); i++ {
		if country, ok := amiandoParticipant.FindUserData(aCitizenshipData[i]); ok {
			person.Citizenship.Set(country.String())
		}
	}

	// Gender
	aGenderData := amiandoData["Gender"]
	for i := 0; i < len(aGenderData); i++ {
		if gender, ok := amiandoParticipant.FindUserData(aGenderData[i], amiando.UserDataGender); ok {
			switch gender.Value.(float64) {
			case amiando.Male:
				person.Gender.Set("Male")

			case amiando.Female:
				person.Gender.Set("Female")
			}

			break
		}
	}

	// Age
	aAgeData := amiandoData["Gender"]
	for i := 0; i < len(aAgeData); i++ {
		if age, ok := amiandoParticipant.FindUserData(aAgeData[i]); ok {
			a, err := strconv.ParseInt(age.String(), 10, 64)
			if err != nil {
				return err
			}
			person.BirthYear.Set(int64(time.Now().UTC().Year()) - a)

			break
		}
	}

	// Date of Birth
	aBirthData := amiandoData["Birthday"]
	for i := 0; i < len(aBirthData); i++ {
		if dateOfBirth, ok := amiandoParticipant.FindUserData(aBirthData[i]); ok {
			date, err := amiandoToModelDate(dateOfBirth.String())
			if err != nil {
				return err
			}
			err = person.BirthDate.Set(date)
			if err != nil {
				return err
			}
			t, err := time.Parse(amiando.DateFormat, dateOfBirth.String())
			if err != nil {
				return err
			}
			person.BirthYear.Set(int64(t.Year()))
		}
	}

	// University
	aUniData := amiandoData["University"]
	for i := 0; i < len(aUniData); i++ {
		if university, ok := amiandoParticipant.FindUserData(aUniData[i]); ok {
			person.University.Set(university.String())

			break
		}
	}

	// University
	aCompanyData := amiandoData["Company"]
	for i := 0; i < len(aCompanyData); i++ {
		if company, ok := amiandoParticipant.FindUserData(aCompanyData[i]); ok {
			person.Company.Set(company.String())

			break
		}
	}

	return person.Save()
}

// sync event participant data
func (self *Event) updateEventParticipantFromAmiandoParticipant(logger *log.Logger, person *Person, eventParticipant *EventParticipant, amiandoParticipant *amiando.Participant) error {
	eventParticipant.Event.Set(self)
	eventParticipant.Person.Set(person)

	date, err := amiandoToModelDateTime(amiandoParticipant.CreatedDate)
	if err != nil {
		return err
	}
	err = eventParticipant.AppliedDate.Set(date)
	if err != nil {
		return err
	}

	// Background
	aBackgroundData := amiandoData["Background"]
	for i := 0; i < len(aBackgroundData); i++ {
		if background, ok := amiandoParticipant.FindUserData(aBackgroundData[i]); ok {
			eventParticipant.Background.Set(background.String())

			break
		}
	}

	// Other Background
	aOtherData := amiandoData["Other"]
	for i := 0; i < len(aOtherData); i++ {
		if bgother, ok := amiandoParticipant.FindUserData(aOtherData[i]); ok {
			eventParticipant.Background2.Set(bgother.String())

			break
		}
	}

	// Pitching?
	aIsPichtingData := amiandoData["IsPitching"]
	for i := 0; i < len(aIsPichtingData); i++ {
		if isPitching, ok := amiandoParticipant.FindUserData(aIsPichtingData[i]); ok {
			if isPitching.String() == "yes" || isPitching.String() == "Yes" {
				eventParticipant.PresentsIdea = true
				logger.Printf("Presents an idea")

				// aTeamNameData := amiandoData["TeamName"]
				// for j:=0; j < len(aTeamNameData); j++ {
				// 	if teamname, ok := amiandoParticipant.FindUserData(aTeamNameData[j]); ok {
				// 		team, created := createTeamFromAmiando(amiandoParticipant)

				// 		if created {
				// 			person.Team = team.Ref()
				// 		}

				// 	}
				// }
			}
			break
		}
	}

	// Startup Name
	aStartupNameData := amiandoData["StartupName"]
	for i := 0; i < len(aStartupNameData); i++ {
		if startupname, ok := amiandoParticipant.FindUserData(aStartupNameData[i]); ok {
			eventParticipant.Startup.Name.Set(startupname.String())

			break
		}
	}

	// Startup Website
	aStartupWebsiteData := amiandoData["StartupWebsite"]
	for i := 0; i < len(aStartupWebsiteData); i++ {
		if startupWebsite, ok := amiandoParticipant.FindUserData(aStartupWebsiteData[i]); ok {
			eventParticipant.Startup.Website.Set(startupWebsite.String())

			break
		}
	}

	// Startup #Employee
	aStartupEmployeeData := amiandoData["StartupNrEmployee"]
	for i := 0; i < len(aStartupEmployeeData); i++ {
		if startupNrEmployees, ok := amiandoParticipant.FindUserData(aStartupEmployeeData[i]); ok {
			eventParticipant.Startup.NrEmployees.Set(startupNrEmployees.String())

			break
		}
	}

	// Startup Years active
	aStartupYearsData := amiandoData["StartupYears"]
	for i := 0; i < len(aStartupYearsData); i++ {
		if startupYearsActive, ok := amiandoParticipant.FindUserData(aStartupYearsData[i]); ok {
			eventParticipant.Startup.YearsActive.Set(startupYearsActive.String())

			break
		}
	}

	// Accommodation
	aAccomodationData := amiandoData["Accommodation"]
	for i := 0; i < len(aAccomodationData); i++ {
		accommodation, ok := amiandoParticipant.FindUserData(aAccomodationData[i])
		if ok {

			eventParticipant.Accommodation.Set(accommodation.String())

			break
		}

	}

	// Ticket Data
	eventParticipant.Ticket.AmiandoTicketID.Set(amiandoParticipant.TicketID.String())
	eventParticipant.Ticket.Type.Set(string(amiandoParticipant.TicketType))
	eventParticipant.Ticket.InvoiceNumber.Set(amiandoParticipant.InvoiceNumber)
	eventParticipant.Ticket.RegistrationNumber.Set(amiandoParticipant.RegistrationNumber)
	if amiandoParticipant.CheckedDate != "" {
		date, err := amiandoToModelDateTime(amiandoParticipant.CheckedDate)
		if err != nil {
			return err
		}
		err = eventParticipant.Ticket.CheckedDate.Set(date)
		if err != nil {
			return err
		}
	}
	if amiandoParticipant.CancelledDate != "" {
		date, err := amiandoToModelDateTime(amiandoParticipant.CancelledDate)
		if err != nil {
			return err
		}
		err = eventParticipant.Ticket.CancelledDate.Set(date)
		if err != nil {
			return err
		}
	}

	return eventParticipant.Save()
}

const (
	eventTeamNameQuestion               = "Awesome! Name it!"
	eventTeamLogoURLQuestion            = "URL of logo"
	eventTeamAbstractQuestion           = "Abstract for the masses..."
	eventTeamPitchtrainingQuestion      = "Will you attend the pitchtraining?"
	eventTeamProblemOpportunityQuestion = "Problem which is solved? Opportunity which is addressed?"
	eventTeamSolutionQuestion           = "How do you solve the problem?"
	eventTeamHavesQuestion              = "Have?"
	eventTeamNeedTechiesQuestion        = "Need techies?"
	eventTeamNeedBizPeopleQuestion      = "Need biz people?"
	eventTeamNeedDesignersQuestion      = "Need designers?"
	eventTeamNeedOtherQuestion          = "Other needs?"
)

func (self *Event) setEventTeamFromAmiandoParticipant(logger *log.Logger, person *Person, eventParticipant *EventParticipant, amiandoParticipant *amiando.Participant) error {
	if eventParticipant.PresentsIdea {

		idea, err := amiandoParticipant.FindRequiredUserData(eventTeamNameQuestion)
		if err != nil {
			return nil
		}
		teamName := idea.String()

		logger.Printf("Presenting idea: %s", teamName)

		teamQuery := EventTeams.FilterEqualCaseInsensitive("Name", teamName)
		count, err := teamQuery.Count()
		if err != nil {
			return nil
		}

		var eventTeam EventTeam
		if count > 0 {
			logger.Print("models.EventTeam for idea already exists")
			err := teamQuery.OneDocument(&eventTeam)
			if err != nil {
				return nil
			}
		} else {
			logger.Printf("Creating models.EventTam for idea: %s", teamName)
			eventTeam.Event.Set(self)
			eventTeam.Leader.Set(person)
			eventTeam.Name.Set(teamName)

			if logoUrl, ok := amiandoParticipant.FindUserData(eventTeamLogoURLQuestion); ok {
				image, err := media.NewImageFromURL(logoUrl.String())
				if err == nil {
					eventTeam.Logo.Set(image)
					eventTeam.LogoURL.Set("")
				} else {
					// Maybe the URL works later on:
					eventTeam.Logo.Set(nil)
					eventTeam.LogoURL.Set(logoUrl.String())
				}
			}

			if abstract, ok := amiandoParticipant.FindUserData(eventTeamAbstractQuestion); ok {
				eventTeam.Abstract.Set(abstract.String())
			}

			if pitchtraining, ok := amiandoParticipant.FindUserData(eventTeamPitchtrainingQuestion, amiando.UserDataCheckbox); ok {
				err = eventTeam.PitchtrainingBooked.SetString(pitchtraining.String())
				if err != nil {
					return nil
				}
			}

			if problemOpportunity, ok := amiandoParticipant.FindUserData(eventTeamProblemOpportunityQuestion); ok {
				eventTeam.ProblemOpportunity.Set(problemOpportunity.String())
			}

			if solution, ok := amiandoParticipant.FindUserData(eventTeamSolutionQuestion); ok {
				eventTeam.Solution.Set(solution.String())
			}

			if haves, ok := amiandoParticipant.FindUserData(eventTeamHavesQuestion); ok {
				eventTeam.Haves.Set(haves.String())
			}

			if needTechies, ok := amiandoParticipant.FindUserData(eventTeamNeedTechiesQuestion); ok {
				err = eventTeam.NeedTechies.SetString(needTechies.String())
				if err != nil {
					return nil
				}
			}

			if needBizPeople, ok := amiandoParticipant.FindUserData(eventTeamNeedBizPeopleQuestion); ok {
				err = eventTeam.NeedBizPeople.SetString(needBizPeople.String())
				if err != nil {
					return nil
				}
			}

			if needDesigners, ok := amiandoParticipant.FindUserData(eventTeamNeedDesignersQuestion); ok {
				err = eventTeam.NeedDesigners.SetString(needDesigners.String())
				if err != nil {
					return nil
				}
			}

			if needOther, ok := amiandoParticipant.FindUserData(eventTeamNeedOtherQuestion); ok {
				eventTeam.NeedOther.Set(needOther.String())
			}

			eventTeam.Pitching = true

			err = EventTeams.InitAndSaveDocument(&eventTeam)
			if err != nil {
				return err
			}
		}

		eventParticipant.Team.Set(&eventTeam)
		return eventParticipant.Save()
	}

	return nil
}
