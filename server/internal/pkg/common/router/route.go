package router

import "net/http"

type Method string

type Route struct {
	Name        string
	Method      Method
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route
