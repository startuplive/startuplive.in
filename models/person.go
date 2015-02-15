package models

import (
	"bytes"

	"github.com/ungerik/go-gravatar"
	"github.com/ungerik/go-start/config"
	"github.com/ungerik/go-start/media"
	"github.com/ungerik/go-start/model"
	"github.com/ungerik/go-start/user"
	"github.com/ungerik/go-start/utils"
	"github.com/ungerik/go-start/view"
)

// We are using user.NewCollection here instead of mongo.NewCollection
// because user.NewCollection sets the correct mongo.Collection.DocLabelSelectors
// so that mongo.Collection.DocumentLabel(id) returns a label for
// the document with id composed of the name modelext.Name components
// Prefix + First + Middle + Last + Postfix + Organization.
var People = user.NewCollection("people")

func init() {
	People.DocLabelSeparator = " "
}

func NewPerson() *Person {
	var doc Person
	People.InitDocument(&doc)
	return &doc
}

type PersonImages struct {
	URL_50x50   model.Url
	URL_100x100 model.Url
	URL_160x160 model.Url
	URL_320x320 model.Url
	URL_284x144 model.Url
}

func (self *PersonImages) SetAllEmptyURLs(url string) {
	if self.URL_50x50.IsEmpty() {
		self.URL_50x50.Set(url)
	}
	if self.URL_100x100.IsEmpty() {
		self.URL_100x100.Set(url)
	}
	if self.URL_160x160.IsEmpty() {
		self.URL_160x160.Set(url)
	}
	if self.URL_320x320.IsEmpty() {
		self.URL_320x320.Set(url)
	}
	if self.URL_284x144.IsEmpty() {
		self.URL_284x144.Set(url)
	}
}

///////////////////////////////////////////////////////////////////////////////
// Person

type Person struct {
	user.User `bson:",inline"`

	Image  media.ImageRef
	Images PersonImages `view:"label=Old Image"`

	SuperAdmin  model.Bool
	Gender      model.Choice `model:"options=Male,Female"`
	University  model.String
	BirthDate   model.Date
	BirthYear   model.Int
	Citizenship model.Country
	Languages   []model.Language // first is native language

	EventOrganiser           model.Bool
	OrganiserInfo            model.Text
	OrganiserEmail           model.Email `view:"placeholder=firstname.lastname@startuplive.in" model:"label=Your startuplive.in organiser email - type: firstname.lastname@startuplive.in"`
	OrganiserForwardingEmail model.Email `view:"placeholder=the private email address" model:"label=Organiser forwarding email - the private email address"`

	Judge     model.Bool
	JudgeInfo model.Text

	Mentor     model.Bool
	MentorInfo model.Text

	FeaturedMentor     model.Bool
	FeaturedMentorInfo model.Text

	Company  model.String // todo replace with ref to organisations
	Position model.String // todo replace
	Tags     model.String // todo replace with array of strings

	TermsAndConditions model.Bool `view:"auth=admin|label=Yes, I accept the <a href='starteurope.at/termsandconditions/' target='_blank'>Terms and Conditions</a>" model:"required"`
	TaCDate            model.Date `view:"auth=admin"`
	/*
		{"CV", &Array{
			Of: &Document{
				Nodes: Nodes{
					// external without starteurope event roles
					{"Organisation", &Ref{To: &Organisation}},
					{"Joined", &Date{}},
					{"Left", &Date{}},
					{"Role", &Str{}},                          // todo defined roles
					{"Recommendations", &Array{Of: &Str{}}}, // todo ?
				},
			},
		}},
		{"Tags", &Array{Of: &Ref{To: &Tag}}},
	*/

}

func (self *Person) String() string {
	return self.User.Name.String()
}

// "Position at Company"
func (self *Person) JobDescription() string {
	var buf bytes.Buffer
	buf.WriteString(self.Position.Get())
	if self.Position != "" && self.Company != "" {
		buf.WriteString(" at ")
	}
	buf.WriteString(self.Company.Get())
	return buf.String()
}

func (self *Person) WasMentor() (count int, err error) {
	return Events.Filter("Mentors", self.ID).Count()
}

func (self *Person) WasJudge() (count int, err error) {
	return Events.Filter("Judges", self.ID).Count()
}

func (self *Person) GetOrganisedEvents() model.Iterator {
	return Events.Filter("Organisers", self.ID).Sort("From").Iterator()
}

