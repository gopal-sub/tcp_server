package server

import (
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
	"tcp-chat/internal/client"
	"tcp-chat/internal/framer"
	"tcp-chat/internal/room"
	"tcp-chat/internal/processQueue"
	"time"
)


type Server struct{
	address string
	listner net.Listener
	rooms []*room.Room
	Client map[*client.Client]struct{} //global list of all clients and client can exist with out being in a room
	mu sync.Mutex
	Waitgrp sync.WaitGroup
}

func NewServer(address string) (*Server, error){
	listener, err:= net.Listen("tcp", address)

	if err != nil {
		return nil, err
	}

	return &Server{address: address, listner: listener}, nil
}

func (s *Server) Start() error{
	TimeoutFrequency := 3*time.Minute

	s.CreateGlobalClientList()
	go s.TimeoutChecker(TimeoutFrequency)
	go s.PingClients(TimeoutFrequency)
	// test graceful shutdown
	// go s.Shutdown()
	
	for {
		// go s.TestClose()
		conn, err := s.listner.Accept()
		if err != nil{
			if errors.Is(err, net.ErrClosed){
				return ServerShutdown
			}
			// netErr.Temporary() is depricated
			// if netErr, ok := err.(net.Error); ok && netErr.Temporary(){
			// 	time.Sleep(10 * time.Millisecond)
			// 	continue
			// }
			return err
		}
		// https://www.youtube.com/watch?v=f6kdp27TYZs
		go s.HandleConnection(conn)
		

	}
}
func (s *Server)Shutdown(){
	s.Waitgrp.Add(1)
	time.Sleep(10*time.Second)
	
	s.listner.Close()
	fmt.Println("shutting down server")
	for c := range s.Client{
		s.RemoveClientByServer(c)
	}

}


func (s *Server) HandleConnection(conn net.Conn){
	fmt.Println("handleconn")
	// for graceful removal of client
	s.Waitgrp.Add(1)
	defer s.Waitgrp.Done()

	feed := framer.NewFramer()
	defer conn.Close()
	// client_id := conn.RemoteAddr()
	buffer := make([]byte, 1024)
	client := client.CreateClient(conn)
	s.AddToGlobalClientList(client)

	for{

		n, err := conn.Read(buffer)
		if err == io.EOF{
			fmt.Println("client disconnected");
			s.RemoveClientFromAllRooms(client)
			return
		}
		if !client.Active{
			return
		}
		if err != nil {
			fmt.Println("connection error:", err)
			s.RemoveClientFromAllRooms(client)
			return
		}
		messages := feed.Feed(buffer[:n])
		fmt.Println(messages)
		// newJob := NewJob(messages, client, s)
		// newJob.JobProcessor()

	}

}

