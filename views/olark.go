package views

import (
	"github.com/ungerik/go-start/user"
	. "github.com/ungerik/go-start/view"

	"github.com/STARTeurope/startuplive.in/models"
)

var RenderOlark = Render(
	func(ctx *Context) (err error) {

		return nil // deactivated for now

		ctx.Response.WriteString(`<!-- begin olark code --><script type='text/javascript'>/*{literal}<![CDATA[*/window.olark||(function(i){var e=window,h=document,a=e.location.protocol=="https:"?"https:":"http:",g=i.name,b="load";(function(){e[g]=function(){(c.s=c.s||[]).push(arguments)};var c=e[g]._={},f=i.methods.length; while(f--){(function(j){e[g][j]=function(){e[g]("call",j,arguments)}})(i.methods[f])} c.l=i.loader;c.i=arguments.callee;c.f=setTimeout(function(){if(c.f){(new Image).src=a+"//"+c.l.replace(".js",".png")+"&"+escape(e.location.href)}c.f=null},20000);c.p={0:+new Date};c.P=function(j){c.p[j]=new Date-c.p[0]};function d(){c.P(b);e[g](b)}e.addEventListener?e.addEventListener(b,d,false):e.attachEvent("on"+b,d); (function(){function l(j){j="head";return["<",j,"></",j,"><",z,' onl'+'oad="var d=',B,";d.getElementsByTagName('head')[0].",y,"(d.",A,"('script')).",u,"='",a,"//",c.l,"'",'"',"></",z,">"].join("")}var z="body",s=h[z];if(!s){return setTimeout(arguments.callee,100)}c.P(1);var y="appendChild",A="createElement",u="src",r=h[A]("div"),G=r[y](h[A](g)),D=h[A]("iframe"),B="document",C="domain",q;r.style.display="none";s.insertBefore(r,s.firstChild).id=g;D.frameBorder="0";D.id=g+"-loader";if(/MSIE[ ]+6/.test(navigator.userAgent)){D.src="javascript:false"} D.allowTransparency="true";G[y](D);try{D.contentWindow[B].open()}catch(F){i[C]=h[C];q="javascript:var d="+B+".open();d.domain='"+h.domain+"';";D[u]=q+"void(0);"}try{var H=D.contentWindow[B];H.write(l());H.close()}catch(E){D[u]=q+'d.write("'+l().replace(/"/g,String.fromCharCode(92)+'"')+'");d.close();'}c.P(2)})()})()})({loader:(function(a){return "static.olark.com/jsclient/loader0.js?ts="+(a?a[1]:(+new Date))})(document.cookie.match(/olarkld=([0-9]+)/)),name:"olark",methods:["configure","extend","declare","identify"]});`)
		defer ctx.Response.WriteString("\nolark.identify('5131-170-10-8312');/*]]>{/literal}*/\n</script><!-- end olark code -->")

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
				ctx.Response.Printf("\nolark('api.visitor.updateEmailAddress', {emailAddress: '%s'});", email)
			}

			name := person.Name.String()
			if name != "" {
				ctx.Response.Printf("\nolark('api.visitor.updateFullName', {fullName: '%s'});", name)
			}
		}

		return nil
	},
)
