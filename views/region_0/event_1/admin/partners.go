package admin

import (
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/STARTeurope/startuplive.in/views/admin"
	// "github.com/ungerik/go-start/utils"
	"github.com/ungerik/go-start/debug"
	"github.com/ungerik/go-start/media"
	"github.com/ungerik/go-start/model"
	"github.com/ungerik/go-start/mongo"
	. "github.com/ungerik/go-start/view"
	// "image"
	"image/color"
	"strconv"
)

func init() {
	Region0_Event1_Admin_Partners = &Page{
		OnPreRender: SetEventPageData,
		Title:       EventAdminTitle("Partners"),
		CSS:         IndirectURL(&Region0_DashboardCSS),
		Scripts: Renderers{
			admin.PageScripts,
			JQueryUI,
			Render(
				func(ctx *Context) (err error) {

					// if !Config.IsProductionServer {
					// 	return nil
					// }

					ctx.Response.Write([]byte(`<script type="text/javascript">
		  		$(function() {
		  			$('.partner-success').hide();
			        $( ".sortable" ).sortable({
			            placeholder: "ui-state-highlight",
			            update: function(event, ui){
					        // console.log(event);
					        console.log(ui);
					        // var savebtn = $($(".savebtn").get(0)).
					        var idsInOrder = $(this).sortable("toArray");
					        console.log(idsInOrder);
					    }
			        });
			        $( ".sortable" ).disableSelection();
			        $(document).ready(function() {
			        	
       					$('.saveOrder').click(function() {
            				var order = "";
            				var that = this;
            				var i = 0;
            				$($(this).prev()).children().each(function() {
            					order += $(this).attr("pos") + ",";	
							});
							order = order.slice(0,order.length-1);
							var neworder = {
								Category: $(this).prev().attr("category"),
								Order: order,
							}
							// console.log(neworder);
							// $.get('http://` + ctx.Request.Host + ctx.Request.RequestURI + `order', {"order": order});
							$.ajax({
							  type: 'POST',
							  url: 'http://` + ctx.Request.Host + ctx.Request.RequestURI + `order',
							  data: neworder,
							  success: function() {
							  		$(that).next().show().fadeOut(5000);
							  	},
							  dataType: 'json'
							});
							
            			});
        			});
			    });
				</script>`))

					return nil
				},
			),

			// JQueryUIAutocompleteFromURL(".add-mentor", IndirectURL(&API_People), 2),
		},
		Content: Views{
			eventadminHeader(),
			DIV("content",
				H2("Partner Categories:"),
				DynamicView(
					func(ctx *Context) (view View, err error) {

						event := ctx.Data.(*PageData).Event
						// event.EventPartners = []models.EventPartner{}
						// event.Save()
						var views Views
						event.OrderEventPartner()
						// partners := models.GetPartnersByEventAndCategoryIterator(event, models.PARTNERS)
						partners, err := event.GetPartnersByCategory(models.PARTNERS)
						if err != nil {
							return nil, err
						}
						views = append(views, renderPartnersByCategory(ctx, partners, models.PARTNERS))
						supporters, err := event.GetPartnersByCategory(models.SUPPORTERS)
						if err != nil {
							return nil, err
						}
						views = append(views, renderPartnersByCategory(ctx, supporters, models.SUPPORTERS))
						mediapartners, err := event.GetPartnersByCategory(models.MEDIAPARTNERS)
						if err != nil {
							return nil, err
						}
						views = append(views, renderPartnersByCategory(ctx, mediapartners, models.MEDIAPARTNERS))

						return views, nil
					},
				),
				BR(),
				HTML("<span style='color:darkred; font-size:20px'>For best results please upload images with 200x200 pixels.</span>"),
				DynamicView(
					func(ctx *Context) (view View, err error) {
						evt := ctx.Data.(*PageData).Event

						return &Form{
							SubmitButtonText:  "Add new Partner",
							SubmitButtonClass: "button",
							FormID:            "newpartner",
							GetModel: func(form *Form, ctx *Context) (interface{}, error) {
								return &PartnerForm{Event: evt.Ref()}, nil
							},
							OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
								m := formModel.(*PartnerForm)
								partner := models.NewPartner()

								partner.Name = m.Name
								partner.Website = m.Website
								partner.Logo = m.Logo
								partner.Events = append(partner.Events, m.Event)

								err := partner.Save()
								if err != nil {
									return "", StringURL("."), err
								}

								if err != nil {
									return "", StringURL("."), err
								}
								debug.Nop()

								// evt.EventPartners.[m.Category.Get()] = append(evt.EventPartners[m.Category.Get()], eventpartner)

								return "", StringURL("."), evt.AddEventPartnerToCategory(m.Category.Get(), partner)
							},
						}, nil
					}),
			),
		},
	}
}

type PartnerForm struct {
	Name     model.String   `model:"required"`
	Website  model.String   `model:"required"`
	Logo     media.ImageRef `model:"required"`
	Event    mongo.Ref      `model:"required|to=events"`
	Category model.Choice   `model:"options=Partners,Supporters,Media Partners"`
}

func renderPartnersByCategory(ctx *Context, partners []models.Partner, cat string) Views {
	var views Views
	event := ctx.Data.(*PageData).Event
	// var doc models.Partner
	views = append(views, TitleBar(cat))

	var listitem Views
	for j := 0; j < len(partners); j++ {
		i := j
		partner := partners[j]
		listitem = append(listitem,
			&Tag{
				Tag:   "li",
				Class: "",
				Attribs: map[string]string{
					"pos": strconv.Itoa(i),
				},
				Content: Views{
					DynamicView(
						func(ctx *Context) (view View, err error) {
							return partner.Logo.VersionTouchOrigFromOutsideView(100, 100, media.HorCenter, media.VerCenter, false, color.White, "")
						},
					),
					DIV("partner-name", HTML(partner.Name.String())),
					DIV("partner-website", A_blank(partner.Website.String(), partner.Website.String())),
					&Form{
						SubmitButtonText:    "Delete",
						SubmitButtonClass:   "partner-delete",
						FormID:              "deletePartner" + partner.ID.Hex(),
						SubmitButtonConfirm: "Are you sure you want to delete this comment: " + partner.Name.Get() + "?",
						OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {

							return "", StringURL("."), event.RemoveEventPartner(&partner)

						},
					},
					DivClearBoth(),
				},
			},
		)
	}

	views = append(views,
		&If{
			Condition: len(listitem) > 0,
			Content: &Tag{
				Tag:     "ul",
				Class:   "sortable",
				Content: listitem,
				Attribs: map[string]string{
					"category": cat,
				},
			},
		},
		&If{
			Condition: len(listitem) > 0,
			Content: &Button{
				Class:   "button saveOrder ",
				Name:    "saveOrder",
				Content: Escape("Save List Order"),
			},
		},
		HTML("<span class='partner-success'>...successfully changed order</span>"),
		BR(),
	)

	return views
}
