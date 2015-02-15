package event_1

import (
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/ungerik/go-start/debug"
	. "github.com/ungerik/go-start/view"
)

func init() {
	debug.Nop()
	Region0_Event1_Voting_Dialog = &Page{
		OnPreRender:  SetEventPageData,
		Title:        EventTitle("Voting"),
		MetaViewport: "width=device-width, initial-scale=1",
		PostCSS:      StylesheetLink("/jquerymobile/jquery.mobile.css"),
		Scripts: Renderers{
			JQuery,
			ScriptLink("http://code.jquery.com/mobile/1.0.1/jquery.mobile-1.0.1.min.js"),
			HTML("<script>$(document).bind('mobileinit', function(){$.mobile.ajaxEnabled = false;});</script>"),
		},
		Content: Views{
			&Tag{
				Tag: "div",
				Attribs: map[string]string{
					"data-role": "page",
				},
				Class: "type-interior ui-page ui-body-c ui-page-active",
				Content: Views{
					&Tag{
						Tag:   "div",
						Class: "ui-header ui-bar-f",
						Attribs: map[string]string{
							"data-role":  "header",
							"data-theme": "c",
							"role":       "banner",
						},
						Content: H1("Success"),
					},
					&Tag{
						Tag: "div",
						Attribs: map[string]string{
							"data-role": "content",
							"role":      "main",
						},
						Content: DynamicView(votedSuccessView),
					},
				},
			},
		},
	}
}

func votedSuccessView(ctx *Context) (view View, err error) {

	//voteURL := Region0_Event1_Admin_Judgements_Team2.URL(response, ctx.URLArgs[0], ctx.URLArgs[1], team.ID.Hex())

	view = Views{
		&Tag{
			Tag: "div",
			Content: Views{
				//Printf("Yout voted for: "),
				Printf("Thanks for Voting"),
				A(".", "back"),
			},
		},
	}
	return view, nil
}
