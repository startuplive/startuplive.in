package region_0

import (
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	// "github.com/ungerik/go-start/debug"
	. "github.com/ungerik/go-start/view"
)

func init() {
	Region0 = NewViewURLWrapper(
		DynamicView(
			func(ctx *Context) (view View, err error) {
				region, err := EventRegion(ctx.URLArgs)
				if err != nil {
					return nil, err
				}
				event, found, err := region.LatestPublicEvent(models.StartupLive)
				// debug.Print("latest public event: ", event.Name)
				if err != nil {
					return nil, err
				}
				if !found {
				        // TODO: Make it more beautiful! Don't hate me -Lutfi D. 
					//event, found, err := region.LatestPublicEvent(models.LiveAcademy)
					// debug.Print("latest public event: ", event.Name)
					//if err != nil {
					//	return nil, err
					//}
					//if !found {
						return nil, NotFound("404: No event found!")
					//}
				}
				url := Region0_Event1.URL(ctx.ForURLArgs(region.Slug.Get(), event.Number.String()))
				url = event.RoundupURL.GetOrDefault(url)
				return nil, Redirect(url)
			},
		),
	)
}
