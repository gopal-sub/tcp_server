package server

import (
	"fmt"
	"io"
	"net"
	"tcp-chat/internal/framer"
	"tcp-chat/internal/protocol"
	"tcp-chat/internal/client"
	"tcp-chat/internal/room"
)


type Server struct{
	address string
	listner net.Listener
	rooms []*room.Room
}

func NewServer(address string) (*Server, error){
	listener, err:= net.Listen("tcp", address)

	if err != nil {
		return nil, err
	}

	return &Server{address: address, listner: listener}, nil
}


func (s *Server) Start() error{
	
	for {
		conn, err := s.listner.Accept()
		if err != nil{
			return err
		}
		// https://www.youtube.com/watch?v=f6kdp27TYZs
		go s.HandleConnection(conn)

	}
}


func (s *Server) HandleConnection(conn net.Conn){
	feed := framer.NewFramer()
	defer conn.Close()
	// client_id := conn.RemoteAddr()
	buffer := make([]byte, 1024)
	client := client.CreateClent(conn)
	for{

		n, err := conn.Read(buffer)
		if err == io.EOF{
			fmt.Println("client disconnected");
			return
		}
		if err != nil {
			fmt.Println("client crashed");
			return
		}
		messages := feed.Feed(buffer[:n])
		for _, message := range messages{
			command, err := protocol.Parser(message)
			if err != nil{
				conn.Write([]byte("invalid command"))
			}
			switch command.Type{
				case protocol.JOIN:
					// add user to room
					_ , exists := s.GetRoom(command.Arg)
					if !exists{
						// create the room
						room := room.CreateRoom(command.Arg)
						room.AddClientToRoom(client)
					}else {
						// add the client to the room
					}
				case protocol.MESSAGE:

				case protocol.QUIT:
				
			}


		}
		



	}

}

func (s *Server) GetRoom(room string) {
	for _, val := range s.rooms{
		if val.Name == room{
			
		} 

	}
	return val, exists
	
}