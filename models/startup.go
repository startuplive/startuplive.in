package models

import (
	"github.com/ungerik/go-start/media"
	"github.com/ungerik/go-start/model"

	"github.com/ungerik/go-start/mongo"

	// "github.com/ungerik/go-start/view"
)

// Save Revisions?

// Generate Milestones for achievements/events like:
// Funding, Mentions on Top Blogs and Printing Press, TV, 1st Employee, C Level Employees, 10, 50, 100 Employees
var Startups = mongo.NewCollection("startups", "Name")

func NewStartup() *Startup {
	var doc Startup
	Startups.InitDocument(&doc)
	return &doc
}

type Startup struct {
	mongo.DocumentBase `bson:",inline"`
	Founder            mongo.Ref    `model:"to=people"`
	Name               model.String `view:"label=The name of your startup|size=25" model:"required"`
	Logo               media.ImageRef
	BizCategory        model.Choice `view:"label=Your startups business category" model:"required|options=Aerospace and Defense,Arts & Entertainment,Automotive,Biotechnology and Pharmaceuticals,Business & Professional Services,Chemicals,Clothing & Accessoires,Community & Government,Construction & Contractors,Consumer Electronics,Education,Energy,Food & Dining,Health & Medicine,Home & Garden,Industry & Agriculture,Legal & Financial,Media & Communications,Personal Care & Services,Real Estate,Shopping,Sports & Recreation,Travel & Transportation"`
	Tagline            model.String
	Abstract           model.Text   `view:"label=Short and crisp description of what your startup is doing" model:"required"`
	Website            model.Url    `view:"size=25|label=Your startups website" model:"required"`
	TeamMember         []TeamMember `view:"label=Team Members"`
	// Founded            model.String `view:"label=When was your startup founded?" model:"required"`
	FoundingYear model.Choice `view:"label=When was your startup founded?" model:"required|options=2003,2004,2005,2006,2007,2008,2009,2010,2011,2012"`
	Stage        model.Choice `view:"label=What stage is your startup in?" model:"required|options=Concept Stage (got an idea),Seed Stage (Working on product),Early Stage (Close to market),Growth Stage (we're out there and making some cash),Sustainable Business (we already made it to a sustainable business)"`
	Located      model.String `view:"size=25|label=Where is your startup located?" model:"required"`
	// Offices            []modelext.PostalAddress
	PressArticles []PressArticles `view:"label=Any Press Mentions?"`
	FundingAmount model.Choice    `view:"hidden|Did you got funding - How much?" model:"options=,0,1-25k,26-75k,76-125k,126-175k,176-250k,251-325k,326-500k,500k+"`
	Financing     model.Choice    `model:"required|options=bootstrapping,investors,grants,investors and grants"`
	Feedback      model.Text      `view:"label=What is your feedback on the Startup Live Event"`
	LiveHelped    model.Text      `view:"label=In what areas did Startup Live helped you and how?"`
	Testimonial   []model.Text    `view:"label=Write a short testimonial"`

	CreationDate model.Date `view:"hidden"`
}

type TeamMember struct {
	FirstName  model.String `view:"label=First Name|size=20" model:"required"`
	LastName   model.String `view:"label=Last Name|size=20" model:"required"`
	Background model.Choice `view:"label=Background|size=20" model:"options=Business, Design, Technical"`
}
type PressArticles struct {
	Source model.String `view:"label=Name of the site|size=20" model:"required"`
	URL    model.Url    `view:"label=URL|size=20" model:"required"`
}

func (self *Startup) String() string {
	return self.Name.String()
}
