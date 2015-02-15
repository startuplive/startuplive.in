package admin

import (
	"github.com/ungerik/go-start/debug"
	"github.com/ungerik/go-start/model"
	. "github.com/ungerik/go-start/view"

	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/STARTeurope/startuplive.in/views/admin"
)

func init() {
	Region0_Event1_Admin_About = &Page{
		OnPreRender: SetEventPageData,
		Title:       EventAdminTitle("About"),
		CSS:         IndirectURL(&Region0_DashboardCSS),
		Scripts:     admin.PageScripts,
		Content: Views{
			eventadminHeader(),
			DynamicView(
				func(ctx *Context) (view View, err error) {
					debug.Nop()
					isAdmin := SessionUserIsSuperAdmin(ctx)
					event := ctx.Data.(*PageData).Event

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

					eventForm := &Form{
						FormID:                   "eventdata",
						SubmitButtonText:         "Save Event Data",
						SubmitButtonClass:        "button",
						GeneralErrorOnFieldError: true,
						GetModel:                 FormModel(event),
						OnSubmit:                 OnFormSubmitSaveModelAndRedirect(StringURL(".")),
						ExcludedFields:           []string{"Show"},
					}

					eventForm.ExcludedFields = []string{"AmiandoTicketCategories"}

					if !isAdmin {
						eventForm.DisabledFields = []string{"Status", "Type"}
						eventForm.ExcludedFields = append(eventForm.ExcludedFields, []string{
							"Type",
							"Number",
							"Show",
							"Start",
							"End",
							"TimeZone",
							"FAQ_HTML",
							"DashboardInfo_HTML",
							"MentorsJudgesTab_Title",
							"MentorsJudgesTab_RenameMentors",
							"ExtraTab_Title",
							"ExtraTab_HTML",
							"Started",
							"RoundupURL",
							"AmiandoAccEmail",
							"AmiandoHostId",
							"AmiandoEventIdentifier",
							"AmiandoEventApiKey",
							"AmiandoEventId",
							"AmiandoIframeCode",
							"AmiandoTicketCategories",
							"AmiandoTicketCategoriesAdded",
							"AmiandoPaymentData",
							"AmiandoPaymentAdded",
							"AmiandoBillingData",
							"AmiandoBillingAdded",
							"AmiandoPromoCodeAdded",
							"AmiandoAttendeeDataAdded",
							"AmiandoEventActivated",
							"PitcherRegistrationWelcomeText",
						}...)
					}

					view = Views{
						H3("About Site"),
						eventForm,
					}

					return view, nil
				},
			),
		},
	}
}
