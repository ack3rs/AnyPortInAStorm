package PortController

import (
	l "github.com/acky666/ackyLog"
	"golang.org/x/net/context"
)

type Server struct {
	Save func(string, PortInfo) (bool, error)
}

func (s *Server) UpdatePort(ctx context.Context, port *Port) (*Port, error) {
	l.INFO("Received UpdatePort Request %s %+v", port.PortName, port.PortInfo)

	// Decode
	P, err := UnMarshalPortInfo([]byte(port.PortInfo))
	if err != nil {
		l.ERROR("Error Saving %v failed to decode", err)
	}

	// Save
	result, err := s.Save(port.PortName, P)
	if err != nil {
		l.ERROR("Error Saving %v", err)
	}
	l.INFO("✅ %v %v", port.PortName, result)

	return &Port{PortName: port.PortName}, nil
}

func (s *Server) GetPort(ctx context.Context, port *Port) (*Port, error) {
	l.INFO("Received GetPort Request %s", port.PortName)
	return &Port{PortName: "Test"}, nil
}

func (s *Server) mustEmbedUnimplementedPortQueryServiceServer() {
	l.WARNING("Unimplemented")
}
