package model

import "net/http"

// Route defines properties of the HTTP route
type Route struct {
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
