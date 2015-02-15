package models

import (
	"github.com/ungerik/go-start/model"
)

///////////////////////////////////////////////////////////////////////////////
// ColorScheme

func NewColorScheme(primaryIndex, secondaryIndex, pattern int) *ColorScheme {
	scheme := new(ColorScheme)

	primaryColor := Colors[primaryIndex]
	secondaryColor := Colors[secondaryIndex]

	scheme.Primary.Set(primaryColor)
	scheme.Secondary.Set(secondaryColor)
	scheme.Pattern.SetInt(pattern)

	primaryButtonColorSchemes := ButtonColorSchemes[primaryColor]
	secondaryButtonColorSchemes := ButtonColorSchemes[secondaryColor]

	scheme.ButtonColorScheme = primaryButtonColorSchemes
	scheme.SecondaryButton = secondaryButtonColorSchemes

	scheme.TimetablePrimary = TimetableColorScheme{
		Border:      primaryButtonColorSchemes.Button.Border,
		InnerShadow: primaryButtonColorSchemes.Button.InnerShadow,
		Background:  primaryButtonColorSchemes.Button.Top,
	}
	scheme.TimetableSecondary = TimetableColorScheme{
		Border:      secondaryButtonColorSchemes.Button.Border,
		InnerShadow: secondaryButtonColorSchemes.Button.InnerShadow,
		Background:  secondaryButtonColorSchemes.Button.Top,
	}

	return scheme
}

type ColorScheme struct {
	Primary   model.String
	Secondary model.String
	Pattern   model.Int
	ButtonColorScheme
	SecondaryButton    ButtonColorScheme
	TimetablePrimary   TimetableColorScheme
	TimetableSecondary TimetableColorScheme
}

type ButtonColors struct {
	Border      model.String
	InnerShadow model.String
	Top         model.String
	Bottom      model.String
}

type ButtonColorScheme struct {
	Button         ButtonColors
	HoverButton    ButtonColors
	PressedButton  ButtonColors
	DisabledButton ButtonColors
}

type TimetableColorScheme struct {
	Border      model.String
	InnerShadow model.String
	Background  model.String
}

///////////////////////////////////////////////////////////////////////////////
// Color Values

var Colors = [...]string{
	"#fcc329",
	"#f37822",
	"#d01f58",
	"#ef61a3",
	"#9170aa",
	"#1b9cd7",
	"#8cd7f7",
	"#6dc7be",
	"#98c93c",
	"#c1c2c4",
}