func (self *Person) findAlternativeProfileURL(size int, defaultURL string) string {
	if facebook := self.PrimaryFacebookIdentity(); facebook != nil {
		return facebook.ProfileImageURL()
	}
	if twitter := self.PrimaryTwitterIdentity(); twitter != nil {
		return twitter.ProfileImageURL()
	}
	if defaultURL == "" {
		defaultURL = gravatar.MysteryMan
	}
	return gravatar.UrlSizeDefault(self.PrimaryEmail(), size, defaultURL)
}

func (self *Person) Image_50x50() (*view.Image, error) {
	if self.Image.IsEmpty() {
		url := self.Images.URL_50x50.Get()
		if url == "" {
			url = self.findAlternativeProfileURL(50, "http://startuplive.in/images/avatar50x50.jpg")
		}
		return view.IMG(url, 50, 50), nil
	}
	version, err := self.Image.VersionCentered(50, 50, false)
	if err != nil {
		config.Logger.Println("Propably an invalid image ref: ", err)
		self.Image.Set(nil)
		return self.Image_50x50()
	}
	return version.View(""), nil
}

func (self *Person) Image_100x100() (*view.Image, error) {
	if self.Image.IsEmpty() {
		url := self.Images.URL_100x100.Get()
		if url == "" {
			url = self.findAlternativeProfileURL(100, "http://startuplive.in/images/avatar100x100.jpg")
		}
		return view.IMG(url, 100, 100), nil
	}
	version, err := self.Image.VersionCentered(100, 100, false)
	if err != nil {
		config.Logger.Println("Propably an invalid image ref: ", err)
		self.Image.Set(nil)
		return self.Image_100x100()
	}
	return version.View(""), nil
}

func (self *Person) Image_160x160(class string) (*view.Image, error) {
	if self.Image.IsEmpty() {
		url := self.Images.URL_160x160.Get()
		if url == "" {
			url = self.findAlternativeProfileURL(160, "http://startuplive.in/images/avatar160x160.jpg")
		}
		return &view.Image{
			Class:  class,
			Src:    url,
			Width:  160,
			Height: 160,
		}, nil
	}
	version, err := self.Image.VersionCentered(160, 160, false)
	if err != nil {
		config.Logger.Println("Propably an invalid image ref: ", err)
		self.Image.Set(nil)
		return self.Image_160x160(class)
	}
	return version.View(class), nil
}

func (self *Person) HasImage_160x160_and_284x144() bool {
	return self.Image != "" || // self.Image provides all sizes
		(self.Images.URL_160x160 != "" && self.Images.URL_284x144 != "")
}

func (self *Person) Image_320x320(class string) (*view.Image, error) {
	if self.Image.IsEmpty() {
		url := self.Images.URL_320x320.Get()
		if url == "" {
			url = self.findAlternativeProfileURL(320, "http://startuplive.in/images/avatar320x320.jpg")
		}
		return &view.Image{
			Class:  class,
			Src:    url,
			Width:  320,
			Height: 320,
		}, nil
	}
	version, err := self.Image.VersionCentered(320, 320, false)
	if err != nil {
		config.Logger.Println("Propably an invalid image ref: ", err)
		self.Image.Set(nil)
		return self.Image_320x320(class)
	}
	return version.View(class), nil
}

func (self *Person) Image_284x144(class string) (*view.Image, error) {
	if self.Image.IsEmpty() {
		url := self.Images.URL_320x320.Get()
		if url == "" {
			url = self.findAlternativeProfileURL(284, "http://startuplive.in/images/avatar284x144.jpg")
		}
		return &view.Image{
			Class:  class,
			Src:    self.Images.URL_284x144.GetOrDefault("http://startuplive.in/images/avatar284x144.jpg"),
			Width:  284,
			Height: 144,
		}, nil
	}
	version, err := self.Image.VersionCentered(284, 144, false)
	if err != nil {
		config.Logger.Println("Propably an invalid image ref: ", err)
		self.Image.Set(nil)
		return self.Image_284x144(class)
	}
	return version.View(class), nil
}

///////////////////////////////////////////////////////////////////////////////
// Startup 

func (self *Person) GetStartups() model.Iterator {
	return Startups.Filter("Founder", self.ID).SortFunc(
		func(a, b *Startup) bool {
			return utils.CompareCaseInsensitive(a.Name.Get(), b.Name.Get())
		},
	)
}
