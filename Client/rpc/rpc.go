package rpc

import (
	"context"
	"os"

	"github.com/acky666/AnyPortInAStorm/Client/config"
	pc "github.com/acky666/AnyPortInAStorm/PortController"
	l "github.com/acky666/ackyLog"
	"google.golang.org/grpc"
)

var conn *grpc.ClientConn
var connectionerr error

func init() {

	conn, connectionerr = grpc.Dial(config.Get("RPC"), grpc.WithInsecure())
	if connectionerr != nil {
		l.ERROR("Error when calling pdss rpc server: %s", connectionerr)
		os.Exit(1)
	}

}

func Shutdown() {
	conn.Close()
}

func UpdatePort(portName string, portInfo string) error {

	rpcClient := pc.NewPortQueryServiceClient(conn)

	response, err := rpcClient.UpdatePort(context.Background(), &pc.Port{PortName: portName, PortInfo: portInfo})
	if err != nil {
		l.ERROR("Error when calling UpdatePort: %v", err)
		return err
	}

	l.INFO("RPC Response was %v", response)
	return err
}

func GetPort(portName string) (pc.PortInfo, error) {

	return pc.PortInfo{}, nil
}
