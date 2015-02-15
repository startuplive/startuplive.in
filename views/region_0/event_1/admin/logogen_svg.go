package admin

import (
	"fmt"
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	// "github.com/ungerik/go-start/debug"
	. "github.com/ungerik/go-start/view"
	"strconv"
)

func init() {
	Region0_Event1_Admin_ChooseLogo_SVG = &Page{}
	Region0_Event1_Admin_ChooseLogo_SVG.Template.Filename = "logo_ohneschrift.svg"
	Region0_Event1_Admin_ChooseLogo_SVG.GetContext = func(ctx *Context) (context interface{}, err error) {
		region, err := EventRegion(ctx.URLArgs)
		if err != nil {
			return nil, err
		}

		regionName := region.Name.Get()

		primaryColorIndex, _ := strconv.Atoi(ctx.URLArgs[2])
		secondaryColorIndex, _ := strconv.Atoi(ctx.URLArgs[3])
		fillPatternIndex, _ := strconv.Atoi(ctx.URLArgs[4])

		var Ctx struct {
			*models.ColorScheme
			Region      string
			Initial     string
			FillPattern string
		}

		Ctx.ColorScheme = models.NewColorScheme(primaryColorIndex, secondaryColorIndex, fillPatternIndex)
		Ctx.Region = regionName
		Ctx.Initial = regionName[0:1]
		//Ctx.FillPattern = "FillPattner" + strconv.Itoa(fillPatternIndex)
		Ctx.FillPattern = fmt.Sprintf("FillPattern%d", fillPatternIndex+1)
		return Ctx, nil
	}
}
