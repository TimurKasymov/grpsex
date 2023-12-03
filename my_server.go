package main

import (
	"github.com/TimurKasymov/grpsex/myPkgName"
	"golang.org/x/net/context"
	"log"
)

// should implement the interface myPkgName.InvoicerServer
type myGRPCServer struct {
	// type embedded to comply with Google lib
	myPkgName.UnimplementedInvoicerServer
}

func (m *myGRPCServer) Create(ctx context.Context, request *myPkgName.CreateRequest) (*myPkgName.CreateResponse, error) {
	log.Println("Create called")
	return &myPkgName.CreateResponse{Pdf: []byte("TODO")}, nil
}
