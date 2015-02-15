package api

import (
	"strings"
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
)

func init() {
	API_Organisers = SearchPeopleJSONView(
		func(person *models.Person, searchTerm string) (ok bool) {
			if !person.EventOrganiser {
				return false
			}
			name := strings.ToLower(person.Name.String())
			return person.Blocked == false && name != "admin" && (searchTerm == "" || strings.Contains(name, searchTerm))
		},
	)
}