package event_1

import (
	enc "encoding/hex"
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	// "github.com/ungerik/go-start/debug"
	. "github.com/ungerik/go-start/view"
	"net/http"
	// "strings"
	"time"
)

func init() {
	//debug.Nop()
	Region0_Event1_Voting = &Page{
		OnPreRender: SetEventPageData,
		Title:       EventTitle("Voting"),
		Scripts: Renderers{
			HTML("<meta name='viewport' content='width=device-width, initial-scale=1'>"),
			StylesheetLink("/jquerymobile/jquery.mobile.css"),
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
							"data-theme": "b",
							"role":       "banner",
						},
						Content: H1("Voting"),
					},
					&Tag{
						Tag: "div",
						Attribs: map[string]string{
							"data-role": "content",
							"role":      "main",
						},
						Content: DynamicView(votingListView),
					},
					DynamicView(votingFooterView),
				},
			},
		},
	}
}

func votingListView(ctx *Context) (view View, err error) {
	//debug.Print("********* in list view")

	//voteURL := Region0_Event1_Admin_Judgements_Team2.URL(response, ctx.URLArgs[0], ctx.URLArgs[1], team.ID.Hex())

	view = Views{
		&Tag{
			Tag:   "ul",
			Class: "ui-listview ui-corner-all ui-shadow",
			Attribs: map[string]string{
				"data-role":        "listview",
				"data-split-theme": "c",
			},
			Content: Views{
				DynamicView(votingListItemsView),
			},
		},
	}
	return view, nil
}

func votingListItemsView(ctx *Context) (view View, err error) {

	event := ctx.Data.(*PageData).Event

	iterator := event.PitchingTeamIterator()
	var views Views

	userIP := ctx.Request.RemoteAddr
	// userIPandPort := ctx.Request.RemoteAddr
	// userIP := strings.Split(userIPandPort, ":")[0]

	mycookie, _ := ctx.Request.Cookie("sul_v6")

	hasVoted, votedteam, err := hasVoted(userIP, mycookie, event)
	if err != nil {
		return nil, err
	}
	// hasVoted, votedteam := hasVoted("myip", mycookie, event)
	teamID := ""
	if votedteam != nil {
		teamID = votedteam.ID.String()
	}

	var t models.EventTeam
	for iterator.Next(&t) {
		team := t // copy by value because it will be used in a closure later on
		views = append(views,
			&Tag{
				Tag:   "li",
				Class: "ui-btn-inner ui-li ui-li-static ui-body-c",
				Content: Views{
					&Tag{
						Tag:   "div",
						Class: "ui-li ui-li-has-alt",
						Content: Views{
							Printf("<h5 class='ui-li-heading'>%s</h5>", team.Name.Get()),
							Printf("<small class='ui-li-desc'>%s</small>", team.Tagline.Get()),
						},
					},
					&Tag{
						Tag:   "span",
						Class: "ui-li-link-alt ui-btn-up-c-vote",
						Attribs: map[string]string{
							"data-ajax": "false",
						},
						Content: &If{
							Condition: !hasVoted,
							Content: &Form{
								SubmitButtonText:  "Vote",
								SubmitButtonClass: "button",
								FormID:            "voteform" + team.ID.Hex(),
								Class:             "ui-btn",
								Action:            "./",
								OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
									//debug.Print("----------------- VOTING -------------------")
									currtime := time.Now().UTC().String()
									enccurrtime := enc.EncodeToString([]byte(currtime))
									// expiration := time.Now().UTC().AddDate(1, 0, 0)

									// Erik: @Alex todo cookies
									// cookie := http.Cookie{Name: "sul_v6", Value: enccurrtime, Expires: expiration}
									// // debug.Print("++++++++++++++++++ ", enccurrtime)
									// response.SetCookie(&cookie)

									vote := models.NewVote()
									vote.Created.SetNowUTC()
									vote.IP.Set(userIP)
									vote.Cookie.Set("sul_v6=" + enccurrtime)
									vote.Event.ID = event.ID
									vote.Team.ID = team.ID
									//debug.Print("----------------- voted for ", team.Name.Get())
									return "", StringURL("."), vote.Save()
								},
							},
							ElseContent: &If{
								Condition: teamID == team.ID.String(),
								Content:   IMG("/images/voting-voted.png", 60, 60),
							},
						},
					},
				},
			},
		)
	}

	return views, iterator.Err()
}

func votingFooterView(ctx *Context) (view View, err error) {

	event := ctx.Data.(*PageData).Event

	userIP := ctx.Request.RemoteAddr
	// userIPandPort := ctx.Request.RemoteAddr
	// userIP := strings.Split(userIPandPort, ":")[0]

	mycookie, _ := ctx.Request.Cookie("sul_v6")

	hasVoted, team, err := hasVoted(userIP, mycookie, event)
	if err != nil {
		return nil, err
	}

	teamName := ""
	if team != nil {
		teamName = team.Name.Get()
	}

	view = Views{
		&Tag{
			Tag:   "div",
			Class: "ui-footer ui-bar-a ui-footer-fixed slideup in",
			Attribs: map[string]string{
				"data-role":     "footer",
				"data-position": "fixed",
			},
			Content: &If{
				Condition:   hasVoted,
				Content:     H5("You voted for '" + teamName + "'"),
				ElseContent: H5("You have 1 vote"),
			},
		},
	}

	return view, nil
}

