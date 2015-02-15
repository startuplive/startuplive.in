package models

import (
	"image/color"

	"github.com/ungerik/go-start/media"
	"github.com/ungerik/go-start/model"
	"github.com/ungerik/go-start/modelext"
	"github.com/ungerik/go-start/mongo"
	"github.com/ungerik/go-start/view"
)

var EventLocations = mongo.NewCollection("eventlocations", "Name")

func NewEventLocation() *EventLocation {
	var doc EventLocation
	EventLocations.InitDocument(&doc)
	return &doc
}

type EventLocation struct {
	mongo.DocumentBase `bson:",inline"`
	Name               model.String `model:"required"`
	ShortName          model.String
	Description        model.Text `model:"required"`
	PublicTransport    model.Text
	PrivateCar         model.Text
	Address            modelext.PostalAddress
	GeoLocation        model.GeoLocation

	Header_Image_Top_300x200  media.ImageRef
	Header_Image_Left_149x98  media.ImageRef
	Header_Image_Right_149x98 media.ImageRef

	Image struct {
		Top_URL_300x200  model.Url `view:"label=Old Header Image Top 300x200"`
		Left_URL_149x98  model.Url `view:"label=Old Header Image Left 149x98"`
		Right_URL_149x98 model.Url `view:"label=Old Header Image Right 149x98"`
	}
}

func (self *EventLocation) HasHeaderImages() bool {
	return self.Header_Image_Top_300x200.IsEmpty() == false ||
		self.Header_Image_Left_149x98.IsEmpty() == false ||
		self.Header_Image_Right_149x98.IsEmpty() == false ||
		self.Image.Top_URL_300x200.IsEmpty() == false ||
		self.Image.Left_URL_149x98.IsEmpty() == false ||
		self.Image.Right_URL_149x98.IsEmpty() == false
}

func (self *EventLocation) GetHeaderImages() (top, left, right *view.Image, err error) {
	if self.Header_Image_Top_300x200.IsEmpty() {
		if self.Image.Top_URL_300x200.IsEmpty() {
			top = view.IMG(media.ColoredImageDataURL(color.White), 300, 200)
		} else {
			top = view.IMG(self.Image.Top_URL_300x200.Get(), 300, 200)
		}
	} else {
		top, err = self.Header_Image_Top_300x200.VersionCenteredView(300, 200, false, "")
		if err != nil {
			return nil, nil, nil, err
		}
	}

	if self.Header_Image_Left_149x98.IsEmpty() {
		if self.Image.Left_URL_149x98.IsEmpty() {
			left = view.IMG(media.ColoredImageDataURL(color.White), 149, 98)
		} else {
			left = view.IMG(self.Image.Left_URL_149x98.Get(), 149, 98)
		}
	} else {
		left, err = self.Header_Image_Left_149x98.VersionCenteredView(149, 98, false, "")
		if err != nil {
			return nil, nil, nil, err
		}
	}

	if self.Header_Image_Right_149x98.IsEmpty() {
		if self.Image.Right_URL_149x98.IsEmpty() {
			right = view.IMG(media.ColoredImageDataURL(color.White), 149, 98)
		} else {
			right = view.IMG(self.Image.Right_URL_149x98.Get(), 149, 98)
		}
	} else {
		right, err = self.Header_Image_Right_149x98.VersionCenteredView(149, 98, false, "")
		if err != nil {
			return nil, nil, nil, err
		}
	}

	return top, left, right, nil
}

func (self *EventLocation) String() string {
	return self.Name.Get()
}

func (self *EventLocation) GetRequiredFields() []string {
	return []string{
		"Name",
		"ShortName",
		"Description",
		"PublicTransport",
		"PrivateCar",
		"Address.City",
		"Address.State",
		"Address.FirstLine",
		"Address.ZIP",
	}
}
