package views

import (
	. "github.com/ungerik/go-start/view"
)

var RenderCrazyEgg = Render(
	func(ctx *Context) (err error) {
		// i := strings.LastIndex(context.Request.Host, ":")
		// if context.Request.Host[0:i] == "0.0.0.1" {
		// 	return nil
		// }
		if !Config.IsProductionServer {
			return nil
		}

		ctx.Response.Write([]byte(`<script type="text/javascript">
			setTimeout(function(){var a=document.createElement("script");
			var b=document.getElementsByTagName("script")[0];
			a.src=document.location.protocol+"//dnn506yrbagrg.cloudfront.net/pages/scripts/0013/0417.js?"+Math.floor(new Date().getTime()/3600000);
			a.async=true;a.type="text/javascript";b.parentNode.insertBefore(a,b)}, 1);
			</script>`))

		// usertype := "guest" // admin/guest/host/user
		// defer func() { fmt.Fprintf(writer, "\nmpq.register({user_type:'%s'});", usertype) }()

		// if context.User == nil {
		// 	return nil
		// }
		// if user, ok := ctx.Session.User.(*models.Person); ok {
		// 	email := user.PrimaryEmail()
		// 	if email != "" {
		// 		fmt.Fprintf(writer, "\nmpq.identify('%s');", email)
		// 	}

		// 	name := user.Name.String()
		// 	if name == "" {
		// 		name = email
		// 	}
		// 	if name != "" {
		// 		fmt.Fprintf(writer, "\nmpq.name_tag('%s');", name)
		// 	}

		// 	if user.Admin.Get() {
		// 		usertype = "admin"
		// 	} else if user.EventOrganiser.Get() {
		// 		usertype = "host"
		// 	} else {
		// 		usertype = "user"
		// 	}
		// }

		return nil
	},
)
