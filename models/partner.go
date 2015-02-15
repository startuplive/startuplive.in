package models

import (
	"github.com/ungerik/go-start/media"
	"github.com/ungerik/go-start/model"
	// "github.com/ungerik/go-start/modelext"
	"github.com/ungerik/go-start/mongo"
	// "github.com/ungerik/go-start/user"
	// "github.com/ungerik/go-start/view"
)

var Partners = mongo.NewCollection("partners", "Name")

func NewPartner() *Partner {
	var doc Partner
	Partners.InitDocument(&doc)
	return &doc
}

const (
	PARTNERS      = "Partners"
	SUPPORTERS    = "Supporters"
	MEDIAPARTNERS = "Media Partners"
)

type Partner struct {
	mongo.DocumentBase `bson:",inline"`
	Name               model.String   `model:"required"`
	Website            model.String   `model:"required"`
	Logo               media.ImageRef `model:"required"`
	Events             []mongo.Ref    `model:"required|to=events"`
}

func (self *Partner) String() string {
	return self.Name.String()
}
