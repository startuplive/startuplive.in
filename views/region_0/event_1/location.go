package event_1

import (
	. "github.com/STARTeurope/startuplive.in/views"
	. "github.com/ungerik/go-start/view"
)

//var googleMaps PageWriteFunc = GoogleMaps(GoogleMapsApiKey, false, "")

func init() {
	Region0_Event1_Location = newPublicEventPage(EventTitle("Location"), nil, DynamicView(
		func(ctx *Context) (view View, err error) {
			location := ctx.Data.(*PageData).Location

			locationName := location.Name.GetOrDefault("[Enter event location name!]")
			description := location.Description.Get()
			publicTransport := location.PublicTransport.Get()
			privateCar := location.PrivateCar.Get()
			address := location.Address.String()

			var images View

			if !location.HasHeaderImages() {
				images = GoogleMapsIframe(604, 300, address)
			} else {
				top, left, right, err := location.GetHeaderImages()
				if err != nil {
					return nil, err
				}
				images = Views{
					DIV("", top, left, right),
					GoogleMapsIframe(300, 300, address),
				}
			}

			view = &Div{
				Class: "main location",
				Content: Views{
					TitleBar("About the venue"),
					DIV("location-image-box", images),
					H2(locationName),
					H4(DIV("location-icon"), address),
					P(HTML(description)),
					&If{
						Condition: publicTransport != "",
						Content: Views{
							H3("Getting there by public transport"),
							P(HTML(publicTransport)),
						},
					},
					&If{
						Condition: privateCar != "",
						Content: Views{
							H3("Getting there by car"),
							P(HTML(privateCar)),
						},
					},
				},
			}
			return view, nil
		},
	))
}
