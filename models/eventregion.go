package models

import (
	"fmt"
	"github.com/ungerik/go-start/debug"
	"github.com/ungerik/go-start/model"
	"github.com/ungerik/go-start/mongo"
	"time"
)

var EventRegions = mongo.NewCollection("eventregions", "Name")

func NewEventRegion() *EventRegion {
	var doc EventRegion
	EventRegions.InitDocument(&doc)
	return &doc
}

type EventRegion struct {
	mongo.DocumentBase  `bson:",inline"`
	Name                model.String
	Slug                model.Slug
	Admins              []mongo.Ref `model:"to=people"`
	PrimaryColorIndex   model.Int   `model:"min=0|max=9"`
	SecondaryColorIndex model.Int   `model:"min=0|max=9"`
	FillPatternIndex    model.Int   `model:"min=0|max=8"`
	HeaderLogoURL       model.Url
	PublicHeaderLogoURL model.Url
	HoverLogoURL        model.Url
	InitialURL          model.Url
	InitialURL_60x0     model.Url
	CircleURL           model.Url
	Favicon16x16URL     model.Url
	Favicon57x57URL     model.Url
	Favicon72x72URL     model.Url
	Favicon114x114URL   model.Url
	Favicon129x129URL   model.Url
	LogoSVG             model.String
}

func (self *EventRegion) String() string {
	return self.Name.Get()
}

func (self *EventRegion) ColorScheme() *ColorScheme {
	return NewColorScheme(self.PrimaryColorIndex.GetInt(), self.SecondaryColorIndex.GetInt(), self.FillPatternIndex.GetInt())
}

func (self *EventRegion) Event(eventType string, number int64) (event *Event, found bool, err error) {
	found, err = Events.Filter("Region", self.ID).Filter("Type", eventType).Filter("Number", number).TryOneDocument(&event)
	if !found {
		return nil, false, err
	}
	return event, true, nil
}

func (self *EventRegion) LatestPublicEvent(eventType string) (event *Event, found bool, err error) {
	found, err = Events.Filter("Region", self.ID).Filter("Type", eventType).Filter("Status", EventPublic).Sort("-Number").TryOneDocument(&event)
	// count, err := Events.Filter("Region", self.ID).Filter("Type", eventType).Filter("Status", EventPublic).Count()
	// if err == nil {
	// 	debug.Print("count: ", count)
	// }
	// .SortFunc(
	// 		func(a, b *Event) bool {

	// 			dateA := a.End
	// 			dateB := b.End
	// 			if !dateA.IsEmpty() && !dateB.IsEmpty() {
	// 				dateAtime := dateA.Time()
	// 				dateBtime := dateB.Time()
	// 				if dateAtime.Before(dateBtime) {
	// 					return false
	// 				} else {
	// 					return true
	// 				}
	// 			}
	// 			return false
	// 		},
	// 	),
	if !found {
		return nil, false, err
	}
	return event, true, nil
}

func (self *EventRegion) EventIterator(eventType string) model.Iterator {
	return Events.Filter("Region", self.ID).Filter("Type", eventType).Sort("Number").Iterator()
}

func (self *EventRegion) RegionEventIterator() model.Iterator {
	return Events.Filter("Region", self.ID).Sort("Type").Iterator()
}

func (self *EventRegion) EventCount(eventType string) (count int, err error) {
	return Events.Filter("Region", self.ID).Filter("Type", eventType).Count()
}

func (self *EventRegion) UpcomingPublicStartupLiveEventIterator() model.Iterator {
	now := time.Now()
	return &model.FilterIterator{
		Iterator: Events.Filter("Region", self.ID).Sort("Number").Iterator(),
		PassFilter: func(resultPtr interface{}) bool {
			event := resultPtr.(*Event)
			localNow := now.Add(time.Hour * time.Duration(event.TimeZone.Get()))
			return event.Type == StartupLive && event.Status == EventPublic && event.Start.Time().After(localNow)
		},
	}
}

func (self *EventRegion) CreateEvent(eventType model.Choice) (*Event, error) {
	debug.Nop()
	// todo use location of last event in region
	var location EventLocation
	err := EventLocations.InitAndSaveDocument(&location)
	if err != nil {
		return nil, err
	}

	number, err := Events.Filter("Region", self.ID).Filter("Type", eventType.String()).Count()
	if err != nil {
		return nil, err
	}
	// debug.Print(number)

	eventnumber := int64(number) + 1

	event := NewEvent()
	event.Type = eventType
	event.Status = EventPlanned
	event.Region.Set(self)
	event.Number.Set(eventnumber)

	// debug.Print("event type: ", event.Type)
	var title string
	if eventType == StartupLive {
		title = "Startup Live"
	} else if eventType == StartupLounge {
		title = "Startup Lounge"
	} else if eventType == PioneersFestival {
		title = "Pioneers Festival"
	}
	event.Name.Set(fmt.Sprintf("%s %s #%d", title, self.Name, event.Number))
	event.Location.Set(&location)
	return event, nil
}

func CreateSlug(regionname string) string {
	// result := make([]byte, utf8.RuneCountInString(regionname))
	// i := 0
	// for _, c := range regionname {
	// 	if c >= 'a' && c <= 'z' || c >= '0' && c <= '9' || c == '-' || c == '_' || c == '.' || c == '~' {
	// 		result[i] = byte(c)
	// 	} else if c >= 'A' && c <= 'Z' {
	// 		result[i] = byte(unicode.ToLower(c))
	// 	} else {
	// 		result[i] = '_'
	// 	}
	// 	i++
	// }
	return "string(result)"
}
