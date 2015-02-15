package admin

import (
	"github.com/ungerik/go-start/media"
	. "github.com/ungerik/go-start/view"

	. "github.com/STARTeurope/startuplive.in/views"
)

func init() {
	Admin_Images = &Page{
		Title:   Escape("Images | Admin"),
		CSS:     IndirectURL(&Admin_CSS),
		Scripts: PageScripts,
		Content: Views{
			adminHeader(),
			media.ImagesAdmin(),
		},
	}
}
