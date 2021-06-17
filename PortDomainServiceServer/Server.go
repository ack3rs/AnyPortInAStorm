package main

import (
	"net"
	"os"

	"github.com/acky666/AnyPortInAStorm/PortController"
	config "github.com/acky666/AnyPortInAStorm/PortDomainServiceServer/config"
	"github.com/acky666/AnyPortInAStorm/PortDomainServiceServer/database"
	"google.golang.org/grpc"

	l "github.com/acky666/ackyLog"
)

var VERSION = "v0.0.1"
var GitCommit = "Development"

func main() {

	l.INFO("Starting RPC Server VERSION: [F-CYAN]%s[F-NORMAL] ENV: [F-CYAN]%s[F-NORMAL] on port [F-CYAN]%s[F-NORMAL]", config.VERSION, config.GitCommit, config.Get("SERVERPORT"))

	// Define a listner
	Listener, err := net.Listen("tcp", config.Get("SERVERPORT"))
	if err != nil {
		l.ERROR("failed to listen (on %s): %v", config.Get("SERVERPORT"), err)
		os.Exit(1)
	}

	// Define the RPC Service
	s := PortController.Server{}

	// Since we don't really want to include the logic in the PortController on how to save the details to the database,  that's the servers job.
	// We Tell the Server how to save.

	s.Save = database.Save
	grpcServer := grpc.NewServer()

	//TODO:  Catch OS signals and shutdown nicely.

	// Go Go Go....
	PortController.RegisterPortQueryServiceServer(grpcServer, &s)
	ServerError := grpcServer.Serve(Listener)
	if ServerError != nil {
		l.ERROR("failed to Serve: %v", err)
		os.Exit(1)
	}

}
