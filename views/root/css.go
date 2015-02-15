package root

import (
	. "github.com/ungerik/go-start/view"
	. "github.com/STARTeurope/startuplive.in/views"
)

func init() {
	CSS = NewHTML5BoilerplateCSSTemplate(DefaultCSSContext, "css/common.css", "css/public.css")
}
