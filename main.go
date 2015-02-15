package main

import (
	"bytes"

	"github.com/ungerik/go-mail"
	"github.com/ungerik/go-start/config"
	"github.com/ungerik/go-start/debug"
	"github.com/ungerik/go-start/errs"
	"github.com/ungerik/go-start/media"
	"github.com/ungerik/go-start/mongo"
	"github.com/ungerik/go-start/mongomedia"
	"github.com/ungerik/go-start/user"
	"github.com/ungerik/go-start/view"

	"github.com/STARTeurope/startuplive.in/models"
	"github.com/STARTeurope/startuplive.in/views"

	// Dummy-import view packages for initialization:
	_ "github.com/STARTeurope/startuplive.in/views/admin"
	_ "github.com/STARTeurope/startuplive.in/views/admin/region_0"
	_ "github.com/STARTeurope/startuplive.in/views/admin/region_0/logo"
	_ "github.com/STARTeurope/startuplive.in/views/api"
	_ "github.com/STARTeurope/startuplive.in/views/events"
	_ "github.com/STARTeurope/startuplive.in/views/region_0"
	_ "github.com/STARTeurope/startuplive.in/views/region_0/event_1"
	_ "github.com/STARTeurope/startuplive.in/views/region_0/event_1/admin"
	_ "github.com/STARTeurope/startuplive.in/views/region_0/event_1/dashboard"
	_ "github.com/STARTeurope/startuplive.in/views/root"
)

///////////////////////////////////////////////////////////////////////////
// Extend email.Configuration with a Name() method
// to make it compatible with config.Load()

type EmailConfig struct {
	*email.Configuration
}

func (self *EmailConfig) Name() string {
	return "email"
}

func main() {
	debug.Nop()

	///////////////////////////////////////////////////////////////////////////
	// Load configuration

	defer config.Close() // Close all packages on exit

	config.Load("config.json",
		&EmailConfig{&email.Config},
		&mongo.Config,
		&user.Config,
		&view.Config,
		&media.Config,
		&mongomedia.Config,
	)

	///////////////////////////////////////////////////////////////////////////
	// Ensure that an admin user exists

	var admin models.Person
	_, err := user.EnsureExists("admin", "ungerik@gmail.com", "admin", true, &admin)
	errs.PanicOnError(err)
	admin.Admin = true
	admin.SuperAdmin = true
	err = admin.Save()
	errs.PanicOnError(err)

	///////////////////////////////////////////////////////////////////////////
	// Config view

	view.Config.LoginSignupPage = &views.LoginSignup

	view.Config.Page.DefaultAdditionalHead = view.RSSLink("Startup Live Blog Feed", view.StringURL(views.StartupLiveBlogFeedURL))
	view.Config.Page.DefaultHeadScripts = views.RenderMixpanel
	view.Config.Page.DefaultScripts = view.Renderers{
		view.ProductionServerRenderer(view.GoogleAnalytics(views.GoogleAnalyticsID)),
		views.RenderOlark,
		views.RenderCrazyEgg,
	}

	var toolbar bytes.Buffer
	view.RenderTemplate("wysihtml5-toolbar.html", &toolbar, nil)
	view.Config.RichText.DefaultToolbar = toolbar.String()
	view.Config.RichText.SetStylesheet("")
	view.Config.RichText.EditorCSS = "/style.css"

	//view.Config.GlobalAuth = view.NewBasicAuth("statuplive.in", "gostart", "gostart")

	view.Config.NamedAuthenticators["admin"] = views.Admin_Auth
	view.Config.NamedAuthenticators["super-admin"] = views.SuperAdmin_Auth
	view.Config.NamedAuthenticators["region-admin"] = views.Admin_Region0_Auth
	view.Config.NamedAuthenticators["event-admin"] = views.Region0_Event1_Admin_Auth
	view.Config.NamedAuthenticators["event-dashboard"] = views.Region0_Event1_Dashboard_Auth

	// view.Config.Debug.Mode = true
	// view.Config.Debug.LogPaths = true
	// view.Config.Debug.LogRedirects = true
	// view.Config.DisableCachedViews = true

	///////////////////////////////////////////////////////////////////////////
	// Run server

	view.RunServer(views.Paths())
}
