package region_0

import (
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/ungerik/go-start/debug"
	. "github.com/ungerik/go-start/view"
)

func Header() View {
	return &Div{
		Class: "header",
		Content: Views{
			&Link{
				Class: "title",
				Model: &PageLink{
					Page: &Admin_Regions,
					Content: H1(
						DynamicView(
							func(ctx *Context) (View, error) {
								region := ctx.Data.(*PageData).Region
								debug.Print("image-url: ", region.HeaderLogoURL.Get())
								return Views{
									&Image{Class: "logo", Src: region.HeaderLogoURL.Get()},
									Escape(region.Name.Get() + " Admin Panel"),
								}, nil
							},
						),
					),
				},
			},
			HeaderUserNav(nil),
			&Div{
				Class: "menu-frame",
				Content: Views{
					&Menu{
						Class:           "menu",
						ItemClass:       "menu-item",
						ActiveItemClass: "active",
						BetweenItems:    " &nbsp;/&nbsp; ",
						Items: []LinkModel{
							NewPageLink(&Admin_Region0, "Events"),
							NewPageLink(&Admin_Region0_Logo, "Logo"),
							NewPageLink(&Admin_Regions, "Other Regions"),
						},
					},
					DivClearBoth(),
				},
			},
		},
	}
}
