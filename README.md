# GoCaast REST api (users)

This Application reads and write to Cassandra by gocql library.
this is api for reading/writing to Cassandra database.

- Api listens on port 8088 and localhost. (see /main.go)
- Basic Auth is implemented (-u or authorization: basic header - See Auth/basicAuth.go)
- There is a router based on gorilla/mux and one router.go file with all the endpoints defined.
- GET and POST methods are supported as for now
- things are logged with unix timestamps (every REST request)
- localhost:8088 endpoint is for HearBeat status check.

prerequisites:

	- Cassandra Node has to be created in the system
	- single Cassandra Keyspace needs to be created 
	(in Cass/main.go it's looking for Cluster name "usersmessages")


EXAMPLE USAGE (under assumption basic auth is sent along):
- hearbeat check
- getting user's list from Cluster
- Creating new one by http POST request
- Getting users list once again

Go with:

	$ go run *.go
	$ curl http://localhost:8088 -u user:pass
	$ curl http://localhost:8088/v1/users  -H"Authorization: Basic: dXNlcjpwYXNz"
	$ curl  -H"Content-Type: application/x-www-form-urlencoded" -H"Authorization: Basic: dXNlcjpwYXNz" \
	"http://localhost:8088/v1/users/new" -X POST -d'name="test"&lastname="test2"&email=test@mail.com&city=SanFrancisco&birthyear=1990'
	$ curl http://localhost:8088/v1/users  -H"Authorization: Basic: dXNlcjpwYXNz"


TODO
- channels, routines - > Streams (messages sent from user to user?)
- JWT Support
- some MiddleWare for Logging in maybe?
- mutlicluster architecture support

