package server

import (
	"fmt"
	"io"
	"net"
	"tcp-chat/internal/framer"
	"tcp-chat/internal/protocol"
)


type Server struct{
	address string
	listner net.Listener
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
		go HandleConnection(conn)

	}
}


func HandleConnection(conn net.Conn){
	feed := framer.NewFramer()
	defer conn.Close()
	// client_id := conn.RemoteAddr()
	buffer := make([]byte, 1024)
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
			if command.Type == protocol.JOIN{
				// add user to room
			}else if command.Type == protocol.MESSAGE{
				// send message
			}else if command.Type == protocol.QUIT{
				// remove the user
			}

		}
		



	}

}