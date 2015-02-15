package dashboard

import (
	"fmt"

	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/ungerik/go-start/debug"
	. "github.com/ungerik/go-start/view"
	"labix.org/v2/mgo/bson"
)

func init() {
	debug.Nop()

	Region0_Event1_Dashboard_Mentors = &Page{
		OnPreRender: SetEventPageData,
		Title:       EventDashboardTitle("Mentors"),
		CSS:         IndirectURL(&Region0_DashboardCSS),
		Content: Views{
			DashboardHeader(),
			&Div{
				Class: "content",
				Content: Views{
					&Div{
						Class: "main mentors",
						Content: Views{
							&Div{
								Class: "main-header",
								Content: Views{
									DynamicView(
										func(ctx *Context) (view View, err error) {
											region := ctx.Data.(*PageData).Region.Name.Get()
											mentors := ctx.Data.(*PageData).Event.MentorsJudgesTab_RenameMentors.GetOrDefault("Mentors")
											return H1(mentors + " in " + region), nil
										},
									),
									//							&Form{
									//								Class: "search-mentor",
									//								Content: &TextField{},
									//							},
								},
							},
							&ModelIteratorView{
								GetModel: func(ctx *Context) (interface{}, error) {
									return new(models.Person), nil
								},
								GetModelIterator: EventMentorIterator,
								GetModelView:     eventMentorView,
							},
							DivClearBoth(),
						},
					},
					eventDashboardSidebar(),
				},
			},
			eventDashboardFooter(),
		},
	}
}

func bookMentor(form *Form, ctx *Context) (err error) {
	var eventTeam models.EventTeam
	err = models.EventTeams.OneDocument(&eventTeam)
	if err != nil {
		return err
	}

	mentorID := bson.ObjectIdHex(form.FormID)
	eventTeam.BookedMentors = append(eventTeam.BookedMentors, models.People.RefForID(mentorID))
	return eventTeam.Save()
}

func eventMentorView(ctx *Context, model interface{}) (view View, err error) {
	person := model.(*models.Person)
	mentorURL := Region0_Event1_Dashboard_Mentor2.URL(ctx.ForURLArgs(ctx.URLArgs[0], ctx.URLArgs[1], person.ID.Hex()))
	mentorURL = "#" // disable link
	company := ""
	if person.Company.Get() != "" {
		company = fmt.Sprintf("<span class='company'>&nbsp;(%s)</span>", person.Company.Get())
	}

	//email := person.Contact.Email.Get()
	view = &Div{
		Class: "preview",
		Content: Views{
			A(mentorURL, ViewOrError(person.Image_160x160("framed-image"))),
			&Div{
				Class: "info",
				Content: Views{
					&Div{
						Class: "info-header",
						Content: Views{
							Printf("<div class='name'>%s %s</div>", person.Name.String(), company),
							&If{
								Condition: person.Tags != "",
								Content:   Printf("<div class='tags'>%s</div>", person.Tags),
							},
						},
					},
					// &Paragraph{
					// 	Content: &TextPreview{
					// 		PlainText: person.CV.String(),
					// 		MaxLength: 370,
					// 		MoreLink:  NewLinkModel(mentorURL, ReadMoreArrow),
					// 	},
					// },
					P(HTML(person.MentorInfo.String())),
					//Text("<table class='buttons'><tr><td>"),

					//Printf("<a class='submodal-340-500 button' href='%s'>Show Details</a>", mentorURL),

					//					Text("</td><td>"),
					//					&Form{
					//						FormID:         person.ID(),
					//						SubmitButtonText:     "Book as Mentor",
					//						SubmitButtonClass:    "button",
					//						SuccessMessage: "Mentor booked!",
					//						Save:           BookMentor,
					//						SaveRedirect:   StringURL("."),
					//					},
					//					Text("</td><td>"),
					//					Text("<ul>"),
					//					Text("<li><div class='button'>Get in Touch <div class='dropdown-arrow'></div></div>"),
					//					Text("<ul class='dropdown'>"),
					//					Text("<li><a target='_blank' href='http://twitter.com/moritzplassnig'><img src='/images/icons/twitter-mini.png'/> @moritzplassnig</a></li>"),
					//					Text("<li><a target='_blank' href='http://facebook.com/moritzplassnig'><img src='/images/icons/facebook-mini.png'/>/moritzplassnig</a></li>"),
					//					Text("<li><a target='_blank' href='http://twitter.com/moritzplassnig'><img src='/images/icons/xing-mini.png'/>Moritz Plassnig</a></li>"),
					//					Text("<li><a target='_blank' href='http://twitter.com/moritzplassnig'><img src='/images/icons/linkedin-mini.png'/>Moritz Plassnig</a></li>"),
					//					Text("<li><a target='_blank' href='mailto:moritz.plassnig@starteurope.at'><img src='/images/icons/mail-mini.png'/>moritz.plassnig@start...</a></li>"),
					//					Text("</ul></li></ul>"),
					//Text("</td></tr></table>"),
				},
			},
		},
	}
	return view, nil
}
