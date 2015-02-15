package event_1

import (
	// "github.com/STARTeurope/startuplive.in/models"

	. "github.com/STARTeurope/startuplive.in/views"
	. "github.com/ungerik/go-start/view"
	// "io"
	// "github.com/ungerik/go-start/debug"

)

func init() {
	Region0_Event1_Registration_Success = newPublicEventPage(
		EventTitle("Success"),
		nil,
		HTML("<b>Thank you!</b><br><br>We successfully received your information.<br><br> Like us on <a href='https://www.facebook.com/pages/Startup-Live/441836922496996'>Facebook</a>!."),
	)
}
