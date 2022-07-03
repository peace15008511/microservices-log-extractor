package main

import (
	"io"
	adminUsers "microservices-log-extractor/admin"
	"microservices-log-extractor/users"
	"net/http"
)

func testserver(response http.ResponseWriter, Request *http.Request) {
	io.WriteString(response, "Server is up and running!!")
}
func addUser(response http.ResponseWriter, Request *http.Request) {
	io.WriteString(response, adminUsers.CreateUserProfile())
}
func loginUser(response http.ResponseWriter, Request *http.Request) {
	io.WriteString(response, users.Login("Peace", "1234"))
}
func getMicroservices(response http.ResponseWriter, Request *http.Request) {
	io.WriteString(response, "get a list of microservices on the specified server should be on one server at the moment")
}
func getServers(response http.ResponseWriter, Request *http.Request) {
	io.WriteString(response, "get all servers stored on mongodb")
}
func getLogs(response http.ResponseWriter, Request *http.Request) {
	io.WriteString(response, "get logs directly by running ssh to the server and running commands")
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", testserver)
	mux.HandleFunc("/signup", addUser)
	mux.HandleFunc("/signin", loginUser)
	mux.HandleFunc("/services", getMicroservices)
	mux.HandleFunc("/servers", getServers)
	mux.HandleFunc("/logs", getLogs)

	http.ListenAndServe(":2019", mux)
}
