package admin

import (
	"github.com/ungerik/go-start/user"
	. "github.com/ungerik/go-start/view"

	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
)

type Admin_Authenticator struct{}

func (self *Admin_Authenticator) Authenticate(ctx *Context) (ok bool, err error) {
	var person models.Person
	found, err := user.OfSession(ctx.Session, &person)
	if !found {
		return false, err
	}
	return person.Admin.Get() || person.SuperAdmin.Get(), nil
}

type SuperAdmin_Authenticator struct{}

func (self *SuperAdmin_Authenticator) Authenticate(ctx *Context) (ok bool, err error) {
	var person models.Person
	found, err := user.OfSession(ctx.Session, &person)
	if !found {
		return false, err
	}
	return person.SuperAdmin.Get(), nil
}

func init() {
	Admin_Auth = new(Admin_Authenticator)
	SuperAdmin_Auth = new(SuperAdmin_Authenticator)
}
