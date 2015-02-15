package admin

import (
	"errors"
	"fmt"
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/ungerik/go-start/model"
	. "github.com/ungerik/go-start/view"
	"strings"
)

func init() {
	Admin_Regions = &Page{
		Title:   Escape("Regions | Admin"),
		CSS:     IndirectURL(&Admin_CSS),
		Scripts: PageScripts,
		Content: Views{
			adminHeader(),
			&ModelIteratorView{
				GetModelIterator: func(ctx *Context) model.Iterator {
					return models.EventRegions.Sort("Name").Iterator()
				},
				GetModel: func(ctx *Context) (interface{}, error) {
					return new(models.EventRegion), nil
				},
				GetModelView: func(ctx *Context, model interface{}) (view View, err error) {
					region := model.(*models.EventRegion)
					return Views{
						DIV("", Printf("<h3><a href='../%s/'>%s</a></h3>", region.Slug, region.Name)),
						// &Form{
						// 	SubmitButtonText:    "Delete",
						// 	SubmitButtonClass:   "delete",
						// 	FormID:              "deleteRegion" + region.Slug.Get(),
						// 	SubmitButtonConfirm: "Are you sure you want to delete this comment: " + region.Name.Get() + "?",
						// 	OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {

						// 		return "", StringURL("."), region.Delete()

						// 	},
						// },
					}, nil
				},
			},
			&Form{
				FormID: "addregion",
				GetModel: func(form *Form, ctx *Context) (interface{}, error) {
					return &addRegionnFormModel{}, nil
				},
				SubmitButtonText:  "Add Region",
				SubmitButtonClass: "button",
				OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
					m := formModel.(*addRegionnFormModel)
					if m.Name == "" {
						return "", nil, errors.New("Empty Name")
					}

					slug := strings.ToLower(m.Name.Get())
					count, err := models.EventRegions.Filter("Slug", slug).Count()
					if err != nil {
						return "", nil, err
					}
					if count > 0 {
						return "", nil, fmt.Errorf("Event region with slug '%s' already exists", slug)
					}
					var region models.EventRegion
					models.EventRegions.InitDocument(&region)
					region.Name = m.Name
					region.Slug.Set(slug)
					return "", StringURL("."), region.Save()
				},
			},
		},
	}
}

type addRegionnFormModel struct {
	Name model.String `model:"minlen=4"`
}
