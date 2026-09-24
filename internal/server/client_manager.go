package server

import (
	"tcp-chat/internal/client"
)



func (s *Server)AddToGlobalClientList(client *client.Client){
	s.Client[client] = struct{}{}
}

func (s *Server)RemoveClientTimeout(client *client.Client){
	
	
	s.RemoveClientFromAllRooms(client)
	delete(s.Client, client)
}