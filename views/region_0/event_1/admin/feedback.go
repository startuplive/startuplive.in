package admin

import (
	// "fmt"
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/STARTeurope/startuplive.in/views/admin"
	// "github.com/ungerik/go-start/model"
	// "github.com/ungerik/go-start/debug"
	. "github.com/ungerik/go-start/view"
	"strconv"
)

func init() {
	Region0_Event1_Admin_Feedback = &Page{
		OnPreRender: SetEventPageData,
		Title:       EventAdminTitle("Feedback"),
		CSS:         IndirectURL(&Region0_DashboardCSS),
		PostCSS: Renderers{
			StylesheetLink("/css/ui-lightness/jquery-ui-1.8.17.custom.css"),
			StylesheetLink("/avgrund/avgrund.css"),
		},
		Scripts: Renderers{
			admin.PageScripts,
			JQueryUI,
			ScriptLink("/avgrund/avgrund.js"),
			HTML("<script>function openDialog(id) {Avgrund.show( '#'+id );}function closeDialog() {Avgrund.hide();}</script>"),
		},
		Content: Views{
			eventadminHeader(),
			// HTML("<div class='avgrund-cover'></div>"),
			H3("Feedback Forms"),
			DynamicView(
				func(ctx *Context) (View, error) {

					event := ctx.Data.(*PageData).Event
					// models.FeedbackParticipants.Filter("Event", event.ID).RemoveAll()

					var region models.EventRegion
					err := event.Region.Get(&region)
					if err != nil {
						return nil, err
					}

					return Views{
						// @Alex direkt die URL von der Page verwenden ist viel einfacher:
						// (wenn die URLArgs anders sind als vom aktuellen ctx, dann ctx.ForURLArgs verwenden)
						// A("http://" + ctx.Request.Host + "/" + region.Slug.String() + "/" + event.Number.String() + "/participant-feedback"),
						A(Region0_Event1_FeedbackParticipants.URL(ctx)),
						BR(),
						A("http://" + ctx.Request.Host + "/" + region.Slug.String() + "/" + event.Number.String() + "/mentorjudge-feedback"),
						BR(),
						A("http://" + ctx.Request.Host + "/" + region.Slug.String() + "/" + event.Number.String() + "/organiser-feedback"),
					}, nil
				},
			),
			HR(),
			H2("Participants Feedback"),
			&ModelIteratorTableView{
				Class:            "visual-table",
				GetModelIterator: EventParticipantFeedbackIterator,
				GetRowModel: func(ctx *Context) (interface{}, error) {
					return new(models.FeedbackParticipant), nil
				},
				GetHeaderRowViews: TableHeaderRowEscape("Nr", "Team", "Communication", "Organisation", "Motivation", "Ideas", "Workshops", "Best Workshop", "Mentors", "Best Mentors", "Teambuilding", "Location", "Food", "Comments", "Overall"),
				GetRowViews: func(row int, rowModel interface{}, ctx *Context) (views Views, err error) {

					feedback := rowModel.(*models.FeedbackParticipant)

					teamName := "no team set"
					var team models.EventTeam
					found, err := feedback.Team.TryGet(&team)
					if err != nil {
						return nil, err
					}
					if found {
						teamName = team.Name.String()
					}

					showComments := Views{
						HTML("<button onclick='javascript:openDialog(\"participant" + strconv.Itoa(row) + "\");'>Show it</button>"),

						HTML("<aside id='participant" + strconv.Itoa(row) + "' class='avgrund-popup'><h4>Comments</h4>"),

						P(feedback.Comments),

						HTML("<button onclick='javascript:closeDialog();'>Close</button></aside>"),
					}
					showMentors := Views{
						HTML("<button onclick='javascript:openDialog(\"participantmentors" + strconv.Itoa(row) + "\");'>Show it</button>"),

						HTML("<aside id='participantmentors" + strconv.Itoa(row) + "' class='avgrund-popup'><h4>Best Mentors</h4>"),
						P(feedback.MentorsBest),
						HTML("<button onclick='javascript:closeDialog();'>Close</button></aside>"),
					}

					views = Views{
						Printf("%d", row+1),
						Escape(teamName),
						Escape(feedback.Communication.String()),
						Escape(feedback.Organisation.String()),
						Escape(feedback.Motivation.String()),
						Escape(feedback.Ideas.String()),
						Escape(feedback.Workshops.String()),
						Escape(feedback.WorkshopsBest.String()),
						Escape(feedback.Mentors.String()),
						showMentors,
						// &ModalDialog(HTML())
						Escape(feedback.Teambuilding.String()),
						Escape(feedback.Location.String()),
						Escape(feedback.Food.String()),
						showComments,
						Escape(feedback.Overall.String()),
						// A(editURL, "Edit"),						
					}
					return views, nil
				},
			},
			HR(),
			H2("Mentors/Judges Feedback"),
			&ModelIteratorTableView{
				Class:            "visual-table",
				GetModelIterator: EventMentorJudgeFeedbackIterator,
				GetRowModel: func(ctx *Context) (interface{}, error) {
					return new(models.FeedbackMentorJudge), nil
				},
				GetHeaderRowViews: TableHeaderRowEscape("Nr", "Communication", "Organisation", "Cooperation with Organisers", "Sessions", "Quality of Teams", "Expectations met", "Networking", "Likes/Dislikes", "Comments", "Overall"),
				GetRowViews: func(row int, rowModel interface{}, ctx *Context) (views Views, err error) {

					feedback := rowModel.(*models.FeedbackMentorJudge)
					// editURL := Region0_Event1_Admin_Organiser2.URL(ctx.ForURLArgs(ctx.URLArgs[0], ctx.URLArgs[1], person.ID.Hex()))
					showComments := Views{
						HTML("<button onclick='javascript:openDialog(\"mentorscomments" + strconv.Itoa(row) + "\");'>Show it</button>"),
						HTML("<aside id='mentorscomments" + strconv.Itoa(row) + "' class='avgrund-popup'><h4>Comments</h4>"),
						HTML("<p>" + feedback.Comments.String() + "</p>"),
						HTML("<button onclick='javascript:closeDialog();'>Close</button></aside>"),
					}
					showLikes := Views{
						HTML("<button onclick='javascript:openDialog(\"mentorslikes" + strconv.Itoa(row) + "\");'>Show it</button>"),
						HTML("<aside id='mentorslikes" + strconv.Itoa(row) + "' class='avgrund-popup'><h4>Likes / Dislikes</h4>"),
						HTML("<p>" + feedback.LikesDislikes.String() + "</p>"),
						HTML("<button onclick='javascript:closeDialog();'>Close</button></aside>"),
					}

					views = Views{
						Printf("%d", row+1),
						Escape(feedback.Communication.String()),
						Escape(feedback.Organisation.String()),
						Escape(feedback.OrganiserTeam.String()),
						Escape(feedback.Sessions.String()),
						Escape(feedback.QualityOfTeams.String()),
						Escape(feedback.ExpectationsMet.String()),
						Escape(feedback.Networking.String()),
						showLikes,
						showComments,
						Escape(feedback.Overall.String()),
						// A(editURL, "Edit"),						
					}
					return views, nil
				},
			},
			HR(),
			H2("The Organiser Teams Feedback"),
			&ModelIteratorTableView{
				Class:            "visual-table",
				GetModelIterator: EventHostFeedbackIterator,
				GetRowModel: func(ctx *Context) (interface{}, error) {
					return new(models.FeedbackHost), nil
				},
				GetHeaderRowViews: TableHeaderRowEscape("Nr", "Organisation", "Feeling at Event", "Quality of Ideas", "Workshops", "Best Workshop", "Mentors", "Teambuilding", "Location", "Food", "Likes Dislikes", "Comments", "Overall"),
				GetRowViews: func(row int, rowModel interface{}, ctx *Context) (views Views, err error) {

					feedback := rowModel.(*models.FeedbackHost)
					// editURL := Region0_Event1_Admin_Organiser2.URL(ctx.ForURLArgs(ctx.URLArgs[0], ctx.URLArgs[1], person.ID.Hex()))
					showComments := Views{
						HTML("<button onclick='javascript:openDialog(\"hostscomments" + strconv.Itoa(row) + "\");'>Show it</button>"),
						HTML("<aside id='hostscomments" + strconv.Itoa(row) + "' class='avgrund-popup'><h4>Comments</h4>"),
						HTML("<p>" + feedback.Comments.String() + "</p>"),
						HTML("<button onclick='javascript:closeDialog();'>Close</button></aside>"),
					}
					showLikes := Views{
						HTML("<button onclick='javascript:openDialog(\"hostslikes" + strconv.Itoa(row) + "\");'>Show it</button>"),
						HTML("<aside id='hostslikes" + strconv.Itoa(row) + "' class='avgrund-popup'><h4>Likes / Dislikes</h4>"),
						HTML("<p>" + feedback.LikesDislikes.String() + "</p>"),
						HTML("<button onclick='javascript:closeDialog();'>Close</button></aside>"),
					}

					views = Views{
						Printf("%d", row+1),
						Escape(feedback.Organisation.String()),
						Escape(feedback.Feeling.String()),
						Escape(feedback.Ideas.String()),
						Escape(feedback.Workshops.String()),
						Escape(feedback.WorkshopsBest.String()),
						Escape(feedback.Mentors.String()),
						Escape(feedback.Teambuilding.String()),
						Escape(feedback.Location.String()),
						Escape(feedback.Food.String()),
						showLikes,
						showComments,
						Escape(feedback.Overall.String()),
						// A(editURL, "Edit"),						
					}
					return views, nil
				},
			},
		},
	}
}
