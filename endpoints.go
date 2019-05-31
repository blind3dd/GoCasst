package main

import (
	"net/http"

	"github.com/blind3dd/gocasst/messaging"
	"github.com/blind3dd/gocasst/users"
)

type ResponseStatus struct {
	Status string `json:"status"`
	Code   int    `json:'code'`
}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Heartbeat",
		"GET",
		"/",
		Heartbeat,
	},
	Route{
		"users.GetAllUsersRequest",
		"GET",
		"/v1/users",
		users.GetAllUsersRequest,
	},
	Route{
		"messages.GetAllMessagesRequest",
		"GET",
		"/v1/messages",
		messaging.GetAllMessagesRequest,
	},
	Route{
		"users.GetOneUserRequest",
		"GET",
		"/v1/users/{users_uuid}",
		users.GetOneUserRequest,
	},
	Route{
		"messages.GetOneMsgRequest",
		"GET",
		"/v1/messages/{message_uuid}",
		messaging.GetOneMsgRequest,
	},
	Route{
		"users.PostTheUser",
		"POST",
		"/v1/users/new",
		users.PostTheUser,
	},
	Route{
		"messages.PostTheMessage",
		"POST",
		"/v1/messages/new",
		messaging.PostTheMessage,
	},
}
