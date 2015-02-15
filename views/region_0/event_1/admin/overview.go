package admin

import (
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/STARTeurope/startuplive.in/views/admin"
	// "github.com/STARTeurope/startuplive.in/views/admin/region_0/logo"
	"github.com/ungerik/go-mail"
	"github.com/ungerik/go-start/debug"
	"github.com/ungerik/go-start/model"
	"github.com/ungerik/go-start/mongo"
	// "github.com/AlexTi/go-amiando"
	. "github.com/ungerik/go-start/view"
)

func init() {
	debug.Nop()
	Region0_Event1_Admin = &Page{
		OnPreRender: SetEventPageData,
		Title: Render(
			func(ctx *Context) (err error) {
				// Can be called from Page.LinkTitle with a dummy Reponse wiht Data == nil
				// bad idea, but not better solution yet
				if ctx.Data == nil {
					err = SetEventPageData(Region0_Event1_Admin, ctx)
					if err != nil {
						return err
					}
				}
				event := ctx.Data.(*PageData).Event
				ctx.Response.Printf("%s | Admin", event.Name)
				return nil
			},
		),
		CSS: IndirectURL(&Region0_DashboardCSS),
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
			DynamicView(
				func(ctx *Context) (view View, err error) {
					// is
					// Admin := ctx.Session.User.(*models.Person).Admin.Get()
					event := ctx.Data.(*PageData).Event

					// event.EventPartners = []models.EventPartner{}
					// event.Save()

					numAppliedParticipants := 0
					numTechParticipants := 0
					numDesignParticipants := 0
					numBizParticipants := 0
					numOtherParticipants := 0
					numCancelledParticipants := 0
					numPresentParticipants := 0
					partIter := event.ParticipantIterator()
					var participant models.EventParticipant
					for partIter.Next(&participant) {
						numAppliedParticipants++
						if participant.Cancelled() {
							numCancelledParticipants++
						}
						if participant.PresentsIdea.Get() {
							numPresentParticipants++
						}
						background := participant.Background.String()
						if background == "Tech" || background == "Technical" {
							numTechParticipants++
						} else if background == "Design" {
							numDesignParticipants++
						} else if background == "Business" {
							numBizParticipants++
						} else {
							numOtherParticipants++
						}
					}
					if partIter.Err() != nil {
						return nil, partIter.Err()
					}

					numActiveTeams := 0
					numCancelledTeams := 0
					teamIter := models.EventTeams.Filter("Event", event.ID).Iterator()
					var team models.EventTeam
					for teamIter.Next(&team) {
						if team.Cancelled() {
							numCancelledTeams++
						} else {
							numActiveTeams++
						}
					}
					if teamIter.Err() != nil {
						return nil, teamIter.Err()
					}

					// dashboardURL := Region0_Event1_Dashboard.URL(ctx, ctx.Request.URLArgs[0], ctx.Request.URLArgs[1])
					exportEmailsURL := Region0_Event1_Admin_ExportEmails.URL(ctx)
					exportPresentEmailsURL := Region0_Event1_Admin_ExportPresentEmails.URL(ctx)
					exportJudgeEmailsURL := Region0_Event1_Admin_Judges_Export.URL(ctx)
					exportMentorEmailsURL := Region0_Event1_Admin_Mentors_Export.URL(ctx)

					var startedFormModel struct {
						Started model.Bool
					}
					startedFormModel.Started = event.Started

					//	startedForm := &Form{
					//		SubmitButtonText:  "Save Event Started",
					//		SubmitButtonClass: "button",
					//		FormID:      "eventstarted",
					//		GetModel:    FormModel(&startedFormModel),
					//		OnSubmit: func(form *Form, formModel interface{}, ctx *Context) os.Error {
					//			event.Started = startedFormModel.Started
					//			return event.Save()
					//		},
					//		Redirect: StringURL("."),
					//	}

					if event.RegistrationButton.Get() == "" {
						event.RegistrationButton.Set("Register Now")
					}

					pagesettings := &Form{
						FormID:                   "eventpagesettings",
						SubmitButtonText:         "Save Visibility Settings",
						SubmitButtonClass:        "button",
						GeneralErrorOnFieldError: false,
						ExcludedFields: []string{
							"ExtraTab",
						},
						GetModel: FormModel(&event.Show),
						OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
							m := formModel.(*models.ShowPages)
							event.Show = *m
							err = event.Save()
							if err != nil {
								return "", nil, err
							}
							return "", StringURL("."), nil
						},
					}

					view = Views{
						// &If {
						// 	Condition: isAdmin,
						// 	Content: DynamicView(amiandoSyncView),
						// },
						// Printf("<h3><a href='%s'>Dashboard</a></h3>", dashboardURL),
						DynamicView(eventCrationProcess),
						&If{
							Condition: len(event.Mentors) == 0 && (event.IsPublishedOrApproved() || event.Status == "Approved"),
							Content: Views{
								HTML("You can add <a href='./mentors' target='_blank'>mentors</a> by yourself or send them a <a href='../mentor-registration' target='_blank'>link to a form</a> to fill out"),
								HR(),
							},
						},
						&If{
							Condition: len(event.Judges) == 0 && (event.IsPublishedOrApproved() || event.Status == "Approved"),
							Content: Views{
								HTML("You can add <a href='./judges' target='_blank'>judges</a> by yourself or send them a <a href='../judge-registration' target='_blank'>link to a form</a> to fill out"),
								HR(),
							},
						},
						&If{
							Condition: (event.IsPublishedOrApproved() || event.Status == "Approved"),
							Content: Views{
								HR(),
								Printf("<h3>Participants: %d applied, %d cancelled (<a href='%s'>emails of not cancelled</a>)</h3>", numAppliedParticipants, numCancelledParticipants, exportEmailsURL),
								Printf("Business: %d, Tech: %d, Design: %d, Other: %d", numBizParticipants, numTechParticipants, numDesignParticipants, numOtherParticipants),
								Printf("<h3>Pitchers: %d (<a href='%s'>emails</a>)</h3>", numPresentParticipants, exportPresentEmailsURL),
								Printf("<h3>Teams: %d active, %d cancelled</h3>", numActiveTeams, numCancelledTeams),
								Printf("<h3>Mentors: %d (<a href='%s'>emails</a>)</h3>", len(event.Mentors), exportMentorEmailsURL),
								Printf("<h3>Judges: %d (<a href='%s'>emails</a>)</h3>", len(event.Judges), exportJudgeEmailsURL),
								HR(),
							},
						},
						HR(),
						DynamicView(removeMissingMentorRefs),
						DynamicView(removeAllPartners),
						H2("Set which pages you want to display"),
						pagesettings,
						//Printf("<h3>Amiando Identifier: <a href='https://www.amiando.com/mycenter/event/%s/overview' target='_blank'>%s</a></h3>", event.AmiandoEventIdentifier, event.AmiandoEventIdentifier),
						//startedForm,

					}

					return view, nil
				},
			),
		},
	}
}

