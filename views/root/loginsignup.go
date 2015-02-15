package root

import (
	"errors"

	"github.com/ungerik/go-start/model"
	gostartuser "github.com/ungerik/go-start/user"
	. "github.com/ungerik/go-start/view"

	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
)

func init() {
	Logout = NewViewURLWrapper(gostartuser.LogoutView(nil))

	main := DIV("public-content",
		DIV("main",
			TitleBar("Email confirmation"),
			DIV("main-content", gostartuser.EmailConfirmationView(IndirectURL(&Profile))),
		),
		// DynamicView(),
	)

	ConfirmEmail = &Page{
		Title:   HTML("Email Confirmation | Startup Live"),
		Content: PublicPageStructure("menu-area", PublicPageLogo(), HeaderMenu(), main, nil),
		Scripts: Renderers{
			JQuery,
			IndirectRenderer(&Config.Page.DefaultScripts),
		},
	}

	LoginSignup = NewPublicPage("Login or Sign up | Startup Live",
		DIV("public-content",
			DynamicView(
				func(ctx *Context) (view View, err error) {
					_, hasFrom := ctx.Request.Params["from"]
					id := ctx.Session.ID()
					if id == "" && hasFrom {
						view = DIV("main",
							DIV("main-content",
								H3("Your account doesn't have sufficient rights to view this page"),
								Printf("You may <a href='%s'>logout</a> and login with a different account", Logout.URL(ctx)),
							),
						)
					} else {
						view = DIV("row",
							DIV("cell right-border",
								TitleBar("Log in"),
								DIV("main-content",
									gostartuser.NewLoginForm("Log in", "login", "error", "success", IndirectURL(&Profile)),
								),
							),
							DIV("cell right",
								TitleBarRight("Sign up"),
								DIV("main-content",
									NewSignupForm("Sign up", "signup", "error", "success", IndirectURL(&ConfirmEmail), nil),
								),
							),
							DivClearBoth(),
						)
					}
					return view, nil
				},
			),
		),
	)
}

func NewSignupForm(buttonText, class, errorMessageClass, successMessageClass string, confirmationURL, redirectURL URL) *Form {
	return &Form{
		Class:               class,
		ErrorMessageClass:   errorMessageClass,
		SuccessMessageClass: successMessageClass,
		SuccessMessage:      gostartuser.Config.ConfirmationMessage.Sent,
		SubmitButtonText:    buttonText,
		FormID:              "gostart_user_signup",
		GetModel: func(form *Form, ctx *Context) (interface{}, error) {
			return &EmailPasswordFormModel{}, nil
		},
		OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
			m := formModel.(*EmailPasswordFormModel)
			email := m.Email.Get()
			password := m.Password1.Get()

			var person models.Person
			models.People.InitDocument(&person)
			found, err := gostartuser.WithEmail(email, &person)
			if err != nil {
				return "", nil, err
			}
			if found {
				if person.EmailPasswordConfirmed() {
					return "", nil, errors.New("A user with that email and a password already exists")
				}
				person.Password.SetHashed(password)
			} else {
				models.People.InitDocument(&person)
				err = person.SetEmailPassword(email, password)
				if err != nil {
					return "", nil, err
				}
			}
			err = <-person.Email[0].SendConfirmationEmail(ctx, confirmationURL)
			if err != nil {
				return "", nil, err
			}

			person.TermsAndConditions = m.TermsAndConditions
			person.TaCDate.SetTodayUTC()
			return "", redirectURL, person.Save()
		},
	}
}

type EmailPasswordFormModel struct {
	Email                         model.Email `model:"required" view:"size=20"`
	gostartuser.PasswordFormModel `bson:",inline" view:"size=20"`
	TermsAndConditions            model.Bool `view:"label=Yes, I accept the <a href='http://starteurope.at/termsandconditions/' target='_blank'>Terms and Conditions</a>" model:"required"`
}
