package admin

import (
	"github.com/ungerik/go-start/media"
	. "github.com/ungerik/go-start/view"

	. "github.com/STARTeurope/startuplive.in/views"
)

func init() {
	Admin_Files = &Page{
		Title: Escape("Files | Admin"),
		CSS:   IndirectURL(&Admin_CSS),
		Scripts: Renderers{
			PageScripts,
			JQueryUI,
		},
		Content: Views{
			adminHeader(),
			media.BlobsAdmin(),
		},
	}
}
