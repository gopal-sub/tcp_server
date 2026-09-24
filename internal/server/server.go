package server

import (
	"fmt"
	"io"
	"net"
	"sync"
	"tcp-chat/internal/client"
	"tcp-chat/internal/framer"
	"tcp-chat/internal/protocol"
	"tcp-chat/internal/room"
)


type Server struct{
	address string
	listner net.Listener
	rooms []*room.Room
	Client map[*client.Client]struct{} //global list of all clients and client can exist with out being in a room
	mu sync.Mutex
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
	s.AddToGlobalClientList(client)

	for{

		n, err := conn.Read(buffer)
		if err == io.EOF{
			fmt.Println("client disconnected");
			s.RemoveClientFromAllRooms(client)
			return
		}
		if err != nil {
			fmt.Println("connection error:", err)
			s.RemoveClientFromAllRooms(client)
			return
		}
		messages := feed.Feed(buffer[:n])
		fmt.Println("message")
		fmt.Println(messages)
		for _, message := range messages{
			command, err := protocol.Parser(message)
			if err != nil{
				conn.Write([]byte("invalid command\n"))
				continue
			}
			switch command.Type{
				case protocol.JOIN:
					
					// add user to room
					roomExists := s.GetRoom(command.Arg)

					// case 1 => room does not exist
					if roomExists == nil{
					// room does not exist create room and add conn to room
						newRoom := room.CreateRoom(command.Arg)
						newRoom.AddClientToRoom(client)
						s.AddRoomToServer(newRoom)
						break
					}
					// case 2 => client join room the client already is in 
					if roomExists.DoesClientExistInRoom(client){
						client.SendMessage("Already in room")
       					break
					}


					//case 3 => cleint joins new room
					roomsClientExistsIn := s.RoomsClientExistsIn(client)

					
					// remove client from any existing room
					for _, room_val := range roomsClientExistsIn {
						room_val.RemoveClientFromRoom(client)
					}
					roomExists.AddClientToRoom(client)	
					
					


					
				case protocol.MESSAGE:
					roomsClientExistsIn := s.RoomsClientExistsIn(client)
					if len(roomsClientExistsIn) ==0 {
						client.SendMessage("You are not in any room")
						break
					}

					for _,room := range roomsClientExistsIn{
						room.BroadcastMessageExceptClient(client, command.Arg)
					}
					

				case protocol.QUIT:
					s.RemoveClientFromAllRooms(client)

				
			}


		}
		



	}

}

