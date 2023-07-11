package decision

import (
	"github.com/mikhailbolshakov/decision/http"
)

func GetRoutes(c Controller) []*http.Route {
	return []*http.Route{
		// non authorize zone
		http.R("/guests/decisions", c.MakeDecisionGuest).POST(),

		// authorized zone
		http.R("/users/{userId}/decisions", c.MakeDecision).POST(),
	}
}
