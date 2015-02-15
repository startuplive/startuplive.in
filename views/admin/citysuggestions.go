package admin

import (
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	// "github.com/ungerik/go-start/model"
	// "github.com/ungerik/go-start/utils"
	// "github.com/ungerik/go-start/debug"
	. "github.com/ungerik/go-start/view"
	"strings"
)

func init() {
	Admin_CitySuggestions = &Page{
		Title: Escape("City Suggestions | Admin"),
		CSS:   IndirectURL(&Admin_CSS),
		Scripts: Renderers{
			PageScripts,
			SCRIPT(`
				$(document).ready(function() {
					$($('.toggle').parent().next()).hide();
					$('.toggle').click(function(){
						$($(this).parent().next()).toggle(500);
					});
				});
	   		`),
		},
		Content: Views{
			adminHeader(),
			DynamicView(cityList),
			HR(),
			&Form{
				SuccessMessageClass: "success",
				SuccessMessage:      "Successfull.",
				SubmitButtonText:    "Add City",
				SubmitButtonClass:   "button",
				FormID:              "addcity",
				GetModel: func(form *Form, ctx *Context) (interface{}, error) {
					var suggestion models.CitySuggestion
					models.CitySuggestions.InitDocument(&suggestion)
					return &suggestion, nil
				},
				OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
					city := formModel.(*models.CitySuggestion)
					city.Date.SetNowUTC()
					city.Type.Set("Startup Live Event")
					return "", StringURL("."), city.Save()
				},
			},
		},
	}
}

func cityList(ctx *Context) (view View, err error) {

	suggestions := CitySuggestionIterator()
	var views Views
	var cities = map[string][]*models.CitySuggestion{}

	var c models.CitySuggestion
	for suggestions.Next(&c) {
		city := c
		cityname := strings.ToLower(city.Name.String())
		cities[cityname] = append(cities[cityname], &city)
	}
	if suggestions.Err() != nil {
		return nil, suggestions.Err()
	}

	for k, v := range cities {
		var subviews Views
		mysuggestions := v

		for i := 0; i < len(mysuggestions); i++ {
			subviews = append(subviews, HTML(cities[k][i].Date.String()+" - "+cities[k][i].Email.String()), BR())
		}

		views = append(views, Printf("<h3><span class='toggle' style='cursor:pointer; background:gray; padding:5px; color:white'>v</span> %s (%v)</h3><div>", k, len(cities[k])),
			subviews,
			Printf("</div>"),
		)
	}

	return views, nil
}
