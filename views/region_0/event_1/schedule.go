package event_1

import (
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	. "github.com/ungerik/go-start/view"
)

func init() {
	Region0_Event1_Schedule = newPublicEventPage(EventTitle("Schedule"), nil, DynamicView(
		func(ctx *Context) (view View, err error) {
			event := ctx.Data.(*PageData).Event
			region := ctx.Data.(*PageData).Region
			days, err := event.Days()
			if err != nil {
				return nil, err
			}
			var views Views
			scheduleHasItems := false
			for _, day := range days {

				var daySponsor View
				if region.Slug == "split" && event.Number == 1 && day.Day() == 2 {
					daySponsor = A_blank("http://www.netokracija.com", IMG("http://dl.dropbox.com/u/5565424/Netokracija.jpg", 0, 50))
				}

				tableModel := ViewsTableModel{{HTML("Time"), HTML(""), HTML("What")}}
				i := event.DaySchedulItemIterator(day)
				var item models.EventScheduleItem
				for i.Next(&item) {
					scheduleHasItems = true
					from := item.From.Format("15:04")
					until := item.Until.Format(" - 15:04")
					tableModel = append(tableModel, Views{HTML(from), HTML(until), HTML(item.Title)})
				}
				if i.Err() != nil {
					return nil, i.Err()
				}
				views = append(views,
					DIV("schedule-day",
						TitleBar(day.Format("Monday, January 2")),
						daySponsor,
						&Table{Class: "schedule", Model: tableModel, HeaderRow: true},
					),
				)
			}
			if !scheduleHasItems {
				views = Views{
					DIV("schedule-day",
						TitleBar("Schedule"),
						H1("Schedule will be available during the next days"),
					),
				}
			}
			return DIV("main", views), nil
		},
	))
}
