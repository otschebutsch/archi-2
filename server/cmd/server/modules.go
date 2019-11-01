package main

import (
	//"github.com/google/wire"
	"test/server/computers"
)

// ComposeApiServer will create an instance of ComputerServer according to providers defined in this file.
func ComposeApiServer(port HttpPortNumber) (*ComputerServer, error) {
	db, err := NewDbConnection()
	if err != nil {
		return nil, err
	}
	dataBase := computers.NewStore(db)
	httpHandlerFunc := computers.HttpHandler(dataBase)
	computerServer := &ComputerServer{
		Port:           port,
		ComputersHandler: httpHandlerFunc,
	}
	return computerServer, nil
}
