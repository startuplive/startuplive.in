package admin

import (
	. "github.com/STARTeurope/startuplive.in/views"
	. "github.com/ungerik/go-start/view"
)

func init() {
	Admin_CSS = NewViewURLWrapper(
		NewHTML5BoilerplateCSSTemplate(DefaultCSSContext, "css/common.css", "css/dashboard.css", "css/wysihtml5.css"),
	)
}
