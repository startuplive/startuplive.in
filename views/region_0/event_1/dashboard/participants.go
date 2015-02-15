package dashboard

import (
	"github.com/ungerik/go-start/debug"
	. "github.com/ungerik/go-start/view"

	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
)

/*
	<link rel="stylesheet" type="text/css" href="/submodal/submodal.css" />
	<script type="text/javascript" src="/submodal/submodal.js"></script>
*/

func init() {
	debug.Nop()

	Region0_Event1_Dashboard_Participants = &Page{
		OnPreRender: SetEventPageData,
		Title:       EventDashboardTitle("Participants"),
		CSS:         IndirectURL(&Region0_DashboardCSS),
		Content: Views{
			DashboardHeader(),
			&Div{
				Class: "content",
				Content: Views{
					&Div{
						Class: "main participants",
						Content: Views{
							&Div{
								Class: "main-header",
								Content: Views{
									DynamicView(func(ctx *Context) (view View, err error) {
										region := ctx.Data.(*PageData).Region
										return H1("Participants in "+region.Name.Get(), HTML(" (upload your photo <a href='http://gravatar.com' target='_blank'>here</a>)")), nil
									}),
									//							&Form{
									//								Class: "search-mentor",
									//								Content: &TextField{},
									//							},
								},
							},
							&ModelIteratorView{
								GetModel: func(ctx *Context) (interface{}, error) {
									return new(models.EventParticipant), nil
								},
								GetModelIterator: EventParticipantIterator,
								GetModelView:     eventParticipantPreview,
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

func eventParticipantPreview(ctx *Context, model interface{}) (view View, err error) {
	participant := model.(*models.EventParticipant)

	if participant.Cancelled() {
		return nil, nil
	}

	var person models.Person
	err = participant.Person.Get(&person)
	if err != nil {
		return nil, err
	}

	// event := ctx.Data.(*PageData).Event

	// if event.Started.Get() && !participant.PresentsIdea.Get() {
	// 	return nil, nil
	// }

	participantUrl := Region0_Event1_Dashboard_Participant2.URL(ctx.ForURLArgs(ctx.URLArgs[0], ctx.URLArgs[1], participant.ID.Hex()))
	participantUrl = "#"

	// change to 100x100 ?
	image, err := person.Image_160x160("")
	if err != nil {
		return nil, err
	}
	image.Width = 110
	image.Height = 110

	var lastNameInitial string
	if person.Name.Last != "" {
		lastNameInitial = person.Name.Last.String()[:1] + "."
	}
	shortName := person.Name.First.String() + " " + lastNameInitial

	view = &Div{
		Class: "participant-preview framed-image",
		Content: Views{
			&Link{
				//Class: "submodal-340-500",
				Model: &StringLink{
					Url:   participantUrl,
					Title: person.Name.String(),
					Content: Views{
						image,
						//Text("<div class='image-gray'></div>"),
						Printf("<div class='name'>%s</div>", shortName),
					},
				},
			},
		},
	}
	return view, nil
}
