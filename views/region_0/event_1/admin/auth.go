package admin

import (
	"labix.org/v2/mgo/bson"

	"github.com/ungerik/go-start/mongo"
	. "github.com/ungerik/go-start/view"

	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/STARTeurope/startuplive.in/views/admin"
	"github.com/STARTeurope/startuplive.in/views/admin/region_0"
)

func init() {
	Region0_Event1_Admin_Auth = AnyAuthenticator{
		new(admin.Admin_Authenticator),
		new(region_0.Admin_Region0_Authenticator),
		new(Region0_Event1_Admin_Authenticator),
	}
}

type Region0_Event1_Admin_Authenticator struct{}

func (self *Region0_Event1_Admin_Authenticator) Authenticate(ctx *Context) (ok bool, err error) {
	sessionID := ctx.Session.ID()
	if sessionID == "" {
		return false, nil
	}
	_, event, err := RegionAndEvent(ctx.URLArgs)
	if err != nil {
		return false, err
	}
	_, ok = mongo.FindRefWithID(event.Organisers, bson.ObjectIdHex(sessionID))
	return ok, nil
}

var OnlyRegionAdminOrEventOrga AnyAuthenticator = AnyAuthenticator{
	new(region_0.Admin_Region0_Authenticator),
	new(Region0_Event1_Admin_Authenticator),
}
