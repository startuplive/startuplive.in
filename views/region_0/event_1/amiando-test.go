package event_1

import (

	//	"github.com/STARTeurope/startuplive.in/models"
	"github.com/AlexTi/go-amiando"
	. "github.com/STARTeurope/startuplive.in/views"
	. "github.com/ungerik/go-start/view"
)

func init() {
	Region0_Event1_AmiandoTest = newPublicEventPage(EventTitle("Amiando Test"), nil, DynamicView(
		func(ctx *Context) (view View, err error) {
			region, event, err := RegionAndEvent(ctx.URLArgs)
			if err != nil {
				return nil, err
			}
			amiando.TestAmiandoWebHook("http://0.0.0.0:8080/" + region.Slug.String() + "/" + event.Number.String() + "/amiando-tracking")

			return HTML("sent"), nil
		},
	))
}
