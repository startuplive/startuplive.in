package dashboard

import (
	"labix.org/v2/mgo/bson"

	. "github.com/ungerik/go-start/view"

	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/STARTeurope/startuplive.in/views/admin"
	event_admin "github.com/STARTeurope/startuplive.in/views/region_0/event_1/admin"
)

func init() {
	Region0_Event1_Dashboard_Auth = AnyAuthenticator{
		new(admin.Admin_Authenticator),
		new(event_admin.Region0_Event1_Admin_Authenticator),
		new(Region0_Event1_Dashboard_Authenticator),
	}
}

type Region0_Event1_Dashboard_Authenticator struct{}

func (self *Region0_Event1_Dashboard_Authenticator) Authenticate(ctx *Context) (ok bool, err error) {

	return true, nil

	sessionID := ctx.Session.ID()
	if sessionID == "" {
		return false, nil
	}

	_, event, err := RegionAndEvent(ctx.URLArgs)
	if err != nil {
		return false, err
	}

	var participant models.EventParticipant
	found, err := models.EventParticipants.Filter("Event", event.ID).Filter("Person", bson.ObjectIdHex(sessionID)).TryOneDocument(&participant)
	if !found || err != nil {
		return false, err
	}

	ok = !participant.Cancelled() && (!event.Started.Get() || participant.PresentsIdea.Get())
	return ok, nil
}
