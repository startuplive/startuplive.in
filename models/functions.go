package models

import (
	"github.com/ungerik/go-start/model"
	"github.com/ungerik/go-start/utils"
	"labix.org/v2/mgo/bson"
	"net/url"
	"strings"
	"time"
)

///////////////////////////////////////////////////////////////////////////////
// Event

// PastPublicStartupLiveEventIterator iterates models.Event, use Next(*models.Event)
func PastPublicStartupLiveEventIterator() model.Iterator {
	now := time.Now()
	return &model.FilterIterator{
		// Iterator: Events.Filter("Status", "Public").SortFunc(
		// 	func(a, b *Event) bool {
		// 		dateA := a.End
		// 		dateB := b.End
		// 		if !dateA.IsEmpty() && !dateB.IsEmpty() {
		// 			dateAtime := dateA.Time()
		// 			dateBtime := dateB.Time()
		// 			if dateAtime.Before(dateBtime) {
		// 				return false
		// 			} else {
		// 				return true
		// 			}
		// 		}
		// 		return false
		// 	},
		// ),
		Iterator: Events.Filter("Status", "Public").Sort("-Start").Iterator(),
		PassFilter: func(resultPtr interface{}) (ok bool) {
			event := resultPtr.(*Event)
			localNow := now.Add(time.Hour * time.Duration(event.TimeZone.Get()))
			return event.Type == StartupLive && event.Status == EventPublic && event.End.Time().Before(localNow)
		},
	}
}

// UpcomingPublicStartupLiveEventIterator iterates models.Event, use Next(*models.Event)

func UpcomingPublicStartupLiveEventIterator() model.Iterator {
	now := time.Now()
	return &model.FilterIterator{
		Iterator: Events.Sort("End").Iterator(),
		PassFilter: func(resultPtr interface{}) (ok bool) {
			event := resultPtr.(*Event)
			localNow := now.Add(time.Hour * time.Duration(event.TimeZone.Get()))
			return event.Type == StartupLive && event.Status == EventPublic && event.Start.Time().After(localNow)
		},
	}
}


func UpcomingPublicStartupLiveEventIteratorForHomePage() model.Iterator {
	now := time.Now()
	return &model.FilterIterator{
		Iterator: Events.Sort("End").Iterator(),
		PassFilter: func(resultPtr interface{}) (ok bool) {
			event := resultPtr.(*Event)
			localNow := now.Add(time.Hour * time.Duration(event.TimeZone.Get()))
			return event.Type == StartupLive && event.Status == EventPublic && event.Start.Time().After(localNow)
		},
	}
}


// CurrentPublicStartupLiveEventIterator iterates models.Event, use Next(*models.Event)
func CurrentPublicStartupLiveEventIterator() model.Iterator {
	now := time.Now()
	return &model.FilterIterator{
		Iterator: Events.Sort("Start").Iterator(),
		PassFilter: func(resultPtr interface{}) (ok bool) {
			event := resultPtr.(*Event)
			localNow := now.Add(time.Hour * time.Duration(event.TimeZone.Get()))
			return event.Type == StartupLive && event.Status == EventPublic && utils.TimeInRange(localNow, event.Start.Time(), event.End.Time())
		},
	}
}

// YearPublicStartupLiveEventIterator iterates models.Event, use Next(*models.Event)
func YearPublicStartupLiveEventIterator(year int) model.Iterator {
	return &model.FilterIterator{
		Iterator: Events.Sort("End").Iterator(),
		PassFilter: func(resultPtr interface{}) (ok bool) {
			event := resultPtr.(*Event)
			return event.Type == StartupLive && event.Status == EventPublic && event.Start.Time().Year() == year
		},
	}
}

func GetEventRegion(slug string) (region *EventRegion, found bool, err error) {
	s, err := url.QueryUnescape(slug)
	if err != nil {
		return nil, false, err
	}
	found, err = EventRegions.Filter("Slug", strings.ToLower(s)).TryOneDocument(&region)
	if !found {
		return nil, false, err
	}
	return region, true, nil
}

func GetRegionAndEvent(eventType string, regionSlug string, number int64) (region *EventRegion, event *Event, found bool, err error) {
	region, found, err = GetEventRegion(regionSlug)
	if !found {
		return nil, nil, false, err
	}
	event, found, err = region.Event(eventType, number)
	if !found {
		return nil, nil, false, err
	}
	return region, event, true, nil
}

///////////
// Partner Categories

func GetPartnersIterator() model.Iterator {
	// return PartnerCategories.Filter("Event", event.ID).Sort("Order").Iterator()
	return Partners.Iterator()
}

func GetPartnersByEventIterator(event *Event) model.Iterator {
	// return PartnerCategories.Filter("Event", event.ID).Sort("Order").Iterator()
	return Partners.Filter("Events", event.ID).Iterator()
}

///////////
// WIKI

func GetAllWikiEntries() model.Iterator {

	return Wiki.Sort("CreatedAt").Iterator()

}

func GetAllPublicWikiEntries() model.Iterator {

	return Wiki.Filter("Public", true).Sort("CreatedAt").Iterator()

}

///////////
// EVENTPARTICIPANT

func GetParticipantById(id bson.ObjectId) (EventParticipant, error) {
	var resultRef EventParticipant
	_, err := EventParticipants.Filter("_id", id).TryOneDocument(&resultRef)
	return resultRef, err
}
