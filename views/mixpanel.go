package views

import (
	"fmt"

	"github.com/ungerik/go-start/user"
	. "github.com/ungerik/go-start/view"

	"github.com/STARTeurope/startuplive.in/models"
)

var RenderMixpanel = Render(
	func(ctx *Context) (err error) {
		if !Config.IsProductionServer {
			return nil
		}

		ctx.Response.WriteString(`<!-- start Mixpanel --><script type="text/javascript">(function(c,a){window.mixpanel=a;var b,d,h,e;b=c.createElement("script");b.type="text/javascript";b.async=!0;b.src=("https:"===c.location.protocol?"https:":"http:")+'//cdn.mxpnl.com/libs/mixpanel-2.1.min.js';d=c.getElementsByTagName("script")[0];d.parentNode.insertBefore(b,d);a._i=[];a.init=function(b,c,f){function d(a,b){var c=b.split(".");2==c.length&&(a=a[c[0]],b=c[1]);a[b]=function(){a.push([b].concat(Array.prototype.slice.call(arguments,0)))}}var g=a;"undefined"!==typeof f?
		g=a[f]=[]:f="mixpanel";g.people=g.people||[];h="disable track track_pageview track_links track_forms register register_once unregister identify name_tag set_config people.identify people.set people.increment".split(" ");for(e=0;e<h.length;e++)d(g,h[e]);a._i.push([b,c,f])};a.__SV=1.1})(document,window.mixpanel||[]);
		mixpanel.init("499ad0e09a4fd025bf95d54b67589f29");
		mixpanel.track('page viewed', {'page name': document.title, 'url': window.location.pathname});`)
		defer ctx.Response.WriteString("\n</script><!-- end Mixpanel -->")

		usertype := "guest" // admin/guest/host/user
		defer func() {
			fmt.Fprintf(ctx.Response, "\nmixpanel.register({user_type:'%s'});", usertype)
		}()

		if !user.LoggedIn(ctx.Session) {
			return nil
		}
		var person models.Person
		found, err := user.OfSession(ctx.Session, &person)
		if err != nil {
			return err
		}
		if found {
			email := person.PrimaryEmail()
			if email != "" {
				fmt.Fprintf(ctx.Response, "\nmixpanel.people.identify('%s');", email)
			}

			name := person.Name.String()
			if name == "" {
				name = email
			}
			if name != "" {
				fmt.Fprintf(ctx.Response, "\nmixpanel.people.set({ $name: 'John Smith', $last_login: new Date() });")
			}

			if person.Admin.Get() {
				usertype = "admin"
			} else if person.EventOrganiser.Get() {
				usertype = "host"
			} else {
				usertype = "user"
			}
		}

		return nil
	},
)