func eventCrationProcess(ctx *Context) (view View, err error) {
	event := ctx.Data.(*PageData).Event
	region := ctx.Data.(*PageData).Region

	var views Views
	nextTodo := false
	if !event.IsPublishedOrApproved() {
		views = Views{
			H2("Data you need to insert to go 'live' with your event."),
		}

		// //=== Step 0: Choose your logo ===
		// view = &ShortTag{Tag: "h4", Content: HTML("2. Add Lead Organiser")}
		// views = append(views, view)
		// if event.HasLeadOrganiser() {
		// 	view.(*ShortTag).Class = "todo-done"
		// } else if !nextTodo {
		// 	view.(*ShortTag).Class = "todo-next"
		// 	nextTodo = true
		// }

		//=== Step 0: Choose logo ===
		if region.LogoSVG != "" {
			view = &ShortTag{Tag: "h4", Content: HTML("1. Choose the logo of your city.")}
			views = append(views, view)
			view.(*ShortTag).Class = "todo-done"
		} else if !nextTodo {
			debug.Print(ctx.URLArgs[0])
			debug.Print(ctx.URLArgs[1])
			editorURL := Region0_Event1_Admin_ChooseLogo.URL(ctx.ForURLArgs(ctx.URLArgs[0], ctx.URLArgs[1], region.PrimaryColorIndex.String(), region.SecondaryColorIndex.String(), region.FillPatternIndex.String()))
			view = &ShortTag{Tag: "h4", Content: Views{
				HTML("1. Choose <a style='color:purple' href='" + editorURL + "'>the logo</a> of your city."),
			},
			}
			views = append(views, view)
			view.(*ShortTag).Class = "todo-next"
			nextTodo = true
		}

		//=== Step 1: Add more organiser ===
		view = &ShortTag{Tag: "h4", Content: HTML("2. Add your team on the <a href='./organisers' target='_blank'>organiser page</a>")}
		views = append(views, view)
		if len(event.Organisers) > 1 {
			view.(*ShortTag).Class = "todo-done"
		} else if !nextTodo {
			view.(*ShortTag).Class = "todo-next"
			nextTodo = true
		}

		//=== Step 2: Fill in about ===
		view = &ShortTag{Tag: "h4", Content: HTML("3. Add the basic event information on the <a href='./about' target='_blank'>about page</a>")}
		views = append(views, view)
		if event.AboutDone() {
			view.(*ShortTag).Class = "todo-done"
		} else if !nextTodo {
			view.(*ShortTag).Class = "todo-next"
			nextTodo = true
		}

		//=== Step 3: Fill in location ===
		view = &ShortTag{Tag: "h4", Content: HTML("4. Add your location on the <a href='./location' target='_blank'>location page</a>")}
		views = append(views, view)
		if event.LocationSet() {
			view.(*ShortTag).Class = "todo-done"
		} else if !nextTodo {
			view.(*ShortTag).Class = "todo-next"
			nextTodo = true
		}

		//=== Step 4: request going live ===
		view = &ShortTag{Tag: "h4", Content: HTML("5. Check your <a href='../' target='_blank'>event site</a> (still hidden to public) and then request to go live.")}
		views = append(views, view)
		if event.GoLiveRequested {
			view.(*ShortTag).Class = "todo-done"
		} else if !nextTodo {
			view.(*ShortTag).Class = "todo-next"
			nextTodo = true
			views = append(views, HTML("Click on the button below. We will check the event site and go live if everything is ok."))
			views = append(views, &Form{
				Class:               "signup",
				SubmitButtonText:    "I want to go 'live', please check and approve.",
				SubmitButtonClass:   "button",
				SuccessMessageClass: "success",
				SuccessMessage:      "Yay, great! A request was sent, we will check and go live as soon as possible.",
				FormID:              "golive",
				OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
					subject := event.Name.Get() + " - wants to go live"
					message := "please confirm if everything is ok or contact the host"

					go func() {
						err = email.NewBriefMessage(subject, message, SupportEmail).Send()

					}()

					event.GoLiveRequested.Set(true)

					return "", nil, event.Save()
				},
			})
		}

		//=== Step 5: Display event checking ===
		if event.GoLiveRequested {
			view = &ShortTag{Tag: "h4", Content: HTML("*** We are currently checking the event, stay tuned ***")}
			view.(*ShortTag).Class = "todo-next"
			views = append(views, view)
		}

	} else {
		if !event.SetupAmiandoEventRequest && !event.AmiandoEventActivated {
			view = &ShortTag{Tag: "h2", Content: HTML("Congratulations, you are online. Here are the next steps to accomplish")}
			views = append(views, view)

			//=== Step 1: Add amaindo data ===
			view = &ShortTag{Tag: "h4", Content: HTML("1. Please fill out the forms for the registration (ticket selling) on the <a href='./amiando-data' target='_blank'>amiando data page</a>")}
			views = append(views, view)
			if event.AmiandoDataSetup() {
				view.(*ShortTag).Class = "todo-done"
			} else if !nextTodo {
				view.(*ShortTag).Class = "todo-next"
				nextTodo = true
			}

			//=== Step 2: request creating event on amiando ===
			if event.AmiandoDataSetup() {

				view = &ShortTag{Tag: "h4", Content: HTML("2. Check your <a href='./amiando-data' target='_blank'>amiando data</a> if they are right.")}
				views = append(views, view)

				view.(*ShortTag).Class = "todo-next"
				nextTodo = true
				views = append(views, HTML("Click on the button below to request amiando event creation. We will create the event on amiando and set everything up."))
				views = append(views, &Form{
					Class:               "signup",
					SubmitButtonText:    "Yes, please create the event on amiando",
					SubmitButtonClass:   "button",
					SuccessMessageClass: "success",
					SuccessMessage:      "Yay, great! A request was sent, we will come back to you.",
					FormID:              "createamaindoevent",
					OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
						subject := event.Name.Get() + " - create amiando event"
						message := ""

						go func() {
							err = email.NewBriefMessage(subject, message, SupportEmail).Send()

						}()

						event.SetupAmiandoEventRequest.Set(true)

						return "", nil, event.Save()
					},
				})

			}

		} else {
			if !event.AmiandoEventActivated {
				view = &ShortTag{Tag: "h2", Content: HTML("*** We are currently setting up the event on amiando ***")}
				view.(*ShortTag).Class = "todo-next"
				views = append(views, view)
				views = append(views, HR())
			}
		}
	}

	return views, nil
}

