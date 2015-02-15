package dashboard

import (
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/ungerik/go-start/utils"
	. "github.com/ungerik/go-start/view"
	"strings"
	"time"
)

var ReadMoreArrow HTML = "read more &rarr;"

func init() {
	Region0_Event1_Dashboard = &Page{
		OnPreRender: SetEventPageData,
		Title: Render(
			func(ctx *Context) (err error) {
				event := ctx.Data.(*PageData).Event

				ctx.Response.Printf("Dashboard | %s", event.Name)
				return nil
			},
		),
		CSS: IndirectURL(&Region0_DashboardCSS),
		Content: Views{
			DashboardHeader(),
			&Div{
				Class: "content",
				Content: Views{
					&Div{
						Class: "main timetable",
						Content: Views{
							&Div{
								Class: "main-header",
								Content: Views{
									DynamicView(func(ctx *Context) (view View, err error) {
										region := ctx.Data.(*PageData).Region
										return H1(region.Name.Get() + " Timetable"), nil
									}),
									//					schedule/		&Form{
									//								Class: "search-mentor",
									//								Content: &TextField{},
									//							},
								},
							},

							DynamicView(
								func(ctx *Context) (view View, err error) {
									event := ctx.Data.(*PageData).Event

									minHour := 8
									maxHour := 25
									i := event.ScheduleItemIterator()
									var item models.EventScheduleItem
									for i.Next(&item) {
										fromHour := item.From.Time().Hour()
										if fromHour < minHour {
											minHour = fromHour
										}
										untilHour := item.Until.Time().Hour()
										if untilHour < fromHour {
											untilHour += 24
										}
										if untilHour > maxHour {
											maxHour = untilHour
										}
									}
									if i.Err() != nil {
										return nil, i.Err()
									}

									timeRows := Views{HTML("<div class='row'></div>")}
									for i := minHour; i < maxHour; i++ {
										hour := i % 24
										timeRows = append(timeRows, Printf("<div class='row'>%02d:00</div>", hour))
									}

									views := Views{
										Printf("<style>.timetable .col {height: %dpx;}</style>", (maxHour-minHour+1)*60),
										DIV("col time-col", timeRows),
									}

									for day := utils.DayBeginningTime(event.Start.Time()); day.Before(event.End.Time()); day = day.Add(time.Hour * 24) {
										var itemViews Views
										i := event.DaySchedulItemIterator(day)
										var item models.EventScheduleItem
										for i.Next(&item) {
											var class string
											if len(itemViews)%2 == 0 {
												class = "primary"
											} else {
												class = "secondary"
											}
											from := item.From.Time()
											until := item.Until.Time()
											itemViews = append(itemViews, newTimetableItem(from, until, item.Location.Get(), item.Title.Get(), "", class))
										}
										if i.Err() != nil {
											return nil, i.Err()
										}

										view := DIV("col day-col",
											DIV("row header-row",
												H2(HTML(day.Format("Monday"))),
												HTML(strings.ToUpper(day.Format("January 2, 2006"))),
											),
											itemViews,
										)
										views = append(views, view)
									}

									return views, nil
								},
							),
							DivClearBoth(),
						},
					},
					DivClearBoth(),
				},
			},
			eventDashboardFooter(),
		},
	}
}

func newTimetableItem(from, until time.Time, where, title, description, class string) View {
	top := from.Hour()*60 + from.Minute()
	height := until.Hour()*60 + until.Minute() - top - 3
	if height < 0 { // next day
		height += 24 * 60
	}
	top -= 7 * 60

	lineHeight := 10
	if description != "" {
		lineHeight += 16
	}
	marginTop := (height - lineHeight) / 2
	if marginTop < 12 {
		marginTop = 12
	}

	if height < 60 {
		class += " small"
		if height <= 45 {
			class += "er"
		}
	}

	//where = html.EscapeString(where)
	//title = html.EscapeString(title)
	//description = html.EscapeString(description)

	return Views{
		Printf("<div class='item %s' style='top:%dpx;height:%dpx'>", class, top, height),
		Printf("<div class='time'>%02d:%02d - %02d:%02d</div><div class='place'>%s</div>", from.Hour(), from.Minute(), until.Hour(), until.Minute(), where),
		Printf("<h3 style='margin-top:%dpx'>%s</h3>%s", marginTop, title, description),
		DivClearBoth(),
		HTML("</div>"),
	}
}
