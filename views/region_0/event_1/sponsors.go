package event_1

import (
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
	"github.com/ungerik/go-start/debug"
	// "github.com/ungerik/go-start/media"
	. "github.com/ungerik/go-start/view"
	// "image/color"
	"strings"
)

func localEventSponsors() View {
	return DynamicView(
		func(ctx *Context) (view View, err error) {
			var boxes Views
			region := ctx.Data.(*PageData).Region
			event := ctx.Data.(*PageData).Event
			debug.Nop()
			if event.HasEventPartners() {
				boxes, err = renderEventPartner(event)
				if err != nil {
					return nil, err
				}
			} else {

				if region.Slug == "vienna" && event.Number == 6 {
					boxes = Views{
						featuredBox(
							"LOCAL SPONSORS",
							H3(DIV("", "Main partner")),
							A_blank("http://www.ibm.com/isv/startup", IMG("http://dl.dropbox.com/u/5565424/IBM.png", 260)),
							DivClearBoth(),
							H3(DIV("", "Big partners")),
							A_blank("http://speedinvest.com/", IMG("http://dl.dropbox.com/u/5565424/speedinvest.png", 260)),
							A_blank("http://www.hvk.at", IMG("http://dl.dropbox.com/u/5565424/HVK.jpg", 300)),
							DivClearBoth(),
							H3(DIV("", "Medium partners")),
							A_blank("http://www.tecnet.co.at/", IMG("/images/sponsors/tecnet-210x104.png", 220)),
							A_blank("http://www.laudaair.com/", IMG("/images/sponsors/LaudaAir.jpg", 200)),
							DivClearBoth(),
							H3(DIV("", "Small partners")),
							A_blank("https://www.mingo.at/", IMG("http://dl.dropbox.com/u/5565424/MINGO.jpg", 200)),
							DivClearBoth(),
							H3(DIV("", "Location partner")),
							A_blank("http://www.sektor5.at/", IMG("/images/sponsors/sektor5-210x80.png", 160)),
							DivClearBoth(),
							H3(DIV("", "Award partner")),
							A_blank("http://www.a1.net/", IMG("/images/sponsors/A1-105.png", 80)),
							A_blank("http://betahaus.de/", IMG("/images/sponsors/betahaus-210.png", 150)),
						),
					}
				} else if region.Slug == "vienna" && event.Number == 7 {
					boxes = Views{
						featuredBox(
							"LOCAL SPONSORS",
							// H3(DIV("", "Main partner")),
							// DivClearBoth(),
							H3(DIV("", "Big partners")),
							A_blank("http://speedinvest.com/", IMG("http://dl.dropbox.com/u/5565424/speedinvest.png", 260)),
							A_blank("http://www.hvk.at", IMG("http://dl.dropbox.com/u/5565424/HVK.jpg", 300)),
							A_blank("http://www.microsoft.com", IMG("http://dl.dropbox.com/u/5565424/microsoft.png", 300)),
							DivClearBoth(),
							H3(DIV("", "Medium partners")),
							A_blank("http://www.tecnet.co.at/", IMG("/images/sponsors/tecnet-210x104.png", 210, 104)),
							DivClearBoth(),
							// H3(DIV("", "Small partners")),
							// DivClearBoth(),
							// H3(DIV("", "Location partner")),
							// DivClearBoth(),
							// H3(DIV("", "Award partner")),
							H3(DIV("", "Supporters")),
							A_blank("http://www.green-cup.de", IMG("http://dl.dropbox.com/u/5565424/greencupcoffee.jpg", 120)),
							A_blank("http://www.mymuesli.com", IMG("http://dl.dropbox.com/u/5565424/mymuesli.jpg", 250)),
							A_blank("http://www.tedxvienna.at/", IMG("https://dl.dropbox.com/u/8425169/tedxvienna.png", 300)),
						),
					}
				} else if region.Slug == "krems" && event.Number == 1 {
					boxes = Views{
						featuredBox(
							"LOCAL SPONSORS",
							A_blank("http://speedinvest.com/", IMG("http://dl.dropbox.com/u/5565424/speedinvest.png", 210)),
							A_blank("http://www.accent.at/", IMG("/images/sponsors/accent-210x46.png", 210, 46)),
							A_blank("http://www.tecnet.co.at/", IMG("/images/sponsors/tecnet-210x104.png", 210, 104)),
							A_blank("http://www.sektor5.at/", IMG("/images/sponsors/sektor5-210x80.png", 210, 80)),
						),
					}
				} else if region.Slug == "split" && event.Number == 1 {
					boxes = Views{
						featuredBox(
							"LOCAL SPONSORS",
							A_blank("http://www.exportboomers.com", IMG("http://dl.dropbox.com/u/5565424/hrvatski.jpg")),
							A_blank("http://speedinvest.com/", IMG("http://dl.dropbox.com/u/5565424/speedinvest.png", 300)),
							DivClearBoth(),
							A_blank("http://www.domain.me", IMG("http://dl.dropbox.com/u/5565424/domain.me.jpg", 200)),
							A_blank("http://www.h1telekom.hr/", IMG("/images/sponsors/h1-hor_4c_pnb.jpg", 300)),
							DivClearBoth(),
							A_blank("http://www.slicejack.com", IMG("http://dl.dropbox.com/u/5565424/slicejack.png")),
						),
						featuredBox(
							"PREMIUM PARTNERS",
							A_blank("http://www.hp.com", IMG("http://i.crn.com/logos/hp.jpg", 160)),
							A_blank("http://www.split.hr", IMG("/images/sponsors/grad-split.png")),
							DivClearBoth(),
							A_blank("http://www.hgk.hr", IMG("/images/sponsors/HGKlogo.png")),
							A_blank("http://www.utt.hr", IMG("http://dl.dropbox.com/u/5565424/enterpriseeuropenetwork.png", 300)),
							DivClearBoth(),
							A_blank("http://www.eleven.bg", IMG("http://dl.dropbox.com/u/5565424/eleven.png", 160)),
							A_blank("http://www.split.hr", IMG("/images/sponsors/pitagora.JPG", 250)),
						),
						featuredBoxSmall(
							"LOCAL PARTNERS",
							// A_blank("http://www.allaboutapps.at/", IMG("/images/sponsors/logo-aaa-210x172.png", 160)),
							A_blank("http://42.com.hr", IMG("http://dl.dropbox.com/u/5565424/42.png", 160)),
							A_blank("http://maximphoto.tumblr.com", IMG("http://dl.dropbox.com/u/5565424/maxim.jpg", 160)),
							A_blank("http://www.walloftweets.net", IMG("http://dl.dropbox.com/u/5565424/walloftweets.png", 160)),
							A_blank("http://www.codeanywhere.net", IMG("http://dl.dropbox.com/u/5565424/codeanywhere.jpg", 160)),
							DivClearBoth(),
							A_blank("http://www.spotie.com", IMG("http://dl.dropbox.com/u/5565424/spotie.jpg", 160)),
							A_blank("http://www.locastic.com", IMG("http://dl.dropbox.com/u/5565424/locastic.jpg", 160)),
							A_blank("http://www.kutjevacki-vinari.hr/index.php/udruga/o-clanu/camak/", IMG("http://dl.dropbox.com/u/5565424/camak.png", 160)),
							A_blank("http://skica.org", IMG("http://dl.dropbox.com/u/5565424/skica.png")),
							DivClearBoth(),
							A_blank("http://akcija.com.hr", IMG("/images/sponsors/logo_veliki.png", 160)),
							A_blank("http://www.4sec.hr", IMG("/images/sponsors/4sec.png", 140)),
							A_blank("http://www.narodni.net/", IMG("/images/sponsors/narodni468x60.png", 200)),
							A_blank("http://www.seekandhit.com/", IMG("/images/sponsors/1000x1000+linkx2.png", 200)),
						),
						featuredBoxSmall(
							"AWARD PARTNERS",
							A_blank("http://www.schoolforstartups.ro", IMG("http://dl.dropbox.com/u/5565424/school4startups.jpg", 160)),
							A_blank("http://www.pioneersfestival.com", IMG("http://dl.dropbox.com/u/5565424/pioneersfestival.png", 160)),
							A_blank("http://www.plus.hr", IMG("http://dl.dropbox.com/u/5565424/plushosting.png", 160)),
						),
						featuredBoxSmall(
							"MEDIA PARTNERS",
							A_blank("http://thenextweb.com/", IMG("http://dl.dropbox.com/u/5565424/tnw.png", 160)),
							A_blank("http://www.netokracija.com", IMG("http://dl.dropbox.com/u/5565424/Netokracija.jpg", 160)),
							A_blank("http://www.infozona.hr/", IMG("/images/sponsors/infozona.png")),
							A_blank("http://www.bug.hr", IMG("http://dl.dropbox.com/u/5565424/mreza.jpg", 160)),
							DivClearBoth(),
							A_blank("http://planb.tportal.hr", IMG("http://dl.dropbox.com/u/5565424/planb.png", 160)),
							A_blank("http://goaleurope.com/", IMG("http://dl.dropbox.com/u/5565424/goaleurope.jpg", 160)),
							A_blank("http://www.gadgeterija.net", IMG("http://dl.dropbox.com/u/5565424/gadgeterija.jpg", 160)),
							A_blank("http://racunalo.com", IMG("http://dl.dropbox.com/u/5565424/racunalo.png", 160)),
							DivClearBoth(),
							A_blank("http://www.eakademik.com", IMG("http://dl.dropbox.com/u/5565424/eakademik.png", 160)),
							A_blank("http://24sata.hr", IMG("http://dl.dropbox.com/u/5565424/styria.jpg", 160)),
							A_blank("http://www.radiodalmacija.hr", IMG("http://dl.dropbox.com/u/5565424/radiodalmacija.jpg", 160)),
							A_blank("http://www.poslovnipuls.com", IMG("http://dl.dropbox.com/u/5565424/poslovnipuls.jpg", 160)),
							DivClearBoth(),
							A_blank("http://www.rep.hr", IMG("http://dl.dropbox.com/u/5565424/rep.png", 160)),
							A_blank("http://www.totalnifm.hr", IMG("http://dl.dropbox.com/u/5565424/totalnifm.jpg", 160)),
							A_blank("http://zimo.co", IMG("/images/sponsors/zimo.png")),
						),
					}
				} else if region.Slug == "belgrade" && event.Number == 1 {
					boxes = Views{
						featuredBox(
							"LOCAL SPONSORS",
							A_blank("http://www.rda-backa.rs/en", IMG("http://dl.dropbox.com/u/5565424/trda.jpg", 250)),
							A_blank("http://www.kombeg.org.rs", IMG("http://dl.dropbox.com/u/5565424/bcc.jpg", 120)),
							DivClearBoth(),
							A_blank("http://www.sban.eu", IMG("http://dl.dropbox.com/u/5565424/sban.jpg", 150)),
							A_blank("http://www.piratskapartija.com", IMG("http://dl.dropbox.com/u/5565424/piratskapartijasrbije.png", 150)),
						),
						featuredBox(
							"MEDIA PARTNERS",
							A_blank("http://www.kursevi.com/", IMG("http://dl.dropbox.com/u/5565424/kursevi.png", 210)),
							A_blank("http://www.itdogadjaji.com/", IMG("http://dl.dropbox.com/u/5565424/itd.png")),
							A_blank("http://www.netokracija.com", IMG("http://dl.dropbox.com/u/5565424/Netokracija.jpg", 210)),
						),
					}
				} else if region.Slug == "hamburg" && event.Number == 1 {
					boxes = Views{
						featuredBox(
							"LOCAL SPONSORS",
							//A_blank("http://schanzendach.de/", IMG("http://dl.dropbox.com/u/5565424/schanzendach.png")),
							A_blank("http://www.fastbill.com/", IMG("http://dl.dropbox.com/u/5565424/fastbill.png", 300)),
							A_blank("http://www.dailydeal.de/", IMG("http://dl.dropbox.com/u/5565424/dailydeal.jpg", 210)),
							A_blank("http://www.youisnow.de", IMG("http://dl.dropbox.com/u/5565424/youisnow.png", 300)),
							A_blank("http://www.ece.com/", IMG("/images/sponsors/ECE-Logo.jpg", 180)),
							A_blank("http://www.pinck.de", IMG("http://dl.dropbox.com/u/5565424/pinck.jpg", 180)),
						),
						featuredBoxSmall(
							"SUPPORTERS",
							A_blank("http://www.kredito.de/", IMG("http://dl.dropbox.com/u/5565424/kredito.png", 140)),
							A_blank("http://www.iversity.org/", IMG("http://dl.dropbox.com/u/5565424/iversity.png", 140)),
							A_blank("http://metrigo.com/", IMG("http://dl.dropbox.com/u/5565424/metrigo.jpg", 140)),
							A_blank("http://hackfwd.com/", IMG("http://dl.dropbox.com/u/5565424/hackfwd.png", 140)),
							DivClearBoth(),
							A_blank("http://www.makeastartup.com", IMG("http://dl.dropbox.com/u/5565424/makeastartup.png")),
							A_blank("http://www.b-n-p.de", IMG("http://dl.dropbox.com/u/5565424/bnp.png", 140)),
							A_blank("http://www.23company.com", IMG("http://dl.dropbox.com/u/5565424/23.gif", 100)),
							A_blank("http://mymuesli.com/", IMG("http://mymuesli.com/images/start-neu/mm-logo.gif", 140)),
							DivClearBoth(),
							A_blank("http://www.liquidlabs.de", IMG("http://dl.dropbox.com/u/5565424/liquidlabs.jpg", 140)),
							A_blank("http://www.digitalmediawomen.de", IMG("http://dl.dropbox.com/u/5565424/digitalmediawomen.jpg", 140)),
							A_blank("http://www.wunschheim.com", IMG("http://dl.dropbox.com/u/5565424/futurevents.gif", 140)),
							A_blank("http://infernoragazzi.com", IMG("http://dl.dropbox.com/u/5565424/infernoragazzi.jpg", 140)),
							DivClearBoth(),
							A_blank("http://www.knusperreich.de", IMG("http://dl.dropbox.com/u/5565424/knusperreich.png", 140)),
							A_blank("http://www.loftville.com", IMG("http://dl.dropbox.com/u/5565424/loftville.jpg", 140)),
						),
						featuredBox(
							"MEDIA PARTNERS",
							A_blank("http://www.business-punk.com/", IMG("/images/sponsors/business-punk.jpg", 210)),
							A_blank("http://www.hamburg-media.net/", IMG("/images/sponsors/hh_at_work.jpg", 210)),
							DivClearBoth(),
							A_blank("http://www.deutsche-startups.de", IMG("http://dl.dropbox.com/u/5565424/deutschestartups.jpg", 210)),
							A_blank("http://www.gruenderszene.de", IMG("http://dl.dropbox.com/u/5565424/gruenderszene.png", 250)),
							DivClearBoth(),
							A_blank("http://www.maltegoy.de", IMG("http://dl.dropbox.com/u/5565424/maltegoyphotography.jpg", 210)),
						),
					}
				} else if region.Slug == "hamburg" && event.Number == 2 {
					boxes = Views{
						featuredBox(
							"PARTNER",
							A_blank("http://google.de/", IMG("https://dl.dropbox.com/u/8425169/logos/hamburg/google.png", 250)),
							A_blank("http://youisnow.de/", IMG("https://dl.dropbox.com/u/8425169/logos/hamburg/yin.png", 250)),
							A_blank("http://www.it-agile.de/", IMG("https://dl.dropbox.com/u/8425169/logos/hamburg/it-agile.jpg", 250)),
							A_blank("http://www.ebmedien.de/", IMG("https://dl.dropbox.com/u/8425169/logos/hamburg/ebm.png", 250)),
							A_blank("http://www.99designs.com", IMG("https://dl.dropbox.com/u/8425169/logos/hamburg/99designs.png", 250)),
							A_blank("http://www.liquidlabs.de", IMG("https://dl.dropbox.com/u/8425169/logos/hamburg/liquidlabs.jpg", 140)),
						),
						featuredBoxSmall(
							"SUPPORTERS",
							A_blank("http://www.fastbill.com/", IMG("http://dl.dropbox.com/u/5565424/fastbill.png", 140)),
							A_blank("http://mymuesli.com/", IMG("http://mymuesli.com/images/start-neu/mm-logo.gif", 140)),
							A_blank("http://www.knusperreich.de", IMG("http://dl.dropbox.com/u/5565424/knusperreich.png", 140)),
							A_blank("http://www.b-n-p.de", IMG("https://dl.dropbox.com/u/8425169/logos/hamburg/bnp.png", 140)),
							A_blank("http://www.whyown.it", IMG("https://dl.dropbox.com/u/8425169/logos/hamburg/logo_whyownit.png", 90)),
							A_blank("http://infernoragazzi.com", IMG("http://dl.dropbox.com/u/5565424/infernoragazzi.jpg", 90)),
							A_blank("", IMG("https://dl.dropbox.com/u/8425169/logos/hamburg/weeklyhero.png", 140)),
							A_blank("http://www.leuphana.de/professional-school/existenzgruendung.html", IMG("https://dl.dropbox.com/u/8425169/logos/hamburg/EGS_logo.jpg", 140)),
							A_blank("http://www.kreditech.com/", IMG("https://dl.dropbox.com/u/8425169/logos/hamburg/kreditech.png", 140)),
							A_blank("http://www.projektwerk.com/de/", IMG("https://dl.dropbox.com/u/8425169/logos/hamburg/logo_projektwerk.jpg", 140)),
							A_blank("http://www.yelp.de/hamburg", IMG("https://dl.dropbox.com/u/8425169/logos/hamburg/Hamburg_FB_logo.png", 140)),
							A_blank("http://www.loftville.com", IMG("https://dl.dropbox.com/u/8425169/logos/hamburg/loftville_fb_180x180.png", 90)),
							A_blank("http://www.digitalmediawomen.de", IMG("https://dl.dropbox.com/u/8425169/logos/hamburg/dmw.jpg", 140)),
							A_blank("http://www.schweppes.de/", IMG("https://dl.dropbox.com/u/8425169/logos/hamburg/schweppes.jpg", 130)),
							A_blank("http://www.porsche-hamburg.de/portal/loader.php", IMG("https://dl.dropbox.com/u/8425169/logos/hamburg/porsche.png", 90)),
						),
						featuredBoxSmall(
							"MEDIA PARTNERS",
							A_blank("http://www.business-punk.com/", IMG("/images/sponsors/business-punk.jpg", 210)),
							A_blank("http://www.hamburg-media.net/", IMG("/images/sponsors/hh_at_work.jpg", 210)),
							A_blank("http://www.gruenderszene.de", IMG("http://dl.dropbox.com/u/5565424/gruenderszene.png", 250)),
							A_blank("http://griffel-co.com", IMG("https://dl.dropbox.com/u/8425169/logos/hamburg/griffel.png", 210)),
							A_blank("http://www.blog.gruenderplus.de/", IMG("https://dl.dropbox.com/u/8425169/logos/hamburg/gruenderplus.jpg", 210)),
						),
					}
				} else if region.Slug == "cluj-napoca" && event.Number == 1 {
					boxes = Views{
						featuredBox(
							"EVENT PARTNERS",
							A_blank("http://www.businessdays.ro", IMG("http://dl.dropbox.com/u/5565424/adesco.jpg", 210)),
							A_blank("http://www.tedxcluj.com", IMG("http://dl.dropbox.com/u/5565424/tedxcluj.jpg", 210)),
							DivClearBoth(),
							A_blank("http://www.jcicluj.ro", IMG("http://dl.dropbox.com/u/5565424/jcicluj.jpg", 160)),
						),
						featuredBox(
							"LOGISTICS PARTNERS",
							A_blank("http://ro.dallmayr.com", IMG("http://dl.dropbox.com/u/5565424/dallmayrkaffee.jpg", 150)),
							A_blank("http://www.apartamenteregimhoteliercluj.ro", IMG("http://dl.dropbox.com/u/5565424/jonathanapartments.jpg", 150)),
							DivClearBoth(),
							A_blank("http://www.lemnulverde.ro/pages/cofetaria/", IMG("http://dl.dropbox.com/u/5565424/lemnulverde.jpeg", 180)),
							A_blank("http://rapidcatering.ro", IMG("http://dl.dropbox.com/u/5565424/rapidcatering.jpg", 180)),
							DivClearBoth(),
							IMG("http://dl.dropbox.com/u/5565424/clujhubb.jpg", 210),
						),
						featuredBox(
							"MEDIA PARTNERS",
							A_blank("http://www.startups.ro", IMG("http://dl.dropbox.com/u/5565424/startups.ro.jpg", 210)),
							A_blank("http://www.clujlife.com/", IMG("http://dl.dropbox.com/u/5565424/clujlife.png", 140)),
							DivClearBoth(),
							A_blank("http://www.integrahr.ro", IMG("http://dl.dropbox.com/u/5565424/integrahr.png", 210)),
							A_blank("http://www.ilovecluj.ro/", IMG("http://dl.dropbox.com/u/5565424/ilovecluj.png", 160)),
							DivClearBoth(),
							A_blank("http://myaiesec.ro", IMG("http://dl.dropbox.com/u/5565424/aiesec.png", 210)),
							DivClearBoth(),
							A_blank("http://osecluj.ro", IMG("http://dl.dropbox.com/u/5565424/osecluj.png", 160)),
							A_blank("http://cluj.info", IMG("http://dl.dropbox.com/u/5565424/cluj.info.png", 210)),
							DivClearBoth(),
							A_blank("http://cluj-napoca.zilesinopti.ro", IMG("http://dl.dropbox.com/u/5565424/zilesinopti.jpg", 200)),
							A_blank("http://cluj.24fun.ro", IMG("http://dl.dropbox.com/u/5565424/24fun.jpg")),
						),
					}
				} else if region.Slug == "copenhagen" && event.Number == 1 {
					boxes = Views{
						featuredBox(
							"LOCAL SPONSORS",
							A_blank("http://www.carlsberggroup.com", IMG("http://dl.dropbox.com/u/5565424/carlsberg.jpg", 210)),
							A_blank("http://www.podio.com", IMG("http://dl.dropbox.com/u/5565424/podio.png", 210)),
							DivClearBoth(),
							A_blank("http://www.5te.dk", IMG("http://dl.dropbox.com/u/5565424/symbion.jpg", 210)),
							IMG("http://dl.dropbox.com/u/5565424/sas.jpg", 210),
							DivClearBoth(),
							IMG("http://dl.dropbox.com/u/5565424/itu.jpg", 210),
							A_blank("http://gignal.com", IMG("http://dl.dropbox.com/u/5565424/gignal.png", 210)),
							DivClearBoth(),
							IMG("http://dl.dropbox.com/u/5565424/scanomat.jpg", 210),
							IMG("http://dl.dropbox.com/u/5565424/foundershouse.jpg", 210),
							DivClearBoth(),
							IMG("http://dl.dropbox.com/u/5565424/ida.jpg", 210),
							IMG("http://dl.dropbox.com/u/5565424/ciel.jpg", 210),
						),
					}
				} else if region.Slug == "athens" && event.Number == 2 {
					boxes = Views{
						featuredBox(
							"PLATINUM PARTNER",
							A_blank("http://www.ericsson.com", IMG("http://dl.dropbox.com/u/5565424/ericsson.jpg", 200)),
						),
						featuredBox(
							"MAIN PARTNERS",
							A_blank("http://www.randstad.gr", IMG("http://dl.dropbox.com/u/5565424/randstad.jpg", 140)),
						),
						featuredBox(
							"LOCATION PARTNERS",
							//IMG("http://dl.dropbox.com/u/5565424/republic.jpg"),
							A_blank("http://www.alba.edu.gr", IMG("http://dl.dropbox.com/u/5565424/alba.jpg", 300)),
						),
						featuredBox(
							"SUPPORTERS",
							A_blank("http://www.papaki.gr", IMG("http://dl.dropbox.com/u/5565424/pagaki.png", 210)),
							A_blank("http://www.cafetaf.gr", IMG("http://dl.dropbox.com/u/5565424/taf.jpg", 120)),
							DivClearBoth(),
							IMG("http://dl.dropbox.com/u/5565424/paco.png", 210),
							IMG("http://dl.dropbox.com/u/5565424/spiritofwonder.png", 210),
							A_blank("http://www.pepsico.com/", IMG("http://dl.dropbox.com/u/5565424/pepsico.jpg", 210)),
							IMG("http://dl.dropbox.com/u/5565424/athenianbrewery.jpg"),
						),
						featuredBoxSmall(
							"AWARD PARTNERS",
							A_blank("http://www.pioneersfestival.com", IMG("http://dl.dropbox.com/u/5565424/pioneersfestival.png", 210)),
							A_blank("http://www.metavallon.org", IMG("http://dl.dropbox.com/u/5565424/metavallon.png", 210)),
							A_blank("http://www.kariera.gr/", IMG("http://dl.dropbox.com/u/5565424/kariera.png", 200)),
							A_blank("http://www.loft2work.gr", IMG("http://dl.dropbox.com/u/5565424/loft2work.jpg", 100)),
						),
						featuredBoxSmall(
							"UNDER THE AEGIS OF",
							A_blank("http://www.weforum.org/community/global-shapers", IMG("http://dl.dropbox.com/u/5565424/gsc.jpg", 75)),
							A_blank("http://www.britishcouncil.org/gr/greece.htm", IMG("http://dl.dropbox.com/u/5565424/britishcouncil.jpg", 150)),
							A_blank("http://www.startupgreece.gov.gr", IMG("http://dl.dropbox.com/u/5565424/startupgreece.jpg", 75)),
							DivClearBoth(),
							A_blank("http://www.ahk.de/en", IMG("http://dl.dropbox.com/u/5565424/ahk.jpg", 175)),
							A_blank("http://www.repowergreece.com", IMG("http://dl.dropbox.com/u/5565424/repowergreece.jpg", 150)),
							A_blank("http://www.sandbox-network.com", IMG("http://dl.dropbox.com/u/5565424/sandbox.jpg", 150)),
							DivClearBoth(),
							IMG("http://dl.dropbox.com/u/5565424/siepartner.jpg", 175),
							A_blank("http://www.equalsociety.gr", IMG("http://dl.dropbox.com/u/5565424/equalsociety.jpg", 175)),
						),
						featuredBoxSmall(
							"COMMUNICATION PARTNERS",
							A_blank("http://www.dailysecret.com/landingpage/", IMG("http://dl.dropbox.com/u/5565424/dailysecret.png", 120)),
							A_blank("http://www.epixeiro.gr", IMG("http://dl.dropbox.com/u/5565424/epixeiro.jpg", 105)),
							A_blank("http://www.growing.gr", IMG("http://dl.dropbox.com/u/5565424/growing.png", 120)),
							A_blank("http://www.citypress.gr", IMG("http://dl.dropbox.com/u/5565424/citypress.png", 120)),
							DivClearBoth(),
							A_blank("http://www.ypovrixio.gr", IMG("http://dl.dropbox.com/u/5565424/ypovrichio.jpg", 105)),
							A_blank("http://www.pathfinder.gr", IMG("http://dl.dropbox.com/u/5565424/pathfinder.jpg", 105)),
							A_blank("http://goodnews.gr", IMG("http://dl.dropbox.com/u/5565424/goodnews.png", 120)),
							A_blank("http://www.elculture.gr", IMG("http://dl.dropbox.com/u/5565424/elculture.jpg", 120)),
							DivClearBoth(),
							A_blank("http://www.projectyou.gr", IMG("http://dl.dropbox.com/u/5565424/projectyou.png", 120)),
						),
					}
				} else if region.Slug == "hagenberg" && event.Number == 2 {
					boxes = Views{
						featuredBox(
							"LOCAL SPONSORS",
							H3(DIV("", "Main partners")),
							A_blank("http://www.fh-ooe.at", IMG("http://dl.dropbox.com/u/5565424/fhooe.jpg", 210)),
							A_blank("http://www.softwarepark.at", IMG("http://dl.dropbox.com/u/5565424/SoftwareparkHagenberg.png", 210)),
							DivClearBoth(),
							A_blank("http://www.international-incubator.com", IMG("http://dl.dropbox.com/u/5565424/iih.png", 210)),
							DivClearBoth(),
							H3(DIV("", "Big partners")),
							A_blank("http://www.speedinvest.com", IMG("http://dl.dropbox.com/u/5565424/speedinvest.png", 300)),
							A_blank("http://www.ibm.com/isv/startup", IMG("http://dl.dropbox.com/u/5565424/IBM.png", 200)),
							A_blank("http://hvk.at", IMG("http://dl.dropbox.com/u/5565424/HVK.jpg", 300)),
							DivClearBoth(),
							H3(DIV("", "Medium partners")),
							A_blank("http://www.tecnet.co.at", IMG("http://dl.dropbox.com/u/5565424/tecnet.jpg", 120)),
							A_blank("http://www.akostart.at", IMG("http://dl.dropbox.com/u/5565424/akostart.jpg", 150)),
							DivClearBoth(),
							H3(DIV("", "Small partners")),
							A_blank("http://www.studentenwerk.at", IMG("http://dl.dropbox.com/u/5565424/studentenwerk.JPG", 100)),
						),
					}
				} else if region.Slug == "stuttgart" && event.Number == 1 {
					boxes = Views{
						featuredBox(
							"BIG PARTNER",
							A_blank("http://www.brose.com", IMG("http://dl.dropbox.com/u/5565424/brose.jpg", 200)),
						),
						featuredBox(
							"LOCAL PARTNERS",
							A_blank("http://www.region-stuttgart.de", IMG("http://dl.dropbox.com/u/5565424/wirtschaftsfoerderungregionstuttgart.jpg", 210)),
							A_blank("http://www.wac1899.de", IMG("http://dl.dropbox.com/u/5565424/wac.jpg", 150)),
							DivClearBoth(),
							A_blank("http://www.business-chance.de", IMG("http://dl.dropbox.com/u/5565424/businessangelsregionstuttgart.jpg", 210)),
							A_blank("http://www.push-stuttgart.de", IMG("http://dl.dropbox.com/u/5565424/push.png", 150)),
							DivClearBoth(),
							A_blank("http://www.ipogo.de", IMG("http://dl.dropbox.com/u/5565424/ipogo.jpg", 130)),
							A_blank("http://www.wayra.org/de", IMG("http://dl.dropbox.com/u/5565424/wayra.jpg", 150)),
						),
						featuredBox(
							"SUPPORTERS",
							A_blank("http://www.rockstarenergydrink.de", IMG("http://dl.dropbox.com/u/5565424/rockstar.jpg", 200)),
							A_blank("http://www.mymuesli.com", IMG("http://dl.dropbox.com/u/5565424/mymuesli.jpg", 220)),
							DivClearBoth(),
							A_blank("http://www.true-fruits.com", IMG("http://dl.dropbox.com/u/5565424/truefruitsjuice.jpg", 200)),
							A_blank("http://www.fastbill.com", IMG("http://dl.dropbox.com/u/5565424/fastbill.png", 220)),
							DivClearBoth(),
							A_blank("http://www.leihdirwas.de", IMG("http://dl.dropbox.com/u/5565424/leihdirwas.jpg", 200)),
							A_blank("http://www.conceptboard.com", IMG("http://dl.dropbox.com/u/5565424/conceptboard.png", 200)),
							DivClearBoth(),
							A_blank("http://www.iao.fraunhofer.de", IMG("http://dl.dropbox.com/u/5565424/fraunhofer.jpg", 200)),
							A_blank("http://www.wiesingermedia.de", IMG("http://dl.dropbox.com/u/5565424/wiesingermedia.jpg", 200)),
							DivClearBoth(),
							A_blank("http://www.copy-shop-druck.de", IMG("http://dl.dropbox.com/u/5565424/copy-shop-druck.de.jpg", 200)),
							A_blank("http://www.emobility2go.de", IMG("http://dl.dropbox.com/u/5565424/emobility2g.png", 200)),
							A_blank("http://www.stuttgarter-hofbraeu.de", IMG("http://dl.dropbox.com/u/5565424/stuttgarterhofbraeu.jpg", 200)),
						),
						featuredBox(
							"MEDIA PARTNERS",
							A_blank("http://www.foerderland.de", IMG("http://dl.dropbox.com/u/5565424/foerderland.jpg", 180)),
							A_blank("http://www.deutsche-startups.de", IMG("http://dl.dropbox.com/u/5565424/deutschestartups.jpg", 200)),
							DivClearBoth(),
							A_blank("http://www.ideacamp.de", IMG("http://dl.dropbox.com/u/5565424/ideacamp.jpg", 200)),
							A_blank("http://www.startup-stuttgart.de", IMG("http://dl.dropbox.com/u/5565424/startupstuttgart.png", 200)),
							A_blank("http://startupdigest.com", IMG("http://dl.dropbox.com/u/5565424/startupdigeststuttgart.JPG", 200)),
						),
					}
				} else if region.Slug == "alicante" && event.Number == 1 {
					boxes = Views{
						featuredBox(
							"LOCAL PARTNERS",
							A_blank("http://www.alicante-ayto.es/empleo/home.html", IMG("http://dl.dropbox.com/u/5565424/agencialocal.jpg", 210)),
							A_blank("http://www.bbooster.org", IMG("http://dl.dropbox.com/u/5565424/businessbooster.jpg", 210)),
							DivClearBoth(),
							A_blank("http://www.tainforma.net/", IMG("http://dl.dropbox.com/u/5565424/tainforma.jpg", 210)),
							IMG("http://dl.dropbox.com/u/5565424/catedrabancaja.jpg", 210),
							DivClearBoth(),
							IMG("http://dl.dropbox.com/u/5565424/compromisosocialbanaja.jpg", 210),
							A_blank("http://www.laverdad.es", IMG("http://dl.dropbox.com/u/5565424/laverdad.jpg", 210)),
							DivClearBoth(),
							A_blank("http://web.ua.es/dccia/", IMG("http://dl.dropbox.com/u/5565424/escuela.jpg", 210)),
							IMG("http://dl.dropbox.com/u/5565424/observatoriouniversitario.jpg", 210),
						),
					}
				} else if region.Slug == "tirana" && event.Number == 1 {
					boxes = Views{
						featuredBox(
							"LOCAL SPONSORS",
							IMG("http://dl.dropbox.com/u/5565424/abissnet.png", 210),
							IMG("http://dl.dropbox.com/u/5565424/shekulli.jpg", 210),
						),
						featuredBox(
							"LOCAL PARTNERS",
							IMG("http://dl.dropbox.com/u/5565424/fakultetiekonomise.png", 210),
							IMG("http://dl.dropbox.com/u/5565424/ksfe.png", 210),
						),
					}
				} else if region.Slug == "graz" && event.Number == 2 {
					boxes = Views{
						featuredBox(
							"BIG PARTNERS",
							A_blank("http://speedinvest.com/", IMG("http://dl.dropbox.com/u/5565424/speedinvest.png", 260)),
							A_blank("http://www.hvk.at", IMG("http://dl.dropbox.com/u/5565424/HVK.jpg", 300)),
							A_blank("http://www.sciencepark.at", IMG("http://dl.dropbox.com/u/5565424/sciencepark.gif")),
						),
						featuredBox(
							"MEDIUM PARTNERS",
							A_blank("http://www.tecnet.co.at/", IMG("/images/sponsors/tecnet-210x104.png", 210, 104)),
							A_blank("http://www.styria.com", IMG("http://dl.dropbox.com/u/5565424/styria.png", 250)),
							DivClearBoth(),
							A_blank("http://www.bertl-fattinger.at/", IMG("http://dl.dropbox.com/u/5565424/bfp.jpg", 200)),
							A_blank("http://www.austin.at/", IMG("http://dl.dropbox.com/u/5565424/austin.JPG", 200)),
						),
						featuredBox(
							"LOCAL PARTNERS",
							A_blank("http://www.engarde.at", IMG("http://dl.dropbox.com/u/5565424/engarde.jpg", 200)),
							A_blank("http://www.eco.at", IMG("http://dl.dropbox.com/u/5565424/ecoworldstyria.gif", 200)),
							DivClearBoth(),
							IMG("http://dl.dropbox.com/u/5565424/brauereigratzer.jpg", 200),
							IMG("http://dl.dropbox.com/u/5565424/stadtgraz.jpg", 200),
							DivClearBoth(),
							IMG("http://dl.dropbox.com/u/5565424/managerie.jpg", 200),
							IMG("http://dl.dropbox.com/u/5565424/softnet.png", 200),
						),
						featuredBox(
							"SUPPORTERS",
							A_blank("http://www.green-cup.de", IMG("http://dl.dropbox.com/u/5565424/greencupcoffee.jpg", 160)),
							A_blank("http://www.mymuesli.com", IMG("http://dl.dropbox.com/u/5565424/mymuesli.jpg", 250)),
						),
					}
				} else if region.Slug == "klagenfurt" && event.Number == 1 {
					boxes = Views{
						featuredBox(
							"BIG PARTNERS",
							A_blank("http://speedinvest.com/", IMG("http://dl.dropbox.com/u/5565424/speedinvest.png", 260)),
							A_blank("http://www.hvk.at", IMG("http://dl.dropbox.com/u/5565424/HVK.jpg", 300)),
						),
						featuredBox(
							"MEDIUM PARTNERS",
							A_blank("http://www.tecnet.co.at/", IMG("/images/sponsors/tecnet-210x104.png", 210, 104)),
						),
						featuredBox(
							"LOCAL PARTNERS",
							A_blank("http://www.build.or.at", IMG("http://dl.dropbox.com/u/5565424/build!.jpg", 210)),
							A_blank("http://www.sparkasse.at/kaernten", IMG("http://dl.dropbox.com/u/5565424/kaerntnersparkasse.jpg", 210)),
							DivClearBoth(),
							A_blank("https://www.inismo.com", IMG("http://dl.dropbox.com/u/5565424/inismo.jpg", 180)),
							A_blank("http://www.sgz.at", IMG("http://dl.dropbox.com/u/5565424/sgz.png", 210)),
							DivClearBoth(),
							A_blank("http://www.kwf.at", IMG("http://dl.dropbox.com/u/5565424/kwf.jpg", 210)),
							A_blank("http://www.aiesec.ata", IMG("http://dl.dropbox.com/u/5565424/aiesec.png", 210)),
							DivClearBoth(),
							A_blank("http://www.websafari.eu", IMG("http://dl.dropbox.com/u/5565424/websafari.jpg", 210)),
						),
						featuredBox(
							"SUPPORTERS",
							A_blank("http://www.green-cup.de", IMG("http://dl.dropbox.com/u/5565424/greencupcoffee.jpg", 160)),
							A_blank("http://www.mymuesli.com", IMG("http://dl.dropbox.com/u/5565424/mymuesli.jpg", 250)),
						),
					}
				} else if region.Slug == "bratislava" && event.Number == 1 {
					boxes = Views{
						featuredBox(
							"PARTNERS",
							A_blank("http://www.nadsme.sk/", IMG("https://dl.dropbox.com/u/8425169/logos/bratislava/logo-nadsme.jpg", 150)),
							A_blank("http://www.neulogy.com/", IMG("https://dl.dropbox.com/u/8425169/logos/bratislava/Neulogy.jpg", 274)),
							A_blank("", IMG("https://dl.dropbox.com/u/8425169/logos/bratislava/patria_logo.jpg", 200)),
						),
						featuredBox(
							"MEDIUM PARTNERS",
							A_blank("http://www.tecnet.co.at/", IMG("/images/sponsors/tecnet-210x104.png", 210, 104)),
							A_blank("http://www.accent.at/", IMG("/images/sponsors/accent-210x46.png", 210, 46)),
						),
						featuredBox(
							"SUPPORTERS",
							A_blank("http://www.tvojaskusenost.sk", IMG("https://dl.dropbox.com/u/8425169/logos/bratislava/aiesec.png", 220)),
							A_blank("http://www.avis.sk/", IMG("https://dl.dropbox.com/u/8425169/logos/bratislava/avis.jpg", 220)),
							A_blank("http://www.hollandia.sk", IMG("https://dl.dropbox.com/u/8425169/logos/bratislava/hollandia.jpg", 220)),
							A_blank("http://www.karriere.at/", IMG("https://dl.dropbox.com/u/8425169/logos/bratislava/karriereAT_4c.JPG", 220)),
							A_blank("http://www.meinl.sk", IMG("https://dl.dropbox.com/u/8425169/logos/bratislava/meinl.jpg", 220)),
							A_blank("http://www.fei.stuba.sk/", IMG("https://dl.dropbox.com/u/8425169/logos/bratislava/STUFEIlogo.jpg", 180)),
							A_blank("http://www.startupawards.sk/", IMG("https://dl.dropbox.com/u/8425169/logos/bratislava/sua.jpg", 220)),
							A_blank("http://www.technikum-wien.at/", IMG("http://dl.dropbox.com/u/8425169/logos/bratislava/FHTechnikum.jpg", 200)),
							A_blank("http://www.wilddragon.at/", IMG("http://dl.dropbox.com/u/8425169/logos/bratislava/WildDragonLogo.jpg", 110)),
						),
						featuredBox(
							"MEDIA PARTNERS",
							A_blank("http://www.best-bratislava.sk", IMG("https://dl.dropbox.com/u/8425169/logos/bratislava/bestbratislava.png", 220)),
							A_blank("http://www.egoodwill.sk", IMG("https://dl.dropbox.com/u/8425169/logos/bratislava/egoodwill.jpg", 220)),
							A_blank("http://futurezone.at/", IMG("https://dl.dropbox.com/u/8425169/logos/bratislava/fuzo-logo.jpg", 220)),
							A_blank("http://www.itnews.sk/tituly/infoware", IMG("https://dl.dropbox.com/u/8425169/logos/bratislava/infoware.jpg", 220)),
							A_blank("http://www.mc2.sk", IMG("https://dl.dropbox.com/u/8425169/logos/bratislava/dotoho.png", 220)),
							A_blank("http://www.mladiinfo.sk/", IMG("https://dl.dropbox.com/u/8425169/logos/bratislava/mladiinfo.png", 220)),
							A_blank("http://www.my-career.at/", IMG("https://dl.dropbox.com/u/8425169/logos/bratislava/mc.jpg", 220)),
							A_blank("http://www.eurocampus.sk/", IMG("https://dl.dropbox.com/u/8425169/logos/bratislava/student-life.jpg", 220)),
							A_blank("http://www.etrend.sk/", IMG("https://dl.dropbox.com/u/8425169/logos/bratislava/TREND_logo.jpg", 220)),
							A_blank("http://www.tyinternety.cz/", IMG("http://dl.dropbox.com/u/8425169/logos/prague/tyinternety.png", 220)),
							A_blank("http://www.zajtra.sk", IMG("https://dl.dropbox.com/u/8425169/logos/bratislava/logozajtra.jpg", 220)),
							A_blank("http://www.zive.sk/", IMG("https://dl.dropbox.com/u/8425169/logos/bratislava/zive.jpg", 220)),
							A_blank("http://www.zones.sk/", IMG("https://dl.dropbox.com/u/8425169/logos/bratislava/zones.png", 220)),
						),
					}
				} else if region.Slug == "prague" && event.Number == 1 {
					boxes = Views{
						featuredBox(
							"PARTNERS",
							A_blank("http://www.credoventures.com/", IMG("http://dl.dropbox.com/u/8425169/logos/prague/credo-ventures.png", 265)),
							A_blank("http://www.seznam.cz/", IMG("http://dl.dropbox.com/u/8425169/logos/prague/seznam_claim.jpg", 300)),
							A_blank("http://www.allegrogroup.cz/", IMG("https://dl.dropbox.com/u/8425169/logos/prague/allegrogroup.png", 300)),
							A_blank("http://www.havelholasek.cz/en", IMG("https://dl.dropbox.com/u/8425169/logos/prague/hhp.jpg", 300)),
							A_blank("http://www.csas.cz/", IMG("https://dl.dropbox.com/u/8425169/logos/prague/Ceska_sporitelna.jpg", 250)),
							A_blank("#", IMG("https://dl.dropbox.com/u/8425169/logos/prague/tomascupr.jpg", 250)),
						),
						featuredBox(
							"MEDIUM PARTNERS",
							A_blank("http://www.tecnet.co.at/", IMG("/images/sponsors/tecnet-210x104.png", 210, 104)),
							A_blank("http://www.accent.at/", IMG("/images/sponsors/accent-210x46.png", 210, 46)),
						),
						featuredBox(
							"SUPPORTERS",
							A_blank("http://www.rozjezdyroku.cz/", IMG("http://dl.dropbox.com/u/8425169/logos/prague/rozjezdy.png", 265)),
							A_blank("http://www.limeandtonic.com/prague/en/", IMG("http://dl.dropbox.com/u/8425169/logos/prague/limeandtonic.png", 150)),
							A_blank("http://www.google.com/", IMG("https://dl.dropbox.com/u/8425169/logos/prague/Google_Logo_292X94.png", 294)),
							A_blank("http://geekshop.cz/", IMG("https://dl.dropbox.com/u/8425169/logos/prague/geekshop.png", 294)),
						),
						featuredBox(
							"MEDIA PARTNERS",
							A_blank("http://www.tyinternety.cz/", IMG("http://dl.dropbox.com/u/8425169/logos/prague/tyinternety.png", 300)),
						),
					}
				} else if region.Slug == "bucharest" && event.Number == 1 {
					boxes = Views{
						featuredBox(
							"PARTNERS",
							A_blank("http://www.securitas.com/ro/ro/", IMG("https://dl.dropbox.com/u/8425169/logos/bucharest/securitas.jpg", 200)),
							A_blank("http://www.brandsandgangs.com/", IMG("https://dl.dropbox.com/u/8425169/logos/bucharest/brandsandgangs.jpg", 300)),
							A_blank("http://www.staropramen.com/ro", IMG("https://dl.dropbox.com/u/8425169/logos/bucharest/staropramen.jpg", 220)),
							A_blank("http://www.aquacarpatica.com", IMG("https://dl.dropbox.com/u/8425169/logos/bucharest/aqua-carpatica.png", 220)),
						),
						featuredBox(
							"MEDIA PARTNERS",
							A_blank("http://www.wall-street.ro/", IMG("https://dl.dropbox.com/u/8425169/logos/bucharest/wall-street-rgb.jpg", 200)),
							A_blank("http://www.jaderomania.org/jade-romania/", IMG("https://dl.dropbox.com/u/8425169/logos/bucharest/jade.jpg", 200)),
							A_blank("http://sisc.ro/", IMG("https://dl.dropbox.com/u/8425169/logos/bucharest/sisc.jpg", 200)),
							A_blank("http://leap.ro/", IMG("https://dl.dropbox.com/u/8425169/logos/bucharest/leap.jpg", 200)),
							A_blank("http://www.bestresource.ro", IMG("https://dl.dropbox.com/u/8425169/logos/bucharest/bestresource.png", 200)),
							A_blank("http://www.hipo.ro/", IMG("https://dl.dropbox.com/u/8425169/logos/bucharest/hipo.jpg", 200)),
							A_blank("http://think-business.ro/", IMG("https://dl.dropbox.com/u/8425169/logos/bucharest/think%20business.jpg", 200)),
							A_blank("http://www.4career.ro/", IMG("https://dl.dropbox.com/u/8425169/logos/bucharest/4career.jpg", 200)),
							A_blank("http://www.advicestudents.ro/", IMG("https://dl.dropbox.com/u/8425169/logos/bucharest/advice.png", 200)),
							A_blank("http://www.startups.ro/", IMG("https://dl.dropbox.com/u/8425169/logos/bucharest/startupsro.jpg", 200)),
							A_blank("http://www.lsrs.ro/", IMG("https://dl.dropbox.com/u/8425169/logos/bucharest/lsrs.jpg", 200)),
						),
					}
				} else if region.Slug == "istanbul" && event.Number == 1 {
					boxes = Views{
						featuredBox(
							"PARTNERS",
							A_blank("http://bilgiligirisimciler.com/", IMG("https://dl.dropbox.com/u/8425169/logos/istanbul/Partners/Bilgili.png", 180)),
							A_blank("http://www.bilgi.edu.tr/", IMG("https://dl.dropbox.com/u/8425169/logos/istanbul/Partners/bilgi.jpg", 220)),
						),
						featuredBox(
							"Bosphorus Sponsors (Main Sponsors)",
							A_blank("http://www.microsoft.com/tr-tr/default.aspx", IMG("https://dl.dropbox.com/u/8425169/logos/istanbul/BosphorusSponsors/microsoft.png", 200)),
							A_blank("http://www.webrazzi.com/", IMG("https://dl.dropbox.com/u/8425169/logos/istanbul/BosphorusSponsors/webrazzi.jpg", 200)),
							A_blank("http://www.galatabusinessangels.com/ana-sayfa", IMG("https://dl.dropbox.com/u/8425169/logos/istanbul/BosphorusSponsors/gba.png", 200)),
							A_blank("http://www.endeavor.org.tr/tr-TR/anasayfa.aspx", IMG("https://dl.dropbox.com/u/8425169/logos/istanbul/Supporters/endeavor.png", 200)),
						),
						featuredBox(
							"SUPPORTERS",
							A_blank("http://launchub.com/beta/", IMG("https://dl.dropbox.com/u/8425169/logos/istanbul/Supporters/launchub.png", 200)),
							A_blank("http://www.inventram.com/", IMG("https://dl.dropbox.com/u/8425169/logos/istanbul/Supporters/inventram.jpg", 200)),
							A_blank("http://shiftdelete.net/", IMG("https://dl.dropbox.com/u/8425169/logos/istanbul/Supporters/sdn.png", 200)),
							A_blank("http://www.networkingakademi.com/", IMG("https://dl.dropbox.com/u/8425169/logos/istanbul/Supporters/akademi.jpg", 200)),
							A_blank("http://curatingturkey.com/", IMG("https://dl.dropbox.com/u/8425169/logos/istanbul/Supporters/sonkesikopya.jpg", 200)),
							A_blank("http://agile.ir/", IMG("https://dl.dropbox.com/u/8425169/logos/istanbul/Supporters/agile_community.png", 200)),
							A_blank("http://www.intermediateknoloji.com/index.html", IMG("https://dl.dropbox.com/u/8425169/logos/istanbul/Supporters/intermedia.png", 200)),
							A_blank("https://www.venturro.com/Default", IMG("https://dl.dropbox.com/u/8425169/logos/istanbul/Supporters/venturro.png", 200)),
							A_blank("http://turkey.enjoyurbanstation.com/tr/", IMG("https://dl.dropbox.com/u/8425169/logos/istanbul/Supporters/urbanstation.jpg", 150)),
							A_blank("http://www.girisimhaber.com/", IMG("https://dl.dropbox.com/u/8425169/logos/istanbul/Supporters/gh.jpg", 200)),
							A_blank("http://www.sislisosyetepazari.com.tr/", IMG("https://dl.dropbox.com/u/8425169/logos/istanbul/Supporters/ssplogomavi.jpg", 120)),
							A_blank("http://www.isikozalit.com/", IMG("https://dl.dropbox.com/u/8425169/logos/istanbul/Supporters/isik.jpg", 150)),
							A_blank("http://www.tamindir.com/", IMG("https://dl.dropbox.com/u/8425169/logos/istanbul/Supporters/tamindir-logo.png", 200)),
						),
					}
				}
			}

			/*
                        if region.Slug == "vienna" {
				boxes = append(boxes,
					Views{
						featuredBox(
							"STRATEGIC PARTNERS",
							A_blank("http://www.bmwfj.gv.at/", IMG("http://pioneersfestival.com/wp-content/uploads/2012/07/Logo-BMWFJ-CMYK-01-e1342697673507.png")),
						),
						featuredBox(
							"SUBSIDISED BY",
							A_blank("http://www.impulse-awsg.at/", IMG("http://pioneersfestival.com/wp-content/uploads/2012/07/impulseaws_evolvebmwfj_ohnelogo_4c-01-e1342697653827.png")),
						),
					})
			}
                        */

			return &Div{
				Class:   "local-sponsors",
				Content: boxes,
			}, nil
		},
	)
}

