package main

type Routes []Route

var routes = Routes{
	Route{
		"Test",
		"GET",
		"/deku-gate/v1/test",
		Test,
	},
}
