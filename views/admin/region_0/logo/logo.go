package logo

import (
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/STARTeurope/startuplive.in/views/admin"
	"github.com/STARTeurope/startuplive.in/views/admin/region_0"
	. "github.com/ungerik/go-start/view"
)

func init() {
	Admin_Region0_Logo = &Page{
		Title: RegionTitle("Logo"),
		Scripts: Renderers{
			admin.PageScripts,
			ScriptLink("http://canvg.googlecode.com/svn/trunk/rgbcolor.js"),
			ScriptLink("/js/canvg1.2.js"),
			ScriptLink("/js/jszip.js"),
			ScriptLink("/js/swfobject.js"),
			ScriptLink("/js/downloadify.min.js"),
			ScriptLink("/js/downloadCD.js"),
		},
		CSS:         IndirectURL(&Region0_DashboardCSS),
		OnPreRender: SetRegionPageData,
		Content: Views{
			region_0.Header(),
			//DynamicView(showLogo),
			DynamicView(
				func(ctx *Context) (view View, err error) {
					region := ctx.Data.(*PageData).Region
					editorURL := Admin_Region0_Logo_Gen.URL(ctx.ForURLArgs(ctx.URLArgs[0], region.PrimaryColorIndex.String(), region.SecondaryColorIndex.String(), region.FillPatternIndex.String()))
					return H3(A(editorURL, "Logo Generator")), nil
				},
			),
		},
	}
}

// @Alex siehe EmbeddLogo() fuer neues format
// func showLogo(ctx *Context) (View, error) {
// 	region := ctx.Data.(*PageData).Region
// 	debug.Nop()

// 	var logoSVG string

// 	// debug.Print("logo svg: ")
// 	// debug.Print(region.LogoSVG.IsEmpty())

// 	if !region.LogoSVG.IsEmpty() {
// 		logoSVGbytes, err := hex.DecodeString(region.LogoSVG.String())
// 		if err != nil {
// 			return nil, err
// 		}
// 		logoSVG = string(logoSVGbytes)
// 		//logoURL = "data:image/png;base64," + logoURL
// 	} else {
// 		return HTML(""), nil
// 	}

// 	writer.OpenTag("div").Attrib("class", "button downloadCD").Attrib("style", "width:", "190px; ", "text-align:", "center")
// 	writer.Content("Download Corporate Design")
// 	writer.CloseTag()

// 	writer.OpenTag("h2")
// 	writer.Content("Your Logo:")
// 	writer.CloseTag()

// 	writer.OpenTag("div").Attrib("class", "svgcontainer").Attrib("style", "display:", "block")
// 	writer.Content(logoSVG)
// 	writer.CloseTag()
// 	//debug.Print(" huhuhuhuhuhuhu : " + buffer.String())

// 	return HTML(writer.String()), nil
// 	//return Text("<h3 style='font-family:Startup-Heavy'>test</h3>"), nil
// }

//deprecated
// func AddJS(ctx *Context) (View, error) {
// 	//load javascript for rendering svg to canvas - canvg library
// 	writer.OpenTag("script").Attrib("type", "text/javascript").Attrib("src", "http://canvg.googlecode.com/svn/trunk/rgbcolor.js")
// 	writer.Content("").CloseTag()

// 	writer.OpenTag("script").Attrib("type", "text/javascript").Attrib("src", "http://canvg.googlecode.com/svn/trunk/canvg.js")
// 	writer.Content("").CloseTag()

// 	writer.OpenTag("script").Attrib("type", "text/javascript").Attrib("src", "/js/canvas2image.js")
// 	writer.Content("").CloseTag()

// 	writer.OpenTag("script").Attrib("type", "text/javascript").Attrib("src", "/js/jszip.js")
// 	writer.Content("").CloseTag()

// 	writer.OpenTag("script").Attrib("type", "text/javascript").Attrib("src", "/js/swfobject.js")
// 	writer.Content("").CloseTag()

// 	writer.OpenTag("script").Attrib("type", "text/javascript").Attrib("src", "/js/downloadify.min.js")
// 	writer.Content("").CloseTag()

// 	writer.OpenTag("script").Attrib("type", "text/javascript").Attrib("src", "/js/downloadCD.js")
// 	writer.Content("").CloseTag()

// 	return HTML(writer.String()), nil
// }
