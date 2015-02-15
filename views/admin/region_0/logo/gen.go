package logo

import (
	"fmt"
	"strconv"

	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/STARTeurope/startuplive.in/views/admin"
	"github.com/ungerik/go-start/debug"
	"github.com/ungerik/go-start/model"
	. "github.com/ungerik/go-start/view"
)

func init() {
	Admin_Region0_Logo_Gen = &Page{
		Title: RegionTitle("Logo Generator"),
		Scripts: Renderers{
			admin.PageScripts,
			ScriptLink("http://canvg.googlecode.com/svn/trunk/rgbcolor.js"),
			ScriptLink("/js/canvg1.2.js"),
			ScriptLink("/js/base64.js"),
			ScriptLink("/js/canvas2image.js"),
			ScriptLink("/js/jszip.js"),
			ScriptLink("/js/logogen.js"),
			RenderViewBindURLArgs(
				func(ctx *Context, region string, primary, secondary int) {
					script := "logogenSetColors('" + models.Colors[primary] + "','" + models.Colors[secondary] + "');"
					ctx.Response.WriteString("<script>")
					ctx.Response.WriteString(script)
					ctx.Response.WriteString("</script>\n")
				},
			),
			//Script("logogenSetColors('" + models.Colors[primary] + "','" + models.Colors[secondary] + "');"),
		},
		CSS:         IndirectURL(&Region0_DashboardCSS),
		OnPreRender: SetRegionPageData,
		Content: Views{
			HTML("<span style='font-family:Startup-Initials'>.</span>"),
			EmbeddLogo(),
			//Text("<canvas class='c' height='200px'></canvas>"),
			HTML("<h3><a href='../../..'>Back</a></h3>"),
			//HTML("<div class='button ziplogo'>Download CD</div>"),
			&Form{
				Class:  "saveform",
				FormID: "save",
				GetModel: func(form *Form, ctx *Context) (interface{}, error) {
					return &addLogoFormModel{}, nil
				},
				SubmitButtonClass: "button",
				OnSubmit:          saveLogo,
			},
			&Button{
				Class:   "button downloadCD",
				Name:    "downloadCD",
				Content: Escape("Download Corporate Design"),
			},
			LogoGeneratorColorGrid(),
		},
	}
}

type addLogoFormModel struct {
	Png            model.String `view:"hidden"`
	PublicPng      model.String `view:"hidden"`
	HoverPng       model.String `view:"hidden"`
	Initial60      model.String `view:"hidden"`
	Initial        model.String `view:"hidden"`
	Circle         model.String `view:"hidden"`
	Favicon16x16   model.String `view:"hidden"`
	Favicon57x57   model.String `view:"hidden"`
	Favicon72x72   model.String `view:"hidden"`
	Favicon114x114 model.String `view:"hidden"`
	Favicon129x129 model.String `view:"hidden"`
	Svg            model.String `view:"hidden"`
}

func saveLogo(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
	model := formModel.(*addLogoFormModel)

	region := ctx.Data.(*PageData).Region

	//	debug.Print("png model len: ")
	//	debug.Print(len(model.Png))

	//debug.Print("png: " + model.HoverPng.Get())
	//  debug.Print(model.PublicPng);

	region.HeaderLogoURL.SetDataUrl(model.Png.Get())
	region.PublicHeaderLogoURL.SetDataUrl(model.PublicPng.Get())
	region.HoverLogoURL.SetDataUrl(model.HoverPng.Get())
	region.InitialURL_60x0.SetDataUrl(model.Initial60.Get())
	region.InitialURL.SetDataUrl(model.Initial.Get())
	region.CircleURL.SetDataUrl(model.Circle.Get())
	region.Favicon16x16URL.SetDataUrl(model.Favicon16x16.Get())
	region.Favicon57x57URL.SetDataUrl(model.Favicon57x57.Get())
	region.Favicon72x72URL.SetDataUrl(model.Favicon72x72.Get())
	region.Favicon114x114URL.SetDataUrl(model.Favicon114x114.Get())
	region.Favicon129x129URL.SetDataUrl(model.Favicon129x129.Get())
	region.LogoSVG.Set(model.Svg.Get())

	primaryColorIndex, err := strconv.ParseInt(ctx.URLArgs[1], 10, 64)
	if err != nil {
		return "", nil, err
	}
	region.PrimaryColorIndex.Set(primaryColorIndex)

	secondaryColorIndex, err := strconv.ParseInt(ctx.URLArgs[2], 10, 64)
	if err != nil {
		return "", nil, err
	}
	region.SecondaryColorIndex.Set(secondaryColorIndex)

	fillPatternIndex, err := strconv.ParseInt(ctx.URLArgs[3], 10, 64)
	if err != nil {
		return "", nil, err
	}
	region.FillPatternIndex.Set(fillPatternIndex)

	return "", StringURL("."), region.Save()
}

