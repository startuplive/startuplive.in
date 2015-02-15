package views

import (
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/ungerik/go-start/view"
	"time"
)

const (
	ContactEmail = "contact@startuplive.in"
	SupportEmail = "support@startuplive.in"

	PublicPageCacheDuration = time.Hour
	HomepageNumBlogPosts    = 3

	HiddenStartupLiveBlogURL = "http://blog.startuplive.in/"
	StartupLiveBlogFeedURL   = "http://blog.startuplive.in/feed/"

	STARTeuropeFacebookURL = "https://www.facebook.com/pages/Startup-Live/441836922496996"
	STARTeuropeTwitterURL  = "https://twitter.com/startuplive"
	STARTeuropeLinkedInURL = "http://linkedin.com/company/starteurope/"
	STARTeuropeYoutubeURL  = "http://youtube.com/starteurope"
	STARTeuropeFlickrURL   = "http://flickr.com/starteurope"

	GoogleAnalyticsID = "UA-26690424-1"
	GoogleMapsApiKey  = "AIzaSyDyIKmzKMhF8d2Yk4tI6Ht9KjZU4y85wQ4"

	HackernewsCode = `(function(d, t) {
    var g = d.createElement(t),
        s = d.getElementsByTagName(t)[0];
    g.src = '//hnbutton.appspot.com/static/hn.js';
    s.parentNode.insertBefore(g, s);
}(document, 'script'));`

	PleskUrl  = "https://host.internetkultur.at:8443/enterprise/control/agent.php"
	PleskUser = "starteurope"
	PleskPW   = "Vhiluoomhag2"

	AmiandoAdminEmail = "amiando@startuplive.in"

	RichTextToolbar = ``
)

const (
	DefaultPrimaryColorIndex   = 8
	DefaultSecondaryColorIndex = 6
	DefaultPatternIndex        = 0
)

// Pseudo constants:
var (
	DefaultColorScheme = models.NewColorScheme(DefaultPrimaryColorIndex, DefaultSecondaryColorIndex, DefaultPatternIndex)
	DefaultCSSContext  = TemplateContext(DefaultColorScheme)
)
