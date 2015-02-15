package api

import (
	"strings"
	"github.com/STARTeurope/startuplive.in/models"
	. "github.com/STARTeurope/startuplive.in/views"
)

func init() {
	API_Mentors = SearchPeopleJSONView(
		func(person *models.Person, searchTerm string) (ok bool) {
			if !person.Mentor {
				return false
			}
			name := strings.ToLower(person.Name.String())
			return person.Blocked == false && name != "admin" && (searchTerm == "" || strings.Contains(name, searchTerm))
		},
	)
}