package region_0

import (
	"bytes"
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/STARTeurope/startuplive.in/views/admin"
	. "github.com/ungerik/go-start/view"
	"labix.org/v2/mgo/bson"
	"log"
)

func init() {
	Admin_Region0_SyncEvent1 = &Page{
		Title:       RegionTitle("Sync Event"),
		CSS:         IndirectURL(&Region0_DashboardCSS),
		OnPreRender: SetRegionPageData,
		Scripts:     admin.PageScripts,
		Content: Views{
			Header(),
			DynamicView(SyncEventView),
		},
	}
}

func SyncEventView(ctx *Context) (view View, err error) {
	id := bson.ObjectIdHex(ctx.URLArgs[1])
	var event models.Event
	err = models.Events.DocumentWithID(id, &event)
	if err != nil {
		return nil, err
	}

	var region models.EventRegion
	err = event.Region.Get(&region)
	if err != nil {
		return nil, err
	}

	regionURL := Admin_Region0.URL(ctx.ForURLArgs(region.Slug.Get()))

	var syncLog string
	var buf bytes.Buffer
	err = event.SyncAllDataWithAmiando(log.New(&buf, "", 0))
	if err == nil {
		syncLog = buf.String()
	} else {
		syncLog = err.Error()
	}

	views := Views{
		Printf("<h3><a href='%s'>Back to %s</a></h3>", regionURL, region.Name.Get()),
		Printf("<pre>%s</pre>", syncLog),
		Printf("<h3><a href='%s'>Back to %s</a></h3>", regionURL, region.Name.Get()),
	}
	return views, nil
}
