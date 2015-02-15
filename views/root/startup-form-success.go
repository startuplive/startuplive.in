package root

import (
	// "github.com/STARTeurope/startuplive.in/models"

	. "github.com/STARTeurope/startuplive.in/views"
	. "github.com/ungerik/go-start/view"
	// "io"
	// "github.com/ungerik/go-start/debug"

)

func init() {
	StartupFormSuccess = NewPublicPage("Success | Startup Live", DynamicView(
		func(ctx *Context) (view View, err error) {
			ctx.Response.RequireScript("mixpanel.track('startup form submitted')", 0)

			view = DIV("public-content",
				DIV("main",
					TitleBar("Thank you!"),
					DIV("main-content",
						P(HTML("<b>Thank you!</b><br><br>We successfully received your information.<br>You can alwas change the data of your startup on your <a href='"+Profile.URL(ctx)+"'>profile</a>. <br><br><br>Join the community on <a href='https://www.facebook.com/pages/Startup-Live/441836922496996'>Facebook</a>!.")),
					),
				),
			)
			return view, nil

		},
	))
}
