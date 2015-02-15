package admin

import (
	"strings"

	"github.com/AlexTi/go-plesk"
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/STARTeurope/startuplive.in/views/admin"
	"github.com/ungerik/go-start/debug"
	user "github.com/ungerik/go-start/user"
	. "github.com/ungerik/go-start/view"
)

func init() {
	Region0_Event1_Admin_Organiser2 = &Page{
		OnPreRender: SetEventPersonPageData,
		Title: Render(
			func(ctx *Context) (err error) {
				person := ctx.Data.(*PageData).Person
				return EventAdminTitle("Organiser " + person.Name.String()).Render(ctx)
			},
		),
		CSS: IndirectURL(&Region0_DashboardCSS),
		Scripts: Renderers{
			admin.PageScripts,
		},
		Content: Views{
			eventadminHeader(),
			DynamicView(eventAdminOrganiserView),
		},
	}
}

func eventAdminOrganiserView(ctx *Context) (view View, err error) {
	event := ctx.Data.(*PageData).Event
	// isAdmin := SessionUserIsSuperAdmin(ctx)

	person := ctx.Data.(*PageData).Person
	person.EventOrganiser = true

	organisermail := person.OrganiserEmail.String()
	organiserforwardmail := person.OrganiserForwardingEmail.String()

	exludedFields, err := ExcludedPersonFormFields(ctx)
	exludedFields = append(exludedFields, []string{
		"Name.Prefix",
		"Name.Postfix",
		"PostalAddress",
		"Phone",
		"Web",
		"Xing",
		"GitHub",
		"BirthDate",
		"Images",
		"Judge",
		"JudgeInfo",
		"Mentor",
		"MentorInfo",
		"FeaturedMentor",
		"FeaturedMentorInfo",
	}...)
	if err != nil {
		return nil, err
	}

	requiredFields := []string{}
	if !Config.IsProductionServer {
		requiredFields = append(requiredFields, []string{
			"Name.First",
			"Name.Last",
		}...)
	} else {
		requiredFields = append(requiredFields, []string{
			"Name.First",
			"Name.Last",
			"OrganiserEmail",
			"OrganiserForwardingEmail",
		}...)
	}
	// person.OrganiserEmail.Set("firstname.lastname@startuplive.in")
	//mentorsURL := Region0_Event1_Admin_Organisers.URL(response, ctx.URLArgs[0], ctx.URLArgs[1])

	organiserform := &Form{
		SubmitButtonText:  "Save Person Data",
		SubmitButtonClass: "button",
		FormID:            "person" + person.ID.Hex(),
		GetModel:          FormModel(person),
		ExcludedFields:    append(AlwaysExcludedPersonFormFields, exludedFields...),
		DisabledFields: []string{
			"EventOrganiser",
		},
		RequiredFields: requiredFields,
		OnSubmit: func(form *Form, formModel interface{}, ctx *Context) (string, URL, error) {
			found := false

			for i := 0; i < len(person.User.Email); i++ {
				if person.User.Email[i].Address.String() == person.OrganiserEmail.String() {
					found = true

				}
			}

			if !found {
				if person.OrganiserEmail.String() != "" {
					person.User.Email = append(person.User.Email, user.EmailIdentity{
						Address:     person.OrganiserEmail,
						Description: "Official Organiser Email",
					})
				}
			}

			if organisermail != person.OrganiserEmail.String() ||
				organiserforwardmail != person.OrganiserForwardingEmail.String() {

				manageOrganiserEmail(person, event)
			}

			return "", Region0_Event1_Admin_Organisers, person.Save()
		},
	}

	// if !isAdmin {
	// 	organiserform.ExcludedFields = append(organiserform.ExcludedFields, "OrganiserEmail", "OrganiserForwardingEmail")
	// }

	views := Views{
		H2("Organiser " + person.Name.String()),
		//PersonForm(person, Region0_Event1_Admin_Organisers, []string{"EventOrganiser"}),
		organiserform,
	}
	return views, nil
}

func manageOrganiserEmail(person *models.Person, event *models.Event) (err error) {

	plesk := plesk.NewPleskApi(PleskUrl, PleskUser, PleskPW, false)

	// POST requets
	debug.Nop()

	// headers := []string{
	// 	"HTTP_AUTH_LOGIN:starteurope",
	// 	"HTTP_AUTH_PASSWD:Vhiluoomhag2",
	// 	"Content-Type: text/xml",
	// }

	name := person.OrganiserEmail.String()
	name = strings.Split(name, "@")[0]

	// alias := ""

	forward := []string{person.OrganiserForwardingEmail.String()}

	///////////////////////// 
	// CHECK for mail account

	mailacc, err := plesk.EmailExists(name)
	if err != nil {
		return err
	}

	if mailacc.Status != "error" {
		// debug.Print("update organsier email")
		plesk.UpdateEmail(name, forward)

		///////////////////////// 
		// CHECK for mailing list 

		var region models.EventRegion
		err := event.Region.Get(&region)
		if err != nil {
			return err
		}

		list := region.Slug.Get() + "" + event.Number.String()
		// debug.Print("mailing list: " + list)

		mailingList, err := plesk.EmailExists(list)
		if err != nil {
			return err
		}

		name = name + "@startuplive.in"
		listforwarding := []string{name}
		if mailingList.Status != "error" {
			// debug.Print("update organsier email: " + list + " forward to: " + name)

			listforwarding = append(listforwarding, mailingList.ForwardingAdresses...)
			// debug.Print("forwarding: " + mailingList.ForwardingAdresses[0])
			plesk.UpdateEmail(list, listforwarding)
		} else {
			// debug.Print("create organsier email: " + list + " forward to: " + name)
			plesk.CreateEmail(list, listforwarding)
		}

		// mailingList, err := mailingListExists(plesk, list)
		// if err != nil {
		// 	return err
		// }
		// // debug.Print("list status: " + mailingList.Status)

		// name = name + "@startuplive.in"
		// if mailingList.Status != "error" {
		// 	addEmailToList(plesk, list, name)
		// } else {
		// 	createMailingList(plesk, list)
		// 	addEmailToList(plesk, list, name)
		// }

	} else {
		// debug.Print("create organsier email")
		plesk.CreateEmail(name, forward)

		///////////////////////// 
		// CHECK for mailing list 

		var region models.EventRegion
		err := event.Region.Get(&region)
		if err != nil {
			return err
		}

		list := region.Slug.Get() + "" + event.Number.String()
		// debug.Print("mailing list: " + list)

		mailingList, err := plesk.EmailExists(list)
		if err != nil {
			return err
		}
		// debug.Print("list status: " + mailingList.Status)

		name = name + "@startuplive.in"
		listforwarding := []string{name}
		if mailingList.Status != "error" {
			// debug.Print("update organsier email: " + list + " forward to: " + name)

			listforwarding = append(listforwarding, mailingList.ForwardingAdresses...)
			// debug.Print("forwarding: " + mailingList.ForwardingAdresses[0])
			plesk.UpdateEmail(list, listforwarding)
		} else {
			// debug.Print("create organsier email: " + list + " forward to: " + name)
			plesk.CreateEmail(list, listforwarding)
		}

	}

	return nil
}
