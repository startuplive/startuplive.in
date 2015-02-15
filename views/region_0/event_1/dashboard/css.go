package dashboard

import (
	. "github.com/STARTeurope/startuplive.in/views"
	. "github.com/ungerik/go-start/view"
)

func init() {
	Region0_DashboardCSS = NewViewURLWrapper(
		NewHTML5BoilerplateCSSTemplate(RegionCSSContext, "css/common.css", "css/dashboard.css", "css/wysihtml5.css"),
	)
	Region0_DashboardSubmodalCSS = NewViewURLWrapper(
		NewHTML5BoilerplateCSSTemplate(RegionCSSContext, "css/common.css", "css/dashboard_submodal.css", "css/wysihtml5.css"),
	)
}
