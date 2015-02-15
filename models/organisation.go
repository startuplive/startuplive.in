package models

import (
	"github.com/ungerik/go-start/media"
	"github.com/ungerik/go-start/model"
	"github.com/ungerik/go-start/modelext"
	"github.com/ungerik/go-start/mongo"
	"github.com/ungerik/go-start/user"
	// "github.com/ungerik/go-start/view"
)

// Save Revisions?

// Generate Milestones for achievements/events like:
// Funding, Mentions on Top Blogs and Printing Press, TV, 1st Employee, C Level Employees, 10, 50, 100 Employees
var Organisations = mongo.NewCollection("organisations", "Name.Organization")

func NewOrganisation() *Organisation {
	var doc Organisation
	Organisations.InitDocument(&doc)
	return &doc
}

type Organisation struct {
	mongo.DocumentBase `bson:",inline"`
	Type               model.Choice `model:"options=School,University,Non Profit,Club,Event Orga,Startup,Company,Big Enterprise,VC Fund,Angel Network,Incubator,Seed Accelerator,Online News,Print News,Television,Business Consultancy,Ad Agency,Legal Agency"`
	Founded            model.String
	User               user.User
	LegalName          model.String
	Logo               media.ImageRef
	Tagline            model.String
	Description        model.Text
	BusinessAreas      model.Choice `view:"label=Your startups business category" model:"required|options=Aerospace and Defense,Arts & Entertainment,Automotive,Biotechnology and Pharmaceuticals,Business & Professional Services,Chemicals,Clothing & Accessoires,Community & Government,Construction & Contractors,Consumer Electronics,Education,Energy,Food & Dining,Health & Medicine,Home & Garden,Industry & Agriculture,Legal & Financial,Media & Communications,Personal Care & Services,Real Estate,Shopping,Sports & Recreation,Travel & Transportation"`
	NumEmployees       model.Int
	Images             []model.Url
	Videos             []model.Url
	Mentions           []model.Url
	Parent             mongo.Ref `model:"to=organisations"`
	Products           []model.String
	Offices            []modelext.PostalAddress
	Status             model.Choice `model:"options=Founding,Active,Idle,Dead,Acquired"`
	Funding            []struct {
		Date                model.Date
		Round               model.Choice `model:"options=Seed,Angel,A,B,C,D,E,F,Grant,Dept,Unattributed"`
		Amount              model.Int
		Currency            model.Choice `model:"options=EUR,USD,GBP"`
		SourceOrganizations []mongo.Ref  `model:"to=organisations"`
		SourcePeople        []mongo.Ref  `model:"to=people"`
		Mentions            []model.Url
	}
}

func (self *Organisation) String() string {
	return self.User.Name.String()
}
