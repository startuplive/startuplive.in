package root

import (
	. "github.com/STARTeurope/startuplive.in/views"
	. "github.com/ungerik/go-start/view"
)

func init() {
	Contact = NewPublicPage("Contact | Startup Live",
		DIV("public-content",
			DIV("row",
				DIV("cell right-border"), /*TitleBar("Contact us"),
				DIV("main-content",
					P("If you have any questions about the Startup Live events or how to participate, please don’t hesitate to contact us. We’d love to hear from you!"),
					NewContactForm(ContactEmail, "StartupLive Contact ", "contact-form", "button", "contactform"),
				),*/

				DIV("cell right",
					TitleBarRight("Postal Address"),
					DIV("main-content",
						GoogleMapsIframe(425, 425, "Praterstraße 62-64 A-1020 Vienna, Austria"),
						P(
							HTML("Startup Live<br/>"),
							HTML("Praterstraße 62-64 A-1020<br/>"),
							HTML("1060 Vienna<br/>"),
							HTML("Austria<br/>"),
						),
					),
				),
				DivClearBoth(),
			),
		),
	)
}
