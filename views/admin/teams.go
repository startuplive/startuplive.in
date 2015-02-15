package admin

import (
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	// "github.com/ungerik/go-start/model"
	// "github.com/ungerik/go-start/utils"
	"github.com/ungerik/go-start/user"
	. "github.com/ungerik/go-start/view"
)

func init() {
	Admin_Teams = &Page{
		Title: Escape("Teams | Admin"),
		CSS:   IndirectURL(&Admin_CSS),
		Scripts: Renderers{
			PageScripts,
			RequireScript(
				`$(document).ready(function() {
					$('.toggle').click(function(){
						$($(this).parent().next()).toggle(500);
					});
				});`,
				0,
			),
		},
		Content: Views{
			adminHeader(),
			DynamicView(regionsList),
		},
	}
}

func regionsList(ctx *Context) (view View, err error) {
	regions := RegionsIterator()
	var views Views

	var region models.EventRegion
	var totalcount = 0
	for regions.Next(&region) {
		eventslist, eventscount, err := eventsList(ctx, &region)
		totalcount += eventscount
		if err != nil {
			return nil, nil
		}
		views = append(views,
			Printf("<h3><span class='toggle' style='cursor:pointer; background:gray; padding:5px; color:white'>v</span> <a href='../%s/'>%s</a> (%v)</h3><div>", region.Slug, region.Name, eventscount),
			eventslist,
			Printf("</div>"),
		)
	}
	if regions.Err() != nil {
		return nil, regions.Err()
	}

	var topview Views
	topview = append(topview, Printf("<h3>Total amount of (registered) teams: %d</h3>", totalcount))

	return append(topview, views...), nil
}

func eventsList(ctx *Context, region *models.EventRegion) (view View, count int, err error) {
	events := region.EventIterator(models.StartupLive)
	var views Views

	count = 0

	var event models.Event
	for events.Next(&event) {
		slug := region.Slug.String()
		number := event.Number.String()

		eventPublicURL := Region0_Event1.URL(ctx.ForURLArgs(slug, number))
		eventDashboardURL := Region0_Event1_Dashboard.URL(ctx.ForURLArgs(slug, number))
		eventAdminURL := Region0_Event1_Admin.URL(ctx.ForURLArgs(slug, number))

		teamslist, teamcount, err := teamsList(ctx, &event, region)
		count = count + teamcount
		if err != nil {
			return nil, 0, nil
		}
		views = append(views,
			Printf("<h4>%s<a href='%s'>Public Site</a>, <a href='%s'>Dashboard</a>, <a href='%s'>Admin</a></h4>", event.Name, eventPublicURL, eventDashboardURL, eventAdminURL),
			teamslist,
		)
	}
	if events.Err() != nil {
		return nil, 0, events.Err()
	}

	return views, count, nil
}

func teamsList(ctx *Context, event *models.Event, region *models.EventRegion) (view View, teamcount int, err error) {

	teams := event.TeamIterator()
	var views Views

	teamcount = 0
	views = append(views,
		HTML("<table>"),
	)
	var t models.EventTeam
	for teams.Next(&t) {
		team := t // copy by value because it will be used in a closure later on
		editURL := Region0_Event1_Admin_Team2.URL(ctx.ForURLArgs(region.Slug.String(), event.Number.String(), team.ID.Hex()))
		var person models.Person
		found, err := user.OfSession(ctx.Session, &person)
		if err != nil {
			return nil, 0, err
		}
		if !found {
			return HTML("You have to be logged in"), 0, nil
		}
		teamcount++
		views = append(views,
			Printf("<tr><td>%v</td><td>%s</td><td>%s</td><td>%s</td><td><a href='%s'>  &nbsp;  &nbsp;  &nbsp; &nbsp;  &nbsp;  &nbsp; edit  &nbsp;  &nbsp;  &nbsp; &nbsp;  &nbsp;  &nbsp;</a></td><td>", teamcount, team.Name, team.Tagline, team.LeaderName(), editURL),
			&If{
				Condition: person.SuperAdmin.Get(),
				Content: &Form{
					SubmitButtonText:    "Delete",
					SubmitButtonConfirm: "Are you sure you want to delete team " + team.Name.Get() + "?",
					FormID:              "delete" + team.ID.Hex(),
					OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
						return "", StringURL("."), team.Delete()
					},
				},
			},
			HTML("</td></tr>"),
		)
	}
	if teams.Err() != nil {
		return nil, 0, teams.Err()
	}
	views = append(views,
		HTML("</table>"),
	)

	return views, teamcount, nil
}