var ButtonColorSchemes = map[string]ButtonColorScheme{
	"#fcc329": {
		Button: ButtonColors{
			Border:      "#b08e31",
			InnerShadow: "#fddc81",
			Top:         "#fcc93e",
			Bottom:      "#e4b026",
		},
		HoverButton: ButtonColors{
			Border:      "#b08e31",
			InnerShadow: "#eccf7d",
			Top:         "#e2b538",
			Bottom:      "#cd9e22",
		},
		PressedButton: ButtonColors{
			Border:      "#91721c",
			InnerShadow: "#9c7d27",
			Top:         "#d6ab35",
			Bottom:      "#c29620",
		},
		DisabledButton: ButtonColors{
			Border:      "#8e8e8e",
			InnerShadow: "#dcdcdc",
			Top:         "#c9c9c9",
			Bottom:      "#b0b0b0",
		},
	},
	"#f37822": {
		Button: ButtonColors{
			Border:      "#b06631",
			InnerShadow: "#f8b07d",
			Top:         "#f48538",
			Bottom:      "#db6d20",
		},
		HoverButton: ButtonColors{
			Border:      "#b06631",
			InnerShadow: "#e8a67a",
			Top:         "#db7732",
			Bottom:      "#c5621d",
		},
		PressedButton: ButtonColors{
			Border:      "#914d1c",
			InnerShadow: "#985223",
			Top:         "#d07130",
			Bottom:      "#ba5d1b",
		},
		DisabledButton: ButtonColors{
			Border:      "#767676",
			InnerShadow: "#c0c0c0",
			Top:         "#9e9e9e",
			Bottom:      "#868686",
		},
	},
	"#d01f58": {
		Button: ButtonColors{
			Border:      "#b03159",
			InnerShadow: "#e37c9d",
			Top:         "#d43569",
			Bottom:      "#bc1d50",
		},
		HoverButton: ButtonColors{
			Border:      "#b03159",
			InnerShadow: "#d57896",
			Top:         "#be305e",
			Bottom:      "#a91a48",
		},
		PressedButton: ButtonColors{
			Border:      "#911c41",
			InnerShadow: "#832141",
			Top:         "#b42d59",
			Bottom:      "#a01944",
		},
		DisabledButton: ButtonColors{
			Border:      "#5b5b5b",
			InnerShadow: "#9f9f9f",
			Top:         "#6a6a6a",
			Bottom:      "#525252",
		},
	},
	"#ef61a3": {
		Button: ButtonColors{
			Border:      "#b0316c",
			InnerShadow: "#f5a3c9",
			Top:         "#f071ac",
			Bottom:      "#d85894",
		},
		HoverButton: ButtonColors{
			Border:      "#b0316c",
			InnerShadow: "#e69bbd",
			Top:         "#d8659a",
			Bottom:      "#c24f85",
		},
		PressedButton: ButtonColors{
			Border:      "#921b53",
			InnerShadow: "#95466a",
			Top:         "#cc6092",
			Bottom:      "#b84b7e",
		},
		DisabledButton: ButtonColors{
			Border:      "#5e5e5e",
			InnerShadow: "#c0c0c0",
			Top:         "#9e9e9e",
			Bottom:      "#858585",
		},
	},
	"#9170aa": {
		Button: ButtonColors{
			Border:      "#7931b0",
			InnerShadow: "#bfabcd",
			Top:         "#9c7eb2",
			Bottom:      "#83669a",
		},
		HoverButton: ButtonColors{
			Border:      "#7931b0",
			InnerShadow: "#b4a3c1",
			Top:         "#8c71a0",
			Bottom:      "#765c8a",
		},
		PressedButton: ButtonColors{
			Border:      "#5e1c91",
			InnerShadow: "#614e6e",
			Top:         "#856b97",
			Bottom:      "#6f5783",
		},
		DisabledButton: ButtonColors{
			Border:      "#555555",
			InnerShadow: "#b5b5b5",
			Top:         "#8d8d8d",
			Bottom:      "#747474",
		},
	},
	"#1b9cd7": {
		Button: ButtonColors{
			Border:      "#3188b0",
			InnerShadow: "#7ac5e8",
			Top:         "#32a6db",
			Bottom:      "#198dc2",
		},
		HoverButton: ButtonColors{
			Border:      "#3188b0",
			InnerShadow: "#76bad9",
			Top:         "#2d95c5",
			Bottom:      "#167fae",
		},
		PressedButton: ButtonColors{
			Border:      "#1c6c91",
			InnerShadow: "#1f6788",
			Top:         "#2b8dba",
			Bottom:      "#1578a5",
		},
		DisabledButton: ButtonColors{
			Border:      "#727272",
			InnerShadow: "#b2b2b2",
			Top:         "#898989",
			Bottom:      "#707070",
		},
	},
	"#8cd7f7": {
		Button: ButtonColors{
			Border:      "#629cb5",
			InnerShadow: "#bbe8fa",
			Top:         "#97dbf7",
			Bottom:      "#7fc2df",
		},
		HoverButton: ButtonColors{
			Border:      "#629cb5",
			InnerShadow: "#b2d9ea",
			Top:         "#88c5de",
			Bottom:      "#72aec8",
		},
		PressedButton: ButtonColors{
			Border:      "#629cb5",
			InnerShadow: "#5d8899",
			Top:         "#80bad2",
			Bottom:      "#6ca5be",
		},
		DisabledButton: ButtonColors{
			Border:      "#8d8d8d",
			InnerShadow: "#dcdcdc",
			Top:         "#cacaca",
			Bottom:      "#b1b1b1",
		},
	},
	"#6dc7be": {
		Button: ButtonColors{
			Border:      "#31b0a3",
			InnerShadow: "#a9ded9",
			Top:         "#7bccc4",
			Bottom:      "#63b4ac",
		},
		HoverButton: ButtonColors{
			Border:      "#31b0a3",
			InnerShadow: "#a1d0cc",
			Top:         "#6eb7b0",
			Bottom:      "#59a29a",
		},
		PressedButton: ButtonColors{
			Border:      "#1c9186",
			InnerShadow: "#4d7f7a",
			Top:         "#69aea7",
			Bottom:      "#549992",
		},
		DisabledButton: ButtonColors{
			Border:      "#888888",
			InnerShadow: "#cecece",
			Top:         "#b3b3b3",
			Bottom:      "#9b9b9b",
		},
	},
	"#98c93c": {
		Button: ButtonColors{
			Border:      "#84b031",
			InnerShadow: "#c2df8c",
			Top:         "#a2ce4f",
			Bottom:      "#8ab637",
		},
		HoverButton: ButtonColors{
			Border:      "#84b031",
			InnerShadow: "#b7d187",
			Top:         "#91b947",
			Bottom:      "#7ca331",
		},
		PressedButton: ButtonColors{
			Border:      "#68911c",
			InnerShadow: "#658031",
			Top:         "#8aaf43",
			Bottom:      "#759b2f",
		},
		DisabledButton: ButtonColors{
			Border:      "#959595",
			InnerShadow: "#cdcdcd",
			Top:         "#b3b3b3",
			Bottom:      "#9b9b9b",
		},
	},
	"#c1c2c4": {
		Button: ButtonColors{
			Border:      "#959595",
			InnerShadow: "#dbdbdc",
			Top:         "#c7c8ca",
			Bottom:      "#afafb1",
		},
		HoverButton: ButtonColors{
			Border:      "#959595",
			InnerShadow: "#cececf",
			Top:         "#b3b4b5",
			Bottom:      "#9d9d9f",
		},
		PressedButton: ButtonColors{
			Border:      "#959595",
			InnerShadow: "#7b7c7d",
			Top:         "#a9aaac",
			Bottom:      "#959597",
		},
		DisabledButton: ButtonColors{
			Border:      "#b5b5b5",
			InnerShadow: "#e6e6e6",
			Top:         "#d8d8d8",
			Bottom:      "#c7c7c7",
		},
	},
}
