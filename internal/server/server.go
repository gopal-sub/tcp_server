package server

import (
	"fmt"
	"io"
	"net"
	"tcp-chat/internal/client"
	"tcp-chat/internal/framer"
	"tcp-chat/internal/protocol"
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
					for _,room := range roomsClientExistsIn{
						room.BroadcastMessageExceptClient(client, command.Arg)
					}
					

				case protocol.QUIT:
					roomsClientExistsIn := s.RoomsClientExistsIn(client)
					for _, room_val := range roomsClientExistsIn {
						room_val.RemoveClientFromRoom(client)
					}

				
			}


		}
		



	}

}

func (s *Server) GetRoom(room string) *room.Room{
	for _, val := range s.rooms{
		if val.Name == room{
			return val;
		}

	}
	return nil
	
}

func (s *Server) AddRoomToServer(room *room.Room){
	s.rooms = append(s.rooms, room)
}

func (s *Server) RoomsClientExistsIn(client *client.Client)[]*room.Room{
	roomsClientExistsIn := []*room.Room{}
	for _ ,room := range s.rooms{
		if room.DoesClientExistInRoom(client){
			roomsClientExistsIn = append(roomsClientExistsIn, room)
		}
		
	}
	return roomsClientExistsIn
}