func EmbeddLogo() View {
	return RenderView(
		func(ctx *Context) error {
			context, err := Admin_Region0_Logo_Gen_SVG.GetContext(ctx)
			if err != nil {
				return err
			}
			debug.Nop()
			ctx.Response.XML.OpenTag("div").Attrib("class", "svgcontainer").Attrib("style", "display:", "block")
			RenderTemplate(Admin_Region0_Logo_Gen_SVG.Filename, ctx.Response, context)
			ctx.Response.XML.CloseTag()
			//debug.Print(" huhuhuhuhuhuhu : " + buffer.String())

			return nil
			//return Text("<h3 style='font-family:Startup-Heavy'>test</h3>"), nil
		},
	)
}

func LogoGeneratorColorGrid() View {
	return RenderView(
		func(ctx *Context) error {

			primary, _ := strconv.Atoi(ctx.URLArgs[1])
			secondary, _ := strconv.Atoi(ctx.URLArgs[2])
			fillPatternIndex, _ := strconv.Atoi(ctx.URLArgs[3])

			columnWidth := 60
			rowHeight := 40

			pattern := []string{
				"Muster-01.svg",
				"Muster-02.svg",
				"Muster-03.svg",
				"Muster-04.svg",
				"Muster-05.svg",
				"Muster-06.svg",
				"Muster-07.svg",
				"Muster-08.svg",
				"Muster-09.svg",
			}

			ctx.Response.XML.OpenTag("h3").Content("Choose Pattern").CloseTag() //h3

			ctx.Response.XML.OpenTag("table")
			ctx.Response.XML.OpenTag("tr")

			for i := 0; i < len(pattern); i++ {
				ctx.Response.XML.OpenTag("td")
				ctx.Response.XML.OpenTag("a")
				ctx.Response.XML.Attrib("href", "../../../", primary, "/", secondary, "/", i, "/")
				ctx.Response.XML.Attrib("style", "z-index:", "2;", "display:", "block;", "position:", "absolute;", "width:", "90px;", "height:", "90px")
				ctx.Response.XML.Content("&nbsp;")
				ctx.Response.XML.CloseTag() //a
				if fillPatternIndex == i {
					ctx.Response.XML.Printf("<embed class='pattern' src='/pattern/%s' type='image/svg+xml' style='width:90px; height:90px; border-radius:5px; border:solid 4px #1FA22E; z-index:1' />", pattern[i])
				} else {
					ctx.Response.XML.Printf("<embed class='pattern' src='/pattern/%s' type='image/svg+xml' style='width:90px; height:90px; border-radius:5px; z-index:1' />", pattern[i])
				}

				ctx.Response.XML.CloseTag() //td
			}

			ctx.Response.XML.CloseTag() //tr
			ctx.Response.XML.CloseTag() //table

			ctx.Response.XML.OpenTag("h3").Content("Choose Color").CloseTag()

			ctx.Response.XML.OpenTag("table").Attrib("class", "color").Attrib("border", 1)

			regions := models.EventRegions.Iterator()

			regionsMap := make(map[string][]*models.EventRegion)
			//primaries := make([]int, 0, 16)
			//secondaries := make([]int, 0, 16)

			var r models.EventRegion
			for regions.Next(&r) {
				region := r
				//debug.Print(int(region.PrimaryColorIndex))
				regionsMap[fmt.Sprintf("%d %d", int(region.PrimaryColorIndex), int(region.SecondaryColorIndex))] = append(regionsMap[fmt.Sprintf("%d %d", int(region.PrimaryColorIndex), int(region.SecondaryColorIndex))], &region)
				//primaries = append(primaries, int(region.PrimaryColorIndex))
				//secondaries = append(secondaries, int(region.SecondaryColorIndex))

			}
			if regions.Err() != nil {
				return regions.Err()
			}

			// Hardcode width of columns
			ctx.Response.XML.OpenTag("colgroup")
			for i := 0; i < 11; i++ {
				ctx.Response.XML.OpenTag("col").Attrib("width", columnWidth, "px").CloseTag()
			}
			ctx.Response.XML.CloseTag() // colgroup

			// Write first table row with primary colors
			// iterator only till 9 because the 10th element ist gray
			ctx.Response.XML.OpenTag("tr").Attrib("height", rowHeight, "px")
			ctx.Response.XML.OpenTag("td").Content("&nbsp;").CloseTag() // first cell is empty
			for primary := 0; primary < 10; primary++ {
				primaryColor := models.Colors[primary]
				ctx.Response.XML.OpenTag("td").Attrib("style", "background:", primaryColor).Content("&nbsp;").CloseTag()
			}
			ctx.Response.XML.CloseTag() // tr

			for s := 0; s < 10; s++ {
				secondary := s
				secondaryColor := models.Colors[secondary]
				ctx.Response.XML.OpenTag("tr").Attrib("height", rowHeight, "px")
				// First cell is secondary color
				ctx.Response.XML.OpenTag("td").Attrib("style", "background:", secondaryColor).Content("&nbsp;").CloseTag()
				for p := 0; p < 10; p++ {
					primary := p
					if primary == secondary || primary == 8 || secondary == 6 {
						ctx.Response.XML.OpenTag("td").CloseTagAlways()
						continue
					}
					primaryColor := models.Colors[primary]

					ctx.Response.XML.OpenTag("td").Attrib("style", "font-weight:", "bold")
					ctx.Response.XML.OpenTag("a")
					ctx.Response.XML.Attrib("href", "../../../", primary, "/", secondary, "/", fillPatternIndex, "/")
					if reg := regionsMap[fmt.Sprintf("%d %d", primary, secondary)]; reg != nil {
						cities := ""
						for r := 0; r < len(reg); r++ {
							cities += reg[r].Name.Get() + ", "
						}
						ctx.Response.XML.Attrib("title", cities)

					}
					ctx.Response.XML.OpenTag("div")
					ctx.Response.XML.Attrib("style", "color:", primaryColor) //, ";line-height:", size/2, "px")
					ctx.Response.XML.Content("Circle")
					ctx.Response.XML.CloseTag() // div
					ctx.Response.XML.OpenTag("div")
					ctx.Response.XML.Attrib("style", "color:", secondaryColor) // , ";line-height:", size/2, "px")
					ctx.Response.XML.Content("Letter")
					ctx.Response.XML.CloseTag() // div
					ctx.Response.XML.CloseTag() // a
					ctx.Response.XML.CloseTag() // td

				}
				ctx.Response.XML.CloseTag() // tr
			}

			ctx.Response.XML.CloseTag() // table

			//load javascript for rendering svg to canvas - canvg library
			/*ctx.Response.XML.OpenTag("script").Attrib("type", "text/javascript").Attrib("src", "http://canvg.googlecode.com/svn/trunk/rgbcolor.js")
			ctx.Response.XML.ExtraCloseTag()

			ctx.Response.XML.OpenTag("script").Attrib("type", "text/javascript").Attrib("src", "http://canvg.googlecode.com/svn/trunk/canvg1.2.js")
			ctx.Response.XML.Content("").CloseTag()

			ctx.Response.XML.OpenTag("script").Attrib("type", "text/javascript").Attrib("src", "/js/base64.js")
			ctx.Response.XML.Content("").CloseTag()

			ctx.Response.XML.OpenTag("script").Attrib("type", "text/javascript").Attrib("src", "/js/canvas2image.js")
			ctx.Response.XML.Content("").CloseTag()

			ctx.Response.XML.OpenTag("script").Attrib("type", "text/javascript").Attrib("src", "/js/logogen.js")
			ctx.Response.XML.Content("")
			ctx.Response.XML.CloseTag() //script

			ctx.Response.XML.OpenTag("script").Attrib("type", "text/javascript")
			ctx.Response.XML.Content("logogenSetColors('" + models.Colors[primary] + "','" + models.Colors[secondary] + "');")
			ctx.Response.XML.CloseTag() */ //script

			/*ctx.Response.XML.OpenTag("script").Attrib("type", "text/javascript")
				html.Content(" var svg = $('.svgcontainer').get(0); " +
							 " var svgcontent = $('.svgcontainer').html();  " +
							 " console.log('svg: ' + svg); " +
			    			 " var canvas = document.createElement('canvas');" +
			    			 " canvg(canvas, svgcontent);" +
			    			 " var logo = Canvas2Image.saveAsPNG(canvas, true);  " +
			    			 " console.log('hier 1' + logo); " +
			    			 " $(svg).replaceWith(logo);" +
			    			 " console.log('hier 2');")
				ctx.Response.XML.CloseTag()*/ //script

			/*ctx.Response.XML.OpenTag("script").Attrib("type", "text/javascript")
			html.Content("var btn = $('.exportpng')[0]; var c = $('.c')[0]; $(btn).click(function() { var mysvg ='svg/'; canvg(c, mysvg);var logo = c.toDataURL('image/png'); $(c).replaceWith('<img src='+logo+'/>');});")
			ctx.Response.XML.CloseTag() //script*/
			return nil
		},
	)
}
