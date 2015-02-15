package views

import (
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/ungerik/go-start/view"
)

func NewPublicPage(title string, main View) *Page {
	return &Page{
		Title: Escape(title),
		Scripts: Renderers{
			JQuery,
		},
		Content: PublicPageStructure("menu-area", PublicPageLogo(), HeaderMenu(), main, nil),
	}
}

func PublicPageLogo() View {
	return Views{
		IMG("/images/logo-green.png", 146, 162),
		&Image{Class: "logo-hover", Src: "/images/logo-hover.png"},
	}
}

func PublicPageStructure(menuAreaClass string, logo, menu, content, additionalTopNav View) View {
	return Views{
		&Div{
			Class: "header",
			Content: Views{
				HeaderTopNav(additionalTopNav),
				&Div{
					Class: menuAreaClass,
					Content: &Div{
						Class: "center",
						Content: Views{
							&Div{
								Class: "logo-container",
								Content: Views{
									logo,
									MegaDropdown(),
								},
							},
							&Image{
								Class: "select-your-city",
								Src:   "/images/select-your-city.png",
							},
							menu,
							DIV("paperstack-top"),
						},
					},
				},
			},
		},
		&Div{
			Class: "content",
			Content: Views{
				&Div{
					Class: "center",
					Content: Views{
						&Image{Class: "paperstack-left", Src: "/images/paperstack-left.png", Width: 19, Height: 235},
						&Div{
							Class:   "paperstack",
							Content: content,
						},
						&Image{Class: "paperstack-right", Src: "/images/paperstack-right.png", Width: 17, Height: 235},
						DivClearBoth(),
					},
				},
			},
		},
		Footer(),
	}
}

func EventContact(person *models.Person) View {
	if person == nil {
		return nil
	}

	var eventContactTableModel ViewsTableModel
	if id := person.PrimaryEmailIdentity(); id != nil {
		eventContactTableModel = append(eventContactTableModel, Views{HTML("Mail:"), A_blank_nofollow(id)})
	}
	if phone := person.PrimaryPhone(); phone != "" {
		eventContactTableModel = append(eventContactTableModel, Views{HTML("Phone:"), Escape(phone)})
	}
	if id := person.PrimaryTwitterIdentity(); id != nil {
		eventContactTableModel = append(eventContactTableModel, Views{HTML("Twitter:"), A_blank_nofollow(id)})
	}

	return &Div{
		Class: "event-contact",
		Content: Views{
			DIV("image-frame",
				ViewOrError(person.Image_100x100()),
			),
			H4("Event Contact"),
			Escape(person.Name.String()),
			&Table{Model: eventContactTableModel},
		},
	}
}

func EventOrganiserView(person *models.Person) View {
	contactViews := make(Views, 0, 7)
	if id := person.PrimaryEmailIdentity(); id != nil {
		contactViews = append(contactViews, A_blank_nofollow(id, "Mail"))
	}
	if id := person.PrimaryTwitterIdentity(); id != nil {
		if len(contactViews) > 0 {
			contactViews = append(contactViews, HTML("&nbsp;&nbsp;|&nbsp;&nbsp;"))
		}
		contactViews = append(contactViews, A_blank_nofollow(id, "Twitter"))
	}
	if id := person.PrimaryFacebookIdentity(); id != nil {
		if len(contactViews) > 0 {
			contactViews = append(contactViews, HTML("&nbsp;&nbsp;|&nbsp;&nbsp;"))
		}
		contactViews = append(contactViews, A_blank_nofollow(id, "Facebook"))
	}
	if id := person.PrimaryLinkedInIdentity(); id != nil {
		if len(contactViews) > 0 {
			contactViews = append(contactViews, HTML("&nbsp;&nbsp;|&nbsp;&nbsp;"))
		}
		contactViews = append(contactViews, A_blank_nofollow(id, "LinkedIn"))
	}
	return &Div{
		Class: "event-organiser",
		Content: Views{
			ViewOrError(person.Image_284x144("image-box")),
			H3(person.Name.String()),
			P(HTML(person.OrganiserInfo.Get())),
			contactViews,
		},
	}
}
