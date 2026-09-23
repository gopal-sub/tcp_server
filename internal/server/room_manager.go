package server

import (
	"tcp-chat/internal/room"
	"tcp-chat/internal/client"
)

//crud methods add mutux to these methods

func (s *Server) GetRoom(room string) *room.Room{
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, val := range s.rooms{
		if val.Name == room{
			return val;
		}

	}
	return nil
	
}

func (s *Server) AddRoomToServer(room *room.Room){
	s.mu.Lock()
	defer s.mu.Unlock()

	s.rooms = append(s.rooms, room)
}

func (s *Server) RemoveClientFromAllRooms(client *client.Client) {
	s.mu.Lock()
	defer s.mu.Unlock()


	roomsClientExistsIn := s.RoomsClientExistsIn(client)
	for _, room_val := range roomsClientExistsIn {
		room_val.RemoveClientFromRoom(client)
	}
}



//helper methods
// do not add a mutex to these methods

func (s *Server) RoomsClientExistsIn(client *client.Client)[]*room.Room{
	//dont add a lock here as this method is called by
	//helper fn


	roomsClientExistsIn := []*room.Room{}
	for _ ,room := range s.rooms{
		if room.DoesClientExistInRoom(client){
			roomsClientExistsIn = append(roomsClientExistsIn, room)
		}
		
	}
	return roomsClientExistsIn
}