func hasVoted(IP string, cookie *http.Cookie, event *models.Event) (hasVoted bool, team *models.EventTeam, err error) {
	i := models.Votes.Filter("Event", event.ID).Iterator()
	hasVoted = false

	var vote models.Vote
	for i.Next(&vote) {
		if vote.IP.Get() == IP || (cookie != nil && vote.Cookie.Get() == cookie.String()) {
			// if vote.IP.Get() == IP {
			hasVoted = true
			var votedteam models.EventTeam
			err := vote.Team.Get(&votedteam)
			if err != nil {
				return false, nil, err
			}
			return hasVoted, &votedteam, nil
		}
	}
	if i.Err() != nil {
		return false, nil, err
	}

	return hasVoted, nil, nil
}

/*func votingListView(ctx *Context) (view View, err os.Error) {
	debug.Print("********* in list view")

	//voteURL := Region0_Event1_Admin_Judgements_Team2.URL(response, ctx.URLArgs[0], ctx.URLArgs[1], team.ID.Hex())

	view = Views{
		&Tag {
			Tag: "ul",
			Class: "ui-listview",
			Attribs: map[string]string{
				"data-role": "listview",
				"data-split-theme": "d",
				"data-split-icon": "gear",
			},
			Content: Views{
				DynamicView(votingListItemsView),
			},
		},

	}
	return view, nil
}

func votingListItemsView(ctx *Context) (view View, err os.Error) {
	debug.Print("in list items view")
	event := ctx.Data.(*PageData).Event

	iterator := event.TeamIterator()
	views := Views{Printf("fe")}

	for doc := iterator.Next(); doc != nil; doc = iterator.Next() {
		team := doc.(*models.EventTeam)

		views = append(views, 
			&Tag {
				Tag: "li",
				Class: "ui-btn ui-btn-icon-right ui-li ui-li-has-alt ui-li-has-thumb ui-btn-up-c",
				Attribs: map[string]string{
					"data-corners": "false",
					"data-shadow": "false",
					"data-iconshadow": "true",
					"data-inline": "false",
					"data-wrapperels": "div",
					"data-icon": "false",
					"data-iconpos": "right",
					"data-theme": "c",
				},
				Content: Views {
					&Div {
						Class: "ui-btn-inner ui-li ui-li-has-alt",
						Content: &Div {
							Class: "ui-btn-text",
							Content: 
								&Tag{
									Tag: "div",
									Class: "ui-link-inherit",
									Content: Views {
										Printf("<img src='%s' class='ui-li-thumb'/>", team.LogoURL.Get()),
										Printf("<h3 class='ui-li-heading'>%s</h3>", team.Name.Get()),
										Printf("<p class='ui-li-desc'>%s</p>", team.Tagline.Get()),				
									},
								},
							//Printf(team.Name.Get()),
						},
					},
					&Tag {
						Tag: "a",
						Class: "ui-li-link-alt ui-btn ui-btn-up-c",
						Attribs: map[string]string {
							"data-rel": "dialog",
							"data-transition": "slideup",
							"data-iconshadow": "true",
							"data-wrapperels": "span",
							"data-icon": "false",
							"data-iconpos": "false",
							"data-theme": "c",
						},
						Content: &Tag {
							Tag: "span",
							Class: "ui-btn-inner",
							Content: 
								&Tag{
									Tag: "span",
									Class: "ui-btn ui-btn-up-d ui-btn-icon-notext ui-btn-corner-all ui-shadow",
									Attribs: map[string]string {
										"data-corners": "true",
										"data-shadow": "true",
										"data-inline": "false", 
										"data-rel": "dialog",
										"data-transition": "slideup",
										"data-iconshadow": "true",
										"data-wrapperels": "span",
										"data-icon": "gear", 
										"data-iconpos": "notext",
										"data-theme": "d",
									},
									Content: &Tag {
										Tag: "span",
										Class: "ui-btn-inner ui-btn-corner-all",
										Content: &Form {
											SubmitButtonText:  "Vote",
											SubmitButtonClass: "button",
											FormID:      "voteform",
											/*GetModel: func(form *Form, ctx *Context) (interface{}, os.Error) {
												return &addJudgeModel{}, nil
											},*/
/*OnSubmit: addVote,
											Redirect: StringURL("."),
										},
									},
								},
							//Printf(team.Name.Get()),
						},
					},

				},
			},
		)
	}	

	return views, nil
}
*/
