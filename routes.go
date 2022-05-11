package main

import (
)

type Routes []Route

var routes = Routes{
	Route{
		"Test",
		"GET",
		"/dgate/v1/test",
		Test,
	},
}