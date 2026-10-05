package server

import (
	"net"
	"testing"
)






func TestClientTCPConnection(t *testing.T){
	svrPort := ":3000"


	svr, err := StartServer(svrPort)
	defer svr.Shutdown()
	
	if err != nil {
		t.Fatal("The svr could not start")
	}

	_ , err = net.Dial("tcp",svrPort)

	if err != nil {
		t.Fatal("client could not connect")
	}

}

// helper methods

func StartServer(port string) (*Server, error){

	svr, err := NewServer(port)

	

	if err!=nil{
		return nil, err
	}

	go svr.Start()
	
	return svr, nil
}