func renderEventPartner(event *models.Event) (Views, error) {
	views := make([]View, 10)

	partners := event.EventPartners
	for j := 0; j < len(partners); j++ {
		categoryname := partners[j].Name
		// subpartners, err := event.GetPartnersByCategory(categoryname.Get())
		var subviews Views
		for k := 0; k < len(partners[j].Partners); k++ {
			l := k

			if !event.EventPartners[j].Partners[l].IsEmpty() {
				var p models.Partner
				err := event.EventPartners[j].Partners[l].Get(&p)
				if err != nil {
					return nil, err
				}

				subviews = append(subviews,
					DynamicView(
						func(ctx *Context) (view View, err error) {
							var subviews Views

							//img, err := p.Logo.VersionTouchOrigFromOutsideView(200, 200, media.HorCenter, media.VerCenter, false, color.White, "")
							img, err := p.Logo.OriginalVersionView("")

							if err != nil {
								return nil, err
							}
							subviews = append(subviews, A_blank(p.Website.Get(), IMG(img.URL.URL(ctx), 200)))

							return subviews, nil
						},
					),
				)
			}
		}
		views[partners[j].Order] = DIV("featured-box",
			DIV("box-title", Escape(strings.ToUpper(categoryname.Get()))),
			subviews,
			DivClearBoth(),
		)
	}

	return views, nil

}
