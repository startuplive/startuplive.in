package logo

import (
	"fmt"
	"strconv"
	//	"github.com/ungerik/go-start/model"
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	. "github.com/ungerik/go-start/view"
)

func init() {
	Admin_Region0_Logo_Gen_SVG = &Page{}
	Admin_Region0_Logo_Gen_SVG.Template.Filename = "logo_ohneschrift.svg"
	Admin_Region0_Logo_Gen_SVG.GetContext = func(ctx *Context) (context interface{}, err error) {
		region, err := EventRegion(ctx.URLArgs)
		if err != nil {
			return nil, err
		}

		regionName := region.Name.Get()
		primaryColorIndex, _ := strconv.Atoi(ctx.URLArgs[1])
		secondaryColorIndex, _ := strconv.Atoi(ctx.URLArgs[2])
		fillPatternIndex, _ := strconv.Atoi(ctx.URLArgs[3])

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