func removeMissingMentorRefs(ctx *Context) (view View, err error) {
	if SessionUserIsSuperAdmin(ctx) {
		event := ctx.Data.(*PageData).Event
		var views Views

		views = append(views, &Form{
			SubmitButtonText:  "Remove Missing Refs",
			SubmitButtonClass: "button",
			FormID:            "rmmissingrefs",
			OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {

				mentoriterator := EventMentorIterator(ctx)

				var mentors []mongo.Ref

				var p models.Person
				for mentoriterator.Next(&p) {
					m := p

					for i := 0; i < len(event.Mentors); i++ {

						if event.Mentors[i].ID == m.ID {
							mentors = append(mentors, event.Mentors[i])

						}
					}
				}
				event.Mentors = mentors

				judgeiterator := EventJudgeIterator(ctx)
				var judges []mongo.Ref
				var p1 models.Person
				for judgeiterator.Next(&p1) {
					j := p1

					for i := 0; i < len(event.Judges); i++ {

						if event.Judges[i].ID == j.ID {
							judges = append(judges, event.Judges[i])

						}
					}
				}
				event.Judges = judges

				organiseriterator := EventOrganiserIterator(ctx)
				var organisers []mongo.Ref
				var p2 models.Person
				for organiseriterator.Next(&p2) {
					o := p2

					for i := 0; i < len(event.Organisers); i++ {

						if event.Organisers[i].ID == o.ID {
							organisers = append(organisers, event.Organisers[i])

						}
					}
				}
				event.Organisers = organisers
				return "", StringURL("."), event.Save()
			},
		})
		views = append(views, HR())
		return views, nil
	}
	return nil, nil
}

func removeAllPartners(ctx *Context) (view View, err error) {
	if SessionUserIsSuperAdmin(ctx) {
		event := ctx.Data.(*PageData).Event
		var views Views

		views = append(views, &Form{
			SubmitButtonText:  "Remove All Partners",
			SubmitButtonClass: "button",
			FormID:            "rmpartners",
			OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
				var partners []models.EventPartner
				event.EventPartners = partners
				return "", StringURL("."), event.Save()
			},
		})
		views = append(views, HR())
		return views, nil
	}
	return nil, nil
}
