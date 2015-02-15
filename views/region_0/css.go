package region_0

import (
	. "github.com/ungerik/go-start/view"
	. "github.com/STARTeurope/startuplive.in/views"
)

func init() {
	Region0_CSS = NewViewURLWrapper(NewHTML5BoilerplateCSSTemplate(RegionCSSContext, "css/common.css", "css/public.css"))
}
