package models

import (
	"github.com/ungerik/go-start/config"
	// "github.com/ungerik/go-start/debug"
	"github.com/ungerik/go-start/model"
	"github.com/ungerik/go-start/mongo"
)

var EventParticipants = mongo.NewCollection("eventparticipants")

func NewEventParticipant() *EventParticipant {
	var doc EventParticipant
	EventParticipants.InitDocument(&doc)
	return &doc
}

type EventParticipant struct {
	mongo.DocumentBase `bson:",inline"`
	CheckedIn          model.Bool `view:"label=checked in"`
	Team               mongo.Ref  `model:"to=eventteams"`
	Event              mongo.Ref  `model:"to=events"`
	Person             mongo.Ref  `model:"to=people"`
	AppliedDate        model.DateTime
	CancelledDate      model.DateTime
	CancelledBy        mongo.Ref `model:"to=people"`
	PresentsIdea       model.Bool
	Background         model.String
	Background2        model.String
	Accommodation      model.String

	Ticket struct {
		AmiandoTicketID    model.String
		Type               model.String
		CheckedDate        model.DateTime
		CancelledDate      model.DateTime
		InvoiceNumber      model.String
		RegistrationNumber model.String
	}
	Startup struct {
		Name        model.String `view:"placeholder=Name of your startup"`
		Website     model.Url
		NrEmployees model.String
		YearsActive model.String
	}
}

func (self *EventParticipant) TicketChecked() bool {
	return !self.Ticket.CheckedDate.IsEmpty()
}

func (self *EventParticipant) Cancelled() bool {
	return !self.CancelledDate.IsEmpty() || !self.Ticket.CancelledDate.IsEmpty()
}

func (self *EventParticipant) Cancel(by *Person) {
	if by != nil {
		self.CancelledBy.Set(by)
	}
	self.CancelledDate.SetNowUTC()
}

func (self *EventParticipant) Uncancel() {
	self.CancelledBy.Set(nil)
	self.CancelledDate.SetEmpty()
	self.Ticket.CancelledDate.SetEmpty()
}

func (self *EventParticipant) Name() string {
	var person Person
	_, err := self.Person.TryGet(&person)
	if err != nil {
		return err.Error()
	}
	return person.Name.String()
}

func (self *EventParticipant) String() string {
	return self.Name()
}

func (self *EventParticipant) GetPerson() *Person {
	var person Person
	// config.Logger.Println("in GetPerson")
	found, err := self.Person.TryGet(&person)
	if err != nil {
		config.Logger.Println("error: ", err.Error())
	}
	if !found {
		return nil
	}
	return &person
}
