# GoCaast REST api (users)

This Application reads and write to Cassandra db by gocql lib.
this is api for POSTING GETTING / writing to Cassandra database.

- Api listens on port 8088 and localhost. see /main.go
- Basic Auth is implemented (-u or authorization: basic header). See Auth/basicAuth.go
- There is a router and one router.go file with all the endpoints defined.
- GET and POST methods are supported as for now
- things are logged with unix timestamps (every REST request 

prerequisites:
	- Cassandra Keyspace needs to be created - in Cass/main.go it's looking for Cluster name "usersmessages"
	- Cassandra Node has to exist first (certainly)



EXAMPLE USAGE (getting user's list from Cluster and Creating new one with data provided as in example):

$ go run *.go
$ curl http://localhost:8088/v1/users -u user:pass
$ curl -X POST -H"Content-Type: application/x-www-form-urlencoded" -H"Authorization: Basic "
curl  -H"Content-Type: application/x-www-form-urlencoded" -H"Authorization: Basic: dXNlcjpwYXNz" \
"http://localhost:8088/v1/users/new" -X POST -d'name="test"&lastname="test2"&email=test@mail.com&city=SanFrancisco&birthyear=1990'


TODO

- channels, routines - > Streams (messages sent from user to user?)
- mutlicluster architecture support
