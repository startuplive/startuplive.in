package root

import (
	. "github.com/STARTeurope/startuplive.in/views"
	. "github.com/ungerik/go-start/view"
)

func init() {
	Imprint = NewPublicPage("Imprint | Startup Live",
		DIV("public-content",
			DIV("main",
				TitleBar("Imprint"),
				DIV("main-content",
					&Table{
						Class: "imprint",
						Model: EscapeStringsTableModel{
							{"Company Name:", "JFDI GmbH"},
							{"Managing Directors:", "Jürgen Furian, Andreas Tschas"},
							{"Commercial Register:", "FN 375833 x"},
							{"VAT:", "ATU67137779"},
							{"Bank Identifier Code:", "GIBAATWWXXX (Erste Bank)"},
							{"IBAN:", "AT892011182013028500"},
						},
					},
				),
			),
			DIV("main",
				TitleBar("Address"),
				DIV("main-content",
					P(
						HTML("STARTeurope<br/>"),
						HTML("Getreidemarkt 11/17<br/>"),
						HTML("1060 Vienna<br/>"),
						HTML("Austria<br/>"),
					),
				),
			),
		),
	)
}
