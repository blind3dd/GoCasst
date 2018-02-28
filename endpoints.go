package main

import (
	"net/http"

	"GoCasst/Messaging"
	"GoCasst/Users"
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
		"Users.GetAllUsersRequest",
		"GET",
		"/v1/users",
		Users.GetAllUsersRequest,
	},
	Route{
		"Messages.GetAllMessagesRequest",
		"GET",
		"/v1/messages",
		Messaging.GetAllMessagesRequest,
	},
	Route{
		"Users.GetOneUserRequest",
		"GET",
		"/v1/users/{users_uuid}",
		Users.GetOneUserRequest,
	},
	Route{
		"Messages.GetOneMsgRequest",
		"GET",
		"/v1/messages/{message_uuid}",
		Messaging.GetOneMsgRequest,
	},
	Route{
		"Users.PostTheUser",
		"POST",
		"/v1/users/new",
		Users.PostTheUser,
	},
	Route{
		"Messaging.PostTheMessage",
		"POST",
		"/v1/messages/new",
		Messaging.PostTheMessage,
	},
}
