package event_1

import (
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	. "github.com/ungerik/go-start/view"
)

func init() {
	Region0_Event1_Judges = newPublicEventPage(
		Render(
			func(ctx *Context) (err error) {
				event := ctx.Data.(*PageData).Event
				return EventTitle(event.MentorsJudgesTab_Title.GetOrDefault("Mentors/Judges")).Render(ctx)
			},
		),
		nil,
		DynamicView(
			func(ctx *Context) (view View, err error) {
				personView := func(ctx *Context, model interface{}) (View, error) {
					person := model.(*models.Person)
					return &Div{
						Class: "mentor-judge",
						Content: Views{
							ViewOrError(person.Image_50x50()),
							&Div{
								Class: "info",
								Content: Views{
									H4(person.Name.String()),
									Escape(person.Position.Get()),
									BR(),
									HTML(person.Company.Get()),
								},
							},
							DivClearBoth(),
						},
					}, nil
				}

				event := ctx.Data.(*PageData).Event
				mentorsTitle := event.MentorsJudgesTab_RenameMentors.GetOrDefault("Mentors")
				judgesTitle := "Judges / Expert Panel"
				if len(event.Mentors) == 0 && len(event.Judges) == 0 {
					mentorsTitle = ""
					judgesTitle = "Stay tuned for our mentors/judges"
				} else {
					if len(event.Mentors) == 0 {
						m := event.MentorsJudgesTab_RenameMentors.GetOrDefault("mentors")
						mentorsTitle = "Stay tuned for our " + m
					}
					if len(event.Judges) == 0 {
						judgesTitle = "Stay tuned for our judges"
					}
				}

				view = &Div{
					Class: "main",
					Content: Views{
						TitleBar(event.MentorsJudgesTab_Title.GetOrDefault("Mentors & Judges")),
						&Div{
							Class: "mentors-judges",
							Content: Views{
								H3(mentorsTitle),
								&ModelIteratorView{
									GetModelIterator: EventMentorIterator,
									GetModel: func(ctx *Context) (interface{}, error) {
										return new(models.Person), nil
									},
									GetModelView: personView,
								},
							},
						},
						&Div{
							Class: "mentors-judges",
							Content: Views{
								H3(judgesTitle),
								&ModelIteratorView{
									GetModelIterator: EventJudgeIterator,
									GetModel: func(ctx *Context) (interface{}, error) {
										return new(models.Person), nil
									},
									GetModelView: personView,
								},
							},
						},
					},
				}
				return view, nil
			},
		),
	)
}

func titleBar(title string) View {
	return &Div{Class: "title-bar", Content: Escape(title)}
}

func titleBarRight(title string) View {
	return &Div{Class: "title-bar right", Content: Escape(title)}
}
