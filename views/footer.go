package views

import (
	"strings"

	. "github.com/ungerik/go-start/view"

	"github.com/STARTeurope/startuplive.in/models"
)

func Footer() View {
	return Views{
		&Div{
			Class: "footer",
			Content: Views{
				&Div{Class: "footer-border"},
				&Div{
					Class: "past-events",
					Content: &Div{
						Class: "center",
						Content: Views{
							HTML("<span>PAST EVENTS:</span>"),
							DynamicView(
								func(ctx *Context) (view View, err error) {
									const maxChars = 90
									var numChars int
									views := Views{}
									i := models.PastPublicStartupLiveEventIterator()
									var event models.Event
									for i.Next(&event) && numChars < maxChars {
										var region models.EventRegion
										err := event.Region.Get(&region)
										if err != nil {
											return nil, err
										}
										url := Region0_Event1.URL(ctx.ForURLArgs(region.Slug.Get(), event.Number.String()))
										name := region.Name.Get()
										numChars += 2 + len(name)
										views = append(views, A(url, HTML("&nbsp;&nbsp;"+strings.ToUpper(name))))
									}
									if i.Err() != nil {
										return nil, i.Err()
									}
									return views, nil
								},
							),
						},
					},
				},
				&Div{
					Class: "footer-main",
					Content: &Div{
						Class: "center",
						Content: Views{
							// &Template{Filename: "mailchimp.html"},
							&List{
								Class: "sitemap",
								Model: &MultiViewsListModel{
									{
										HTML("Startup Live"),
										&Menu{
											Class:           "sitemap-list",
											ItemClass:       "sitemap-item",
											ActiveItemClass: "active",
											Items: []LinkModel{
												NewLinkModel(&Events, "What is Startup Live?"),
												NewLinkModel(&Events_Where, "Where are the next events?"),
												//NewLinkModel(&Investors, "Investors"),
												//NewLinkModel(&Mentors, "Startup Mentors"),
												NewLinkModel(&Organisers, "Event Organisers"),
												NewLinkModel(&Blog, "Startup Live Blog"),
											},
										},
									},
									/*
										{
											HTML("Misc"),
											&Menu{
												Class:           "sitemap-list",
												ItemClass:       "sitemap-item",
												ActiveItemClass: "active",
												Items: []LinkModel{
													NewPageLink(&PartnersBeAPart, "Be a part"),
													NewPageLink(&Partners, "Partners"),
													NewPageLink(&Press, "Press-kit"),
												},
											},
										},
									*/
									{
										HTML("Social Media"),
										&Menu{
											Class:           "sitemap-list",
											ItemClass:       "sitemap-item",
											ActiveItemClass: "active",
											Items: []LinkModel{
												&StringLink{Title: "follow us on twitter", Url: STARTeuropeTwitterURL, Rel: "me"},
												&StringLink{Title: "find us on facebook", Url: STARTeuropeFacebookURL, Rel: "me"},
												// &StringLink{Title: "find us on linkedin", Url: STARTeuropeLinkedInURL, Rel: "me"},
												&StringLink{Title: "find us on flickr", Url: STARTeuropeFlickrURL, Rel: "me"},
											},
										},
									},
								},
							},
							DIV("logos",
								A_blank(
									"http://pioneersfestival.com/",
									&Image{
										Title:  "Pioneers Festival",
										Src:    "/images/footer-logo-pioneers.png",
										Width:  139,
										Height: 35,
									},
								),
							),
							DivClearBoth(),
							// DIV("sponsor-logos",
							// 	A_blank(
							// 		"http://www.impulse-awsg.at/",
							// 		&Image{
							// 			Description: "impulse / aws",
							// 			URL:         "/images/sponsors/impulse_aws_bw160x37.png",
							// 			Width:       160,
							// 			Height:      37,
							// 		},
							// 	),
							// 	A_blank(
							// 		"http://www.evolve.or.at/",
							// 		&Image{
							// 			Description: "evolve / bmwfi",
							// 			URL:         "/images/sponsors/evolve_bmwfj_bw160x43.png",
							// 			Width:       160,
							// 			Height:      43,
							// 		},
							// 	),
							// ),
							DivClearBoth(),
						},
					},
				},
				&Div{
					Class: "imprint",
					Content: &Div{
						Class: "center",
						Content: Views{
							&Menu{
								Class:        "menu",
								BetweenItems: "&nbsp; | &nbsp;",
								Items: []LinkModel{
									NewLinkModel(&Contact, "Contact"),
									NewLinkModel(&Imprint, "Imprint"),
								},
							},
							A_blank("http://starteurope.at", "Copyright © 2009-2014 Startup Live"),
							A_blank("http://go-start.org", "Proudly built with #golang"),
						},
					},
				},
			},
		},
	}
}
