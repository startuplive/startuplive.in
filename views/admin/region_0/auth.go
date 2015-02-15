package region_0

import (
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/STARTeurope/startuplive.in/views/admin"
	"github.com/ungerik/go-start/mongo"
	"github.com/ungerik/go-start/user"
	. "github.com/ungerik/go-start/view"
)

type Admin_Region0_Authenticator struct{}

func (self *Admin_Region0_Authenticator) Authenticate(ctx *Context) (ok bool, err error) {
	var person models.Person
	found, err := user.OfSession(ctx.Session, &person)
	if err != nil {
		return false, err
	}
	if !found {
		return false, nil
	}
	region, err := EventRegion(ctx.URLArgs)
	if err != nil {
		return false, err
	}

	_, ok = mongo.FindRefWithID(region.Admins, person.ID)
	return ok, nil
}

func init() {
	Admin_Region0_Auth = AnyAuthenticator{
		new(admin.Admin_Authenticator),
		new(Admin_Region0_Authenticator),
	}
}
