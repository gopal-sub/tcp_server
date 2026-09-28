package server

import (
	"fmt"
	"tcp-chat/internal/client"
	"time"
)


func (s *Server)CreateGlobalClientList(){
	s.Client = make(map[*client.Client]struct{})
}
func (s *Server)AddToGlobalClientList(client *client.Client){
	
	s.Client[client] = struct{}{}
}

func (s *Server)RemoveClientTimeout(client *client.Client){
	
	
	s.RemoveClientFromAllRooms(client)
	delete(s.Client, client)
	client.CloseClientByServer()
	client.Conn.Close()
}

func (s *Server) TimeoutChecker(){
	freq := 10 * time.Second


	ticker := time.NewTicker(freq)
	for {
		<-ticker.C
		for c, _ := range s.Client{
			if time.Since(c.LastActivity) > freq {
				fmt.Println("client removed")
				s.RemoveClientTimeout(c)
			}
		}
	}
}