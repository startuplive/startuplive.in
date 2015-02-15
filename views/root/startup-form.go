package root

import (
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	// "github.com/ungerik/go-start/model"
	"github.com/ungerik/go-start/user"
	. "github.com/ungerik/go-start/view"
)

func init() {
	StartupForm = NewPublicPage("Your Startup | Startup Live", DynamicView(
		func(ctx *Context) (view View, err error) {

			excludedFields := []string{}
			excludedFields = append(excludedFields)

			requiredFields := []string{}

			disabledFields := []string{}

			view = DIV("public-content",
				DIV("main",
					DIV("",
						H1("Tell us about your entrepreneurial adventures! "),
						HTML(`<p style="margin:10px 50px">Let’s rewind the clocks back to Startup Live: People, Pitches & Passion. A full weekend launching from mere 
ideas to real business! Some time has passed since your remarkable participation in the Startup Live event 
and we would like to know more about your current status and love to receive your feedback.<br><br>
Please take a few minutes and ﬁll out the following form.<br><br>
Thank you, we greatly appreciate it.  </p>`),
						DynamicView(
							func(ctx *Context) (view View, err error) {
								var person models.Person
								found, err := user.OfSession(ctx.Session, &person)
								if err != nil {
									return nil, err
								}

								if !found {

									view = DIV("",
										DIV("", HTML("<p style='margin-left:50px'>Please login or sign up first.</p>")),
										DIV("row",
											DIV("cell right-border",
												TitleBar("Log in"),
												DIV("main-content",
													user.NewLoginForm("Log in", "login", "error", "success", IndirectURL(&StartupForm)),
												),
											),
											DIV("cell right",
												TitleBarRight("Sign up"),
												DIV("main-content",
													NewSignupForm("Sign up", "signup", "error", "success", IndirectURL(&ConfirmEmail), nil),
												),
											),
											DivClearBoth(),
											DIV("startupform-disclaimer", `You can always come back to edit your feedback on your proﬁle. Please be informed that we treat your data 
with privacy & respect`),
										),
									)
								} else {
									view = DynamicView(
										func(ctx *Context) (view View, err error) {
											ctx.Response.RequireScript(`mixpanel.track('page viewed and loggedin', {'page name': document.title, 'url': window.location.pathname, 'userstatus': 'loggedin'});`, 0)

											return Views{
												&Form{
													SubmitButtonText:  "Submit",
													SubmitButtonClass: "button",
													FormID:            "newStartup",
													Class:             "public-form",
													GetModel: func(form *Form, ctx *Context) (interface{}, error) {
														var startup models.Startup
														models.Startups.InitDocument(&startup)
														return &startup, nil
													},
													ExcludedFields: []string{"CreationDate", "Logo", "Tagline"},
													DisabledFields: disabledFields,
													RequiredFields: requiredFields,
													OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
														startup := formModel.(*models.Startup)
														startup.CreationDate.SetTodayUTC()
														startup.Founder.Set(&person)
														return "", StartupFormSuccess, startup.Save()
													},
												},
												DIV("startupform-disclaimer", `You can always come back to edit your feedback on your proﬁle. Please be informed that we treat your data with privacy & respect`),
											}, nil
										})
								}
								return view, nil
							},
						),
					),
				),
			)
			return view, nil
		},
	))
}

// type startupFormModel struct {
// 	// FirstName   model.String `view:"label=First Name|size=20" model:"required"`
// 	// LastName    model.String `view:"label=Last Name|size=20" model:"required"`
// 	// Email       model.Email  `view:"label=Email Address|size=20" model:"required"`
// 	StartupName model.String `view:"label=The name of your startup|size=20" model:"required"`
// 	BizCategory model.Choice `view:"label=Your startups business category" model:"required|options=Aerospace and Defense,Arts & Entertainment,Automotive,Biotechnology and Pharmaceuticals,Business & Professional Services,Chemicals,Clothing & Accessoires,Community & Government,Construction & Contractors,Consumer Electronics,Education,Energy,Food & Dining,Health & Medicine,Home & Garden,Industry & Agriculture,Legal & Financial,Media & Communications,Personal Care & Services,Real Estate,Shopping,Sports & Recreation,Travel & Transportation"`
// 	Abstract    model.Text   `view:"label="Short and crisp description of what you are doing" model:"required"`
// 	Website     model.Url    `view:"label="Your startups website"`
// 	TeamMember  []TeamMember `view:"label=Team Members"`
// 	Founded     model.String `view:"label=When was your startup founded?" model:"required"`
// 	Stage       model.Choice `view:"label=What stage is your startup in?" model:"required|options=Concept Stage (got an idea),Seed Stage (Working on product),Early Stage (Close to market),Growth Stage (we're out there and making some cash),Sustainable Business (we already made it to a sustainable business)"`
// 	Located     model.String `view:"label=Where is your startup located?" model:"required"`
// 	// AttendedLives model.String //Multiple Ref Choice
// 	PressArticles []PressArticles `view:"label=Any Press Mentions?"`
// 	FundingAmount model.Choice    `view:"Did you got funding - How much?" model:"options=,0,1-25k,26-75k,76-125k,126-175k,176-250k,251-325k,326-500k,500k+"`
// 	Financing     model.Choice    `model:"required|options=bootstrapping,investors,grants,investors and grants"`
// 	// Grants        model.Choice    `model:"options=yes,in talks,no"`
// 	Feedback          model.Text `view:"label=What is your feedback on the Startup Live Event"`
// 	LiveHelped        model.Text `view:"label=In what areas did Startup Live helped you and how?"`
// 	Testimonial       model.Text `view:"label=Write a short testimonial"`
// 	TermsOfConditions model.Bool `view:"label=Yes, I accept the <a href='starteurope.at/termsandconditions/' target='_blank'>Terms and Conditions</a>" model:"required"`
// }
