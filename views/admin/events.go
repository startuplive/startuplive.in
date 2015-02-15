package admin

import (
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/ungerik/go-start/config"
	// "github.com/ungerik/go-start/debug"
	"github.com/ungerik/go-start/model"
	. "github.com/ungerik/go-start/view"
	"strconv"
)

func init() {
	Admin_Events = &Page{
		Title: Escape("Events | Admin"),
		CSS:   IndirectURL(&Admin_CSS),
		Scripts: Renderers{
			PageScripts,
		},
		Content: Views{
			adminHeader(),
			H3("Current Events"),
			DynamicView(
				func(ctx *Context) (view View, err error) {
					events := models.CurrentPublicStartupLiveEventIterator()
					return renderEvents(events)
				},
			),
			HR(),
			H3("Upcoming Events"),
			DynamicView(
				func(ctx *Context) (view View, err error) {
					events := models.UpcomingPublicStartupLiveEventIterator()
					return renderEvents(events)
				},
			),
			HR(),
			H3("Past Events"),
			DynamicView(
				func(ctx *Context) (view View, err error) {
					events := models.PastPublicStartupLiveEventIterator()
					return renderEvents(events)
				},
			),
		},
	}
}

func renderEvents(events model.Iterator) (view View, err error) {
	var views Views
	views = append(views,
		&ModelIteratorTableView{
			Class: "visual-table",
			GetModelIterator: func(ctx *Context) model.Iterator {
				return events
			},
			GetRowModel: func(ctx *Context) (interface{}, error) {
				return new(models.Event), nil
			},
			GetHeaderRowViews: TableHeaderRowEscape("Event", "#P", "#T", "#M", "#J", "Status", "$$", "Eventpage"),
			GetRowViews: func(row int, rowModel interface{}, ctx *Context) (views Views, err error) {
				event := rowModel.(*models.Event)
				var region models.EventRegion
				err = event.Region.Get(&region)
				if err != nil {
					return nil, err
				}

				eventURL := Region0_Event1_Admin.URL(ctx.ForURLArgs(region.Slug.Get(), event.Number.String()))

				/* Statistics */
				//nr. of participants
				nrOfParticipants, _ := event.ParticipantsCount()

				//nr. of judges
				nrOfJudges := len(event.Judges)
				//nr. of mentors
				nrOfMentors := len(event.Mentors)
				//nr. of teams
				nrOfTeams, err := event.TeamsCount()
				if err != nil {
					config.Logger.Println(err.Error())
				}
				//event is live
				eventstatus := "not setup"
				if event.GoLiveRequested.Get() && !event.IsPublished() {
					eventstatus = "ready"
				} else if event.IsPublished() {
					eventstatus = "live"
				}
				//ticketing is running
				ticketing := "not setup"
				if event.Show.Registration {
					ticketing = "live"
				}

				views = Views{
					Escape(event.Name.String() + " | " + event.GetDate()),
					Escape(strconv.Itoa(nrOfParticipants)),
					Escape(strconv.Itoa(nrOfTeams)),
					Escape(strconv.Itoa(nrOfMentors)),
					Escape(strconv.Itoa(nrOfJudges)),
					Escape(eventstatus),
					Escape(ticketing),
					A(eventURL, "Goto"),
				}
				return views, nil
			},
		},
	)

	// for doc := events.Next(); doc != nil; doc = events.Next() {
	// 	event := doc.(*models.Event)

	// 	/* Statistics */
	// 	//nr. of participants
	// 	nrOfParticipants := len(event.ParticipantIterator())
	// 	//nr. of judges
	// 	nrOfJudges := len(event.Judges)
	// 	//nr. of mentors
	// 	nrOfMentors := len(event.Mentors)
	// 	//event is live
	// 	eventstatus := "not setup"
	// 	if event.GoLiveRequested && !event.IsPublished() {
	// 		eventstatus = "ready"
	// 	} else if event.IsPublished() {
	// 		eventstatus = "live"
	// 	}
	// 	//ticketing is running
	// 	ticketing := "not setup"
	// 	if event.Show.Registration {
	// 		ticketing := "live"
	// 	}

	// 	views = append(views,
	// 		Printf("<h4>%s</h4><div>", event.Name, nrOfParticipants, nrOfJudges, nrOfMentors, eventstatus, ticketing),
	// 	)

	// }

	return views, nil
}
