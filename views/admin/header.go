package admin

import (
	. "github.com/STARTeurope/startuplive.in/views"
	. "github.com/ungerik/go-start/view"
)

func adminHeader() View {
	return &Div{
		Class: "header",
		Content: Views{
			&Link{
				Class: "title",
				Model: &PageLink{
					Page: &Admin,
					Content: &Tag{
						Tag: "h1",
						Content: Views{
							&Image{Class: "logo", Src: "/images/logo-startup-live.png"},
							HTML("Admin Panel"),
						},
					},
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
							NewPageLink(&Admin, "Overview"),
							NewPageLink(&Admin_Events, "Events"),
							NewPageLink(&Admin_People, "People"),
							NewPageLink(&Admin_Teams, "Teams"),
							NewPageLink(&Admin_Startups, "Startups"),
							NewPageLink(&Admin_Regions, "Regions"),
							NewPageLink(&Admin_CitySuggestions, "City Suggestions"),
							NewPageLink(&Admin_Wiki, "Wiki"),
							NewPageLink(&Admin_Images, "Images"),
							NewPageLink(&Admin_Files, "Files"),
						},
					},
					DivClearBoth(),
				},
			},
		},
	}
}
