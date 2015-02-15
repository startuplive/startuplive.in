package events

import (

	//	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	. "github.com/ungerik/go-start/view"
)

func init() {
	Events_Application = NewPublicPage("Your City", DynamicView(
		func(ctx *Context) (view View, err error) {
			return view, nil
		},
	))
}
