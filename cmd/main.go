package main

import (
	"RolesBased/gen/restapi"
	"RolesBased/gen/restapi/operations"
	"RolesBased/handlers"
	"flag"
	"github.com/go-openapi/loads"
	"log"
)
var portFlag = flag.Int("port", 3000, "Port to run this service on")
func main() {
	// load embedded swagger file
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}
	// create new service API
	api := operations.NewRoleBaseAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()
	// parse flags
	flag.Parse()
	// set the port this service will be run on
	server.Port = *portFlag
	api.GettingRoleGetroleHandler = handlers.RoleBase()
	// serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

